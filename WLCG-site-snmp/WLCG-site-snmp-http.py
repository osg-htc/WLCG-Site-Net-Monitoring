#!/bin/env python
#
# Example code to query for interface information from network devices,
# add relevant values together and output JSON file match WLCG Mon TF needs
#
#  See https://gitlab.cern.ch/wlcg-doma/site-network-information
#
# Requires easysnmp package (see https://easysnmp.readthedocs.io/en/latest/)
#
#   CentOS example install#   (assuming you have python3)
#   sudo yum install gcc python3-devel net-snmp-devel
#   sudo pip3 install easysnmp
#
from easysnmp import Session
from datetime import datetime, timezone
import time
import json
import os
import logging, sys
import argparse
import subprocess
import csv
from http.server import HTTPServer, BaseHTTPRequestHandler

arg_parser = argparse.ArgumentParser()

arg_parser.add_argument('--snmp-config',
                        dest='snmp_config_file',
                        default='site-snmp-config.json',
                        help='Site snmp config file name. Default config_file.')
arg_parser.add_argument('--install_location',
                        dest='install_loc',
                        default=os.environ['PWD'],
                        help='Install Location. Default: PWD')
arg_parser.add_argument('--output-file',
                        dest='ooutput_file',
                        default='site-output.json',
                        help='Site output file name. Default site-output.json.')
arg_parser.add_argument('--debug',
                        default=False,
                        action='store_true',
                        help='Add debug printing')

args, unknown = arg_parser.parse_known_args()


# Note: interfaces have the ifDescr value holding the right index. This is 
# needed for the INDICES python dictionary below.
#
#    To see what your device provides try:
#       snmpwalk -v2c -c <COMMUNITY> <HOST> IF-MIB::ifDescr
#
#    For the example below, my Cisco provides:
#
#      IF-MIB::ifDescr.436233216 = STRING: Ethernet1/51
#      IF-MIB::ifDescr.436233728 = STRING: Ethernet1/52
#                      ^^^^^^^^^           ^^^^^^^^^^^^
#                      Index               Interface description

# Set up logging and specify level as logging.<LEVEL> with <LEVEL>=DEBUG,INFO,WARN,ERROR...
logging.basicConfig(stream=sys.stderr, level=logging.DEBUG)

                        
# ---------------------------------------------------------------------------
# ================= CUSTOMIZE THIS SECTION ==================================
# ---------------------------------------------------------------------------

# Install location (directory that hosts WLCG-site-snmp.py script)
INSTALL_LOC=args.install_loc
MESSAGE="Install Location INSTALL_LOC: "+INSTALL_LOC
logging.info(MESSAGE)

# Define the set of switches and ports that represent the site "border"
#   You will need to find the correct SNMP indices to use.  See info above
site_snmp_config=json.load(open("{}/{}".format(INSTALL_LOC,args.snmp_config_file)))

# ------------ Define the needed 64-bit OIDs for In/Out Octets --------------
ifHCInOctets = ".1.3.6.1.2.1.31.1.1.1.6"
ifHCOutOctets = ".1.3.6.1.2.1.31.1.1.1.10"

# ----------- For reference, here are the 32-bit OIDs for In/Out Octets -----
# ----------- If the 64-bit OIDs above don't exist you can try these --------
IfInOctets = ".1.3.6.1.2.1.2.2.1.10"
IfOutOctets = ".1.3.6.1.2.1.2.2.1.16"

InStartTime = {}
InEndTime = {}
ifInCntrStart = {}
ifInCntrEnd = {}

OutStartTime = {}
OutEndTime = {}
ifOutCntrStart = {}
ifOutCntrEnd = {}

# Sleep interval between loop executions (in seconds)
INTERVAL = 60

# Announce service start up
MESSAGE=" WLCG site traffic monitor started at " + datetime.now(timezone.utc).isoformat()
print(MESSAGE)
MESSAGE="  -------  traffic monitor directory " + INSTALL_LOC
print(MESSAGE)

def snmp_get_data(INDICES,COMM):
    MESSAGE="INDICES:"+json.dumps(INDICES, indent=4)
    logging.info(MESSAGE)
    MESSAGE="COMM:"+json.dumps(COMM, indent=4)
    logging.info(MESSAGE)

    # ------------ Loop over interfaces, gathering data -------------------------
    MonInterfaces = []
    InBytesPerSec = 0
    OutBytesPerSec = 0
    # Loop over all devices and interfaces, adding up Octets
    for host, interface in INDICES.items():
        MESSAGE=" host: " + host + " comm: " + COMM[host]
        logging.debug(MESSAGE)
        session = Session(hostname=host, community=COMM[host], version=2)
        # ----------- Gather In/Out Octets and Associated time ----------------------
        for desc in interface:
            MESSAGE=" Interface: " + desc + " index: ", interface[desc]
            logging.debug(MESSAGE)
            KEY = host + "_" + desc
            MonInterfaces.append(KEY)
# Get end info for IN
            ifInCntrEnd[KEY] = int(session.get((ifHCInOctets, interface[desc])).value)
            InEndTime[KEY] = datetime.now().isoformat()
            MESSAGE=" Key:" + KEY + "In End Counter:" + str(ifInCntrEnd[KEY]) + " End Time:" + InEndTime[KEY]
            logging.debug(MESSAGE)
            if InStartTime.get(KEY) is not None:
                # ------------------------ Calculate rate and swap variables
                time_diff = datetime.strptime(InEndTime[KEY],'%Y-%m-%dT%H:%M:%S.%f')-datetime.strptime(InStartTime[KEY],'%Y-%m-%dT%H:%M:%S.%f')
                Rate = (ifInCntrEnd[KEY] - ifInCntrStart[KEY]) / time_diff.total_seconds()
                InBytesPerSec = InBytesPerSec + Rate
# Get new start info for In
            InStartTime[KEY] = InEndTime[KEY]
            ifInCntrStart[KEY] = ifInCntrEnd[KEY]
# Get end info for Out 
            ifOutCntrEnd[KEY] = int(session.get((ifHCOutOctets, interface[desc])).value)
            OutEndTime[KEY] = datetime.now().isoformat()
            # ------------------------ Calculate rate and swap variables
            if OutStartTime.get(KEY) is not None:
                time_diff = datetime.strptime(OutEndTime[KEY],'%Y-%m-%dT%H:%M:%S.%f')-datetime.strptime(OutStartTime[KEY],'%Y-%m-%dT%H:%M:%S.%f')
                Rate = (ifOutCntrEnd[KEY] - ifOutCntrStart[KEY]) / time_diff.total_seconds()
                OutBytesPerSec = OutBytesPerSec + Rate
# Get new start info for Out
            OutStartTime[KEY] = OutEndTime[KEY]
            ifOutCntrStart[KEY] = ifOutCntrEnd[KEY]
    print("{}, {}".format(InBytesPerSec,OutBytesPerSec))    
    if InBytesPerSec != 0 or OutBytesPerSec != 0:
        # Need time in ISO 8601 format for UTC
        LastTime_us = datetime.now(timezone.utc).isoformat()
        output = {
            "Description": "Network statistics for {}".format(site_snmp_config['site']),
            "UpdatedLast": LastTime_us,
            "InBytesPerSec": InBytesPerSec,
            "OutBytesPerSec": OutBytesPerSec,
            "UpdateInterval": str(INTERVAL) + " seconds",
            "MonitoredInterfaces": MonInterfaces,
        }
        logging.debug(json.dumps(output))
        return json.dumps(output)
#----------------------------------
# HTTP server section             
#----------------------------------
class WebRequestHandler(BaseHTTPRequestHandler):
    # ...
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-Type", "application/json")
        self.end_headers()
        snmp_output=snmp_get_data(INDICES=site_snmp_config['indices'],COMM=site_snmp_config['comm'])
        self.wfile.write(json.dumps(snmp_output).encode('utf-8')) # Read the file and send the contents 
#----------------------------------
# Main server section             
#----------------------------------

if __name__ == "__main__":

    server = HTTPServer(("0.0.0.0", 8000), WebRequestHandler)
    server.serve_forever()

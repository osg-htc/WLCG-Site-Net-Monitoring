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
import os.path
import logging, sys

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
logging.basicConfig(stream=sys.stderr, level=logging.WARN)

# ---------------------------------------------------------------------------
# ================= CUSTOMIZE THIS SECTION ==================================
# ---------------------------------------------------------------------------

# Install location (directory that hosts WLCG-site-snmp.py script)
INSTALL_LOC="/root/site-network-information/WLCG-site-snmp/"
MESSAGE="Install Location INSTALL_LOC: "+INSTALL_LOC
logging.info(MESSAGE)

# Location of the output JSON (need to add suitable URL to CRIC for access)
JSONOUTFILE = "/var/www/html/aglt2-netmon.json"
MESSAGE="JSON Outfile JSONOUTFILE: "+JSONOUTFILE
logging.info(MESSAGE)

# Sleep interval between loop executions (in seconds)
INTERVAL = 60

MESSAGE="Update INTERVAL: "+str(INTERVAL)
logging.info(MESSAGE)

# Define the set of switches and ports that represent the site "border"
#   You will need to find the correct SNMP indices to use.  See info above
INDICES = {
    "aglt2-rtr-1.local": {
        "Ethernet1/48": 436231680,
        "Ethernet1/51": 436233216,
        "Ethernet1/52": 436233728,
    },
    "aglt2-rtr-2.local": {"Ethernet1/51": 436233216, "Ethernet1/52": 436233728},
}

# Define SNMP community strings for each device as needed
COMM = {"aglt2-rtr-1.local": "mycomm1", "aglt2-rtr-2.local": "mycomm2"}
# ---------------------------------------------------------------------------
# ================= END OF CUSTOMIZATION ====================================
# ---------------------------------------------------------------------------

if os.path.isfile(INSTALL_LOC+"indices.json"):
    with open(INSTALL_LOC+"indices.json", "r") as infile:
        INDICES=json.load(infile)
        MESSAGE="Read INDICES from "+INSTALL_LOC+"indices.json..."
        logging.info(MESSAGE)
else:
    with open(INSTALL_LOC+"indices.json", "w") as outfile:
        json.dump(INDICES, outfile, indent=4)
        MESSAGE="Wrote INDICES to "+INSTALL_LOC+"indices.json..."
        logging.warn(MESSAGE)

if os.path.isfile(INSTALL_LOC+"comm.json"):
    with open(INSTALL_LOC+"comm.json", "r") as infile:
        COMM=json.load(infile)
        MESSAGE="Read COMM from "+INSTALL_LOC+"comm.json..."
        logging.info(MESSAGE)
else:
    with open(INSTALL_LOC+"comm.json", "w") as outfile:
        json.dump(COMM, outfile, indent=4)
        MESSAGE="Wrote COMM to "+INSTALL_LOC+"comm.json..."
        logging.warn(MESSAGE)

MESSAGE="INDICES:"+json.dumps(INDICES, indent=4)
logging.info(MESSAGE)
MESSAGE="COMM:"+json.dumps(COMM, indent=4)
logging.info(MESSAGE)

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

# Announce service start up
MESSAGE=" WLCG site traffic monitor started at " + datetime.now(timezone.utc).isoformat()
print(MESSAGE)
MESSAGE="  -------  traffic monitor directory " + INSTALL_LOC
print(MESSAGE)
MESSAGE="  -------  traffic JSON output " + JSONOUTFILE
print(MESSAGE)

# ------------ Service Loop -------------------------------------------------
while True:
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
    if InBytesPerSec != 0 or OutBytesPerSec != 0:
        # Need time in ISO 8601 format for UTC
        LastTime_us = datetime.now(timezone.utc).isoformat()
        output = {
            "Description": "Network statistics for AGLT2",
            "UpdatedLast": LastTime_us,
            "InBytesPerSec": InBytesPerSec,
            "OutBytesPerSec": OutBytesPerSec,
            "UpdateInterval": str(INTERVAL) + " seconds",
            "MonitoredInterfaces": MonInterfaces,
        }
        logging.info(json.dumps(output))
        # Directly from dictionary
        with open(JSONOUTFILE, "w") as outfile:
            json.dump(output, outfile)
    # Sleep till next loop run
    time.sleep(INTERVAL)

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
import json

# Note: interfaces have the ifDescr value holding the right index
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

# ---------------------------------------------------------------------------
# ================= CUSTOMIZE THIS SECTION ==================================
# ---------------------------------------------------------------------------
# Location of the output JSON file
JSONOUTFILE = "/var/www/html/aglt2-netmon.json"

# Sleep interval between loop executions (in seconds)
INTERVAL = 60

# ------------ Input the set of network devices and interfaces to query -----
INDICES = {
    "aglt2-rtr-1.local": {
        "Ethernet1/48": 436231680,
        "Ethernet1/51": 436233216,
        "Ethernet1/52": 436233728,
    },
    "aglt2-rtr-2.local": {"Ethernet1/51": 436233216, "Ethernet1/52": 436233728},
}

# ------------ Add the snmp community used to access each device ------------
COMM = {"aglt2-rtr-1.local": "mycommunity", "aglt2-rtr-2.local": "mycommunity"}
# ---------------------------------------------------------------------------
# ================= END OF CUSTOMIZATION ====================================
# ---------------------------------------------------------------------------

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

# ------------ Service Loop -------------------------------------------------
while True:
    # ------------ Loop over interfaces, gathering data -------------------------
    MonInterfaces = []
    InBytePerSec = 0
    OutBytesPerSec = 0
    # Loop over all devices and interfaces, adding up Octets
    for host, interface in INDICES.items():
        #               print(" host: " + host + " comm: " + SNMPCOMM)
        session = Session(hostname=host, community=COMM[host], version=2)
        # ----------- Gather In/Out Octets and Associated time ----------------------
        for desc in interface:
            print(" Interface: " + desc + " index: ", interface[desc])
            KEY = host + "_" + desc
            MonInterfaces.append(KEY)
            ifInCntrEnd[KEY] = int(session.get((ifHCInOctets, interface[desc])).value)
            InEndTime[KEY] = datetime.now(timezone.utc).isoformat()
            if InStartTime.get(KEY) is not None:
                # ------------------------ Calculate rate and swap variables
                dTime = 1000
                Rate = (ifInCntrEnd[KEY] - ifInCntrStart[KEY]) / dTime
                InBytesPerSec = InBytesPerSec + Rate
                InStartTime[KEY] = InEndTime[KEY]
                ifInCntrStart[KEY] = ifInCntrEnd[KEY]
            ifOutCntrEnd[KEY] = int(session.get((ifHCOutOctets, interface[desc])).value)
            OutEndTime[KEY] = datetime.now(timezone.utc).isoformat()
            LastTime_us = OutEndTime[KEY]
            # ------------------------ Calculate rate and swap variables
            if OutStartTime.get(KEY) is not None:
                dTime = 10000
                Rate = (ifOutCntrEnd[KEY] - ifOutCntrStart[KEY]) / dTime
                OutBytesPerSec = OutBytesPerSec + Rate
                OutStartTime[KEY] = OutEndTime[KEY]
                ifOutCntrStart[KEY] = ifOutCntrEnd[KEY]
    if InBytesPerSec != 0 or OutBytesPerSec != 0:
        # Need time in ISO 8601 format
        output = {
            "Description": "Network statistics for AGLT2",
            "UpdatedLast": LastTime_us,
            "InBytesPerSec": InBytesPerSec,
            "OutBytesPerSec": OutBytesPerSec,
            "UpdateInterval": INTERVAL + " seconds",
            "MonitoredInterfaces": MonInterfaces,
        }
    # Directly from dictionary
    with open(JSONOUTFILE, "w") as outfile:
        json.dump(output, outfile)
    # Sleep till next loop run
    time.sleep(INTERVAL)

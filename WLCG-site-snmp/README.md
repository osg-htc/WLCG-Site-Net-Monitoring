# Site Traffic Monitoring Service README
This directory contains the **Site Traffic Monitoring Service** example and is intended to be deployed at WLCG sites to gather simple network statistics from the site’s network border device(s).  It contains a python3 script that can be configured to read multiple interfaces via SNMP, add the IN and OUT traffic (In Bytes/sec) up and save the output in a standard formatted JSON file for use by WLCG Monitoring.
## Installation Considerations
This software should be installed on a system that has:
Network access to the site’s border devices via SNMP
A web server (alternatively the output could be copied to a web server via a suitable script)
Python3 and git installed
Pick a location on that system to deploy to which we will call INSTALL_LOC, e.g., 
```
export INSTALL_LOC=~/my-site-monitoring
```

Then your can get the project and copy the relevant director to INSTALL_LOC
```
cd /tmp  
git clone https://gitlab.cern.ch/wlcg-doma/site-network-information.git
cp -ar site-network-information/WLCG-site-snmp/* ${INSTALL_LOC}/
```
After getting the starting files in place in INSTALL_LOC, your should make some working copies of the relevant files:
```
cd ${INSTALL_LOC}
cp site-traffic-monitor.service-example site-traffic-monitor.service
cp WLCG-site-snmp-example.py WLCG-site-snmp.py
```
You will need to edit WLCG-site-snmp.py to customize for your use-case.   The variables to be set are described in the table below.
| Variable name | Description | Example |
| -------------------- | -------------------------------------- | ----------------------- |
| INSTALL_LOC | The location the code is installed at | See above |
| INTERVAL | The update interval in seconds | 60 |
| JSONOUTFILE | The name and location of the output file | /var/www/mysite-net-info.json|
| INDICES | A python dictionary holding the devices and interfaces to monitor | See below |
| COMM | The snmp community which has read access to the devices in INDICES | myreadcomm |

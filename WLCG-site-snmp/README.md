*** This directory contains a python3 script that can be configured to read multiple interfaces via SNMP, add the IN and OUT traffic (In Bytes/sec) up and save the output in a standard formatted JSON file for use by WLCG Monitoring

After checking out this directory, copy the following files to make your own version:

cp site-traffic-monitor.service-example site-traffic-monitor.service
cp WLCG-site-snmp-example.py WLCG-site-snmp.py

You will need to edit WLCG-site-snmp.py to customize for your use-case

Variables to set:

 - INSTALL_LOC
 - INTERVAL
 - JSONOUTFILE
 - INDICES
 - COMM

 

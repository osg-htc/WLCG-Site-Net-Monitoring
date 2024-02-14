# Table of Contents 

1. [Site Traffic Monitoring Service README](#site-traffic-monitoring-service-readme)
2. [Installation Considerations](#installation-considerations)
3. [Install Pre-Requisite](#install-pre-requisites)
4. [Installing](#installing)
5. [Configuring](#configuring)
6. [Implementing as a systemd service](#implementing-as-a-systemd-service)
7. [snmp script with an https server](#snmp-script-with-an-https-server)
8. [Register in CRIC](#register-in-cric)

# Site Traffic Monitoring Service README

This directory contains the **Site Traffic Monitoring Service** example and is intended to be deployed at WLCG sites to gather simple network statistics from the site’s network border device(s).  It contains a python3 script that can be configured to read multiple interfaces via SNMP, add the IN and OUT traffic (In Bytes/sec) up and save the output in a standard formatted JSON file for use by WLCG Monitoring.  We desire relatively fine-grained monitoring, if possible, and **should have an update interval of 60 seconds**, unless there is a technical reason not to. 

## Installation Considerations

This software should be installed on a system that has:

- Network access to the site’s border devices via SNMP
- A web server (alternatively the output could be copied to a web server via a suitable script)
- Python3, net-snmp and git installed

Pick a location on that system to deploy to which we will call `INSTALL_LOC`, e.g.,

```
export INSTALL_LOC=~/my-site-monitoring
mkdir -p ${INSTALL_LOC}
```

### Go version

A Go-based version with no external dependencies (other than SystemD for running as a daemon) can be found
under the `go` directory. The accompanying `README.md` together with the Markdown-formatted manpage shed
much more light on how to configure and deploy this flavour.

A thing to note is this version is also distributed as a ready-to-install RPM package.

### Dockerized Version

Thanks to Justin Balcas / Caltech, we have a dockerized version of the software available at the following links:

Docker templates available here: https://github.com/cmscaltech/docker/blob/master/wlcg-site-mon/
Docker image available here: https://hub.docker.com/repository/docker/cmscaltech/wlcg-site-mon/general
CI/CD (Git auto build) is also here: https://github.com/cmscaltech/docker/blob/master/.github/workflows/build-wlcg-site-mon.yml

## Install Pre-Requisites

You should also make sure you have installed any dependencies if needed:

```
yum install net-snmp git net-snmp-devel python3-devel gcc
pip3 install easysnmp==0.2.5
```
<span style="color:red">**Note:**</span> there is a problem installing easysnmp  0.2.6 via pip3 as of September 5, 2023.  We are working on a solution.  For now, the workaround is to force the use of the previous version as above '==0.2.5' during the pip3 install.

## Installing

Then your can get the project and copy the relevant director to `INSTALL_LOC`

```
cd /tmp  
git clone https://gitlab.cern.ch/wlcg-doma/site-network-information.git
cp -ar site-network-information/WLCG-site-snmp/* ${INSTALL_LOC}/
```

After getting the starting files in place in `INSTALL_LOC`, your should make some working copies of the relevant files:

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
| JSONOUTFILE | The path for the output file | /var/www/html/aglt2-netmon.json |
| INDICES | A python dictionary holding the devices and interfaces to monitor | See below |
| COMM | A python dictionary holding device snmp community strings | See below |

## Configuring

To configure your installation, you will need to identify the network interfaces that represent your site’s border.   There may be more than one interface and more than one device hosting border interfaces.   You will likely need to work with your network team to ensure you correctly identify the appropriate interface(s) and get readonly access via SNMP. 
The next step is to confirm that snmp is working for accessing your devices.  Once your have the SNMP community to use, set it as an environment variable:

```
export SNMPCOMM=myreadonlycomm
```

Then we can test access to the switch(es) or router(s) that host the site border interface(s):

```
snmpwalk -v2c -c ${SNMPCOMM} <device> .sysDescr
```

This should return a value:

```
[root@sysprov02 ~]# snmpwalk -v2c -c ${SNMPCOMM} aglt2-rtr-1.local .sysDescr
SNMPv2-MIB::sysDescr.0 = STRING: Cisco NX-OS(tm) nxos.9.3.5.bin, Software (nxos), Version 9.3(5), RELEASE SOFTWARE Copyright (c) 2002-2020 by Cisco Systems, Inc. Compiled 7/20/2020 20:00:00
```

Please check each device that hosts border interfaces to ensure snmp access works.

Once that is verified, you can configure the COMM dictionary in your `WLCG-site-snmp.py` file.  Here is an example for a site with two border devices:

```
COMM = {"aglt2-rtr-1.local": "mycomm1", "aglt2-rtr-2.local": "mycomm2"}
```

We now need to identify the correct SNMP indices for the border interfaces we want to monitor. Continuing with the example of aglt2-rtr-1.local and aglt2-rtr-2.local, let’s say the border interfaces are Ethernet1/51 and Ethernet1/52 on each host.  We can use the following command to identify the correct index to configure in `WLCG-site-snmp.py` for aglt2-rtr-1.local:

```
snmpwalk -v2c -c ${SNMPCOMM} aglt2-rtr-1.local  IF-MIB::ifDescr
```

The example output lists all the indices and interfaces:

```
...
IF-MIB::ifDescr.436233216 = STRING: Ethernet1/51
IF-MIB::ifDescr.436233728 = STRING: Ethernet1/52
                ^^^^^^^^^           ^^^^^^^^^^^^
                Index               Interface description
```

You should find the appropriate index values for all devices hosting border interfaces.  Below is the example from AGLT2:

```
INDICES = {
    "aglt2-rtr-1.local": {
        "Ethernet1/51": 436233216,
        "Ethernet1/52": 436233728,
    },
    "aglt2-rtr-2.local": {
        "Ethernet1/51": 436233216,
        "Ethernet1/52": 436233728
    },
}
```

Feel free to adjust any of the variables as desired.   Once you have them configured, you can run a test:

```
# python3 WLCG-site-snmp.py
 WLCG site traffic monitor started at 2022-05-09T14:41:45.955477+00:00
  -------  traffic monitor directory /root/site-network-information/WLCG-site-snmp/
  -------  traffic JSON output /var/www/html/aglt2-netmon.json
```

Note that the code is set to run forever.  You need to wait for at least `INTERVAL`, CTRL-C and then verify a new output file `JSONOUTFILE` was created.

For AGLT2 it looks like:

```
{
    "Description": "Network statistics for AGLT2",
    "UpdatedLast": "2022-05-09T14:47:23.057942+00:00",
    "InBytesPerSec": 1275982208.1281207,
    "OutBytesPerSec": 1526111205.30333,
    "UpdateInterval": "60 seconds", 
    "MonitoredInterfaces": [
        "aglt2-rtr-1.local_Ethernet1/51",
        "aglt2-rtr-1.local_Ethernet1/52",
        "aglt2-rtr-2.local_Ethernet1/51",
        "aglt2-rtr-2.local_Ethernet1/52"
        ]
}
```

If things look OK we can set it up as a systemd service

## Implementing as a systemd service

You need to edit the systemd service file we created by copying the example:  site-traffic-monitor.service
The only change needed is edit that file and replace `<INSTALL_LOC>` with the actual install location.

Then you can copy this file (as 'root') to /etc/systemd/system/site-traffic-monitor.service

```
sudo cp ${INSTALL_LOC}/site-traffic-monitor.service /etc/systemd/system/site-traffic-monitor.service'
```

You then need to reload systemd, enable the new service and start it

```
sudo systemctl daemon-reload
sudo systemctl enable site-traffic-monitor.service
sudo systemctl start site-traffic-monitor.service
```

While this service runs it should create a new `JSONOUTFILE` every `INTERVAL` seconds. NOTE: set INTERVAL at 60 seconds unless technical issues preclude this.  
If the location of `JSONOUTFILE` is NOT accessible via a web URL, you will need to have some mechanism to move it to a web accessible location.

## snmp script with an https server

If you don't have an http server where to copy the snmp json output at hand you can use the script in snmp-with-http-example. This will do the same as the main example, only will start also an http server that will produce the snmp json output when queried. It has 3 options to define the INSTALL_LOC, to define the config path (default ./), and to change the debug level. All the rest can be configured in a json configuration file which contains the following information:

- site name: your RC site name
- host_cert: location of your host pem certificate to use to start https (use a certificate that can be generally recognised, or you can run this on a standard grid machine)
- host_key: location of your host pem key to use to start https
- host_port: the port you prefer to start the https server on. (do NOT add quotes)
- interval: minimum interval between one snmp request and the other (default 60s) if the elapsed time is less than interval the previous json values will be returned
- comm: communities to fill with uplink interfaces and communities (i.e. password) for those interfaces
- indices: interfaces and the aliases they have in your switches/routers

for the last two config attributes follow the instructions for COMM and INDICES in the [Configuration](#configuration) section above.

To start the server manually it's like the script without https server, simply

```
python3 WLCG-site-snmp-http.py
```
then to test it go to another machine and run this curl command
```
curl --capath <CA_CERT_DIR> https://<your_host>:<your_port>/
```
The first query after the server starts will be null. But after that it will server the results calculated from the previous query. If you restart the server the first query will be null again.

The script also can make use of systemd and there is a systemd example. To install it follow the instructions  in the [Implementing as a systemd service](#implementing-as-a-systemd-service) section above but using the http example in the snmp-with-http-example directory.

*CAVEATS:* This is operationally easier to run than producing a file and copying it to an https server, it can run on a machine with certificates (grid storage or perfsonar) and is likely more reliable as the numbers are caclulated when requested (i.e. it doesn't go out of sync) and access can be restricted only to CERN machines doing the query. easysnmp doesn't easily install on CetnOS7 so you need at least an EL8 machine and python 3.6.8.

## Register in CRIC

The last thing to do is register the URL for the `JSONOUTFILE` in CRIC.
Each network site in CRIC is shown at [CRIC Network Site](https://wlcg-cric.cern.ch/core/netsite/list/).  Find the appropriate network site and update the monitoring URL with the correct location allowing access to your `JSONOUTFILE`

For example

```
Monitoring URL
Monitoring URL that shows real-time network traffic into and out-of this site
https://head01.aglt2.org/aglt2-netmon.json
```

The WLCG Monitoring Task Force plans to use these URLs to grab the JSON monitoring for all participating sites.

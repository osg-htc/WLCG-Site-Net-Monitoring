
---

# UKI-NORTHGRID-MAN-HEP Network Information
This page describes `UKI-NORTHGRID-MAN-HEP` network information for WLCG use. 

LAST UPDATE (when file is changed, please update): 10-May-2022 09:05 Eastern

## Network Overview [Mandatory; can be brief]
We have a single site at the University of Manchester.  100Gb/s router connected to the university with dedicate bandwidth limited at 40Gb/s and that the university is connected to Janet backbone at 100Gb/s.


### Network Description [Optional]

The storage nodes currently have 2x10Gb/s bonded links to the ToR switches which are a mixture of Dell and HPE brands. The oldest Dell switches have a 2x10Gb/s bonded uplink, the newer HPEs and Dells have 40Gb/s and 100Gb/s uplinks to the Tier2 core switches. The WNs have all 10Gb/s 10GBASE-T. The core switches are Arista DCS-7050CX3-32S with the 100Gb/s uplink to the university router is limited at 40Gb/s.

### Peering Description [Optional]
*Please describe how your site connects by responding to the following questions.*  

**How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?**
- **If so, who do you peer with and what is the bandwidth available for this type of traffic?**
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

**If you are connected to LHCONE, who/where do you peer with and at what bandwidth?**
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

**If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?**   
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

**Do you have a peering for research and education networks for non-LHCONE sites?**  
- **If so, who/where do you peer with and at what bandwidth?** 
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

### Network Equipment Details [Optional]
Our site has a mix of Cisco and Dell switches and routers.   We have X Cisco YYY-ZZZZ systems running NXOS X.XXX and Y Dell YYYYY switches running OS10 X.XX

## Network Monitoring [Mandatory]

Sites will need to create a web reachable JSON file that tracks the IN and OUT traffice (In Bytes / second) for their site.
We have provided example code (in python) that can be turned into a `systemd` service which creates the neccessary file. 
Please read and deploy the site traffic monitor described in the WLCG-site-snmp directory in this project or provide your own equivalent.

The format of the JSON file should look something like this:
```
{
Description: "Network statistics for AGLT2",
UpdatedLast: "2022-05-09T18:04:43.395714+00:00",
InBytesPerSec: 3018122809.4014854,
OutBytesPerSec: 3736689086.9365864,
UpdateInterval: "60 seconds",
MonitoredInterfaces: [
"aglt2-rtr-1.local_Ethernet1/48",
"aglt2-rtr-1.local_Ethernet1/51",
"aglt2-rtr-1.local_Ethernet1/52",
"aglt2-rtr-2.local_Ethernet1/51",
"aglt2-rtr-2.local_Ethernet1/52"
]
}
```
This file should be updated at some reasonable interval (once per minute or similar).

It is important to have this network monitoring information from your site's **border:** the location where all traffic has to pass to get into or out of your site.

**This whole section should be replaced with the publically accessible URL once a site has deployed a service to provide it.  Sites can optionally add any relevant description about the underlying service.**

### Network Monitoring Link Into CRIC [Mandatory]

For the top level monitoring that provides at least IN/OUT for the site, please update CRIC's NetSite Monitoring URL link (see https://wlcg-cric.cern.ch/core/netsite/list/ and pick your site).   If you have separate monitoring for specific networks WLCG CRIC also includes a Monitoring URL containg the above JSON formatted file for each set of NetworkRoutes.  Please update this information in CRIC if relevant.

## Network Diagrams [Optional]

Please provide a link or links to access your most recent network diagrams.  We plan to host these in sub-directories named as SitePages/Diagrams/<SITE>
In this part of the template you can provide descriptions and links to the diagrams.

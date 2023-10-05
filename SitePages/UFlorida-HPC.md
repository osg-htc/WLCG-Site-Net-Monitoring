---

# UFlorida-HPC Network Information
This page describes <br>UFlorida-HPC</br> (aka <br>T2_US_Florida</br> in <br>CMS</br>) network information for WLCG use.

<br>UFlorida-HPC</br> is part of NetSites: https://wlcg-cric.cern.ch/core/netsite/list/

LAST UPDATE : 04-October-2023 15:24 Eastern

## Network Overview
UFlorida-HPC is a CMS Tier2 site embeded in the University of Florida Research Computing (UFRC). UFRC is connected to the Unversity Campus Research Network to 2 100 Gpbs. Topology is modeled on a fat tree 2- or 3-tiers depending on the equipment attached.
1. UFRC is connected to UFnet3 currently at 2x100G
2. UFnet3 backbone is 800G (2x400G)
3. For a single flow, UFnet3 has a connection to the ewan at 100G which is shared with all of campus
4. UF has 100G of capacity to FLRnet-REnet and thus Internet2 or ESnet

2-4 are shared with campus and 1 is just UFRC.

University of Florida will be increasing its connectivity to FLR to 400G in the 2023 December timeframe so 3 and 4 will go to 400G.

### Network Description
All UFlorida-HPC (T2_US_Florida in CMS, a CMS Tier2) equipment at the University of Florida is hosted and maintained by the UFRC in the University Data Center on East Campus which has 100Gbps WAN connection to LHCONE via Florida Lambda Rail(FLR)). Research Computing maintains a pair of Brocade SX1024 WAN switch stack uplinked to the campus Brocade MLXe 16 switch via multiple trunked 40GbE for WAN connections. All Out-facing CMS-related servers such as the CE machiens and the XrootD servers are connected to the SX1024 WAN switches via 10GbE fibers. Internal compute/worker nodes can indirectly access the Internet including LHCONE via a campus NAT server with 40Gbps capacity.

The out-facing CMS-related servers also have 56Gbps FDR Infiniband connections via Mellanox SX6025 switches for local data traffics and 1GbE via a Netgear GS724T switch for slow management/BMC communications. The main storage systems including the dedicated Lustre storage for CMS are connected via 56Gbps FDR Infiniband for data traffic and via various 10GbE/1GbE switches for management/BMC communications. The compute/worker nodes are connected via Mellanox SX6025 56Gbps FDR Infiniband switches for data traffic and via Brocade ICX6650 or Mellanox SX1024 10GbE switches for management/BMC communications.


### Peering Description [Optional]
*Please describe how your site connects by responding to the following questions.*  

**How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?**
UFlorida-HPC uses (TO BE CHECKED) Networks to provide connectivity to the commodity network.

- **If so, who do you peer with and what is the bandwidth available for this type of traffic?** 
UFlorida-HPC shares its commodity connection with the rest of campus ( 100(TO BE CONFIRMED) Gbps)

- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**
UFlorida-HPC does not have a firewall or security device for commodity.

**If you are connected to LHCONE, who/where do you peer with and at what bandwidth?**
UFlorida-HPC connects to LHCONE via FLR in Atlanta(TO BE CONFIRMED) at 100(TO BE CONFIRMED) Gpbs.

- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**
No firewalls or security devices for LHCONE(TO BE CONFIRMED)

**If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?**
UFlorida-HPC has no LHCOPN connection.

- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**
Not applicable

**Do you have a peering for research and education networks for non-LHCONE sites?**  
Currently UFlorida-HPC sees R&E connections via Commodity Network(TO BE CONFIRMED) or FLR
We are working toward peering with Internet2(TO BE CONFIRMED) for these non-LHCONE R&E sites.

- **If so, who/where do you peer with and at what bandwidth?**
We hope to have 100(TO BE CONFIRMED)Gbps

- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**
No firewalls or security devices planned.

### Network Equipment Details [Optional]
UFlorida-HPC uses Brocade SX1024 WAN switch stack uplinked to the campus Brocade MLXe 16 switch. 

## Network Monitoring [Mandatory]
http://cmsio9.rc.ufl.edu/<TO BE UPDATED>
### Network Monitoring Link Into CRIC [Mandatory]

For the top level monitoring that provides at least IN/OUT for the site, please update CRIC's NetSite Monitoring URL link (see https://wlcg-cric.cern.ch/core/netsite/list/ and pick your site).   If you have separate monitoring for specific networks WLCG CRIC also includes a Monitoring URL containg the above JSON formatted file for each set of NetworkRoutes.  Please update this information in CRIC if relevant.

## Network Diagrams [Optional]

Please provide a link or links to access your most recent network diagrams.  We plan to host these in sub-directories named as SitePages/Diagrams/<SITE>
In this part of the template you can provide descriptions and links to the diagrams.


---

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




---

# UFlorida-HPC Network Information
This page describes <br>UFlorida-HPC</br> (aka <br>T2_US_Florida</br> in <br>CMS</br>) network information for WLCG use.

<br>UFlorida-HPC</br> is part of NetSites: https://wlcg-cric.cern.ch/core/netsite/list/

LAST UPDATE : 11-August-2023 10:48 Eastern

## Network Overview
UFlorida-HPC is a CMS Tier2 site embeded in the University of Florida Research Computing (UFRC). UFRC is connected to the Unversity Campus Research Network to 2 100 Gpbs. Topology is modeled on a fat tree 2- or 3-tiers depending on the equipment attached.

### Network Description
All UFlorida-HPC (T2_US_Florida in CMS, a CMS Tier2) equipment at the University of Florida is hosted and maintained by the UFRC in the University Data Center on East Campus which has 100Gbps WAN connection to LHCONE via FLR (Florida Lambda Rail). Research Computing maintains a pair of Brocade SX1024 WAN switch stack uplinked to the campus Brocade MLXe 16 switch via multiple trunked 40GbE for WAN connections. All Out-facing CMS-related servers such as the CE machiens and the XrootD servers are connected to the SX1024 WAN switches via 10GbE fibers. Internal compute/worker nodes can indirectly access the Internet including LHCONE via a campus NAT server with 40Gbps capacity.

The out-facing CMS-related servers also have 56Gbps FDR Infiniband connections via Mellanox SX6025 switches for local data traffics and 1GbE via a Netgear GS724T switch for slow management/BMC communications. The main storage systems including the dedicated Lustre storage for CMS are connected via 56Gbps FDR Infiniband for data traffic and via various 10GbE/1GbE switches for management/BMC communications. The compute/worker nodes are connected via Mellanox SX6025 56Gbps FDR Infiniband switches for data traffic and via Brocade ICX6650 or Mellanox SX1024 10GbE switches for management/BMC communications.


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




---

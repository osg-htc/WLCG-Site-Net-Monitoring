# SITE Network Information
---
This page describes SITE network information for WLCG use.  Our goal is to understand each site's network and track the sites network use (a minimum of IN/OUT total traffic). We would like to get information in three areas: network description, network monitoring and diagrams.  Once this page is created and updated for a specific site, please add it's URL as WLCG CRIC NetSite Info URL (see list of NetSites at https://wlcg-cric.cern.ch/core/netsite/list/)

LAST UPDATE: 15-Feb-2022 08:00 Eastern

## Network Description [Mandatory; can be brief]
In this section of the page, please provide a human readable descrption of your network.  The goal is to give a context for your site's network configuration, including relevant information about the network equipment, peering arrangements, topology, capacity and connectivity.  

### Network Overview [Optional]
**Example:** The network for SITE consists of two Cisco deep buffer router/switches interconnected by 2x100G links, connected to a distribution layer of Dell switches and finally a set of top of rack switches that connect to compute and storage systems.  The uplink connects via 2x40G to our regional connector.  Typical compute nodes connect at 2x1G but some have 10G or 2x10G, while storage nodes are typically connected at 2x25G but some are connected at 100G.

### Peering Description [Optional]
How does your site peer with other networks?  Who provides your default internet connectivity and at what bandwidth?  If you are connected to LHCONE, who do you peer with and at what bandwidth?   If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?   Do you have a peering for research and education networks for non-LHCONE sites?  If so, who do you peer with and at what bandwidth? Finally does you site provide commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?  If so, who do you peer with and what is the bandwidth available for this type of traffic?

TO-DO: Convert this to a table to fill out ?

### Network Equipment Details [Optional]
Our site has a mix of Cisco and Dell switches and routers.   We have X Cisco YYY-ZZZZ systems running NXOS X.XXX and Y Dell YYYYY switches running OS10 X.XX

## Network Monitoring [Mandatory]
---
In this section please provide a brief overview of your site network traffic monitoring along with an appropriate URL that is publically accessible and provides, at a minimum, your sites IN/OUT bandwidth.   **Example:** Our site is monitoring via Grafana which tracks the IN/OUT bandwidth of all our site's Research and Education traffic on two 40 interfaces (which include LHCONE traffic).   Additionally Grafana monitors another 10G interface which provides our sites commodity (non-R&E) traffic.    

It is important to have this network monitoring at your site's **border:** the location where all traffic has to pass to get into or out of your site.

**Note** we prefer monitoring links that are machine consumable.  Our goal is to be able to automate collection of all of our major sites IN/OUT traffic in near real time.    

TO-DO:  Provide explicit examples of monitoring URLs that are machine consumable.  Investigate Prometheus, CheckMK, RRDtool, Cacti, Grafana, others?

### Network Monitoring Link [Mandatory]
For R&E traffic, including LHCONE, please see https://grafana.site.org/d/D2dvElGGz/site-net-monitoring. 

For commodity traffic, please see https://grafana.site.org/d/K2dvEl9Gz/site-net-mon-comm. 

Note, the links identified need to be added to CRIC.   For the top level monitoring that provides at least IN/OUT for the site, please update CRIC's NetSite Monitoring URL link (see https://wlcg-cric.cern.ch/core/netsite/list/ and pick your site).   If you have separate monitoring for specific networks WLCG CRIC also includes a Monitoring URL for each set of NetworkRoutes.  Please update this information in CRIC if relevant.

## Network Diagrams [Optional]
---
Please provide a link to access your most recent network diagrams.  

TO-DO:  Should we create a location in gitlab that can host site network diagrams?   The idea would be that many sites may require that site network diagrams be behind some kind of authorization and will not publish diagrams for the general public to access.

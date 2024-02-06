
---

# SITE Network Information
This page describes `FZK-LCG2` (GridKa/T1_DE_KIT) network information for WLCG use.  # Replace with your site name or suitable introduction line

The sections below are listed as either Mandatory or Optional.  There is descriptive text that should be replaced once the template is copied
to become a site specific document.  Feel free to remove unused or non-relevant sections or text.

*Our goal is to understand each site's network and track the sites network use (a minimum of IN/OUT total traffic). We would like to get information in three areas: network description, network monitoring and diagrams.  Once this page is created and updated for a specific site, please add it's URL as WLCG CRIC NetSite Info URL (see list of NetSites at https://wlcg-cric.cern.ch/core/netsite/list/)*

LAST UPDATE (when file is changed, please update): 06-Feb-2024

_NOTE: this is the template file.  Please see filled out examples in the SitePages directory of this project._

## Network Overview [Mandatory; can be brief]
*In this section of the page, please provide a human readable descrption of your network.  The goal is to give a context for your site's network configuration, including relevant information about the network equipment, peering arrangements, topology, capacity and connectivity.*

As of beginning of 2024, `FZK-LCG2` has two 100Gbit/s connections to `CERN/LHCOPN` and two additional 100Gbit/s connections which are shared between `internet` and `LHCONE` (operated by the German NREN DFN). One additional 100Gbit/s link is available for `LHCOPN` during the time of DC2024

### Network Description [Optional]


### Peering Description [Optional]
*Please describe how your site connects by responding to the following questions.*  

**How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?**
- Connected via DFN (German NREN)
- Traffic is handles by a firewall with very limited throughput (<20Gbit/s). For known WLCG sites, bypasses are in place.

**If you are connected to LHCONE, who/where do you peer with and at what bandwidth?**
- Connected via DFN (see above)
- no firewalls for LHCONE traffic

**If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?**   
- 100Gb/s from KIT to CERN via DFN/GEANT
- 100Gb/s from KIT to CERN via BelWue/SWITCH
- DC2024 only! 100Gb/s from KIT to CERN via DFN/GEANT

**Do you have a peering for research and education networks for non-LHCONE sites?**  
- no

### Network Equipment Details [Optional]

All equipment of the Tier-1 is Cisco HW. Central backbone switches and order routers are Nexus N7100 & N7700. Storage server are connected to Nexus 93600cd-gx switches.

## Network Monitoring [Mandatory]

https://wlcg-netmon-kit.scc.kit.edu/kit-gridka-netmon.json

### Network Monitoring Link Into CRIC [Mandatory]

For the top level monitoring that provides at least IN/OUT for the site, please update CRIC's NetSite Monitoring URL link (see https://wlcg-cric.cern.ch/core/netsite/list/ and pick your site).   If you have separate monitoring for specific networks WLCG CRIC also includes a Monitoring URL containg the above JSON formatted file for each set of NetworkRoutes.  Please update this information in CRIC if relevant.

## Network Diagrams [Optional]




---

# SITE Network Information
This page describes **DESY-ZN** network information for WLCG use.

LAST UPDATE (when file is changed, please update): 10-Nov-2023 08:00 CET


## Network Overview [Mandatory]

DESY-ZN is a Tier 2 site based in Zeuthen (near Berlin), Germany. 

The site is connected to DESY-HH via DFN (the German NREN) through 4 dedicated 10 Gb/s links over 2 independent paths. 2 links build up a port channel, each. So, in fact, the site is connected via 2 independent 20Gbit/s port channels.

Internal networking behind the router is a 160 Gb/s spine, with individual servers mostly connected at 2x10 Gb/s.


### Network Description [Optional]


### Peering Description [Optional]

**How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?**
- **If so, who do you peer with and what is the bandwidth available for this type of traffic?**

Connected via DFN.

- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

Juniper firewall filtering *all* traffic. No science dmz implemented, atm.

**If you are connected to LHCONE, who/where do you peer with and at what bandwidth?**

Connected to LHCONE via DESY-HH over standard 2x20Gb/s port channels

### Network Equipment Details [Optional]

## Network Monitoring [Mandatory]

The DESY-ZN monitoring data is available at https://perfson1.zeuthen.desy.de/wlcg-monitoring/desy-zn.json

### Network Monitoring Link Into CRIC [Mandatory]

https://wlcg-cric.cern.ch/core/netsite/detail/DESY-ZN/

## Network Diagrams [Optional]


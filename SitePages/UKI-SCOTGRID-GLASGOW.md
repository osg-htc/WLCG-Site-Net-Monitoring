# UKI-SCOTGRID-GLASGOW Network Information
This page describes `UKI-SCOTGRID-GLASGOW` network information for WLCG use.

LAST UPDATE (when file is changed, please update): 23-August-2023 12:45 BST

## Network Overview [Mandatory; can be brief]

UKI-SCOTGRID-GLASGOW is a Tier 2 site based at the University of Glasgow.

The site is connected to Janet (the National Research and Education Network in the UK) at 2 × 10 Gbps.  It is possible this will be upgraded in the near future.

The site's internal backbone runs at 80 Gbps.  Hosts are connected at 10 Gbps, 2 × 10 Gbps (where both external and site-local connections are provided), or 40 Gbps in the case of certain storage systems.  A management network on a segregated VLAN provides host connections at 1 Gbps.

### Network Description [Optional]

### Peering Description [Optional]

**How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?**
- **If so, who do you peer with and what is the bandwidth available for this type of traffic?**
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

Connected via Janet, currently without in-line firewall / security devices

**If you are connected to LHCONE, who/where do you peer with and at what bandwidth?**

Not connected to LHCONE

**If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?**   

Not connected to LHCOPN

**Do you have a peering for research and education networks for non-LHCONE sites?**  

No peering for non-LHCONE sites

### Network Equipment Details [Optional]

## Network Monitoring [Mandatory]

[https://monitor.gla.scotgrid.ac.uk:8443/metrics.json](https://monitor.gla.scotgrid.ac.uk:8443/metrics.json)

### Network Monitoring Link Into CRIC [Mandatory]

For the top level monitoring that provides at least IN/OUT for the site, please update CRIC's NetSite Monitoring URL link (see https://wlcg-cric.cern.ch/core/netsite/list/ and pick your site).   If you have separate monitoring for specific networks WLCG CRIC also includes a Monitoring URL containg the above JSON formatted file for each set of NetworkRoutes.  Please update this information in CRIC if relevant.

## Network Diagrams [Optional]

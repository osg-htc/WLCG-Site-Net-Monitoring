# SITE Network Information
---
This page describes `**SITE**` network information for WLCG use.  # Replace with your site name or suitable introduction line

_Our goal is to understand each site's network and track the sites network use (a minimum of IN/OUT total traffic). We would like to get information in three areas: network description, network monitoring and diagrams.  Once this page is created and updated for a specific site, please add it's URL as WLCG CRIC NetSite Info URL (see list of NetSites at https://wlcg-cric.cern.ch/core/netsite/list/)_

LAST UPDATE (when file is changed, please update): 30-Mar-2022 08:45 Eastern

_NOTE: this is the template file.  Please see filled out examples in the SitePages directory of this project._

## Network Overview [Mandatory; can be brief]
In this section of the page, please provide a human readable descrption of your network.  The goal is to give a context for your site's network configuration, including relevant information about the network equipment, peering arrangements, topology, capacity and connectivity.  

### Network Description [Optional]
**Example:** We have two physical sites that make up AGLT2, one at Michigan State University and one at University of Michigan. Michigan State University has a shared 100Gbps path-diverse connection to two MSU owned routers in Chicago (600 West and 710 N Lakeshore drive).   The University of Michigan site also has a shared 100Gbps path to OmniPoP (710 N Lakeshore drive) in Chicago.  Both sites are interconnected via a 100Gbps research triangle between Michigan State, University of Michigan and Wayne State University.

### Peering Description [Optional]
Please describe how your site connects by responding to the following questions:

How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?  
    If so, who do you peer with and what is the bandwidth available for this type of traffic?
    Are there any firewall or security devices in-line with this traffic?  If so, please describe.

If you are connected to LHCONE, who/where do you peer with and at what bandwidth?
    Are there any firewall or security devices in-line with this traffic?  If so, please describe.

If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?   
    Are there any firewall or security devices in-line with this traffic?  If so, please describe.

Do you have a peering for research and education networks for non-LHCONE sites?  
    If so, who/where do you peer with and at what bandwidth? 
    Are there any firewall or security devices in-line with this traffic?  If so, please describe.

### Network Equipment Details [Optional]
Our site has a mix of Cisco and Dell switches and routers.   We have X Cisco YYY-ZZZZ systems running NXOS X.XXX and Y Dell YYYYY switches running OS10 X.XX

## Network Monitoring [Mandatory]
---
In this section please provide a brief overview of your site network traffic monitoring along with an appropriate URL that is publically accessible and provides a JSON formatted file in the following format:

{
  "Description": "Network statistics for AGLT2",
  "UpdatedLast": "2022-03-30T12:48:49.027Z",
  "InBytesPerSec": 0,
  "OutBytesPerSec": 0,
  "UpdateInterval": "1 minute"
}

This file should be updated at some reasonable interval (once per minute or similar).

It is important to have this network monitoring information from your site's **border:** the location where all traffic has to pass to get into or out of your site.

TO-DO:  Provide code example showing the use of SNMP to query two or more interface, add the results and produce the needed JSON file.

### Network Monitoring Link [Mandatory]
For R&E traffic, including LHCONE, please see https://grafana.site.org/d/D2dvElGGz/site-net-monitoring. 

For commodity traffic, please see https://grafana.site.org/d/K2dvEl9Gz/site-net-mon-comm. 

Note, the links identified need to be added to CRIC.   For the top level monitoring that provides at least IN/OUT for the site, please update CRIC's NetSite Monitoring URL link (see https://wlcg-cric.cern.ch/core/netsite/list/ and pick your site).   If you have separate monitoring for specific networks WLCG CRIC also includes a Monitoring URL containg the above JSON formatted file for each set of NetworkRoutes.  Please update this information in CRIC if relevant.

## Network Diagrams [Optional]
---
Please provide a link to access your most recent network diagrams.  We plan to host these in sub-directories named as SitePages/Diagrams/<SITE>
In this part of the template you can provide descriptions and links to the diagrams.

TO-DO:  Should we create a location in gitlab that can host site network diagrams?   The idea would be that many sites may require that site network diagrams be behind some kind of authorization and will not publish diagrams for the general public to access.

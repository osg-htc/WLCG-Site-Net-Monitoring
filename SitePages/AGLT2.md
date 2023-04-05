# SITE Network Information
---
This page describes AGLT2 network information for WLCG use.  

LAST UPDATE: 05-April-2022 15:54 Eastern

## Network Overview
We have two physical sites that make up AGLT2, one at Michigan State University(MSU) and one at University of Michigan(UM). Michigan State University has a shared 100Gbps path-diverse connection to two MSU owned routers in Chicago (600 West and 710 N Lakeshore drive).   The University of Michigan site also has a shared 100Gbps path to OmniPoP (710 N Lakeshore drive) in Chicago.  Both sites are interconnected via a 100Gbps research triangle between Michigan State, University of Michigan and Wayne State University.    

### Network Description
The **Michigan State site** is hosted at the MSU data center and served by the MSU Research Network on its DMZ.   The routing and switching equipment is all from Juniper.  There are two core campus and two core data center routers.   The two spine routers (QFX5120-32C) for AGLT2 connect to the two core datacenter routers with 100Gbps links (each spine has a 100Gbps link to each core router). NOTE that MSU is using EVPN/VXLAN to each data switch. Every MSU rack has two Juniper 48 port 10G/25G data switches (QFX5120-48Y) plus one management switch (EX2300-48T), which is connected to both data switches.  Every node has a bonded link providing both public and private VLAN access (1G, 10G or 25G links by node type).  The management switch is used for iDRAC and provisioning purposes.  A typical storage node is connected at 2x10G, 4x10G or 2x25G and worker nodes are connected at 2x1G, 2x10G or the newest at 2x25G.   The MSU site also redundantly connects to the state-wide 100Gbps research network and on to LHCONE and the rest of the research and education networks via a path diverse 100Gbps infrastructure to Chicaco/ESnet/Internet2.  Commodity network access is via Merit Networks using MSU's campus connection to Merit, with backup via AGLT2 at UM's Merit connection.

The **University of Michigan site** is hosted in the College of Language, Sciences and Arts (LS&A) data center with its network connections outside of the University firewall.  The routing and switching equipment is from Dell and Cisco with two routers inside the data center (Cisco N3K-C36180YC-R). There is typically one data switch in each rack (Dell S5232F-ON) and one management switch (S3048-ON).  Typical compute nodes connect at 2x1G but some have 10G or 2x10G, while storage nodes are typically connected at 2x25G but some are connected at 100G.  The UM site reaches the wide area network and the MSU site via a state-wide 100Gbps resilient research network.   The UM access to this network is via two deep buffer Cisco switches, located at two different campuses.   Each Cisco router at AGLT2 in LSA connects to both of the deep buffer Ciscos.  From these deep buffer Ciscos, UM reaches LHCONE and the rest of the research and education networks via a 2x40Gbps to our 100Gbps path to Chicago/ESnet/Internet2. Commodity network access is via Merit Networks using a 10 Gbps peering with Merit at UM, with backup via AGLT2 at MSU's Merit connection.

Both UM and MSU connect to each other via the state wide 100Gbps research network.   

### Peering Description
Please describe how your site connects by responding to the following questions:

**How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?**  
    AGLT2 uses Merit Networks to provide connectivity to the commodity network for both MSU and UM.  

- **If so, who do you peer with and what is the bandwidth available for this type of traffic?**
    The MSU site shares its commodity connection with the rest of campus (shared 100Gbps)
    The UM site has a dedicated 10Gbps directly to Merit in Ann Arbor

- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**
    Neither site has a firewall or security device for commodity.

**If you are connected to LHCONE, who/where do you peer with and at what bandwidth?**
    AGLT2 connects to LHCONE via ESnet in Chicago at 100Gpbs (two locations? 600W and 710 NLSD).  
    Each site has its connection and peering to ESnet.

- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**
    No firewalls or security devices for LHCONE

**If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?**   
    No AGLT2 is a Tier-2 and has no LHCOPN connection.

- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**
    Not applicable.

**Do you have a peering for research and education networks for non-LHCONE sites?**
    Currently AGLT2 sees R&E connections via Merit or ESnet.  
    We are working toward peering with Internet2 for these non-LHCONE R&E sites.

- **If so, who/where do you peer with and at what bandwidth?** 
    Not yet, but hope to have 100Gpbs via Internet2
    
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**
    No firewalls or security devices planned.

### Network Equipment Details
AGLT2 runs a mix of Juniper, Cisco and Dell switches and routers.  

At Michigan State the current JunOS 20.2R3.9 on the spines and 20.2R2-S3 on the leafs.

At the University of Michigan we have Cisco NXOS 9.3(5) and Dell OS10 10.5.1

## Network Monitoring
---
The UM AGLT2 site has implemented the site traffic monitor.  Results are available at https://head01.aglt2.org/aglt2-netmon.json 

The MSU AGLT2 site has implemented the site traffic monitor.  Results are available at https://www.aglt2.org/aglt2-msu-netmon.json 

### Network Monitoring Link into CRIC

CRIC has been updated for UM AGLT2.  See https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20University%20of%20Michigan/ and https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20Michigan%20State%20University/

## Network Diagrams
---
We are working on gathering suitable diagrams and plan to share them in Gitlab under SitePages/Diagrams.

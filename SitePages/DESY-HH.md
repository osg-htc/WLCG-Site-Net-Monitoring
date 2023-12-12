SITE Network Information
This page describes DESY-HH network information for WLCG use.
LAST UPDATE (when file is changed, please update): 07-Dec-2023 09:00 CET

Network Overview
DESY-HH, located in Hamburg (Germany), operates as LHC Tier 2, Belle II Tier 1 and XFEL Tier 0 site.

DESY (ASN1754) is connected to other NRENs and global upstreams via DFN (the German NREN) through two peerings with a bandwidth of 50GBit/s each, which are terminated on two independent Juniper WAN routers. In addition on of DESY's WAN routers has a symmetric 300MBit/s BGP peering with a commercial ISP to provide access for on-premise startups and commercial traffic in general. 

Both WAN routers operate two 10G links to provide internet access for the DESY Zeuthen site (described in DESY-ZN.md). All traffic to DESY-HH internal networks passes an active-standby SRX5k cluster connected with 100G links to both WAN routers and the four data center core routers. There is no traffic bypassing this perimeter firewall.

The backbone of the internal network is based on portchannels with multiple 40G or 100G links. A monitoring is in place to ensure that this backbone is never limiting the data center performance, but extended if necessary.

So far the vast majority of servers is connected with single 10G links or lacp bonds of 10G interfaces. An increasing number of 100G server links arises, in particular for storage proxies in the context of data acquisition. 

Peering Description
DESY has 
 - Two redundant BGP peerings with the German NREN (DFN Verein e.V.) with 50GBit/s each.
   - this peering consists of three logical links which share the available bandwidth
     1) General purpose internet access
     2) LHCOne 
     3) Hifis Backbone (an optical VPN between the research centers of the Helmholtz Association)
 - Non-redundant BGP peering with a commercial ISP with 300MBit/s
 - A dedicated 100G link to NCBJ (Swierk, Poland) for XFEL data


Are there any firewall or security devices in-line with this traffic? 

There is a active-standby Juniper SRX cluster capable of filtering all traffic without losses or limitation. 

If you are connected to LHCONE, who/where do you peer with and at what bandwidth?
The LHCOne link shares the 2x50G peering with the German NREN (DFN) together with the scientific non-LHCOne traffic. There is no guaranteed bandwith for this connection, however the by far largest fraction of this bandwidth is consumed by LHCOne traffic.


Network Monitoring [Mandatory]

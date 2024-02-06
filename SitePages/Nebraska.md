
---

# Nebraska Network Information
This page describes Nebraska network information for WLCG use.

LAST UPDATE: 10-Nov-2023 09:36 Central

## Network Overview

The T2_US_Nebraska USCMS Tier-2 is hosted within the Holland Computing Center (HCC) at the University of Nebraska-Lincoln in Lincoln, NE. HCC connects to the UNL campus border router via a 2x100Gb LAG and the path from that border router to our regional provider (Great Plains Network, GPN) and Internet2 and ESnet in Kanasas city is also a 2x100Gb LAG. This is a shared link for the campus community. Connectivity between GPN and ESnet in Kansas City is a single 400Gb link. HCC and ESnet peer directly via this path for LHCONE connectivity.

HCC operates in three datacenters with the Lincoln, NE datacenter containing the majority of the USCMS Tier-2 resources. A second datacenter in Omaha, NE is connected via a 100Gbps link and houses general purpose HPC clusters and the third datacenter elsewhere in Lincoln, NE operated by central campus ITS is connected by 2x40Gbps. A small quantity of older donated 1GbE connected hardware acting as CMS worker nodes exists in this third datacenter but is connected directly to the core switches in the main Tier-2 datacenter in Lincoln effectively acting as the same location.

### Network Description

The HCC datacenter hosting T2_US_Nebraska utilizes a Dell S4248FB-ON switch as the edge device connecting to the campus WAN router at 2x100Gbps. This edge switch peers with ESnet in Kanasas City for LHCONE via stretched VLAN through our campus and regional providers.

A pair of Dell Z9264F-ON switches configured with VLT act as the core for this datacenter and connect via 1x100Gbps each for a 2x100Gb total to the above edge switch. All access switches for storage and worker nodes used by the Tier-2 connect to this switch pair via multple 10, 40, or 100Gbps links as appropriate. A variety of Dell, HP, and fs.com 1GbE switches are used for management / BMC / out-of-band network purposes.

### Peering Description [Optional]
*Please describe how your site connects by responding to the following questions.*  

**How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?**

UNL connects to Windstream and ALLO for commodity internet providers.

- **If so, who do you peer with and what is the bandwidth available for this type of traffic?**
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

HCC and the Tier-2 connect to the campus NU-WAN border router at 2x100Gbps and that router connects to the commodity internet providers directly. There is no firewall in line with traffic to commodity networks but there is between our site and other NU system campuses, ie: we are treated as our own campus.

**If you are connected to LHCONE, who/where do you peer with and at what bandwidth?**
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

HCC and the Tier-2 peer with ESnet for LHCONE via a VLAN stretched between our Lincoln datacenter and Kansas City. The full path is 2x100Gb and no firewall or devices are in-line with this.

**If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?**   
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

We do not connect to LHCOPN.

**Do you have a peering for research and education networks for non-LHCONE sites?**  
- **If so, who/where do you peer with and at what bandwidth?** 
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

HCC and the Tier-2 connect to the campus NU-WAN border router at 2x100Gbps and in turn the campus border connects via our regional provider (GPN) to Internet2 at 2x100Gbps in Kansas City in addition to ESnet at 400Gbps in Kansas City.

### Network Equipment Details

Our site runs a mix of mostly Dell 10Gb and 25Gb access switches such as the Dell S4048-ON and Dell 5248F-ON models. Numerous N3200 and S3048 switches are used for 1GbE management networks in addition to some old HP and fs.com switches. The core switches are Dell S9264F-ON and the edge switch is a Dell S4248FB-ON.

## Network Monitoring

The T2_US_Nebraska site has implemented the site traffic monitor. Results are available at https://t2.unl.edu/t2_us_nebraska-netmon.json

### Network Monitoring Link Into CRIC

NetSite Monitoring URL link is updated and available at https://wlcg-cric.cern.ch/core/netsite/detail/US-Nebraska/

## Network Diagrams

We will try to provide network diagrams in the future.
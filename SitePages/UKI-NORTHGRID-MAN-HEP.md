
---

# UKI-NORTHGRID-MAN-HEP Network Information
This page describes `UKI-NORTHGRID-MAN-HEP` network information for WLCG use. 

LAST UPDATE: 08-Aug-2023 13:42 CEST

## Network Overview [Mandatory; can be brief]
We have a single site at the University of Manchester. The Tier2 has 100Gb/s core switches connected to the university router with dedicate bandwidth limited at 40Gb/s. The university is connected to Janet UK backbone at 100Gb/s.

### Network Description [Optional]

The storage nodes currently have 2x10Gb/s bonded links to the ToR switches which are a mixture of sizes going from 20Gb/s bonded to 40 Gb/s. The WNs have all 10Gb/s 10GBASE-T. The core switches are 100Gb/s capable but the uplink to the university router is limited at 40Gb/s. The university network has been outsourced to a commercial company which makes any interaction very difficult.

### Peering Description [Optional]
*Please describe how your site connects by responding to the following questions.*  

**How does your site connect for commodity (non research and education) connectivity (www.google.com, www.github.com, etc)?**
- **If so, who do you peer with and what is the bandwidth available for this type of traffic?**
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

via [JANET](https://beta.jisc.ac.uk/janet)

**If you are connected to LHCONE, who/where do you peer with and at what bandwidth?**
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

N/A

**If you are connected to LHCOPN, how to you connect and peer with CERN and at what bandwidth?**   
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

N/A

**Do you have a peering for research and education networks for non-LHCONE sites?**  

Via [JANET](https://beta.jisc.ac.uk/janet)

- **If so, who/where do you peer with and at what bandwidth?** 
- **Are there any firewall or security devices in-line with this traffic?  If so, please describe.**

No.

### Network Equipment Details [Optional]
The site has a mix of Dell, HPE, and Aristas.

## Network Monitoring [Mandatory]

https://wlcg-nmon.tier2.hep.manchester.ac.uk/snmp-json

### Network Monitoring Link Into CRIC [Mandatory]

https://wlcg-cric.cern.ch/core/netsite/detail/UKI-NORTHGRID-MAN-HEP/

https://wlcg-nmon.tier2.hep.manchester.ac.uk/UKI-NORTHGRID-MAN-HEP-NET.DESC.html

## Network Diagrams [Optional]

Please provide a link or links to access your most recent network diagrams.  We plan to host these in sub-directories named as SitePages/Diagrams/<SITE>
In this part of the template you can provide descriptions and links to the diagrams.

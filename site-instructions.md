# Instructions for Implementing WLCG Site Network Information

> ⚠️ **Note:** This document is intended as an outreach/email summary for WLCG site administrators. For the full authoritative deployment instructions, see [README.md](./README.md).

The WLCG Monitoring Task Force is conducting a campaign to describe and instrument the network for the largest WLCG sites in advance of the upcoming WLCG Data Challenge 2027 (DC27) and as part of establishing this as standard ongoing monitoring.  This document provides a high level view of what site administrators will need to do in response to the campaign.

## Minimum Required for DC27

For a site's traffic to appear on the [WLCG Site Network dashboard](https://monit-grafana-open.cern.ch/d/MwuxgogIk/wlcg-site-network?orgId=16&from=now-7d&to=now&var-site=All), the **required** steps are:

1. Publish a JSON document with your total border IN and OUT traffic (bytes/sec), refreshed every 60 seconds, in the format in [JSON-FORMAT.md](https://github.com/osg-htc/WLCG-Site-Net-Monitoring/blob/master/JSON-FORMAT.md).
2. Make that URL reachable from the CERN MONIT subnets `137.138.0.0/16`, `188.184.0.0/15` and `2001:1458:0D00::/44`.
3. Register the URL as the **Monitoring URL** of your NetSite in [WLCG CRIC](https://wlcg-cric.cern.ch/core/netsite/list/).

**Strongly encouraged:** a site network description registered as the NetSite **Info URL** (a brief overview is enough to start).
**Optional:** separate Monitoring URLs per CRIC NetworkRoute (e.g. LHCOPN / LHCONE / R&E), and the detailed sections of the site description.

## Summary of Required Tasks

There are two primary components for the network site information that need to be provided:
  - A description of the site's network architecture, components and configuration
  - Monitoring of the total site IN and OUT traffic, updated every 60 seconds and accessible from CERN MONIT

We are using GitHub to host the relevant instructions, example code and markdown templates at https://github.com/osg-htc/WLCG-Site-Net-Monitoring  Site administrators will need access to this project to allow them to create and maintain their site network information.

## Site Network Description

Sites will need to create a site network description following our markdown template at https://github.com/osg-htc/WLCG-Site-Net-Monitoring/blob/master/Template/SitePageTemplate.md   Instructions:
  - Copy SitePageTemplate.md to a file named `RCSITE`.md (replace `RCSITE` with your site's WLCG-CRIC "RC Site" name). Existing site pages can be seen as examples in the SitePages directory on CERN GitLab (https://gitlab.cern.ch/wlcg-doma/site-network-information/-/tree/master/SitePages), where you may also submit yours.
  - Edit the `RCSITE`.md file filling in all the MANDATORY sections as well as any OPTIONAL sections you are willing to provide.  Note this file will allow others to help your site identify problems and suggest improvements, so adding more information will allow more effective help.
  - Publish the completed file at some URL (web server), accessible to others in WLCG.  There are various tools that convert markdown to html (see information on `pandoc` in [README.md](./README.md))
  - Last step is to register the URL for your site network description into WLCG-CRIC NetSite Info URL.   For sites with more than one NetSIte, we suggest describing all net site details in one markdown file and using the same Info URL in each NetSite.
      - Example:  AGLT2 RC Site has two NetSites:  US-AGLT2 Michigan State University and US-AGLT2 University of Michigan but both use the same Info URL: https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20Michigan%20State%20University/ and https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20University%20of%20Michigan/

The intent is that sites own the information they provide and will regularly update information as their network evolves.   In the above instructions, we are assuming sites will provide a web server that will host their completed site network description.  In the future we may be able to provide a WLCG web server that can utilize the `RCSITE`.md files in the SitePages directory.

## Site Traffic Monitoring

The second component for site network information is enabling tracking of the total network traffic IN and OUT of EACH NetSite.   A site's connectivity can vary significantly across WLCG, from single WAN (wide area network) connections to many connections of various bandwidth.   For DC27 and ongoing standard monitoring we want sites to add up all input and output traffic and publish suitably formatted JSON with those numbers every 60 seconds.
  - First step is to identify all the network interfaces that connect the site to the wide area network
  - Second step is to programmatically gather the IN and OUT statistics for all interfaces, summing up IN and OUT to get totals.   Example code (a Go binary/RPM, recommended, plus Python3 and Docker versions) using SNMP is available in https://github.com/osg-htc/WLCG-Site-Net-Monitoring/tree/master/site-traffic-monitoring
  - Third step is to publish the statistics in the JSON format described in [JSON-FORMAT.md](https://github.com/osg-htc/WLCG-Site-Net-Monitoring/blob/master/JSON-FORMAT.md) every 60 seconds at a URL accessible from the CERN MONIT subnets listed above.    
  - Last step is to register the JSON statistics URL in WLCG CRIC for the NetSite in the Monitoring URL

  - Optional: if your LHCOPN, LHCONE and general R&E traffic use different links, you can also publish one JSON per route and register each as the Monitoring URL of the matching CRIC NetworkRoute. See "Per-route monitoring" in the [README](https://github.com/osg-htc/WLCG-Site-Net-Monitoring/blob/master/README.md#per-route-monitoring-lhcopn--lhcone--re).

Sites should verify that the [WLCG Site Network dashboard](https://monit-grafana-open.cern.ch/d/MwuxgogIk/wlcg-site-network?orgId=16&from=now-7d&to=now&var-site=All) starts showing their data within about 24 hours. If it does not, see the troubleshooting checklist in the [README](https://github.com/osg-htc/WLCG-Site-Net-Monitoring/blob/master/README.md#troubleshooting-my-site-is-not-on-the-dashboard).

## Additional information

There is additional information on the project GitHub pages (https://github.com/osg-htc/WLCG-Site-Net-Monitoring) as well as previous presentations on this topic:
  - WLCG Operations presentation (2025/2026 — link to be added)
  - CHEP 2023 Poster https://docs.google.com/presentation/d/1yzD3Gm6Ph8lAGf3c0WupL3m4L96zT61mASx5dVpBt-Q/edit?usp=sharing
  - Earlier WLCG Operations presentation (2024) https://docs.google.com/presentation/d/1sB4xPJPLoLbNnfV0mfwfXmHXzX8lZ8VtC4TSqDXWUMw/edit?usp=sharing

Questions can be directed to the WLCG Monitoring Task Force (CERN egroup WLCGMon-TF).

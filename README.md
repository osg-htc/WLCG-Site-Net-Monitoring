# Site Networking Information

The WLCG is trying to organize and collect site networking details as part of the ongoing series of Data Challenges, the first of which was October 2021. This project is intended to provide sites with a *site networking template* that they can fill to provide needed information for both automated data collection and for experts to review and understand so they can advise sites about possible architectural, configuration or hardware changes that might be beneficial.

The aim is to collect centrally Input/Output traffic and match it with a site network description to help understanding the site external traffic. 

> 📧 **For site administrators:** A concise outreach summary is available in [site-instructions.md](./site-instructions.md) — suitable for email campaigns or quick reference.

# Overview of Steps for Sites

To provide this information, sites should start by mentally drawing a circle around their site.  Any network connections that cross the circle need to be monitored.    The goal is to sum up the IN and OUT traffic crossing the circle.  Sites typically do this by querying the relevant network ports using SNMP.  (NOTE: some sites may already have monitoring that can be "harvested" to provide the needed info and it is up to those sites to use that data to create the needed monitoring information.)

Once sites have the data, it needs to be made available for the CERN MONIT system to consume.   The sites need to provide a URL that publishes the data in JSON format (WLCG-site-snmp folder has details on the required format).  This URL can be protected but the following CERN address space needs access:
```
137.138.0.0/16
188.184.0.0/15
2001:1458:0D00::/44
```

We have example code in subdirectories in this project:
  - [WLCG-site-snmp](./WLCG-site-snmp/README.md):  provides example python code, including a containerized version, for site traffic monitoring.  Sites can use this to gather the appropriate data and produce correctly formatted output.
  - [Elastiflow-Example](./Elastiflow-Example/README.md): provides python code to read/parse CRIC network data for Elastiflow (note: not needed for most sites)

After the data is available at a URL, the site needs to update their WLCG CRIC entry (see "Site Network Description" below).  That will cause the CERN MONIT system to start gathering the data from the site.

Questions can be directed to the WLCG Monitoring Task Force (wlcgmon-tf 'at' cern.ch)

The following sections provide further details.

## Site Network description

There are two directories, `Template/` and `SitePages/`, that host, respectively, the site network template and the completed site network information pages.

> **Note:** The `Template/` directory (containing `SitePageTemplate.md`) lives in **this GitHub repository**. The completed site pages (`SitePages/`) are hosted on the **CERN GitLab** mirror at <https://gitlab.cern.ch/wlcg-doma/site-network-information/-/tree/master/SitePages>. See [`Template/README.md`](./Template/README.md) for full instructions on how to fill in and submit a site page.

The plan is to use the completed site network template as the CRIC NetSite **Info URL**  (e.g., for AGLT2 as an example see https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20Michigan%20State%20University/ ).   Sites should clone the example SitePageTemplate.md file, creating a new entry in SitePages named <SITE>.md.  Then they can then edit the md file to provide the information requested.  To allow the file to be accessible in CRIC, the file will need to be converted to HTML, which can be done via `pandoc`.   

```
yum install pandoc
pandoc <SITE>.md -f markdown -t html5 -S -s -o /path/<SITE>.html
```
The site should then copy the HTML file to a web server and register the URL in WLCG CRIC.

The goal is that sites will "own" their site network information pages once they create them. Sites should plan to regularly provide updates, especially when significant changes are made.

Sites can also add a diagram of their network if they want to in the SitePages/Diagrams.

## Traffic Monitoring information 

In addition we have two directories with example code: WLCG-site-snmp contains example code for site traffic monitoring. Elastiflow-Example provides python code that can read and parse CRIC network data and produces network annotations usable inside Elastiflow.

Sites also need to determine the best **Monitoring URL**.  

Each NetSite has a Monitoring URL that should be used to point to network monitoring that shows IN/OUT traffic for **that** NetSite.  We provide multiple working examples sites can deploy in the WLCG-site-snmp directory:

| Deployment Method | Best When... | Directory |
|---|---|---|
| **Go binary / RPM** ⭐ _Recommended_ | No Python env; want a compiled tool, systemd service, and RPM packaging | [`WLCG-site-snmp/go/`](./WLCG-site-snmp/go/README.md) |
| **Python (SNMP only)** | Python already available; simple SNMP-to-file output | [`WLCG-site-snmp/`](./WLCG-site-snmp/README.md) |
| **Python + HTTPS server** | Python available; no separate web server | [`WLCG-site-snmp/snmp-with-http-example/`](./WLCG-site-snmp/snmp-with-http-example/) |
| **Docker** | Container-native environment | [`WLCG-site-snmp/snmp-with-http-example/`](./WLCG-site-snmp/snmp-with-http-example/) |

The recommended deployment method is the **Go binary / RPM** ([`WLCG-site-snmp/go/README.md`](./WLCG-site-snmp/go/README.md)), which ships as an RPM, requires no Python environment, and is managed via systemd.  Full configuration details are documented in the [manpage source](./WLCG-site-snmp/go/wlcg-site-snmp.1.md).

In addition, each NetworkRoute (see https://wlcg-cric.cern.ch/core/networkroute/list/), composed of one or more network subnets *also* provides an opportunity to have a Monitoring URL.  This should be used to provide monitoring specifically for the identified network routes, if it exists.

Once monitoring is in place and published in CRIC, sites can check their results via https://monit-grafana-open.cern.ch/d/MwuxgogIk/wlcg-site-network?orgId=16&from=now-7d&to=now&var-site=All 

Note the CERN monitoring systems need to be able to read the data.   For those with firewalls, please open the following subnets for access to the Monitoring URL:
- 137.138.0.0/16
- 188.184.0.0/15
- 2001:1458:0D00::/44

## CRIC

If a site doesn't have a NetSite in CRIC (see https://wlcg-cric.cern.ch/core/netsite/list/) they should create one. The starting point is the main RC site in CRIC https://wlcg-cric.cern.ch/core/rcsite/list/


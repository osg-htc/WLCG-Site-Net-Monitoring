# Site Networking Information

The WLCG is trying to organize and collect site networking details as part of the ongoing series of Data Challenges, the first of which was October 2021. This project is intended to provide sites with a *site networking template* that they can fill to provide needed information for both automated data collection and for experts to review and understand so they can advise sites about possible architectural, configuration or hardware changes that might be beneficial.

The aim is to collect centrally Input/Output traffic and match it with a site network description to help understanding the site external traffic. 

## Site Network description

There are two directories, Templates and SitePages that host, respectively, the site network template and associated guide, and the completed site network information pages, to be served by a central web server.  

The plan is to use the completed site network template as the CRIC NetSite **Info URL**  (e.g., for AGLT2 as an example see https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20Michigan%20State%20University/ ).   Sites should clone the example SitePageTemplate.md file, creating a new entry in SitePages named <SITE>.md.  Then they can then edit the md file to provide the information requested.  To allow the file to be accessible in CRIC, the file will need to be converted to HTML, which can be done via `pandoc`.   

```
yum install pandoc
pandoc <SITE>.md -f markdown -t html5 -S -s -o /path/<SITE>.html
```
The site should then copy the HTML file to a web server and register the URL in WLCG CRIC.

The goal is that sites will "own" their site network information pages once they create them. Sites should plan to regularly provide updates, especially when significant changes are made.

Sites can also add a diagram of their network if they want to in the SitePages/Diagrams.

## Traffic Monitoring information 

In addition we have two additional directories that host example code that might be useful: WLCG-site-snmp which contains example python code for site traffic monitoring and Elastiflow-Example, which provides python code that can read and parse CRIC network data and produces network annotations usable inside Elastiflow.

Sites also need to determine the best **Monitoring URL**.  

Each NetSite has a Monitoring URL that should be used to point to network monitoring that shows IN/OUT traffic for **that** NetSite.  We provide a working example sites can deploy in the WLCG-site-snmp directory.  

In addition, each NetworkRoute (see https://wlcg-cric.cern.ch/core/networkroute/list/), composed of one or more network subnets *also* provides an opportunity to have a Monitoring URL.  This should be used to provide monitoring specifically for the identified network routes, if it exists.

Once monitoring is in place and published in CRIC, sites can check their results via https://monit-grafana-open.cern.ch/d/MwuxgogIk/wlcg-site-network?orgId=16&from=now-7d&to=now&var-site=All 

## CRIC

If a site doesn't have a NetSite in CRIC (see https://wlcg-cric.cern.ch/core/netsite/list/) they should create one. The starting point is the main RC site in CRIC https://wlcg-cric.cern.ch/core/rcsite/list/


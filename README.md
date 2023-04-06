# Site Networking Information

The WLCG is trying to organize and collect site networking details as part of the ongoing series of Data Challenges, the first of which was October 2021. This project is intended to provide sites with a *site networking template* that they can fill it to provide needed information for both automated data collection and for experts to review and understand so they can advise sites about possible architectural, configuration or hardware changes might be beneficial.

Structure: we have created two directories, Templates and SitePages, that host, respectively, the site network template and associated guide and the completed site network information pages, to be served by a central web server.  In addition we have two additional directories that host example code that might be useful: WLCG-site-snmp which contains example python code for site traffic monitoring and Elastiflow-Example, which provides python code that can read and parse CRIC network data and produces network annotations usable inside Elastiflow.

Regarding the site network information pages, the goal is that sites will "own" their pages once they create them. Sites should plan to regularly provide updates, especially when significant changes are made.

The plan is to use the completed site network template as the CRIC NetSite **Info URL**  (e.g., for AGLT2 as an example see https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20Michigan%20State%20University/ ).   Sites should clone the example SitePageTemplate.md file, creating a new entry in SitePages named <SITE>.md.  Then they can then edit the md file to provide the information requested.  To allow the file to be accessible in CRIC, the file will need to be converted to HTML, which can be done via `pandoc`.   

```
yum install pandoc
pandoc <SITE>.md -f markdown -t html5 -S -s -o /path/<SITE>.html
```
The site should then copy the HTML file to a web server and register the URL in WLCG CRIC.

Sites also need to determine the best **Monitoring URL**.   Each NetSite (see https://wlcg-cric.cern.ch/core/netsite/list/) has a Monitoring URL that should be used to point to network monitoring that shows IN/OUT traffic for **that** NetSite.  We provide a working example sites can deploy in the WLCG-site-snmp directory.  

In addition, each NetworkRoute (see https://wlcg-cric.cern.ch/core/networkroute/list/), composed of one or more network subnets *also* provides an opportunity to have a Monitoring URL.  This should be used to provide monitoring specifically for the identified network routes, if it exists.

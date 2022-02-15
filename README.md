# Site Networking Information

The WLCG is trying to organize and collect site networking details as part of the ongoing series of Data Challenges, the first of which was October 2021.
This project is intended to provide sites with a *site networking template* that they can fill it to provide needed information for both automated data collection and for experts to review and understand so they can advise sites about possible architectural, configuration or hardware changes might be beneficial.

Structure: we have created two directories, Templates and SitePages, that host, respectively, the site network template and associated guide and the completed site network information pages, to be served by a central web server.

The goal is that sites will "own" their pages once they create them and should regularly provide updates when significant changes are made.

We need to discuss if gitlab is the appropriate location to host "site" network pages that we link to from CRIC.   A proposal is to use the complete site network template as the CRIC NetSite **Info URL**  (e.g., for AGLT2 as an example see https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20Michigan%20State%20University/ ).   Sites also need to determine the best **Monitoring URL**.   Each NetSite (see https://wlcg-cric.cern.ch/core/netsite/list/) has a Monitoring URL that should be used to point to network monitoring that shows IN/OUT traffic for **that** NetSite.   

In addition, each NetworkRoute (see https://wlcg-cric.cern.ch/core/networkroute/list/), composed of one or more network subnets *also* provides an opportunity to have a Monitoring URL.  This should be used to provide monitoring specifically for the identified network routes, if it exists.

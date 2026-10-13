# WLCG Site Network Monitoring & Information

This project helps WLCG sites report their network traffic and publish a machine-readable network description so that the CERN MONIT infrastructure can monitor inter-site data flows.

> 📧 **For site administrators:** A concise outreach summary is available in [site-instructions.md](./site-instructions.md) — suitable for email campaigns or quick reference.

## Minimum Required for DC27

To appear on the [WLCG Site Network dashboard](https://monit-grafana-open.cern.ch/d/MwuxgogIk/wlcg-site-network?orgId=16&from=now-7d&to=now&var-site=All) (which will be used during DC27), a site **must** do the following for each of its CRIC NetSites:

| | Requirement | Status |
|---|---|---|
| 1 | Publish a JSON document with the site's **total border IN and OUT traffic** in bytes/sec, in the format defined in [JSON-FORMAT.md](./JSON-FORMAT.md) | **Required** |
| 2 | Refresh it every **60 seconds** (or as close as your equipment allows) | **Required** |
| 3 | Make the URL reachable (HTTP or HTTPS) from the CERN MONIT subnets listed below | **Required** |
| 4 | Register that URL as the **Monitoring URL** of your NetSite in [WLCG CRIC](https://wlcg-cric.cern.ch/core/netsite/list/) | **Required** |
| 5 | Publish a site network description and register it as the NetSite **Info URL** (a short "Network Overview" paragraph is enough to start) | Strongly encouraged |
| 6 | Separate Monitoring URLs per CRIC NetworkRoute (e.g. LHCOPN / LHCONE / general R&E) — see [Per-route monitoring](#per-route-monitoring-lhcopn--lhcone--re) | Optional, useful for larger sites |
| 7 | Detailed peering, equipment and diagram sections of the site page | Optional |

Items 1–4 are all that is needed for your traffic to show up on the dashboard. You do **not** have to use our code: any existing monitoring (SNMP, NMS exports, flow collectors, …) that can produce the JSON is fine.

CERN MONIT polls from these subnets — allow them through any firewall or ACL in front of your Monitoring URL:

```
137.138.0.0/16
188.184.0.0/15
2001:1458:0D00::/44
```

## Quick Start for Site Administrators

1. **Identify your border:** every interface where traffic enters or leaves your site (see [Overview](#overview-of-steps-for-sites) below). Work with your network team to get read-only SNMP access to them.
2. **Deploy the traffic monitoring tool** using the method that best fits your environment (see [Deployment Methods](#deployment-methods) below), or adapt your existing monitoring.
3. **Publish the JSON report** at a URL reachable from the CERN subnets above, and check it is valid with `curl -s <URL> | python3 -m json.tool`.
4. **Register the Monitoring URL** for your NetSite in [WLCG CRIC](https://wlcg-cric.cern.ch/core/netsite/list/).
5. *(Strongly encouraged)* **Fill in `Template/SitePageTemplate.md`**, convert it to HTML with `pandoc`, host it, and register it as the NetSite **Info URL**.
6. **Verify** your site appears in the [Grafana dashboard](https://monit-grafana-open.cern.ch/d/MwuxgogIk/wlcg-site-network?orgId=16&from=now-7d&to=now&var-site=All) within about 24 hours. If not, see [Troubleshooting](#troubleshooting-my-site-is-not-on-the-dashboard).

> Questions? Contact the WLCG Monitoring Task Force: **wlcgmon-tf 'at' cern.ch**

## Deployment Methods

| Deployment Method | Best When... | Directory |
|---|---|---|
| **Go binary / RPM** ⭐ *Recommended* | No Python env; want compiled tool, systemd service, RPM packaging | [`site-traffic-monitoring/go/`](./site-traffic-monitoring/go/README.md) |
| **Python (SNMP only)** | Python already available; simple SNMP-to-file output | [`site-traffic-monitoring/python/`](./site-traffic-monitoring/python/README.md) |
| **Python + HTTPS server** | Python available; no separate web server | [`site-traffic-monitoring/python/snmp-with-http-example/`](./site-traffic-monitoring/python/snmp-with-http-example/) |
| **Docker** | Container-native environment | [`site-traffic-monitoring/python/snmp-with-http-example/`](./site-traffic-monitoring/python/snmp-with-http-example/) |

Full configuration details for the Go tool are in its [manpage source](./site-traffic-monitoring/go/wlcg-site-snmp.1.md).

# Overview of Steps for Sites

To provide this information, sites should start by mentally drawing a circle around their site.  Any network connections that cross the circle need to be monitored.    The goal is to sum up the IN and OUT traffic crossing the circle.  Sites typically do this by querying the relevant network ports using SNMP.  (NOTE: some sites may already have monitoring that can be "harvested" to provide the needed info and it is up to those sites to use that data to create the needed monitoring information.)

Once sites have the data, it needs to be made available for the CERN MONIT system to consume.   The sites need to provide a URL that publishes the data in JSON format (see [JSON-FORMAT.md](./JSON-FORMAT.md) for the required format).  This URL can be protected but the CERN address space needs access — see the [required CERN subnets](#minimum-required-for-dc27) above.

We have example code in subdirectories in this project:
  - [site-traffic-monitoring](./site-traffic-monitoring/README.md):  provides example python and Go code, including a containerized version, for site traffic monitoring.  Sites can use this to gather the appropriate data and produce correctly formatted output.
  - [Elastiflow-Example](./Elastiflow-Example/README.md): provides python code to read/parse CRIC network data for Elastiflow (note: not needed for most sites)

After the data is available at a URL, the site needs to update their WLCG CRIC entry (see "Site Network Description" below).  That will cause the CERN MONIT system to start gathering the data from the site.

The following sections provide further details.

## Site Network description

There are two directories, `Template/` and `SitePages/`, that host, respectively, the site network template and the completed site network information pages.

> **Note:** The `Template/` directory (containing `SitePageTemplate.md`) lives in **this GitHub repository**. The completed site pages (`SitePages/`) are hosted on the **CERN GitLab** mirror at <https://gitlab.cern.ch/wlcg-doma/site-network-information/-/tree/master/SitePages> *(intentionally hosted on CERN GitLab — not a broken link)*. See [`Template/README.md`](./Template/README.md) for full instructions on how to fill in and submit a site page.

The plan is to use the completed site network template as the CRIC NetSite **Info URL**  (e.g., for AGLT2 as an example see https://wlcg-cric.cern.ch/core/netsite/detail/US-AGLT2%20Michigan%20State%20University/ ).   Sites should clone the example SitePageTemplate.md file, creating a new entry in SitePages named <SITE>.md.  Then they can then edit the md file to provide the information requested.  To allow the file to be accessible in CRIC, the file will need to be converted to HTML, which can be done via `pandoc`.   

```
yum install pandoc   # or: apt install pandoc
pandoc <SITE>.md -f markdown -t html5 -s -o /path/<SITE>.html
```
The site should then copy the HTML file to a web server and register the URL in WLCG CRIC.

The goal is that sites will "own" their site network information pages once they create them. Sites should plan to regularly provide updates, especially when significant changes are made.

Sites can also add a diagram of their network if they want to in the SitePages/Diagrams.

## Traffic Monitoring information 

In addition we have two directories with example code: site-traffic-monitoring contains example code for site traffic monitoring. Elastiflow-Example provides python code that can read and parse CRIC network data and produces network annotations usable inside Elastiflow.

Sites also need to determine the best **Monitoring URL**.  

Each NetSite has a Monitoring URL that should be used to point to network monitoring that shows IN/OUT traffic for **that** NetSite.  See [Deployment Methods](#deployment-methods) above for the tools we provide.

## Per-route monitoring (LHCOPN / LHCONE / R&E)

Many sites carry LHCOPN, LHCONE and general R&E / commodity traffic over different links, and that traffic may be treated differently during DC27. The NetSite Monitoring URL should always report the **total** across all border links. In addition, each CRIC [NetworkRoute](https://wlcg-cric.cern.ch/core/networkroute/list/) has its own Monitoring URL field that can report just the traffic for that route.

To provide per-route data:

1. In CRIC, check that your NetSite has a NetworkRoute for each path you want to report (e.g. one for LHCOPN, one for LHCONE). Create any that are missing.
2. Run one additional instance of the monitoring tool per route, each configured with **only** the interfaces carrying that route's traffic, and each publishing to its own URL (e.g. `.../site-netmon-lhcopn.json`, `.../site-netmon-lhcone.json`). Use the same JSON format as for the NetSite.
   - **Go tool:** create one config file per route and run one systemd unit per config, for example by copying `wlcg-site-snmp.service` to `wlcg-site-snmp-lhcone.service` and pointing its `--conf` flag at the route's config file.
   - **Python tool:** keep a separate install location (with its own `INDICES`, `COMM` and `JSONOUTFILE`) and a separate systemd unit per route.
3. Register each URL as the Monitoring URL of the matching NetworkRoute in CRIC.

If one physical link carries more than one route (e.g. LHCONE and general R&E on the same port, separated by VLAN), monitor the per-VLAN sub-interface if your equipment exposes counters for it; otherwise report only the NetSite total.

## Troubleshooting: my site is not on the dashboard

Check, in order:

1. **Is the URL registered?** Open your NetSite in [CRIC](https://wlcg-cric.cern.ch/core/netsite/list/) and confirm the *Monitoring URL* field is set (on the NetSite itself, not only on a NetworkRoute).
2. **Is the JSON valid?** `curl -s <URL> | python3 -m json.tool` must succeed, and the document must contain the required fields in [JSON-FORMAT.md](./JSON-FORMAT.md). A common mistake is unquoted keys copied from old examples.
3. **Is it fresh?** `UpdatedLast` should be within the last couple of minutes. If it is stale, the collector service has stopped or the copy to the web server is failing.
4. **Can CERN reach it?** The URL must be reachable from the CERN subnets in [Minimum Required for DC27](#minimum-required-for-dc27). Check firewalls, ACLs and any authentication in front of the web server. If you use HTTPS, use a certificate that a standard (IGTF or public) CA bundle will accept.
5. **Have you waited long enough?** New sites should appear within about 24 hours of being registered in CRIC.

If all of these check out, contact **wlcgmon-tf 'at' cern.ch** with your NetSite name and Monitoring URL.

## CRIC

If a site doesn't have a NetSite in CRIC (see https://wlcg-cric.cern.ch/core/netsite/list/) they should create one. The starting point is the main RC site in CRIC https://wlcg-cric.cern.ch/core/rcsite/list/


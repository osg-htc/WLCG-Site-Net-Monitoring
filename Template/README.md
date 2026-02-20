# Template Directory

This directory contains the site network description template used by WLCG sites.

## Contents

- **`SitePageTemplate.md`** — The canonical Markdown template that site administrators copy and fill in to describe their site's network configuration, peering, and monitoring.

## How to Use `SitePageTemplate.md`

1. **Copy** `SitePageTemplate.md` to a working location and rename it `<RCSITE>.md`, where `<RCSITE>` is your site's WLCG-CRIC "RC Site" name (e.g., `AGLT2.md`).

2. **Edit** the file and fill in all **Mandatory** sections. Optional sections can be removed or left blank if not applicable. The more detail you provide, the easier it is for WLCG experts to review and advise your site.

3. **Convert** the completed Markdown file to HTML using `pandoc`:

   ```bash
   yum install pandoc   # or: apt install pandoc
   pandoc <RCSITE>.md -f markdown -t html5 -s -o <RCSITE>.html
   ```

4. **Publish** the resulting HTML file on a web server that is publicly accessible within WLCG.

5. **Register** the URL in WLCG CRIC as the **Info URL** for your NetSite entry:
   - Go to <https://wlcg-cric.cern.ch/core/netsite/list/>
   - Find your site's NetSite entry and set the *Info URL* field to the published HTML URL.
   - If your RC Site has multiple NetSites, you can use the same Info URL for all of them.

## Where Completed Site Pages Are Stored

Completed site pages (filled-in `<RCSITE>.md` files and diagrams) are hosted in the **`SitePages/`** directory on the CERN GitLab mirror of this project:

> <https://gitlab.cern.ch/wlcg-doma/site-network-information/-/tree/master/SitePages>

Please submit your completed `<RCSITE>.md` file there (via a GitLab merge request or by contacting the WLCG Monitoring Task Force).

## Questions

Direct questions to the WLCG Monitoring Task Force: **wlcgmon-tf 'at' cern.ch**

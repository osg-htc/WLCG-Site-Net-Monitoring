# Instructions for Implementing WLCG Site Network Information

The WLCG Monitoring Task Force is conducting a campaign to describe and instrument the network for the largest WLCG sites in advance of the upcoming WLCG Network Data Challenge 2024(DC24).  This document provides a high level view of what site administrators will need to do in response to the campaign.

## Summary of Required Tasks

There are two primary components for the network site information that need to be provided:
  - A description of the site's network architecture, components and configuration
  - Monitoring of the total site IN and OUT traffic, updated every 60 seconds and accessible from CERN MONIT

We are using CERN GitLab to host the relevant instructions, example code and markdown templates at https://gitlab.cern.ch/wlcg-doma/site-network-information/  Site administrators will need access to this project to allow them to create and maintain their site network information.

## Site Network Description

Site's will need to create a site network description following our markdown template at https://gitlab.cern.ch/wlcg-doma/site-network-information/-/blob/master/Template/SitePageTemplate.md   Instructions:
  - Copy the SitePageTemplate.md to the SitePages directory (https://gitlab.cern.ch/wlcg-doma/site-network-information/-/tree/master/SitePages), naming the file as <RCSITE>.md (replace <RCSITE> with your sites WLCG-CRIC "RC Site" name)
  - Edit the <RCSITE>.md file filling in all the MANDATORY sections as well as any OPTIONAL sections

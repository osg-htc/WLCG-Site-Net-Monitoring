# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [Unreleased]

### Changed
- README overhaul: added Quick Start checklist, deployment-method comparison table, deduplicated firewall subnet block, updated stale GitLab links
- Clarified Docker deployment documentation in `WLCG-site-snmp/README.md`

### Added
- `Template/README.md` with step-by-step instructions for site page submission
- `site-instructions.md` as a concise outreach document for site administrators
- `CONTRIBUTING.md` with contribution guidelines
- `CHANGELOG.md` (this file)

### Breaking Changes
- ⚠️ **Directory renamed:** `WLCG-site-snmp/` → `site-traffic-monitoring/`
  - Python files moved to `site-traffic-monitoring/python/`
  - Go implementation moved to `site-traffic-monitoring/go/`
  - Any external scripts or documentation referencing the old `WLCG-site-snmp/` path must be updated

### Migration
- Repository migrated from `gitlab.cern.ch/wlcg-doma/site-network-information` to `github.com/osg-htc/WLCG-Site-Net-Monitoring`

## [1.1.0] - 2023-10-22

### Added
- Go reimplementation of the SNMP monitoring tool under `WLCG-site-snmp/go/`
- Go version distributed as an RPM package (`wlcg-site-snmp.spec`)
- Supports SNMPv2c and SNMPv3, built-in HTTP/HTTPS server, SCP output

## [1.0.0] - 2021-10-01

### Added
- Python3 SNMP-based site traffic monitoring script
- Dockerized version contributed by Justin Balcas / Caltech
- Site network description template (`SitePageTemplate.md`)
- Elastiflow-Example: CRIC network data parser for Elastiflow

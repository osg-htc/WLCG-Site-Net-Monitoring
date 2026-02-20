# Changelog

## [Unreleased]

### Migration from CERN GitLab to GitHub
- Repository migrated from `gitlab.cern.ch/wlcg-doma/site-network-information` to `github.com/osg-htc/WLCG-Site-Net-Monitoring`

### Breaking Changes
- ⚠️ **Directory renamed:** `WLCG-site-snmp/` → `site-traffic-monitoring/`
  - Python files moved to `site-traffic-monitoring/python/`
  - Go implementation moved to `site-traffic-monitoring/go/`
  - Any external scripts or documentation referencing the old `WLCG-site-snmp/` path must be updated

### Added
- Go implementation of the SNMP-based site traffic monitor (in `site-traffic-monitoring/go/`)
  - Ships as an RPM package
  - Includes systemd unit file
  - Full manpage (`wlcg-site-snmp.1.md`)
- `site-traffic-monitoring/README.md` with deployment comparison table
- `CHANGELOG.md` (this file)

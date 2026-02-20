# Site Traffic Monitoring

This directory contains example code for monitoring site network traffic for WLCG sites.
Choose the deployment method that best fits your environment:

## Deployment Methods

| Method | Best When... | Directory |
|---|---|---|
| **Go binary / RPM** ⭐ *Recommended* | No Python env; want compiled tool, systemd service, RPM packaging | [`go/`](./go/README.md) |
| **Python (SNMP only)** | Python already available; simple SNMP-to-file output | [`python/`](./python/README.md) |
| **Python + HTTPS server** | Python available; no separate web server | [`python/snmp-with-http-example/`](./python/snmp-with-http-example/) |
| **Docker** | Container-native environment | [`python/snmp-with-http-example/`](./python/snmp-with-http-example/) |

## Questions

Direct questions to the WLCG Monitoring Task Force: **wlcgmon-tf 'at' cern.ch**

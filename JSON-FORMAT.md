# Site Traffic JSON Format

This is the document that a site's CRIC **Monitoring URL** (NetSite or NetworkRoute) must return. It is produced automatically by the [Go](./site-traffic-monitoring/go/README.md) and [Python](./site-traffic-monitoring/python/README.md) tools in this repository; sites using their own monitoring must produce the same structure.

## Example

```json
{
    "Description": "Network statistics for AGLT2",
    "UpdatedLast": "2026-10-13T14:47:23.057942+00:00",
    "InBytesPerSec": 1275982208.1281207,
    "OutBytesPerSec": 1526111205.30333,
    "UpdateInterval": "60 seconds",
    "MonitoredInterfaces": [
        "aglt2-rtr-1.local_Ethernet1/51",
        "aglt2-rtr-1.local_Ethernet1/52",
        "aglt2-rtr-2.local_Ethernet1/51",
        "aglt2-rtr-2.local_Ethernet1/52"
    ]
}
```

## Fields

| Field | Type | Required | Description |
|---|---|---|---|
| `InBytesPerSec` | number | **Yes** | Average rate of traffic entering the site (summed over all monitored border interfaces) during the last interval, in **bytes per second** (not bits). |
| `OutBytesPerSec` | number | **Yes** | Average rate of traffic leaving the site during the last interval, in **bytes per second**. |
| `UpdatedLast` | string | **Yes** | Time the rates were computed, ISO 8601 with an explicit UTC offset, e.g. `2026-10-13T14:47:23.057942+00:00`. |
| `UpdateInterval` | string | Recommended | Interval the rates are averaged over, as `"<N> seconds"`. Should be `"60 seconds"`. |
| `Description` | string | Recommended | Free text identifying the site (and route, if this is a per-route URL). |
| `MonitoredInterfaces` | array of strings | Recommended | The border interfaces included in the sums, conventionally `<device>_<ifDescr>`. Helps others verify what is being counted. |

## Rules

- The document must be **valid JSON**: keys and strings in double quotes. Check with `curl -s <URL> | python3 -m json.tool`.
- "IN" means traffic arriving at your site from the WAN; "OUT" means traffic your site sends to the WAN.
- Use 64-bit counters (`ifHCInOctets` / `ifHCOutOctets`) for any link of 1 Gbps or more; 32-bit counters wrap too quickly at high rates.
- Serve the document with HTTP status 200. `Content-Type: application/json` is preferred.
- Do not publish an empty or `null` document while the collector starts up if you can avoid it; MONIT will record a gap.

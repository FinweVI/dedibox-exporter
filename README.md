# Dedibox Exporter

A Prometheus Exporter for the [Dedibox API](https://console.online.net/fr/api/) (Dedicated Servers from [Online.net](https://www.scaleway.com/en/dedibox/))

## Metrics

Note: for metrics like `abuse` or `ddos`, only the first page (latest infos) is fetched from the API to avoid overloading it.  
The goal of this exporter is to create alerts on new DDoS attacks or unresolved abuses, not provide advanced statistic on your account.

| Name | Description | Labels |
| -------- | -------- | -------- |
| dedibox_pending_abuse | Pending abuses | service, category |
| dedibox_pending_abuse_count | Pending abuse count | None |
| dedibox_ddos | DDoS attacks on your services | target, mitigation_system, attack_type |
| dedibox_ddos_count | DDoS attacks count | None |
| dedibox_dedibackup_quota_space_total_bytes | Dedibackup total space quota | server_id, active |
| dedibox_dedibackup_quota_space_used_bytes | Dedibackup used space quota | server_id, active |
| dedibox_dedibackup_quota_files_total | Dedibackup total files quota | server_id, active |
| dedibox_dedibackup_quota_files_used | Dedibackup used files quota | server_id, active |
| dedibox_plan | Dedibox plans availability | name, datacenter |

## Arguments

| Name | Default | Description |
| -------- | -------- | -------- |
| collector | abuse | List of Collectors to enable (abuse, ddos, plan, dedibackup)<br />This can be used multiple times to select multiple collectors |
| listen-address | 127.0.0.1:9539 | Address to listen on for web interface and telemetry |
| log-level | 1 | Log level: 0=debug, 1=info, 2=warn, 3=error |
| metric-path | /metrics | Path under which to expose metrics |

### Env Var

Require `ONLINE_API_TOKEN` to be set with your Online.net's API token.

## Issues & Contribution
All bug report, packaging requests, features requests or PR are accepted.  
I mainly created this exporter for my personal usage but I'll be happy to hear about your needs.

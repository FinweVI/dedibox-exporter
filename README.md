# Dedibox Exporter
A Prometheus Exporter for the Dedibox API (Dedicated Servers from Online.net)

https://console.online.net/fr/api/

## Metrics

| Name | Description | Labels |
| -------- | -------- | -------- |
| dedibox_abuse | Online.net's abuse and it's resolution status | id, sender, service, type |
| dedibox_plan | Get Dedibox plan availability | name, datacenter |
| dedibox_dedibackup_quota_space_total_bytes | Get Dedibackup total space quota | server_id, active |
| dedibox_dedibackup_quota_space_used_bytes | Get Dedibackup space quota used | server_id, active |
| dedibox_dedibackup_quota_files_total | Get Dedibackup total quota files | server_id, active |
| dedibox_dedibackup_quota_files_used | Get Dedibackup used quota files | server_id, active |

## Arguments

| Name | Default | Description |
| -------- | -------- | -------- |
| listen-address | :9539 | Address to listen on for web interface and telemetry. |
| metric-path | /metrics | Path under which to expose metrics. |

### Env Var

Require `ONLINE_API_TOKEN` to be set with your Online.net's token.


## Issues & Contribution
All bug report, packaging requests, features requests or PR are accepted.
I mainly created this exporter for my personal usage but I'll be happy to hear about your needs.

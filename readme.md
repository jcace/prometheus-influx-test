# Grafana Cloud Prometheus post test
https://grafana.com/docs/grafana-cloud/data-configuration/metrics/metrics-influxdb/push-from-telegraf/#pushing-from-applications-directly

## Usage
First set up a `.env` file with the required credentials ex

```bash
set -a
PROM_USER_ID="1234"
PROM_API_KEY="xxx"
PROM_BASE_URL="https://influx-prod-10-prod-us-central-0.grafana.net/api/v1/push/influx" # See guide above for how to get the correct URL
set +a
```


Then run the script with the following arguments

```bash
go run .
```

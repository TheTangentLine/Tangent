[← Docs index](../README.md)

# Observability

**For the platform itself**

| Metric | Type | Labels |
|---|---|---|
| `tangent_http_requests_total` | counter | `route`, `method`, `status` |
| `tangent_http_request_duration_seconds` | histogram | `route`, `method` |
| `tangent_provisioning_requests_total` | counter | `kind`, `status` |
| `tangent_provisioning_request_duration_seconds` | histogram | `kind` |
| `tangent_provisioning_step_duration_seconds` | histogram | `step`, `status` |
| `tangent_provisioning_step_retries_total` | counter | `step` |
| `tangent_provisioning_queue_depth` | gauge | none |
| `tangent_operator_reconcile_total` | counter | `result` |

Alerts:

- Provisioning success rate under 95% over 1 hour.
- Queue depth above 50 for 10 minutes.
- p95 time-to-first-deploy above 15 minutes.
- Operator reconcile errors above zero for 15 minutes.

Traces: one trace per provisioning request, with a span per step and child spans for GitHub and Git operations. Store the trace ID on the request so the portal can link to it in Grafana.

**For tenant services**

- The `register_dashboard` step commits a ConfigMap labelled `grafana_dashboard: "1"` into the team namespace; Grafana's sidecar loads it automatically.
- A `ServiceMonitor` is included in each template's Helm chart.
- Logs are collected by Promtail or Grafana Alloy and are queryable by `team`, `service`, `tier` labels.

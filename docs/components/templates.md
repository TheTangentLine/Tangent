[← Docs index](../README.md)

# Golden-path templates

Each template is a directory with a manifest and files rendered with Go `text/template`.

```text
templates/go-http-api/
├── template.yaml
└── skeleton/
    ├── cmd/server/main.go.tmpl
    ├── internal/...
    ├── Dockerfile
    ├── Makefile
    ├── .github/workflows/ci.yaml.tmpl   # calls the reusable workflow
    ├── deploy/chart/                    # Helm chart
    ├── catalog-info.yaml.tmpl           # service metadata
    └── README.md.tmpl
```

`template.yaml`:

```yaml
name: go-http-api
version: 1.2.0
description: Go HTTP API with health checks, slog, OpenTelemetry and Prometheus metrics
parameters:              # JSON Schema, also used to render the portal form
  type: object
  required: [port]
  properties:
    port:     { type: integer, default: 8080, minimum: 1024, maximum: 65535 }
    database: { type: string, enum: [none, postgres], default: none }
files:
  exclude_if:
    - path: "internal/db/**"
      when: '{{ eq .Params.database "none" }}'
```

Every template must provide:

- `/healthz` and `/readyz` endpoints.
- Structured JSON logs to stdout.
- OpenTelemetry traces exported to the in-cluster collector via `OTEL_EXPORTER_OTLP_ENDPOINT`.
- Prometheus metrics on `/metrics`.
- Non-root container, read-only root filesystem, resource requests and limits.
- A CI workflow that calls the shared reusable workflow.

Templates are versioned with semver. Services record the version they were created from so you can later build "template upgrade" PRs.

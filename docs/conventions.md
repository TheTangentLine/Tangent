[← Docs index](README.md)

# Conventions

| Thing | Convention | Example |
|---|---|---|
| Team slug | `^[a-z][a-z0-9-]{1,30}$` | `payments` |
| Service name | `^[a-z][a-z0-9-]{1,40}$`, globally unique | `payments-api` |
| Namespace | `{team}-{tier}` | `payments-dev` |
| Repo | `{org}/{service}` | `acme/payments-api` |
| Image | `ghcr.io/{org}/{service}:{git-sha}` | `ghcr.io/acme/payments-api:3f9c2e1` |
| Hostname | `{service}.{tier}.{base-domain}` | `payments-api.dev.example.com` |
| Keycloak group | `tangent:team:{team}:{role}` | `tangent:team:payments:developer` |

Required labels on every tenant resource:

```yaml
tangent.dev/team: payments
tangent.dev/service: payments-api
tangent.dev/tier: dev
app.kubernetes.io/managed-by: tangent
```

Commit trailers written by the platform: `Tangent-Service-ID`, `Tangent-Request-ID`. These make Git history traceable back to requests and let steps check idempotency.

[← Docs index](README.md)

# Configuration

Environment variables for `tangent api` and `tangent worker`:

| Variable | Required | Description |
|---|---|---|
| `TANGENT_DATABASE_URL` | yes | Postgres connection string |
| `TANGENT_HTTP_ADDR` | no | Default `:8080` |
| `TANGENT_OIDC_ISSUER` | yes | Keycloak realm issuer URL |
| `TANGENT_OIDC_AUDIENCE` | yes | Expected `aud` claim |
| `TANGENT_GITHUB_APP_ID` | worker | GitHub App ID |
| `TANGENT_GITHUB_APP_PRIVATE_KEY_PATH` | worker | Path to mounted private key |
| `TANGENT_GITHUB_ORG` | worker | Org where repos are created |
| `TANGENT_GITOPS_REPO` | worker | `owner/tangent-gitops` |
| `TANGENT_ARGOCD_URL` | worker | ArgoCD API URL |
| `TANGENT_ARGOCD_TOKEN_PATH` | worker | Read-only ArgoCD token |
| `TANGENT_OPENCOST_URL` | api | OpenCost API URL |
| `TANGENT_BASE_DOMAIN` | yes | For generated hostnames |
| `TANGENT_WORKER_CONCURRENCY` | no | Default `4` |
| `TANGENT_LOG_LEVEL` | no | `debug`, `info` (default), `warn`, `error` |
| `OTEL_EXPORTER_OTLP_ENDPOINT` | no | OTel collector endpoint |

Secrets are mounted from Kubernetes Secrets populated by External Secrets, never passed as plain env vars in manifests.

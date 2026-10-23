[← Docs index](README.md)

# Local development

## Prerequisites

Docker, Go, Node.js, `kubectl`, `kind` (or `k3d`), `helm`, Tilt, `sqlc`, Terraform, and a GitHub org (a free personal org works) with a GitHub App installed.

## Make targets

| Target | Does |
|---|---|
| `make cluster-up` | Create the kind cluster with ingress port mappings |
| `make bootstrap` | Install ArgoCD and the app-of-apps (Keycloak, monitoring, Kyverno, OpenCost) |
| `make dev` | Run `tilt up`: API, worker, operator and portal with live reload |
| `make migrate` | Apply database migrations |
| `make gen` | Run `sqlc generate`, `controller-gen`, OpenAPI client generation |
| `make test` | Unit and integration tests |
| `make e2e` | Full end-to-end test against the kind cluster |
| `make lint` | `golangci-lint`, `eslint`, `kubeconform` on manifests |
| `make cluster-down` | Delete the cluster |

## First run

```bash
git clone https://github.com/TheTangentLine/tangent && cd tangent
cp .env.example .env          # fill in GitHub App ID, private key path, org
make cluster-up
make bootstrap                # takes a few minutes the first time
make dev
# Portal:  http://portal.localtest.me
# API:     http://api.localtest.me
# ArgoCD:  http://argocd.localtest.me
# Grafana: http://grafana.localtest.me
# Keycloak seeded with users: alice (platform-admin), bob (team-lead, payments), carol (developer, payments)
```

`localtest.me` resolves to `127.0.0.1`, so no `/etc/hosts` edits are needed.

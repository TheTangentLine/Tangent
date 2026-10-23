[← Docs index](README.md)

# Tech stack

| Layer | Choice | Why |
|---|---|---|
| Backend language | Go | Same language for API, worker, operator and CLI |
| HTTP router | `chi` | Small, idiomatic, `net/http` compatible |
| Database | PostgreSQL | Transactions, `SKIP LOCKED` job queue, JSONB |
| DB access | `sqlc` + `pgx` | Type-safe queries from plain SQL |
| Migrations | `goose` or `golang-migrate` | Versioned SQL migrations |
| Identity | Keycloak (OIDC) | Self-hostable, supports groups and SAML later |
| Operator | kubebuilder / controller-runtime | Standard operator framework |
| GitOps | ArgoCD + ApplicationSets | Declarative, UI for debugging |
| CI | GitHub Actions (reusable workflows) | Native to repos created by the platform |
| Registry | GitHub Container Registry (GHCR) | Auth through `GITHUB_TOKEN`, no extra secrets |
| IaC | Terraform | Cluster, DNS, cloud resources |
| Policy | Kyverno | Policies written as Kubernetes YAML |
| Secrets | External Secrets Operator + Vault (or OpenBao) | No secrets in Git |
| Metrics / logs / traces | Prometheus, Loki, Tempo, Grafana, OpenTelemetry | Open source, standard |
| Cost | OpenCost | Kubernetes cost allocation by label |
| Frontend | Next.js + TypeScript + Tailwind | Fast to build, good DX |
| Local cluster | kind or k3d, Tilt for inner loop | Reproducible local setup |
| Demo hosting | k3s on a single ARM VM | Cheap or free, real public demo |

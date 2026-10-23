[← Docs index](README.md)

# Deployment

- This repo's CI builds `ghcr.io/thetangentline/tangent` (API + worker) and `ghcr.io/thetangentline/tangent-operator`, multi-arch.
- Releases are tagged with semver; release notes generated from conventional commits.
- The Helm chart in `deploy/helm/tangent` deploys API (Deployment, 2 replicas), worker (Deployment), a migration Job (ArgoCD `PreSync` hook), Service, Ingress, ServiceMonitor and PodDisruptionBudget.
- The demo environment runs on a single k3s ARM node provisioned by `infra/terraform/envs/demo`, with TLS from cert-manager and Let's Encrypt.

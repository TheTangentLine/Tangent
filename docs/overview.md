[← Docs index](README.md)

# Overview

## Goals and non-goals

### Goals

- **Self-service:** a developer goes from idea to a running service in a dev environment in under 10 minutes with no human in the loop.
- **Golden paths:** every service starts with logging, tracing, health checks, a Dockerfile, a Helm chart and CI already wired.
- **Git as the only write path to the cluster:** the platform never runs `kubectl apply`. It commits to Git and ArgoCD reconciles.
- **Guardrails, not gates:** quotas, network policies and admission policies are applied automatically; approvals are only required for production and for exceeding default quotas.
- **Auditability:** every action is attributable to a user and recorded in an append-only log.
- **Multi-tenancy by team:** teams own services; RBAC is enforced in the API and in Kubernetes.

### Non-goals (for v1)

- Multi-cluster or multi-cloud scheduling.
- Replacing a full service mesh.
- Managing stateful production databases (Terraform modules provision them, but backups and failover are out of scope).
- A plugin marketplace.

## Core concepts

| Concept | Description |
|---|---|
| **User** | A person authenticated through OIDC. Identity comes from the identity provider; Tangent stores only a reference and profile cache. |
| **Team** | The unit of ownership. Owns services, has members with roles, has a cost budget. |
| **Template** | A versioned golden-path project (for example `go-http-api@v1.2.0`) with parameters. |
| **Service** | A deployable unit created from a template. Has one repository and one or more environments. |
| **Environment** | A service deployed to a tier (`dev`, `staging`, `prod`). Maps to a Kubernetes namespace per team per tier. |
| **TeamEnvironment** | The Kubernetes custom resource that the operator reconciles into a namespace with quotas, RBAC and network policies. |
| **Provisioning request** | A record of "do this to the platform" (create service, add environment, delete service). Executed as an ordered list of steps. |
| **Approval** | A gate attached to a provisioning request that requires a decision from a user with the right role. |
| **Audit event** | An immutable record of who did what, when, to which resource, and the outcome. |

### Roles

| Role | Scope | Can |
|---|---|---|
| `platform-admin` | Global | Everything, including managing templates and approving quota increases |
| `team-lead` | Team | Manage members, create services, approve production environments for their team |
| `developer` | Team | Create services and non-production environments, trigger deploys |
| `viewer` | Team | Read-only |

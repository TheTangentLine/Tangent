[← Project README](../README.md)

# Tangent documentation

This folder is the source of truth for how Tangent is designed and built. Start with the overview and architecture, then read the component specs in the order of the [roadmap](roadmap.md).

## Start here

| Doc | What it covers |
|---|---|
| [Overview](overview.md) | Goals, non-goals, core concepts and roles |
| [Architecture](architecture.md) | Component overview, "create service" lifecycle, design principles |
| [Tech stack](tech-stack.md) | Every technology choice and why |
| [Repository layout](repository-layout.md) | Monorepo structure and the separate GitOps repo |

## Component specs

| Doc | What it covers |
|---|---|
| [Platform API](components/platform-api.md) | Auth, RBAC, approval flow, endpoints, middleware, idempotency |
| [Data model](components/data-model.md) | Postgres schema and entity relationships |
| [Provisioning engine](components/provisioning-engine.md) | Step contract, `create_service` workflow, state machines, job queue, SSE |
| [Kubernetes operator](components/operator.md) | `TeamEnvironment` CRD, Go types, reconcile logic |
| [Golden-path templates](components/templates.md) | Template structure, manifest, required features |
| [GitOps layer](components/gitops.md) | `tangent-gitops` layout, ApplicationSet, promotion flow |
| [CI pipelines](components/ci.md) | Reusable workflow and stages |
| [Infrastructure](components/infrastructure.md) | Terraform, platform namespaces, bootstrap order |
| [Security and policy](components/security.md) | Credentials, write paths, Kyverno policies, secrets, supply chain |
| [Observability](components/observability.md) | Platform metrics, alerts, traces, tenant dashboards |
| [Cost tracking](components/cost-tracking.md) | OpenCost integration and budgets |
| [Developer portal](components/portal.md) | Pages and what each does |

## Building and running

| Doc | What it covers |
|---|---|
| [Conventions](conventions.md) | Naming, labels, commit trailers |
| [Local development](local-development.md) | Prerequisites, make targets, first run |
| [Configuration](configuration.md) | Environment variables |
| [Testing strategy](testing.md) | Test levels and the first cases to write |
| [Deployment](deployment.md) | Images, Helm chart, demo environment |
| [Failure modes](failure-modes.md) | What breaks and how the system recovers |

## Planning and decisions

| Doc | What it covers |
|---|---|
| [Roadmap](roadmap.md) | Milestones M0–M6 with "done when" criteria |
| [Architecture decision records](adr/README.md) | ADR format and the initial list to write |

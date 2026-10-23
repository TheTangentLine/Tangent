[← Docs index](README.md)

# Architecture

## Component overview

```mermaid
flowchart TD
    Portal[Developer portal<br/>Next.js] --> API[Platform API<br/>Go]
    API --> PG[(Postgres)]
    API --> IDP[Keycloak<br/>OIDC]
    Worker[Provisioning worker<br/>Go] --> PG
    Worker --> GH[GitHub<br/>repos and Actions]
    Worker --> GitOps[GitOps repo]
    GitOps --> Argo[ArgoCD]
    Argo --> K8s[Kubernetes cluster]
    Operator[Tangent operator<br/>Go] --> K8s
    K8s --> Prom[Prometheus / Loki]
    K8s --> OpenCost[OpenCost]
    Prom --> Grafana[Grafana]
    API --> OpenCost
    API --> Argo
```

The **API** and the **worker** are the same Go binary started in two modes (`tangent api`, `tangent worker`). They share packages and the database, but scale independently.

## Lifecycle of "create service"

```mermaid
sequenceDiagram
    actor Dev as Developer
    participant P as Portal
    participant A as Platform API
    participant DB as Postgres
    participant W as Worker
    participant GH as GitHub
    participant G as GitOps repo
    participant AR as ArgoCD
    participant O as Operator

    Dev->>P: Fill "new service" form
    P->>A: POST /v1/services (Idempotency-Key)
    A->>DB: tx: insert service, request, steps, outbox, audit
    A-->>P: 202 Accepted {request_id}
    P->>A: GET /v1/requests/{id}/events (SSE)
    W->>DB: claim next step (SKIP LOCKED)
    W->>GH: create repo, push rendered template, set protections
    W->>G: commit TeamEnvironment + app manifests
    AR->>G: detect change
    AR->>O: apply TeamEnvironment CR
    O->>O: reconcile namespace, quota, RBAC, netpol
    AR->>AR: sync app, wait for healthy
    W->>AR: poll application health
    W->>DB: mark request succeeded
    A-->>P: SSE: step updates, final status
```

## Design principles

1. **Git is the write path.** All cluster state lives in the GitOps repo. The worker writes to Git; ArgoCD writes to the cluster. Rollback is `git revert`.
2. **Every step is idempotent.** Running a step twice must give the same result as running it once. Steps check current state before acting.
3. **Durable intent first, side effects second.** The API only writes to Postgres. External calls happen in the worker, driven by rows in the database.
4. **The operator owns tenant scaffolding.** Namespaces, quotas, RBAC and network policies are derived from one `TeamEnvironment` resource, never hand-written.
5. **Least privilege everywhere.** The worker's GitHub App, Kubernetes service account and database user each get only what they need.

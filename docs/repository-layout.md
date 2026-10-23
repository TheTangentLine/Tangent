[← Docs index](README.md)

# Repository layout

A single monorepo:

```text
tangent/
├── cmd/
│   └── tangent/               # main.go: `tangent api | worker | migrate`
├── internal/
│   ├── api/                   # HTTP handlers, middleware, routing
│   ├── auth/                  # OIDC verification, RBAC policy
│   ├── audit/                 # audit event writer
│   ├── config/                # env-based config loading
│   ├── db/                    # sqlc generated code + queries/*.sql
│   ├── domain/                # core types: Team, Service, Environment...
│   ├── provisioning/          # engine: workflows, steps, runner
│   │   ├── steps/             # one file per step
│   │   └── workflows/         # create_service.go, add_environment.go...
│   ├── integrations/
│   │   ├── github/            # GitHub App client (interface + impl + fake)
│   │   ├── gitops/            # clone, write, commit, push
│   │   ├── argocd/            # read application health
│   │   └── opencost/          # cost queries
│   ├── templates/             # template rendering
│   └── telemetry/             # OpenTelemetry setup, metrics
├── operator/                  # kubebuilder project (separate go module)
│   ├── api/v1alpha1/          # TeamEnvironment types
│   ├── internal/controller/   # reconciler
│   └── config/                # CRD, RBAC, manager manifests
├── portal/                    # Next.js app
├── templates/                 # golden-path templates
│   ├── go-http-api/
│   ├── node-service/
│   └── python-worker/
├── gitops/                    # bootstrap manifests (platform's own apps)
│   ├── bootstrap/             # argocd, app-of-apps
│   └── platform/              # tangent, keycloak, monitoring, kyverno...
├── infra/
│   └── terraform/
│       ├── modules/
│       └── envs/{local,demo}/
├── .github/workflows/         # CI for this repo + reusable workflows
├── migrations/                # SQL migrations
├── docs/                      # full specification (start at docs/README.md)
│   ├── components/            # one spec per component
│   └── adr/                   # architecture decision records
├── deploy/helm/tangent/       # Helm chart for API + worker
├── Tiltfile
├── Makefile
└── README.md
```

The GitOps repo that holds **tenant** workloads (`tangent-gitops`) is a separate repository, because the platform writes to it automatically and it should have its own history and branch protections.

[← Docs index](../README.md)

# Developer portal

Next.js (App Router) with TypeScript. Authenticates with Keycloak (OIDC + PKCE) and calls the Platform API with the access token.

| Page | Purpose |
|---|---|
| `/` | Your teams, your services, pending approvals, recent activity |
| `/catalog` | Search and filter all services; owner, template, environments, health |
| `/services/new` | Wizard: template → parameters (form generated from JSON Schema) → team → environments → review |
| `/services/[id]` | Overview, environments with links (app, repo, ArgoCD, Grafana), costs, recent deploys |
| `/requests/[id]` | Live step timeline over SSE, error details, retry and cancel buttons |
| `/approvals` | Approve or reject with comment |
| `/teams/[slug]` | Members, roles, budget, cost chart |
| `/audit` | Filterable audit log (admins and team leads) |
| `/templates` | Template catalogue and versions |

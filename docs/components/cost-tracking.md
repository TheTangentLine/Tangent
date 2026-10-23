[← Docs index](../README.md)

# Cost tracking

- OpenCost allocates CPU, memory, storage and network cost per pod. Because every workload carries `tangent.dev/team` and `tangent.dev/service` labels (enforced by Kyverno), allocation by team and service is a single query.
- The API calls the OpenCost allocation API, for example:

  ```text
  GET /allocation?window=30d&aggregate=label:tangent.dev/team,label:tangent.dev/service
  ```

- Results are cached for 15 minutes and returned at `/v1/teams/{slug}/costs`.
- If a team has a `monthly_budget_cents`, the portal shows burn rate and the API emits a `team.budget.exceeded` audit event and notification when projected spend passes the budget.
- On a single-node demo, set custom pricing in OpenCost so the numbers look realistic.

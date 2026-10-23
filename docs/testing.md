[← Docs index](README.md)

# Testing strategy

| Level | Tooling | Covers |
|---|---|---|
| Unit | Go `testing`, table-driven | RBAC policy, template rendering, step logic with fakes, backoff |
| Integration | `testcontainers-go` (real Postgres) | Queries, job claiming under concurrency, idempotency keys, approval flow |
| Operator | controller-runtime `envtest` | Reconcile creates and repairs resources, finalizer and deletion protection |
| Contract | OpenAPI spec + generated client | Portal and API stay in sync |
| End-to-end | kind in CI + test GitHub org | Create a service through the API and assert the pod is running and the dashboard exists |
| Chaos (later) | kill the worker mid-request | Request resumes and completes without duplicate side effects |

Every integration in `internal/integrations/` is an interface with a real implementation and a fake, so the engine can be tested without network access.

Key test cases to write first:

1. Two workers never claim the same step.
2. A step that fails with a retryable error is retried with backoff and eventually succeeds.
3. Running `create_repo` twice does not create two repos.
4. Cancelling after step 6 reverts the GitOps commit and archives the repo.
5. A developer cannot approve their own prod request.
6. Deleting a namespace by hand causes the operator to recreate it.

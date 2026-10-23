[← Docs index](README.md)

# Failure modes and recovery

| Failure | Behaviour |
|---|---|
| GitHub API rate limited or down | Step returns retryable error; backoff; request shows "retrying" in the portal |
| Worker crashes mid-step | Lock expires; reaper returns step to `pending`; idempotent step re-runs safely |
| ArgoCD sync never becomes healthy | `wait_for_sync` times out → request `failed` with ArgoCD's health message; user can retry or cancel |
| Postgres unavailable | API returns `503` on writes and readiness fails; worker pauses claiming |
| Operator down | Existing namespaces keep working; new `TeamEnvironment`s wait; ArgoCD shows them out of sync |
| Someone edits a namespace quota by hand | Operator reverts it on next reconcile; Kubernetes Event explains why |
| Bad template version released | Mark `deprecated` so it disappears from the wizard; existing services unaffected |
| GitOps repo corrupted by a bad commit | `git revert`; ArgoCD reconciles back |

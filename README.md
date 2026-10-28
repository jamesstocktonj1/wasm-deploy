# Wasm Deploy - Push-Based ArgoCD Deployment

**Ship a version, then choose when it goes live.**

This repo shows push-based deployment with ArgoCD. Each release is an immutable, versioned artifact in an OCI registry. A deploy is one command that points ArgoCD at the version you pick. No manifest commits, no waiting for a branch to reconcile.

```
Publish Function                              Deploy Function
────────────────                              ───────────────
build Wasm component ──► function:0.1.1       argocd app set wasm-deploy
package workload.yaml ─► workload:0.1.1  ───►   --revision 0.1.1
        (ghcr.io, OCI)                        ArgoCD pulls workload:0.1.1 and syncs
```

## Why push-based?

- **Deploys are explicit.** Merging to a branch doesn't change production. Someone, or a pipeline, chooses a version and deploys it.
- **Releases are immutable.** The component and its manifests share one version tag, so `0.1.1` always means the same thing.
- **Rollback is a deploy.** Run the same workflow with an older version. There's no revert commit.
- **One registry for everything.** The Wasm component and the Kubernetes manifests live side by side in OCI, with no separate config repo to keep in sync.
- **Still ArgoCD.** You keep its sync, health status, drift correction and UI.

## Publish Function

Run with a `version` (for example `0.1.1`). It:

1. Bumps the version in `main.go` and the image tag in `workload.yaml`.
2. Builds the Wasm component and pushes it to `ghcr.io/jamesstocktonj1/wasm-deploy/function:<version>`.
3. Packages `workload.yaml` as a `.tar.gz` and pushes it to `ghcr.io/jamesstocktonj1/wasm-deploy/workload:<version>`.

The manifests inside `workload:<version>` reference `function:<version>`, so the two can't get out of step.

## Deploy Function

Run with the same `version`. It joins the tailnet, then runs:

```bash
argocd app set wasm-deploy --revision <version>
```

The ArgoCD `Application` ([application.yaml](application.yaml)) uses the workload artifact as its source. Changing its revision makes ArgoCD pull that version and sync it automatically.

To roll back, run **Deploy Function** with an earlier version.

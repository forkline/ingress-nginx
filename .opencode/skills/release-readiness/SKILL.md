---
name: release-readiness
description: The nine-item checklist a routine must pass before it opens a release pull request of forkline/ingress-nginx, plus the post-merge verification and the confidence gate ("almost sure" releases). Load together with the release skill.
---

## Rule

Release only when almost sure. Merging a `release: prepare` pull request IS the release: `auto-tag.yml` creates a signed tag and the image, chart and plugin workflows publish what users pull, with no second gate. A security fix is urgent, a broken release is worse than a late one. Run the checklist before opening a release pull request; if any item fails, do not open it: keep the tracking issue (see `release-prep`) and escalate per the `maintenance` protocol when a human decision is needed. Unknown counts as fail. No routine merges a release pull request in this phase (`merge-gate` skill, `Release mode`).

Always `git fetch --unshallow --tags` first when the clone is shallow (#237 shipped an incomplete changelog from a shallow clone). Each item names the incident it comes from.

## Checklist (items 1 to 9, all must pass)

1. **Worth releasing.** At least one user-facing change since the last tag (`git describe --tags --abbrev=0`): an NGINX bump, a security bump including Go, a CVE fix, a chart or image fix. Pre-commit and digest noise alone is not a reason. Cadence since 2026-05-18 was 5 to 29 days; shorter only for security (#238).
2. **`main` is green on the exact commit that will be released**, not "recently green": every workflow (`pre-commit`, `Go`, `E2E` including Kubernetes, chroot and chart e2e, `Security` with govulncheck, `Helm`) is `success` on `origin/main` HEAD (`git-api actions list --ref main`), and E2E was not red on main in the previous 24 hours (v2026.11.6 was cut after 36 red hours). Helm and Lua are not required checks: check them explicitly; rerun a flaky Helm failure once, a second failure blocks. Pending is not green: `git-api checks wait main --timeout 1500` once, then fail the item. A Scorecards-only Security failure is not blocking; a govulncheck failure is.
3. **Nothing is claimed but unmerged.** No open pull request or draft titled or labelled as a security bump or a `rebase`, no Renovate NGINX or Go pull request waiting on a human, no open `release: prepare` pull request (except the stale bot-created one you regenerate), no open `needs-maintainer` or `security` item that concerns the release content (`git-api issue list --labels needs-maintainer`, `--labels security`).
4. **The release pull request contains only generated release files**: `CHANGELOG.md`, `TAG` and `images/*/TAG`, `NGINX_BASE`, `charts/ingress-nginx/{Chart.yaml,values.yaml,README.md}`, the `README.md` version table, deploy manifests and scripts, the `artifacthub.io/changes` annotation. Any other path blocks (#120 carried four controller changes, never bundle behavior with a release).
5. **No stale version strings.** Grep the repository for the previous tag and version (`README.md.gotmpl`, docs examples, `artifacthub.io/changes`, deploy manifests, chart README) and run every generator followed by `git diff --exit-code` (`make update-version`-based `.ci/release.sh`, `helm-docs --chart-search-root charts`, `bash hack/generate-deploy-scripts.sh`). Evidence: #325, #351, #392, #437, #532.
6. **The changelog is complete and quiet.** Every pull request merged since the last tag appears in the new CHANGELOG.md entry, or is intentionally squashed by `.ci/squash-changelog.py` (digest, pre-commit and plain Go module updates); the NGINX version, the Go version and every CVE or `GO-` id appear by name; no duplicates; `git rev-list --count <last tag>..origin/main` matches the entries you expect. Evidence: #237, #567, #634, and Go security bumps that never reached the changelog.
7. **Image tags are consistent and `v`-less.** `TAG` files, chart `values.yaml` `tag:` fields and the published tags use `YYYY.M.D`; git tags are `vYYYY.M.D`; controller, nginx, kube-webhook-certgen and every helper follow one scheme; the release notes do not contain `v`-prefixed image tags. Evidence: #58, #61, #64, `59835202`.
8. **Base image provenance.** `NGINX_BASE` equals `ghcr.io/forkline/ingress-nginx/nginx:$NGINX_VERSION` from `images/nginx/rootfs/build.sh`, and the E2E run on the release commit built that base image from the current patches (no cache skip: `09aeb3e4a`, #530).
9. **Chart lint with the release version.** `ct lint` or `helm lint` pass with the new version, and the annotation YAML is quoted (an unquoted annotation broke `ct lint` in #237).

## Release pull request mechanics

Follow the `release` skill and `.ci/release.sh`; never edit generated files by hand. Branch `release/prep` from a fresh `origin/main`, run `.ci/release.sh`, rename the branch to `release/<version>`, push, open the pull request titled `release: prepare <version>`. The body lists the checklist with the evidence for each item and the changelog highlights. Only a bot-created release branch with no human commit may be regenerated with `git push --force-with-lease`. Never tag and never create a GitHub release.

## Confidence gate

Do not open the pull request, or give a would-merge verdict, when any of these happened: a rerun of E2E was needed to pass, a bot commit changed non-generated files, the NGINX patch set changed, or checklist item 6 had to be hand-fixed. Historically 8 of 12 release pull requests hit one of these, so expect to escalate often until the generator improves.

## Post-merge verification (mandatory, done by the `release-verify` routine or a human)

Before anything is announced: wait for `Auto tag`, `Release images`, `Helm` and `plugin` runs on the new tag (`git-api actions list --ref <tag>`). Then `helm pull oci://ghcr.io/forkline/helm-charts/ingress-nginx --version <version>`, extract every image reference from its `values.yaml`, and check each one exists AND lists both `linux/amd64` and `linux/arm64` (`docker manifest inspect ghcr.io/forkline/ingress-nginx/<image>:<tag>`, or `skopeo inspect --raw` / `crane manifest` when docker is not available). Confirm the image paths in the GitHub release notes exist. On any failure do not retag, do not push or tag images by hand (the manual retag in #58 and #64 covered only amd64): escalate and recommend cutting a new release; a same-day hotfix with the script's `-N` suffix is the maintainer's decision.

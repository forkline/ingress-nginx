---
name: release-readiness
description: Checklist a routine must pass before it opens or merges a release pull request of forkline/ingress-nginx ("almost sure" releases). Load together with the release skill.
---

## Rule

Release only when almost sure. A release is a tag and a publish of images and charts that users pull: a security fix is urgent, a broken release is worse than a late one. Do the checklist before opening a release pull request, and again before the `merge-gate` merges one. If any item fails: do not open (or merge) it, escalate with the protocol of the `maintenance` skill on the pull request that triggered the release, and name the failed item.

## Checklist

1. Main CI is fully green on the commit that will be released: `git-api checks list main` (or `git-api actions list --ref main`) shows the Go, Pre-commit, E2E, Helm and Security workflows completed with success for `origin/main`'s head. Branch protection requires only pre-commit, govulncheck, Go and E2E: Helm (chart lint and tests, flaky about a third of the time) and Lua are NOT required checks, so check them explicitly on the release commit; rerun a flaky Helm failure once, and a second failure blocks. Pending is not green: wait with `git-api checks wait main --timeout 1500` once, then fail the item. A Security failure that is not govulncheck-related (Scorecards) is not blocking; a govulncheck failure is.
2. The changelog is complete: every pull request merged into main since the previous tag (`git log $(git describe --tags --abbrev=0)..origin/main --oneline`, `git-api pr list --state closed`) appears in the new CHANGELOG.md entry, except the ones `.ci/squash-changelog.py` squashes on purpose (digest updates, pre-commit hooks, go modules).
3. Version files are consistent: `TAG`, every `images/*/TAG` and `charts/ingress-nginx/Chart.yaml` (version and appVersion) carry the new version; `NGINX_BASE` equals `ghcr.io/forkline/ingress-nginx/nginx:$NGINX_VERSION` from `images/nginx/rootfs/build.sh`; the README Supported Versions table has one new row with the right NGINX and Alpine versions; `git diff --stat` touches only release files.
4. Nothing blocks: no open issue or pull request labelled `needs-maintainer` or `security` that concerns the release content (`git-api issue list --labels needs-maintainer`, `--labels security`), no open `fix(security)` pull request that is almost ready and should ship in the same release, no open `release: prepare` pull request (see below).
5. The new version tag does not exist yet (`git tag -l`), and the version is not unchanged (`.ci/release.sh` prints "Version unchanged" otherwise).

## Release pull request mechanics

Follow the `release` skill and `.ci/release.sh`; never edit generated files by hand. Branch `release/prep` from a fresh `origin/main`, run `.ci/release.sh`, rename the branch to `release/<version>`, push, open the pull request titled `release: prepare <version>`. The body lists the checklist with the evidence for each item. Only a bot-created release branch with no human commit may be regenerated with `git push --force-with-lease`. Never tag and never create a GitHub release: `auto-tag.yml` does that when the release pull request merges.

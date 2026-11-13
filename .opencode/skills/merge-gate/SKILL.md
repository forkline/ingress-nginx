---
name: merge-gate
description: Low-risk policy and procedure for the merge-gate routine, which independently reviews green pull requests authored by the Forkline bot (or Renovate pull requests carrying bot commits) and posts a verdict; in enforce mode it merges or escalates. Load before reviewing a pull request for merging.
---

## Mode

`Mode: observe`

`Release mode: observe`

- `observe` (current): post ONE verdict comment per pull request head (first line `<!-- forkline-merge-gate:<head sha> -->`, none if that marker exists) saying `would merge` or `would escalate`, with the criteria evidence and the reasons. Do nothing else: no merge, no label, no assignment, no auto-merge change, no review. The verdicts are the track record that decides promotion (a maintainer compares them with what he did himself).
- `enforce`: merge or escalate for real, as in the procedure. A maintainer promotes the mode by editing the line above in a pull request. Never edit it yourself.
- `Release mode` applies to `release: prepare` pull requests. Release pull requests are never merged by a routine in this phase, whatever the value: the merge IS the release (`auto-tag.yml` tags and publishes). In `observe` the verdict comment says whether the `release-readiness` checklist would pass; nothing else is done. Merging them is a separate decision of the maintainer, recorded here, not an implicit step of `Mode: enforce`.

## Procedure

0. This run checks out `main` (the routine pins `agent.ref: main`), so the skill text and the Mode lines you read are the ones merged on main, not the pull request's. If this skill, the Mode line or the maintenance skill is missing or unreadable, or the Mode line is anything but exactly `enforce`, behave as `observe`: verdict comment only.
1. Load the `maintenance` skill (hard rules, risk lanes, escalation protocol) and the `release-readiness` skill for release pull requests. Define the helper of the maintenance skill (`API`, `gh_api`). Run the "Bootstrap the clone" section, then fetch the head: `git fetch origin "pull/$N/head:pr-$N"` (the diff is `git diff origin/main...pr-$N`; the branch itself is not checked out).
2. Read the pull request: `gh_api "$API/pulls/$N"`. Stop silently unless: `state` is `open`, `draft` is false, `head.repo.fork` is false, `base.ref` is `main`, and either `user.login` is `forkline[bot]`, or it is `renovate[bot]` and the branch carries at least one commit by `forkline[bot]` (`git log origin/main..origin/<branch> --format='%an'`). Renovate pull requests without bot commits merge through Renovate's own restricted automerge; human pull requests are never touched. Branches `pre-commit-*` are filtered out by the routine and ignored here too: no bot commit may ever be pushed to them (maintenance rule 4).
3. Be an independent reviewer: judge the diff (`git diff origin/main...pr-$N`) and the evidence yourself. Do not trust the author's description or an earlier bot review (marker `<!-- forkline-auto-review -->`): read them as input, not as a verdict.
4. Readiness (all of it, not only what branch protection requires). Branch protection on main requires `pre-commit`, `govulncheck`, `Go checks passed` and `E2E checks passed` (not strict), 0 approving reviews, and ALL CONVERSATIONS RESOLVED. The gate needs more:
   - `mergeable_state` is `clean` (from the same `gh_api` read of `pulls/$N`, together with `head.sha`). Not `unstable`, `blocked`, `behind`, `dirty` or `unknown`: wait once and re-read, then give the verdict from what is there.
   - EVERY check on the head commit is `success` or an expected `skipped`: `git-api checks list <head sha>` has nothing failed, pending, cancelled or neutral-by-failure. Helm, Dependency review, Release images and Lua count; they are not required checks and #506 merged with Helm red. A run that needed a rerun to go green is not green.
   - No unresolved review thread (GraphQL `pullRequest.reviewThreads.nodes.isResolved`, verified query through `gh_api -X POST https://api.github.com/graphql`). An unresolved thread you cannot resolve means would-escalate.
   - No `needs-maintainer` label, and no human comment on the pull request that is unanswered.
   - `main` itself is not red: `git-api actions list --ref main`, the latest push run of every workflow. While main CI is red, nothing merges.
5. Classify with the criteria below. Anything not listed as low risk is not low risk.
6. Outcome in `observe`: the verdict comment of the Mode section. Outcome in `enforce`:
   - Every criterion met: post ONE plain comment FIRST (first line `<!-- forkline-merge-gate:<head sha> -->`: `merging`, with the criteria evidence). Then re-read `pulls/$N`: if `head.sha` changed or `mergeable_state` is no longer `clean`, do not merge. Then `git-api pr merge $N --method squash` . The comment must come first: the runner is stopped about 2 seconds after a merge and cannot post afterwards. `TODO(#2954)`: replace the read-and-merge sequence by native `git-api pr merge --if-ready --sha <head sha>` when it exists.
   - Anything else: escalate per the `maintenance` protocol (label `needs-maintainer`, assign `pando85`, one comment: which criterion failed, options, recommendation).
7. Never merge: a head that changed after the CI run you looked at, while a `needs-maintainer` label is on it, from an unclear state, a draft, a release pull request, or anything in the never-auto-merge list.
8. After a merge the next run of `maintenance-digest` reports main CI health and the checks at merge time; a red main after a gate merge is then a maintainer matter. Do not stack fixes.

## Low-risk criteria (all must hold)

1. Author is `renovate[bot]`, or the pull request is a mechanical fix by `forkline[bot]`, and there is no commit by a human and no non-mechanical bot commit on the branch. A Renovate branch with bot commits qualifies only when those commits are in the mechanical list of maintenance rule 4.
2. Changed paths are a subset of: `.pre-commit-config.yaml`; `images/go-grpc-greeter-server/rootfs/go.mod` and `go.sum`; `hack/.tool-versions`; `docs/requirements.txt`; `.github/actions/mkdocs/Dockerfile`; `go.mod` and `go.sum` (and `images/*/rootfs/go.{mod,sum}`) for PATCH or MINOR updates only, with no import-path change; `GOLANG_VERSION` plus the go.mod `go` line as a stand-alone `fix(security)` pull request. At most about 60 changed lines excluding `go.sum`.
3. All checks on the head commit are green or expected-skipped (step 4), including the ones branch protection does not require.
4. The title is true to the diff: a change exists in the diff for each dependency named in the title, nothing unrelated is in it, the diff is not empty (#92, #209, #339, #638).
5. For a GitHub Actions or tool pin in a workflow (not low risk by default, #506, #622): only with `E2E` and `Helm` actually run (not skipped) on the pull request and green, and the post-merge check of step 8; otherwise would-escalate.
6. The dependency is not on the known-bad list: helm-docs 19.x, `go.yaml.in/yaml` v4, `k8s.io/client-go` `v11+incompatible`, chart-testing mismatches.
7. Rate and state: at most one merge per dependency per 24 hours, and none while `main` CI is red.

## Never auto-merge

Evidence and reasons in the `maintenance` skill (risk lanes). Always would-escalate or ignore:

- Major version bumps of anything (#638, #642); NGINX bumps of any kind and everything under `images/nginx/rootfs/patches/` or `build.sh` beyond the version line; a pull request that reverted or removed the dependency it claims to bump (#209, #339).
- `internal/**`, `cmd/**`, `pkg/**`, `rootfs/etc/nginx/lua/**`, `rootfs/etc/nginx/template/**` non-test code; `charts/ingress-nginx/templates/**`.
- Release scripts and generators (`.ci/**`, `cliff.toml`, `magefiles/**`, `.github/workflows/{auto-tag,docker_images,helm}.yml`); anything that defines green: `.github/workflows/e2e.yml` logic, `test/e2e/**`, the pre-commit hook list, `.commitlintrc.json`, `.yamllint`, `renovate.json`, branch protection.
- Workflow logic changes (only pure action version pins may be considered, criterion 5); Dockerfile base images (alpine, golang) and Go toolchain bumps bundled with anything else.
- Release pull requests (`release: prepare ...`), external or human pull requests, anything with an unresolved human comment.
- A pull request whose NGINX, patches or generated files were changed by a bot in a way the title does not state (#346).

## Release pull requests

Observe only (see Mode). The verdict comment contains the 9-item checklist of the `release-readiness` skill re-run with evidence (do not trust the body of the pull request) and the confidence gate: `would escalate` when any item fails, an E2E rerun was needed, a bot commit changed non-generated files, the NGINX patch set changed, or a checklist item had to be hand-fixed (8 of the 12 release pull requests so far needed corrections inside the pull request).

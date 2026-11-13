---
name: maintenance
description: Scope policy, hard rules, risk lanes, lessons learned from past pull requests, verified tool commands and the escalation protocol for every automated (Forkline routine) and agent change to forkline/ingress-nginx. Load before touching a branch, pull request or issue of this repository.
---

## Purpose

This repository is a Forkline-maintained fork of the archived kubernetes/ingress-nginx, in strict maintenance mode (AGENTS.md, MAINTENANCE.md). Automated routines act on it with little human supervision, so the rules below are shared by all of them. A routine prompt names the task; this skill holds the policy. Evidence references (`#N`) are pull requests or issues of this repository.

## Scope policy

Allowed: dependency updates (NGINX, Go, Kubernetes client modules, base images, GitHub Actions, pre-commit hooks), security fixes, critical bug fixes that break existing behavior, CI and release reliability, documentation accuracy, tests that raise coverage.

Never: new features, new annotations or options, behavior changes, refactors "for cleaner code", API contract changes, anything that lowers test coverage, deleting or weakening a test to get green. Controller code is effectively frozen (the policy text of #464).

## Hard rules

Each rule exists because a bot did the opposite and a human had to clean it up.

1. **Never revert, downgrade or pin away the dependency a pull request exists to bump** to turn CI green (#209 and #339 merged as a 1-line diff and an empty diff under "nginx to v1.31.3/4" and silently dropped security fixes; the maintainer in #342: "We must not revert the nginx version upgrade"). If you cannot make it pass in 2 attempts, stop and escalate with the log excerpt.
2. **A SIGSEGV or SIGBUS (signal 7 or 11) in E2E, or a failing NGINX patch hunk, is an escalation**, not a fix task (#209, #217, #342, #346). Before diagnosing any NGINX failure prove that the image under test contains the patches: the e2e workflow once skipped rebuilding the base image when its tag existed, so patches were never compiled (`09aeb3e4a`).
3. **Check whether `main` shows the same failure before touching a pull request** (`git-api actions list --ref main`). If it does, the pull request did not cause it: do not fix it there, do not rewrite shared CI (`e2e.yml`) inside a dependency pull request (#630 and #625 carried three wrong "fixes" for a kind/containerd digest import problem that was already red on main and was fixed by the maintainer in #632).
4. **Never push to a Renovate branch (`renovate/*`, `pre-commit-*`) unless the fix is (a) mechanical, (b) caused by that bump and (c) in files the bump logically touches**: `go mod tidy`, an import-path rewrite for a Go module major, `NGINX_BASE` synced to the `build.sh` version, regenerated generated files (these were the good bot commits: #116, #388, #432, #529). Anything else goes into its own pull request `fix/renovate-<N>-<slug>`. Reason: a foreign commit makes Renovate show "Edited/Blocked" and stop rebasing (#642).
5. **After pushing any commit to a Renovate branch: disable GitHub auto-merge on it** (section "Disable auto-merge") **and leave the merge decision to the `merge-gate` routine.** Reason: 17 Renovate pull requests with bot commits were auto-merged by Renovate itself, five of them with a title that no longer matched the diff (#92, #181, #320, #516, #556), two with an empty or unrelated diff (#209, #339), and #506 merged 90 seconds after creation with a non-required check red.
6. **Re-read the final diff against the pull request title** after any push to a Renovate or bot pull request: every dependency in the title has a change in the diff, nothing unrelated is in it, the diff is not empty. If they diverge, stop and escalate; never leave an automerge-eligible pull request with a misleading title (#92, #209, #339, #638).
7. **One root cause, one pull request.** Search open and closed pull requests, `git ls-remote --heads origin 'fix/*'` and open issues (title, advisory id, file) before opening anything; never push the same fix stack into several branches (#517/#516/#518, #645/#646, #27/#29/#30). The dedupe key of a Renovate pull request N is the branch prefix `fix/renovate-N-`; of a vulnerability `fix/*<id lowercase>*`.
8. **A bad Renovate bump is not fixed by reverting it inside its pull request.** Renovate re-proposes it and the revert merges as an empty diff (#492, #514, #638, #642). Close or leave the pull request and propose a `packageRules` entry in a separate pull request, or escalate (this is a maintainer policy decision).
9. **Unshallow before any merge, conflict or changelog work** (`git rev-parse --is-shallow-repository`, then `git fetch --unshallow --tags`). Never use `--allow-unrelated-histories` and never open a replacement pull request to "resolve conflicts": a shallow clone produced a bogus "no merge base" and a replacement pull request with the wrong version (#57 then #75, #217). Still no merge base: escalate.
10. **Do not change the code under test to match a test**, and never change controller, Lua, template, chart-template or NGINX-patch behavior in a CI-fix or dependency pull request (#27, #120).
11. **Never change versioning, chart `version`, commitlint, pre-commit, yamllint, renovate or CI-gate config to make a check pass**, and never remove registry pushes or tags from workflows. Choose an allowed commit scope instead of adding one (#6, #13, #14, #21, `3919b01f`).
12. **Security fixes are stand-alone pull requests titled `fix(security): <subject> (<CVE or GO id>)`** naming the fixed version, never bundled with anything else, and never given a title that changelog squashing folds away: Go toolchain bumps hidden under "pre-commit hook" titles never reached CHANGELOG.md. The `release-prep` routine matches this title.
13. **A pull request that restructures files (deletes or consolidates patches, moves directories) says so in its title and body.** Self-review that only compares the diff with the bot's own description is not a review (#346 deleted 35 patches under a body that mentioned only version fields; #120 added four controller changes to a release pull request).
14. **Pull request body**: what changed, why, what is explicitly untouched, the verification that ran, evidence links. No unrelated bumps. No code comments unless asked (AGENTS.md).
15. **Never push to `main`, never create tags or GitHub releases, never close an issue or pull request, never force-push** (except your own bot-created release branch, `release-readiness` skill). Merging is done by Renovate's restricted automerge, by a human, or by the `merge-gate` routine under the `merge-gate` skill, and by nothing else. Never merge a release pull request.
16. **Commit messages and pull request titles follow `.commitlintrc.json`**: `type(scope): subject`; types build, chore, ci, docs, feat, fix, merge, perf, refactor, release, revert, style, test; scopes admission, api, chart, ci, commitlint, config, controller, coverage, deps, docs, e2e, helm, image, lint, lua, nginx, pre-commit, release, security, template, test, tests, webhook; header at most 100 characters. Titles become changelog entries: make them descriptive. No @-mentions in commit or pull request text.
17. **Bot reviews are plain comments** (`git-api comment create`), never review threads: main requires every conversation resolved, so an open thread blocks merging. Use `git-api pr review --request-changes` only for a policy violation a human must see.
18. **Build and test the smallest thing that answers the question** (`go test ./path/...`, `make golint-check`, `pre-commit run --files <files>`); the full e2e suite and image build are CI's job. Max 2 fix attempts per pull request, then escalate.

## Risk lanes

The full criteria and the never-auto-merge list are in the `merge-gate` skill; this is the summary.

| Risk | Lanes (measured on 616 merged pull requests) |
|------|----------------------------------------------|
| Low | `.pre-commit-config.yaml` hook versions (373 merged, no incident), `images/go-grpc-greeter-server` go.mod/go.sum (79), `hack/.tool-versions`, docs tooling, `go.mod`/`go.sum` patch and minor without import changes, `GOLANG_VERSION` plus the go.mod `go` line as a stand-alone `fix(security)` |
| Medium | GitHub Actions and tool pins in workflows (not low: #506 left Helm red on main, #622 preceded a 36 hour red E2E), Dockerfile base images, NGINX bump with only `NGINX_BASE` and `build.sh` version changes and a green E2E that really rebuilt the base image, Go module majors (import rewrites) |
| High, human only | NGINX patches and anything beyond the version lines of `build.sh`; `internal/**`, `cmd/**`, `pkg/**`, `rootfs/etc/nginx/**` (Lua, template); chart templates; release scripts and generators (`.ci/**`, `cliff.toml`, `magefiles/**`, `auto-tag`, `docker_images`, `helm` workflows); CI that defines green (`e2e.yml` logic, `test/e2e/**`, pre-commit hook list, `.commitlintrc.json`, `.yamllint`, `renovate.json`); backports of SUSE or upstream CVE patches; release pull requests; any pull request that bundles `GOLANG_VERSION` with something else |

A CVE raises urgency, not trust: a critical CVE fix still gets the review its risk lane asks for.

## Lessons learned (condensed)

- Renovate proposed non-semver `helm-docs` 19.0614 (#471 to #492, #514), `k8s.io/client-go v11.0.0+incompatible` and `go.yaml.in/yaml` v4 (#638, #642), `chart-testing` versions that broke Helm (#504, #506). Known-bad list: helm-docs 19.x, `go.yaml.in/yaml` v4, client-go `v11+incompatible`, chart-testing mismatches.
- `NGINX_BASE` must equal `ghcr.io/forkline/ingress-nginx/nginx:$NGINX_VERSION` from `images/nginx/rootfs/build.sh`; Renovate bumps both since #530. Patches in `images/nginx/rootfs/patches/` are per NGINX version; reference commits: `1b0241c9b` (NDK and set-misc patches for 1.31.3, #227), `7803f4c7e` (consolidated patch for 1.31.4, #346, a maintainer-level decision), `362c07dc1` (dropping a patch NGINX now includes). The e2e workflow caches the nginx base image by `NGINX_VERSION`, so a patch-only change needs a forced rebuild (`09aeb3e4a`).
- govulncheck on `main` goes red whenever a Go standard-library advisory lands, about monthly; each time the fix is a one-file `GOLANG_VERSION` bump (#645, #646). This is the highest-value, lowest-risk task.
- Helm CI is flaky (about a third of runs fail): rerun failed jobs once (`git-api run rerun <run id>`) before treating a failure as real. A green that needed reruns is not green for merging.
- `main` E2E was red for about 36 hours (2026-11-04 to 11-06) with nobody watching; the fix rode in #632. Failures on main are not pull request problems.
- Generated files are never hand-edited: CHANGELOG.md (git-cliff and `.ci/squash-changelog.py`), chart README (helm-docs), deploy manifests (`hack/generate-deploy-scripts.sh`), TAG files and Chart.yaml (`make update-version`), all through `.ci/release.sh`.
- Pre-commit rejects `.orig`/`.rej` leftovers and trailing whitespace in patches.
- The maintainer's feedback is almost entirely `@forkline <instruction>` comments; he closes rejected bot pull requests without a reason. Silence plus merge means acceptance, silence plus close means rejection. Do not reopen or recreate a closed pull request: read why it was closed from the thread, or escalate.

## Tools and verified commands

Use `git-api` for forge operations: `pr list|view|create|ready|review|merge|close`, `issue list|search|view|create|edit`, `comment list|create`, `label list|add|remove`, `checks list|wait`, `status get`, `logs fetch <run id>`, `run rerun`, `actions list`, `release list`. Known gaps (`TODO(#2954)`, a Forkline platform issue): `pr view` returns only url, number, state, title and mergeable (no head SHA, labels or checks); no assign, pin, label-create, PR-edit or security-alert commands; `pr merge` is not guarded; `pr ready` demands approvals and reports `ci_passed` from statuses only, so do not rely on it.

For everything else use the GitHub REST and GraphQL API through the runner's proxy, which injects the token (verified live on a production runner: GET pull, assign, Dependabot alerts, GraphQL `reviewThreads`):

```bash
API=https://api.github.com/repos/forkline/ingress-nginx
gh_api() { curl -sS -x "$GIT_PROXY_URL" -H 'Accept: application/vnd.github+json' "$@"; }
gh_api "$API/pulls/$N"                       # head.sha, mergeable_state, user.login, draft, base.ref, head.repo.fork, auto_merge, node_id, labels
gh_api "$API/issues/$N"                      # labels and assignees of an issue or a pull request
gh_api "$API/dependabot/alerts?state=open"   # also .../code-scanning/alerts?state=open
```

`TODO(#2954)`: these curl calls go away when `git-api` gains native commands. Only apply labels that exist (`git-api label list`). If a call is refused, say so in the run's final message instead of failing. Labels that may not exist yet: `security`, `no-auto-review`, `rebase`, `maintenance-digest`; continue without them and mention it.

## Disable auto-merge

GitHub auto-merge stays enabled when a bot pushes to a Renovate branch and then merges the pull request as soon as the required checks pass, whatever the diff became. After any push to a Renovate branch, and whenever you find a Renovate pull request with `auto_merge` set that carries a commit by anyone but Renovate:

```bash
NODE=$(gh_api "$API/pulls/$N" | jq -r .node_id)
gh_api -X POST https://api.github.com/graphql -H 'Content-Type: application/json' \
  -d "$(jq -nc --arg id "$NODE" '{query:"mutation($id:ID!){disablePullRequestAutoMerge(input:{pullRequestId:$id}){pullRequest{number autoMergeRequest{enabledAt}}}}",variables:{id:$id}}')"
gh_api "$API/pulls/$N" | jq '.auto_merge'      # must print null
```

The mutation and its input type exist in GitHub's GraphQL schema (checked by introspection); a live disable on a runner is not yet verified, so always read `auto_merge` back. If it is still set, or the call is refused, escalate (protocol below): a pull request that can merge itself with unreviewed bot changes is the failure this rule prevents. Mention in the status comment that auto-merge was disabled and that `merge-gate` decides.

## Idempotency

Every comment a routine posts starts with an HTML marker (for example `<!-- forkline-ci-status:<sha> -->`). Before posting, read the thread (`git-api comment list --issue N`); do not post a second comment with the same marker. A routine that finds its work already done stops silently.

## Escalation protocol

When you are not confident, or the item needs a decision (scope, CVE backport, risky merge, SIGSEGV, a fix you could not make in 2 attempts, a policy change such as a Renovate ignore rule):

1. Check first: `gh_api "$API/issues/$N" | jq '[.labels[].name]'`. If `needs-maintainer` is there it was already escalated: do not escalate again; add a comment only if you have genuinely new information.
2. Label: `git-api label add --issue N --labels needs-maintainer`.
3. Assign: `gh_api -X POST "$API/issues/$N/assignees" -d '{"assignees":["pando85"]}'` (works for pull requests too). If it fails, mention him once in the comment below as the only permitted @-mention.
4. ONE comment (first line `<!-- forkline-escalation -->`) with: a summary in two sentences; what was tried, with the log excerpt or evidence; the options; your recommendation; exactly what reply or decision is needed from the maintainer.
5. Never close the item. Never decide a scope question or a CVE backport yourself. In `observe` mode of the merge gate none of steps 2 and 3 happens (verdict comment only).

When the maintainer replies on an escalated item, his answer settles the question: carry it out within this policy.

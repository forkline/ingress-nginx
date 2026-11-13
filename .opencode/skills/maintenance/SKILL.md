---
name: maintenance
description: Scope policy, risk taxonomy, lessons learned, tool limits and the escalation protocol for every automated (Forkline routine) and agent change to forkline/ingress-nginx. Load before touching a branch, pull request or issue of this repository.
---

## Purpose

This repository is a Forkline-maintained fork of the archived kubernetes/ingress-nginx, in strict maintenance mode (AGENTS.md, MAINTENANCE.md). Routines under `.forkline/routines/` act on it with little human supervision, so the rules below are shared by all of them. A routine prompt names the task; this skill holds the policy.

## Scope policy

Allowed: dependency updates (NGINX, Go, Kubernetes client modules, base images, GitHub Actions, pre-commit hooks), security fixes, critical bug fixes that break existing behavior, CI/release reliability, documentation accuracy, tests that raise coverage.

Never: new features, new annotations or options, behavior changes, refactors "for cleaner code", API contract changes, anything that lowers test coverage, deleting or weakening a test to get green.

## Hard rules

- Never push to `main`, never create tags or GitHub releases, never close an issue. Merging is done by Renovate automerge, by a human, or by the `merge-gate` routine under the `merge-gate` skill, and by nothing else.
- Never push a commit to a branch Renovate owns (`renovate/*`, `pre-commit-*`): a commit by anyone but Renovate makes it show "Edited/Blocked" and stop rebasing the pull request (forkline/ingress-nginx#642). Work on your own branch `fix/renovate-<N>-<slug>` and link it.
- Never force-push, except your own bot-created release branch (`release-readiness` skill).
- Commit messages and PR titles follow `.commitlintrc.json`: `type(scope): subject`; types build, chore, ci, docs, feat, fix, merge, perf, refactor, release, revert, style, test; scopes admission, api, chart, ci, commitlint, config, controller, coverage, deps, docs, e2e, helm, image, lint, lua, nginx, pre-commit, release, security, template, test, tests, webhook. PR titles feed the release notes: make them descriptive. No @-mentions in commit or PR text.
- Security fixes are titled `fix(security): <subject> (<CVE or GO id>)`; the `release-prep` routine matches that title when the pull request merges.
- Bot reviews are plain comments (`git-api comment create`), not review threads: main requires every conversation resolved, so an open thread blocks merging. Use `git-api pr review --request-changes` only for a policy violation a human must see.
- Build and test the smallest thing that answers the question (`go test ./path/...`, `make golint-check`, `pre-commit run --files <files>`); the full e2e suite and the image build are CI's job.

## Risk taxonomy

| Risk | Examples | Consequence |
|------|----------|-------------|
| Low | GitHub Actions bump, pre-commit hook bump, docs, mkdocs, helm CLI in CI, Go module patch bumps with green CI, regeneration of generated files | May be merged by `merge-gate` once its criteria say so |
| Medium | Go toolchain patch (GOLANG_VERSION + go.mod), Go/K8s client minor bumps, alpine tag, chart changes, CI workflow logic, a security fix by version bump | Independent review, then merge or escalate per `merge-gate` |
| High | NGINX version bump or patch rework, anything in `internal/ingress/controller`, `internal/nginx`, `rootfs/etc/nginx/template`, annotations, admission, Lua, backports of SUSE or upstream CVE patches, major version bumps, release PRs with anything other than security, dependency and NGINX-patch content | Human decision; escalate |

A CVE raises urgency, not trust: a critical CVE fix still gets the review its risk class asks for.

## Lessons learned

- Bot commits on Renovate branches stop Renovate rebasing them (#642). `rebase` label on a Renovate pull request makes Renovate rebase it and discard foreign commits.
- Duplicate fix PRs for one govulncheck finding (#645, #646): always search open PRs, branches and issues for the vulnerability ID before opening one. The dedupe key for a Renovate pull request N is the branch prefix `fix/renovate-N-`; for a vulnerability it is `fix/*<id lowercase>*`.
- NGINX bumps routinely stop at failing patches (#209, #346, #630). Patches in `images/nginx/rootfs/patches/` are per NGINX version; a failing `patch -p1` means regenerate the hunks against the new tarball from https://nginx.org/download/. Reference commits: `7803f4c7e` (consolidated patch for 1.31.4), `1b0241c9b` (NDK and set-misc patches for 1.31.3), `362c07dc1` (dropping a patch NGINX now includes). `NGINX_BASE` must equal `ghcr.io/forkline/ingress-nginx/nginx:$NGINX_VERSION` from `images/nginx/rootfs/build.sh`. The e2e workflow caches the nginx base image by `NGINX_VERSION`, so a patch-only change needs a forced rebuild (commit `09aeb3e4a`).
- Helm CI is flaky (about a third of runs fail). Rerun failed jobs once (`git-api run rerun <run id>`) before treating a failure as real.
- Generated files are never hand-edited: CHANGELOG.md (git-cliff), chart README (helm-docs), deploy manifests (`hack/generate-deploy-scripts.sh`), TAG files and Chart.yaml (`make update-version`), all through `.ci/release.sh`.
- Pre-commit rejects `.orig`/`.rej` leftovers and trailing whitespace in patches.

## Tools and their limits

Use `git-api` for forge operations: `pr list|view|create|ready|review|merge|close`, `issue list|search|view|create|edit`, `comment list|create`, `label list|add|remove`, `checks list|wait`, `status get`, `logs fetch <run id>`, `run rerun`, `actions list`, `release list`. It has no command for assignees, pinning, creating labels, editing PR descriptions or security alerts. For those use `gh` or the GitHub REST API if a token is available in the runner (the Forkline GitHub App can merge, rerun workflows, read Dependabot and code-scanning alerts, write issues and pull requests); when neither works, say so in the run's final message instead of failing. Only apply labels that exist (`git-api label list`); if a needed label is missing, continue without it and mention that in the final message.

## Idempotency

Every comment a routine posts starts with an HTML marker (for example `<!-- forkline-ci-status:<sha> -->`). Before posting, read the thread (`git-api comment list --issue N`); do not post a second comment with the same marker. A routine that finds its work already done stops silently.

## Escalation protocol

When you are not confident, or the item needs a decision (scope, CVE backport, risky merge, a fix you could not make):

1. Check the labels first. If the item already has `needs-maintainer`, it was escalated: do not escalate again; add a comment only if you have genuinely new information.
2. Add the label `needs-maintainer` (`git-api label add --issue N --labels needs-maintainer`) and assign `pando85` (`gh issue edit N --add-assignee pando85`, or the REST API; if neither is possible, mention him once in the comment below as the only permitted @-mention).
3. Post ONE comment (marker `<!-- forkline-escalation -->`) with: a summary in two sentences; what was tried; the options; your recommendation; and exactly what reply or decision is needed from the maintainer.
4. Never close the item. Never decide a scope question or a CVE backport yourself.

When the maintainer replies on an escalated item, the mention flow (`issue-response`, `pr-review`) continues from his instruction: his answer settles the question; carry it out within this policy and remove nothing but your own label if the work is done and he said so.

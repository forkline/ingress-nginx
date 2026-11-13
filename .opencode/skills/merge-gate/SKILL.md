---
name: merge-gate
description: Low-risk policy and procedure for the merge-gate routine, which independently reviews green pull requests authored by the Forkline bot and either merges them or escalates. Load before reviewing a pull request for merging.
---

## Mode

`Mode: observe`

- `observe`: never merge and never label. Post one comment per pull request (marker `<!-- forkline-merge-gate -->`) saying `would merge` or `would escalate` with the evidence below. This is the evidence for tuning the criteria.
- `enforce`: merge or escalate for real. A maintainer switches the mode by editing this line in a pull request. Do not switch it yourself.

## Procedure

1. Load the `maintenance` skill (risk taxonomy, hard rules, escalation protocol).
2. Identify the pull request from the CI event (`git-api pr view N`). Stop silently unless it is open, not a draft, not from a fork, authored by `forkline[bot]`, and targeting `main`. Renovate pull requests merge themselves through Renovate automerge; human pull requests are never touched.
3. Be an independent reviewer: judge the diff (`git diff origin/main...HEAD`) and the CI evidence yourself. Do not trust the author's description. If the same bot already reviewed it (marker `<!-- forkline-auto-review -->`), read that review as input, not as a verdict.
4. Readiness: `git-api pr ready N` must report `state: open`, `ci_passed: true`, `conflicts: false`; `git-api checks list <head>` has no failed or pending required check (E2E and Helm included when they ran). A rerun-flaky history on the head is not green.
5. Classify the risk with the taxonomy of the `maintenance` skill and apply the criteria below.
6. Release pull requests (`release: prepare <version>`): only under the extra rule below.
7. Outcome, in `enforce` mode:
   - Low risk and every criterion met: approve (`git-api pr review N --approve`; the bot may not approve its own pull request, in which case skip the approval and continue) and squash-merge with the title `<subject> (#N)` (`git-api pr merge N --method squash`). Post one comment (marker) with the criteria evidence.
   - Anything else: escalate with the protocol of the `maintenance` skill (label `needs-maintainer`, assign `pando85`, one comment saying why it is not low risk, the options and the recommendation).
8. Never merge a pull request whose head changed after the CI run you looked at, never merge while a `needs-maintainer` label is on it, never merge from an unclear state.

## Low-risk criteria

TODO (owner): the precise criteria come from the history analysis of merged pull requests. Until they are filled in, use ONLY these provisional criteria and treat everything else as "not low risk":

- The diff touches only: `.github/workflows/*` action version pins, `.pre-commit-config.yaml` hook versions, `docs/` and `mkdocs.yml`, `images/*/TAG` and other generated version files when produced by `.ci/release.sh`.
- No file under `internal/`, `cmd/`, `pkg/`, `rootfs/`, `charts/*/templates`, `images/nginx/rootfs/` (patches, build.sh), `test/` is modified.
- No test is deleted or changed.
- The title follows `.commitlintrc.json`.
- All checks green on the current head.

## Release pull requests

TODO (owner): confirm. Provisional rule: the gate may merge a `release: prepare` pull request only if (a) the `release-readiness` checklist passes (re-run it here, do not trust the pull request body), and (b) the commits it releases since the previous tag are exclusively security fixes, dependency updates and NGINX patch or version changes. Otherwise escalate.

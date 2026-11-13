# Forkline routines for forkline/ingress-nginx

Routine specs that maintain this repository with little human intervention. They are published to the Forkline platform with `forkline --profile default routine push` (production profile); nothing in this directory runs from GitHub. Policy lives in repo skills the runner loads (`.opencode/skills/`), and the prompts stay thin and name the skill:

| Skill | Content |
|-------|---------|
| `maintenance` | scope policy, hard rules, risk taxonomy, lessons learned (#642, #645/#646, NGINX patches), tool limits, **escalation protocol** shared by every routine |
| `issue-response` | reply categories, tone, citations |
| `merge-gate` | procedure and low-risk criteria of the merge gate (mode line `observe` or `enforce`, criteria marked TODO) |
| `release-readiness` | "almost sure" checklist for opening or merging a release pull request |
| `release` (existing) | how `.ci/release.sh` and the auto-tag workflow release |

## Routines

| File | Trigger | What it does | Human touchpoint |
|------|---------|--------------|------------------|
| `ci-auto-fix.yaml` | managed `ci-auto-fix` (`ci.failed`) | Fixes red CI. On Renovate branches it never pushes (Renovate stops rebasing edited PRs, #642): it diagnoses, reruns flakes, opens a `fix/renovate-N-*` PR. Knows how NGINX patch rework was done. Posts its own status comment. | Escalation comment when it cannot fix |
| `merge-conflict-fix.yaml` | managed `merge-conflict-fix` (`pr.updated`) | Resolves conflicts; on Renovate branches only adds the `rebase` label; release PRs are regenerated, not hand-merged. | Escalation when both intents cannot survive |
| `auto-pr-review.yaml` | `pr.opened`, native | Maintainability review of PRs authored by `forkline[bot]` (not drafts, forks, label `no-auto-review`), commits small certain fixes. Replaces the built-in. | none |
| `release-prep.yaml` | `pr.merged` to `main` (security or NGINX title/label) | Runs the release readiness checklist; only if it passes, opens `release: prepare vX`. Checks that no release PR is open. | Merges or approves the release PR (or `merge-gate` under its rules); escalation when the checklist fails |
| `merge-gate.yaml` | `ci.succeeded` | Independent review (a different model, `openai/gpt-5.5`) of green `forkline[bot]` PRs against the low-risk criteria; merges or escalates. Starts in `observe` mode. | Fills the criteria (TODO in the skill), flips the mode to `enforce`; handles escalations |
| `cve-watch.yaml` | daily 05:07 UTC | GitHub security alerts, govulncheck, nginx.org advisories, Alpine secdb, Kubernetes and SUSE ingress-nginx advisories. Critical or high and fixable: PR titled `fix(security): ...`; else an issue. Deduplicates. | Decides backports (it escalates them), merges the fix PR |
| `govulncheck-main.yaml` | `ci.failed` on main, workflow `Security` | One Go toolchain or module bump PR, deduplicated against open PRs, branches and issues. | Merges |
| `renovate-unblock.yaml` | daily 06:30 UTC | Renovate PRs open more than 48 h that are failing, edited or blocked: `rebase` label, rerun, separate fix branch or one diagnosis comment. | Decides ignore rules for major bumps |
| `issue-triage.yaml` | `issue.opened` | Labels; answers easy issues itself (README, MAINTENANCE.md, docs, code, duplicates, out-of-scope policy); escalates the rest. Never closes. | Answers escalations |
| `maintenance-digest.yaml` | weekly Monday 07:00 UTC | Updates the single "Maintenance digest" issue: needs-maintainer items with age, dependencies, bot work, security and releases, CI health, stuck items. | Pins the issue once, reads it |
| `disabled/issue-response.yaml`, `disabled/pr-review.yaml` | managed mention handlers | `@forkline` mention answers with the maintenance policy. **Do not publish yet**: mentions have no per-repository off switch, so the built-in and the routine would both answer. Kept as specs; `push` and `validate` of this directory skip `disabled/` (not recursive). | n/a |

Models: `isidoro/advanced` for NGINX, CVE, release and Renovate work, `isidoro/balanced` for reviews, conflicts, triage and the digest, `openai/gpt-5.5` for the merge gate so the reviewer is not the author's model. The validator only checks the length of the model string, so confirm that the account can use `openai/gpt-5.5` before enabling the gate (a failing gate fails safe: nothing is merged). Runner `advanced` for everything that builds or tests code, `basic` for triage and the digest.

## Escalation protocol (summary; the `maintenance` skill is normative)

Not confident: label `needs-maintainer`, assign `pando85`, ONE comment (summary, what was tried, options, recommendation, decision needed). Never close an item, never escalate twice (the label is checked first). A maintainer reply on an escalated item continues through the mention flow.

## Labels the routines expect (not created by them)

Create once on GitHub; routines apply only labels that exist and mention it in their final message otherwise.

`needs-maintainer`, `security`, `no-auto-review`, `rebase` (Renovate's rebase label), `maintenance-digest`. Existing ones are used as is: `bug`, `documentation`, `question`, `enhancement`, `duplicate`, `wontfix`, `area/dependency`.

## Rollout

1. Merge this PR, create the labels, pin nothing yet.
2. Turn the built-ins off for this repository only (`forkline repo update ... --ci-auto-fix false`, `--auto-pr-review false`, merge-conflict setting), then `routine push` and publish one routine at a time, a week of observation each, in this order: `auto-pr-review`, `ci-auto-fix`, `merge-conflict-fix`, `govulncheck-main`, `renovate-unblock`, `issue-triage`, `cve-watch`, `release-prep`, `maintenance-digest`, `merge-gate` (observe, then enforce).
3. A managed routine does not retire its built-in: both run while both are enabled.

## Known gaps in what the spec can express

- `ci.succeeded` cannot be listed by `routine events` or used with `concurrency.group: pr`, and carries no PR author, title or labels: `merge-gate` starts a run for every green PR (including Renovate's and humans') and decides by reading the forge.
- Managed routines cannot filter by branch or author: the Renovate rule is in the prompt (`managed.ref`).
- `routine test` renders a managed routine with empty `managed.*` values.
- No per-routine rate limit for events without a PR (`maxRunsPerPr` does not apply): deduplication is in the prompts.
- Schedule runs have nowhere to post their result (F9): the prompts post through `git-api` and end with a final message; the digest cannot see Forkline's own run history.
- `git-api` cannot assign, pin, create labels, edit PR descriptions or read security alerts: prompts use `gh` or the REST API when a token allows it and say so when it does not.
- Merging a bot PR as the bot cannot be approved by the same bot account.

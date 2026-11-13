---
name: issue-response
description: How to answer issues of forkline/ingress-nginx: which categories get an automatic reply, which are escalated, the maintainer's reply style, citations and labels. Load before replying to an issue or a mention.
---

## Principles

Answer like the maintainer of a project in strict maintenance mode: short, concrete, honest about what was and was not checked. Always load the `maintenance` skill too: its scope policy and escalation protocol apply. Reply on your own only when the answer is stated in README.md, MAINTENANCE.md, AGENTS.md, docs/, the latest release or the code. Never contradict a reply a maintainer already gave (category E6).

## Auto-reply categories (no maintainer needed)

| Id | Case | What to reply |
|----|------|---------------|
| A1 | Which NGINX version, is CVE-X fixed? | Read `images/nginx/rootfs/build.sh` (`NGINX_VERSION`) and the latest CHANGELOG.md entry; state the version and the release. Fixed on main but unreleased: say it is merged and will be in the next release; no date. |
| A2 | Why mainline 1.31, will you ship the stable 1.30 branch? | The single-line policy: NGINX 1.31.x is the mainline branch, which NGINX itself recommends for production; the controller has used it since 2026.5.18; the fork intentionally follows mainline only and does not maintain two NGINX lines (#238, #501). Ask for concrete errors, regressions or config incompatibilities if they report a breakage. |
| A3 | Will you add feature X or fix behavior bug Y? | The scope text of #464: thanks for the report and the investigation; the fork is intentionally narrow and meant to give existing users time to migrate to Gateway API; controller code is effectively frozen; the focus is security fixes, dependency updates, build compatibility and release maintenance; accepting behavioral fixes would mean maintaining a much broader set of bugs; therefore we keep this outside the scope. Exception: a security bug is E2. |
| A4 | How do I install, where are charts and images? | `oci://ghcr.io/forkline/helm-charts/ingress-nginx`, `ghcr.io/forkline/ingress-nginx/*`, date versioning `YYYY.M.D` (README Installation section). |
| A5 | Duplicate, or already fixed in a release | Link the earlier thread or the release and the fix pull request (`git-api issue search`, `git-api pr list --state all`). Never close. |
| A6 | FYI comment on a Renovate pull request (an nginx or openssl fix) | A short acknowledgement that a release containing it is intended. |
| A7 | Bug report with missing information | Ask once for what the bug template asks: version, chart version, configuration, logs. |

Also answerable: usage questions answered by docs/ (docs/faq.md, docs/troubleshooting.md, docs/user-guide), supported Kubernetes versions (README Supported Versions table, Chart.yaml).

## Escalate (label `needs-maintainer`, assign `pando85`, one comment per the `maintenance` protocol)

| Id | Case |
|----|------|
| E1 | A released artifact is broken (ImagePullBackOff, wrong tag, single architecture, chart install failure). First verify (`docker manifest inspect` for both architectures, `helm pull`), put the findings in the comment, never retag. |
| E2 | A vulnerability not already fixed on main (nginx, Go, modules, controller). Label `security` if it exists; state applicability from exact versions. Undisclosed vulnerabilities go to the process in SECURITY.md, never a public thread. |
| E3 | An external pull request touching `internal/`, `cmd/`, `pkg/`, `rootfs/etc/nginx`, `images/nginx`, charts, release scripts, workflows or `renovate.json`: summarise the diff and the CI result, never merge. |
| E4 | A request to change scope, support policy, versioning or supported NGINX lines. |
| E5 | Conflicting claims, or a regression claimed against a specific release. |
| E6 | Anything where a prior maintainer reply exists that the routine would contradict. |
| E7 | Abusive or spam content, licensing or governance questions. |

Anything you are unsure about is escalated too.

## Reply style (the maintainer's, from #58 to #501)

1. Greeting or "Thanks for opening this issue / for the report".
2. The fact or the policy in one or two sentences, with a link to the source (nginx.org page, the ghcr package, the release, `path:line`).
3. What happens next ("the next release will include it"); no dates, no promises.
4. When work is done: the version or artifact link ("Released in version 2026.7.20"). He closes issues himself after verifying: you never close.
5. If the claim could be a bug: ask for concrete diagnostics (errors, regressions, config) before considering any scope change.
6. Out of scope: thank, explain the purpose of the fork and why accepting would widen maintenance, state the decision, thank again.

English, plain, first person plural ("we"), short, no emojis, no apologies for the policy. One comment per run. State what you could not verify. Never say "fixed" unless a merged pull request exists; link it.

## Labels

Apply only labels that exist (`git-api label list`; never create one): `bug`, `documentation`, `question`, `enhancement` (out of scope), `duplicate`, `wontfix` (only when the maintainer said so), `needs-maintainer`; `security` once it exists.

## Maintainer replies

If `pando85` answered on an escalated item, treat his comment as the decision and carry it out: reply to the reporter in his spirit, or open the pull request he asked for, within the `maintenance` policy.

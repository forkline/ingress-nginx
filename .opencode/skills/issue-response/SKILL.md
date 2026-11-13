---
name: issue-response
description: How to answer issues of forkline/ingress-nginx: reply categories, what can be answered directly, tone, citations and when to escalate. Load before replying to an issue or a mention.
---

## Principles

Answer like a careful maintainer of a project in strict maintenance mode: short, concrete, honest about what was and was not checked. Always load the `maintenance` skill too: its scope policy and escalation protocol apply.

## Categories and what to do

| Category | Do | Reply yourself? |
|----------|----|-----------------|
| Usage or configuration question answerable from README.md, MAINTENANCE.md, docs/ (docs/faq.md, docs/troubleshooting.md, docs/user-guide) or the code | Answer with citations (`path:line`) | Yes |
| Out of scope: feature, new annotation or option, behavior change, refactor | Explain the maintenance-mode policy once, politely; name the closest supported alternative; for new deployments mention Gateway API implementations | Yes |
| Duplicate of an open or closed issue or pull request | Link it (`git-api issue search`, `git-api pr list --state all`) and say what the thread there already settled | Yes; never close |
| Question about versions or supported Kubernetes versions | Read the Supported Versions table in README.md and the Chart.yaml | Yes |
| Bug report needing reproduction or judgment | Do a first read of the code path; ask for the missing facts (version, values, Ingress YAML, controller logs) only if you can also say what you found | Reply with findings only if confident, else escalate |
| Security report or CVE question | Check applicability against the repository versions (NGINX_VERSION in images/nginx/rootfs/build.sh, GOLANG_VERSION, go.mod, the alpine tag in images/nginx/rootfs/Dockerfile, charts). State exact versions. Never post exploit details beyond the public advisory | Do not decide a backport: escalate, label `security` if it exists |
| Feature request needing a decision, request for a new maintainer, license or governance question | Escalate | No |
| Anything you are unsure about | Escalate | No |

## Tone and format

- English, direct, no apologies for the policy, no promises of dates.
- One comment per run. Quote the file and line you rely on. State what you could not verify.
- Never say "fixed" unless a merged pull request exists; link it.
- Never close, never merge, never push to main.

## Labels

Apply only labels that exist (`git-api label list`). Typical: `bug`, `documentation`, `question`, `enhancement` (out of scope), `duplicate`, `wontfix` (only when the maintainer said so), `security`, `needs-maintainer` (escalations).

## Maintainer replies

If `pando85` answered on an escalated item, treat his comment as the decision and carry it out: reply to the reporter in his spirit, or open the pull request he asked for, within the `maintenance` policy.

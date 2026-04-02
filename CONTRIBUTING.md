# Contributing Guidelines

Read the following guide if you're interested in contributing to Ingress. [Make Ingress-Nginx Work for you, and the Community](https://youtu.be/GDm-7BlmPPg) from KubeCon Europe 2018 is a great video to get you started!!

Note that this guide refers to contributing to actual sources of the repository. If you interested in contributing through issue triaging, have a look at [this guide](./ISSUE_TRIAGE.md).

## Fork Maintenance

This repository is maintained as a fork and no longer follows the retired upstream
project's release and CI model.

Before contributing, read:

- `MAINTENANCE.md` for the active maintenance and CI migration plan
- `docs/developer-guide/getting-started.md` for local development details

The fork uses:

- `main` as the only long-lived branch
- simplified `vX.Y.Z` release tags
- GHCR for published artifacts
- Renovate for dependency updates
- pre-commit and commitlint checks modeled after `kaniop`

## Contributing a Patch

1. Read the [Ingress development guide](docs/developer-guide/getting-started.md).
1. Fork the desired repo, develop and test your code changes.
1. Submit a pull request.

Recommended local checks before opening a pull request:

1. `pre-commit run --all-files`
2. `make test`
3. `make verify-docs`

All changes must be code reviewed. Coding conventions and standards are explained in the official [developer docs](https://github.com/kubernetes/community/tree/master/contributors/devel). Expect reviewers to request that you avoid common [go style mistakes](https://github.com/golang/go/wiki/CodeReviewComments) in your PRs.

Note that the fork is maintained conservatively. Priority is given to security
fixes, dependency upkeep, build and release reliability, NGINX updates, and other
changes that reduce long-term maintenance cost.

### Merge Approval

Changes should land through pull requests into `main` after CI passes and a
maintainer reviews the change. The fork no longer depends on Prow-based `/lgtm`
or `/approve` automation.

## Support Channels

Whether you are a user or contributor, official support channels include:

- GitHub issues: https://github.com/forkline/ingress-nginx/issues/new
- Slack: kubernetes-users room in the [Kubernetes Slack](http://slack.kubernetes.io/)
- Post: [Kubernetes Forum](https://discuss.kubernetes.io)

Before opening a new issue or submitting a new pull request, it's helpful to search the project - it's likely that another user has already reported the issue you're facing, or it's a known issue that we're already aware of.

## New Contributor Tips
If you're a new contributor, you can follow the [New Contributor Tips guide](NEW_CONTRIBUTOR.md)

# Forkline Maintained Fork

> **⚠️ Maintenance Mode Notice**
>
> This fork is maintained in **strict maintenance mode**. We do NOT develop new features or change existing behavior. Our sole purpose is to keep this mature, well-tested software functional and secure through controlled dependency updates.
>
> **How we maintain this project:**
> - **Automated dependency updates**: Renovate handles NGINX, Go modules, Kubernetes clients, and security fixes
> - **Human-in-the-loop**: Every update is validated with comprehensive testing and coverage checks
> - **No regressions policy**: Test coverage must never decrease; all behavior preserved exactly
> - **Critical coverage tracking**: Issue #4 implements coverage baseline to prevent regressions
>
> **Date-based Versioning**: We use date-based versions (e.g., `v2026.4.3`) instead of semantic versioning. Since we don't add features or make breaking changes, semantic versioning is meaningless here. Date versions clearly show when the software was last maintained, helping you identify recent security updates and dependency patches.
>
> This project is maintained by Forkline using coding-agent orchestration. Updates are automatic but never unattended—we prioritize security, stability, and community needs.
>
> **For new deployments**: Consider [Gateway API](https://gateway-api.sigs.k8s.io/guides/) implementations instead. This fork serves existing ingress-nginx users who need continued maintenance.

---

> This fork exists because the upstream project was retired, while the software
> itself remains practical to maintain: the codebase is mature, behavior is well
> known, and most future work is concentrated in controlled update lanes such as
> NGINX, Go libraries, Kubernetes dependencies, and release automation.
>
> We maintain this project with our own coding-agent orchestration platform,
> Forkline, the software factory. Much of the maintenance work is designed to be
> automatic, but never unattended: we keep a human in the loop, listen to the
> community, and prioritize fast bug fixes because we depend on this project
> ourselves.

## Upstream Retirement Context

[What You Need to Know about Ingress NGINX Retirement](https://www.kubernetes.io/blog/2025/11/11/ingress-nginx-retirement/):

* Best-effort maintenance continued until March 2026.
* Upstream no longer provides releases, bugfixes, or security updates.
* Existing deployments and historical artifacts remain important for operators and forks like this one.

# Ingress NGINX Controller

[![CII Best Practices](https://bestpractices.coreinfrastructure.org/projects/5691/badge)](https://bestpractices.coreinfrastructure.org/projects/5691)
[![Go Report Card](https://goreportcard.com/badge/github.com/kubernetes/ingress-nginx)](https://goreportcard.com/report/github.com/kubernetes/ingress-nginx)
[![GitHub license](https://img.shields.io/github/license/kubernetes/ingress-nginx.svg)](https://github.com/kubernetes/ingress-nginx/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/kubernetes/ingress-nginx.svg)](https://github.com/kubernetes/ingress-nginx/stargazers)
[![codecov](https://codecov.io/gh/forkline/ingress-nginx/branch/main/graph/badge.svg?token=REPLACE_WITH_CODECOV_TOKEN)](https://codecov.io/gh/forkline/ingress-nginx)

## Overview

ingress-nginx was an Ingress controller for Kubernetes using [NGINX](https://www.nginx.org/) as a reverse proxy and load
balancer.

[Learn more about Ingress on the Kubernetes documentation site](https://kubernetes.io/docs/concepts/services-networking/ingress/).

## Usage warnings

If you are not already using ingress-nginx, you should not be deploying it as it is [not being developed](#retiring). Instead you should identify a [Gateway API](https://gateway-api.sigs.k8s.io/guides/) implementation and use it.

Do not use in multi-tenant Kubernetes production installations. This project assumes that users that can create Ingress objects are administrators of the cluster. See the [FAQ](https://kubernetes.github.io/ingress-nginx/faq/#faq) for more.

## Troubleshooting

If you encounter issues, review the [troubleshooting docs](docs/troubleshooting.md),
[search for an issue](https://github.com/kubernetes/ingress-nginx/issues), or talk to us on the
[#ingress-nginx-users channel](https://kubernetes.slack.com/messages/ingress-nginx-users) on the Kubernetes Slack server.

## Changelog

See [the list of releases](https://github.com/kubernetes/ingress-nginx/releases) for all changes.
For detailed changes for each release, please check the [changelog-$version.md](./changelog) file for the release version.
For detailed changes on the `ingress-nginx` helm chart, please check the changelog folder for a specific version.
[CHANGELOG-$current-version.md](./charts/ingress-nginx/changelog) file.

### Supported Versions table

All images use unified date-based versioning. The version indicates when the software was last maintained.

| Supported | Ingress-NGINX version | k8s supported version        | Alpine Version | NGINX Version | Helm Chart Version |
| :-------: | --------------------- | ---------------------------- | -------------- | ------------- | ------------------ |
|    ✅     | **v2026.5.2**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.5.2           |
|    🔄     | **v2026.4.3-1**       | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.4.3-1         |
|    🔄     | **v2026.4.3**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.3         | 1.27.1        | 2026.4.3           |

## Get Involved

Thanks for taking the time to join our community and start contributing!

- This project adheres to the [Kubernetes Community Code of Conduct](https://git.k8s.io/community/code-of-conduct.md).
  By participating in this project, you agree to abide by its terms.
- **Contributing**: Documentation contributions are welcome.

  - Read [`CONTRIBUTING.md`](CONTRIBUTING.md) for information about the workflow that we
    expect and instructions on the developer certificate of origin that we require.
  - Join our Kubernetes Slack channel for developer discussion : [#ingress-nginx-dev](https://kubernetes.slack.com/archives/C021E147ZA4).
  - Submit GitHub issues for documentation problems.
    - Please make sure to read the [Issue Reporting Checklist](https://github.com/kubernetes/ingress-nginx/blob/main/CONTRIBUTING.md#issue-reporting-guidelines) before opening an issue. Issues not conforming to the guidelines **may be closed immediately**.

- **Support**:

  - Join the [#ingress-nginx-users](https://kubernetes.slack.com/messages/CANQGM8BA/) channel inside the [Kubernetes Slack](http://slack.kubernetes.io/) to ask questions or get support from the maintainers and other users.
  - The [GitHub issues](https://github.com/kubernetes/ingress-nginx/issues) in the repository are **exclusively** for bug reports and feature requests.

## License

[Apache License 2.0](https://github.com/kubernetes/ingress-nginx/blob/main/LICENSE)

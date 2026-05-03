# Ingress NGINX Controller

> **⚠️ Maintenance Mode — Forkline Maintained Fork**
>
> This is a maintained fork of the retired Kubernetes ingress-nginx controller, kept in **strict maintenance mode**: no new features, no behavior changes — only dependency updates and security patches. We use date-based versioning (e.g., `v2026.5.3`) since semantic versioning has no meaning when there are no features or breaking changes.
>
> **For new deployments**, consider [Gateway API](https://gateway-api.sigs.k8s.io/guides/) implementations instead. This fork serves existing ingress-nginx users who need continued maintenance.
>
> Maintained by [Forkline](https://github.com/forkline) using coding-agent orchestration — automated but never unattended.

## Upstream Retirement Context

[What You Need to Know about Ingress NGINX Retirement](https://www.kubernetes.io/blog/2025/11/11/ingress-nginx-retirement/):

* Best-effort maintenance continued until March 2026.
* Upstream no longer provides releases, bugfixes, or security updates.
* Existing deployments and historical artifacts remain important for operators and forks like this one.

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
|    ✅     | **v2026.5.3**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.5.3           |
|    🔄     | **v2026.5.2-1**       | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.5.2-1         |
|    🔄     | **v2026.5.2**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.5.2           |
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

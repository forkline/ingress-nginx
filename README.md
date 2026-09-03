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

## Installation

The Helm chart is distributed as an OCI package via GitHub Container Registry:

```console
helm upgrade --install ingress-nginx oci://ghcr.io/forkline/helm-charts/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

To install a specific version (chart version matches the release tag, e.g. `2026.8.21`):

```console
helm upgrade --install ingress-nginx oci://ghcr.io/forkline/helm-charts/ingress-nginx \
  --version 2026.8.21 \
  --namespace ingress-nginx --create-namespace
```

Container images are published to `ghcr.io/forkline/ingress-nginx/`:

| Image | Location |
|-------|----------|
| Controller | `ghcr.io/forkline/ingress-nginx/controller:<version>` |
| Webhook certgen | `ghcr.io/forkline/ingress-nginx/kube-webhook-certgen:<version>` |

See the [full installation guide](docs/deploy/index.md) for more options.

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
|    ✅     | **v2026.9.3**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.24.1         | 1.31.5        | 2026.9.3           |
|    🔄     | **v2026.8.21**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.24.1         | 1.31.4        | 2026.8.21           |
|    🔄     | **v2026.8.16**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.24.1         | 1.31.3        | 2026.8.16           |
|    🔄     | **v2026.7.20**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.24.1         | 1.31.3        | 2026.7.20           |
|    🔄     | **v2026.6.28**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.24.1         | 1.31.2        | 2026.6.28           |
|    🔄     | **v2026.6.18**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.24.1         | 1.31.2        | 2026.6.18           |
|    🔄     | **v2026.5.24**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.31.1        | 2026.5.24           |
|    🔄     | **v2026.5.18**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.30.1        | 2026.5.18           |
|    🔄     | **v2026.5.14**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.30.1        | 2026.5.14           |
|    🔄     | **v2026.5.3**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.5.3           |
|    🔄     | **v2026.5.2-1**       | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.5.2-1         |
|    🔄     | **v2026.5.2**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.5.2           |
|    🔄     | **v2026.4.3-1**       | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.4         | 1.27.1        | 2026.4.3-1         |
|    🔄     | **v2026.4.3**         | 1.35, 1.34, 1.33, 1.32, 1.31 | 3.23.3         | 1.27.1        | 2026.4.3           |

## License

[Apache License 2.0](https://github.com/kubernetes/ingress-nginx/blob/main/LICENSE)

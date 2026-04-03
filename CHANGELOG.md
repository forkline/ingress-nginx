# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

**Note**: This project uses date-based versioning (e.g., `v2026.04.03`) instead of semantic versioning, since we maintain existing functionality without adding features or making breaking changes. Date versions clearly show when the software was last maintained for security.

## [v2026.04.03](https://github.com/forkline/ingress-nginx/tree/v2026.04.03)

### Added

- versioning: Switch to date-based versioning (v2026.04.03)
- ci: Add automated release workflow with git-cliff
- ci: Integrate kubectl plugin build into release workflow

### Chore

- renovate: Enable forkProcessing for fork support

## [unreleased]

### 404-server

- Graceful shutdown ([2f8e81e](https://github.com/forkline/ingress-nginx/commit/2f8e81e383d58b0b8d35b08c72b1e7d8982c6816))

### Added

- baremetal: Add kustomization.yaml ([7ddb734](https://github.com/forkline/ingress-nginx/commit/7ddb7343faa2286088409a2ea22bd47a7066e66a))
- ci: Add automated release workflow with git-cliff ([bd1691f](https://github.com/forkline/ingress-nginx/commit/bd1691fd2fc09caadd16e8390636b86c6032c836))
- ci: Integrate kubectl plugin build into release workflow ([1c4d1fe](https://github.com/forkline/ingress-nginx/commit/1c4d1fe3d505fa947fd4cdb8881088e5e72a21d8))
- collectors: Added services to collectorLabels and requests Countervec to capture the name of the kubernetes service used to serve the client request. ([c38c66e](https://github.com/forkline/ingress-nginx/commit/c38c66e00ad309dda1d2fa3c29e59a4d4123e344))
- configmap: Expose gzip-disable ([e6dcd68](https://github.com/forkline/ingress-nginx/commit/e6dcd6845e3dc2e314a676cad7107fe75225def0))
- default_backend: TopologySpreadConstraints on default backend ([e9509e2](https://github.com/forkline/ingress-nginx/commit/e9509e27aa9a6660a8bc70e96c988ce4be4aac11))
- geoip2_autoreload: Enable GeoIP2 auto_reload config  ([3c4e78e](https://github.com/forkline/ingress-nginx/commit/3c4e78e6b755eb33821c4477106f424bd8792d8f))
- helm: Optionally use cert-manager instead admission patch ([d7674e4](https://github.com/forkline/ingress-nginx/commit/d7674e43230274a485ff90054383083ea4280ef8))
- helm: Add loadBalancerClass ([0b4c98b](https://github.com/forkline/ingress-nginx/commit/0b4c98b7c31f95e00dc93b7a346bfe3a6526af51))
- helm: Add documentation about metric args ([e805d49](https://github.com/forkline/ingress-nginx/commit/e805d4955d8cf27e1717b8d7d162b341a15b932e))
- leader_election: Flag to disable leader election feature on controller ([9b63559](https://github.com/forkline/ingress-nginx/commit/9b63559cbb492c4ace8641b81c6adc6fc54aa9ce))
- metrics: Add path and method labels to requests counter ([fbdfc65](https://github.com/forkline/ingress-nginx/commit/fbdfc6505b8fd9eea937ae10af641cbc1dfdc5cd))
- template: Wrap IPv6 addresses in [] ([54f6729](https://github.com/forkline/ingress-nginx/commit/54f6729dc83f38c6bd58fb507be697794ddff1b9))
- Feat/proxytimeout support proxy timeout for stream type ([13ab894](https://github.com/forkline/ingress-nginx/commit/13ab894e6fc6be694960e63f99e187630778b4fe))

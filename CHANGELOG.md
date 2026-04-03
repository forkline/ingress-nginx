# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [v2026.4.3-1](https://github.com/forkline/ingress-nginx/tree/v2026.4.3-1) - 2026-04-03

## [v2026.4.3](https://github.com/forkline/ingress-nginx/tree/v2026.4.3) - 2026-04-03

### Added

- **Date-based versioning** — Switched from semantic versioning (`v1.15.x`) to date-based versioning (`vYYYY.M.D` or `vYYYY.M.D-N` for multiple daily releases). Since this project is in maintenance mode (no new features, no breaking changes), date versions more honestly communicate when the software was last maintained for security updates.
- **Release automation** — Added `.ci/release.sh` script and Makefile targets (`make update-version`, `make update-changelog`) to streamline the release process.
- **GPG-signed auto-tagging** — Added `.github/workflows/auto-tag.yml` that automatically creates a signed git tag when a release commit is pushed to main.

### Fixed

- **Release workflow compatibility** — Updated `docker_images.yml` to work with date-based versions: replaced semver-only `changelog-reader-action` with manual CHANGELOG parsing, added `--skip=validate` to GoReleaser for non-semver tags, and fixed changelog body passing to avoid argument length limits.

### Changed

- **Helm chart** — Updated `charts/ingress-nginx/Chart.yaml` with new versioning scheme.

### Documentation

- Added `AGENTS.md` with maintenance mode guidelines and AI coding instructions.
- Updated `README.md` with date-based versioning rationale.

---

### Prior releases (v1.15.x)

Releases prior to `v2026.4.3` used semantic versioning and are listed below for historical reference.

<details>
<summary>Click to expand historical changelog</summary>

## [v1.15.7](https://github.com/forkline/ingress-nginx/tree/v1.15.7) - 2026-04-02

### Fixed

- Fix(release): point krew plugin to forkline artifacts ([71cb52](https://github.com/forkline/ingress-nginx/commit/71cb526e611748fb1540ce9921b93fca7e40d6ed))

## [v1.15.6](https://github.com/forkline/ingress-nginx/tree/v1.15.6) - 2026-04-02

### Added

- Feat(ci): integrate kubectl plugin build into release workflow ([1c4d1f](https://github.com/forkline/ingress-nginx/commit/1c4d1fe3d505fa947fd4cdb8881088e5e72a21d8))

## [v1.15.5](https://github.com/forkline/ingress-nginx/tree/v1.15.5) - 2026-04-02

### Added

- Feat(ci): add automated release workflow with git-cliff ([bd1691](https://github.com/forkline/ingress-nginx/commit/bd1691fd2fc09caadd16e8390636b86c6032c836))

## [v1.15.4](https://github.com/forkline/ingress-nginx/tree/v1.15.4) - 2026-04-02

### Chore

- Chore(renovate): enable forkProcessing for fork support ([1c100e](https://github.com/forkline/ingress-nginx/commit/1c100e6ec))

## [v1.15.3](https://github.com/forkline/ingress-nginx/tree/v1.15.3) - 2026-04-02

### Fixed

- Fix(ci): run image publishing on tags only ([d53299](https://github.com/forkline/ingress-nginx/commit/d532994200475836bedf6e0da447fd92e768dcee))
- Fix(ci): simplify image publishing to amd64 ([d9765f](https://github.com/forkline/ingress-nginx/commit/d9765fe7a759442863ec55f94deaab4fc37489da))
- Fix(ci): limit image publishing to image changes ([ff620a](https://github.com/forkline/ingress-nginx/commit/ff620aa39e6d3a01201f5511b4d07979f786c798))
- Fix(ci): fetch history for diff-based linting ([1a228b](https://github.com/forkline/ingress-nginx/commit/1a228b7714a1f6612ad5c0d64b8b1f4e9ef5f0a9))
- Fix(ci): scope golangci-lint to new issues ([3aa95c](https://github.com/forkline/ingress-nginx/commit/3aa95ce7c534093bb48142131e57a40a03c0a51d))
- Fix(ci): install golangci-lint with current toolchain ([47d9d0](https://github.com/forkline/ingress-nginx/commit/47d9d0e1174155226cde17f0ded11e1d4f23dfba))
- Fix(ci): run checks in project-compatible environments ([072e04](https://github.com/forkline/ingress-nginx/commit/072e04d730bf4f1603c4d9e4b4fe03692ea8975d))
- Fix(ci): stabilize go workflow on GitHub ([e07a65](https://github.com/forkline/ingress-nginx/commit/e07a6588222610abf9c166a649522f7c2ff27c18))
- Fix(ci): allow helm unittest plugin install ([b378ee](https://github.com/forkline/ingress-nginx/commit/b378eedf4e292b068e974286f0a23e8cef8af0b1))
- Fix(chart): mismatch between values.yml and README.md ([0f8234](https://github.com/forkline/ingress-nginx/commit/0f82342aa65f7339d57972452fbba5e562ff0100))

### CI

- Replace upstream automation with fork-owned workflows ([e24441](https://github.com/forkline/ingress-nginx/commit/e24441b70))

## [v1.15.2](https://github.com/forkline/ingress-nginx/tree/v1.15.2) - 2026-04-01

_Initial fork release from upstream kubernetes/ingress-nginx._

</details>

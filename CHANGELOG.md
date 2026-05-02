# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [v2026.5.2-1](https://github.com/forkline/ingress-nginx/tree/v2026.5.2-1) - 2026-05-02

### Fixed

- Remove v prefix from image tags in Helm chart([5983520](https://github.com/forkline/ingress-nginx/commit/59835202ddcf8b5b56f0cce79d767b60b5071638))


## [v2026.5.2](https://github.com/forkline/ingress-nginx/tree/v2026.5.2) - 2026-05-02

### Added

- Unify all images to date-based versioning v2026.5.2 (versioning)([be70d14](https://github.com/forkline/ingress-nginx/commit/be70d143c89a4e568f866dae13b1bedacf04bc27))


### Fixed

- Update golangci-lint to v2.11.4 for Go 1.26.1 compatibility([05d2e7b](https://github.com/forkline/ingress-nginx/commit/05d2e7bcf58ae9ac950f3bb9b31dd91d8bd88a38))

- Remove nolint comments and fix root lint issues([5518f67](https://github.com/forkline/ingress-nginx/commit/5518f678713ed87cbcdc2efc72fe478f1ea0b20a))

- Use bracket notation for dhparam.pem key in secret test (helm-test)([f0d23ec](https://github.com/forkline/ingress-nginx/commit/f0d23ec3bc49e314f644a9e05d044d612fdea99c))

- Update code for k8s.io/client-go v0.36.0 compatibility([5e5220f](https://github.com/forkline/ingress-nginx/commit/5e5220f5523519979614014c3b9ab19bf1bf076a))

- Use format string directly in Eventf call([470cf4b](https://github.com/forkline/ingress-nginx/commit/470cf4b876b1a8dc8bddb632e9a4b52cde4d14a6))

- Rename unused parameter rv to _ in Bookmark method([09d74d1](https://github.com/forkline/ingress-nginx/commit/09d74d13341b301f2ecfc31981257e5de671a7c0))

- Add pre-commit scope and use valid conventional commit prefix (ci)([b6c5639](https://github.com/forkline/ingress-nginx/commit/b6c5639d603a83e0ffd1a5287cb9222eb914fa17))

- Skip commitlint for Renovate pre-commit update commits (ci)([32d2c29](https://github.com/forkline/ingress-nginx/commit/32d2c2903a2307a4f7872c04398169c715c17151))

- Update go-github imports to v85 and remove duplicate in go.mod([f870ba8](https://github.com/forkline/ingress-nginx/commit/f870ba8e1973d24ae6f3054508a04d0a103b3755))

- Build nginx base image locally and regenerate helm-docs (ci)([a705666](https://github.com/forkline/ingress-nginx/commit/a705666252edc0a835a8458e1dda18f8b9ce6abc))

- Install helm-docs in pre-commit workflow (ci)([bdfacc3](https://github.com/forkline/ingress-nginx/commit/bdfacc33f3212840ac21cb056ed8cecc20a8216e))

- Update expected registry to ghcr.io/forkline in helm tests (test)([1c800f2](https://github.com/forkline/ingress-nginx/commit/1c800f24102cc703189f1351e878c30b51d64c2c))

- Build and load kube-webhook-certgen image for e2e tests (ci)([d74fc65](https://github.com/forkline/ingress-nginx/commit/d74fc65e63b3185ad2fc6501ecd1c5b477fe4f91))

- Correct path to kube-webhook-certgen TAG in e2e scripts (ci)([40e6509](https://github.com/forkline/ingress-nginx/commit/40e65092e70e6bd00aa08a8c0741ae186ba789e9))

- Update WaitForEndpoints to use EndpointSlice API (e2e)([b59f19f](https://github.com/forkline/ingress-nginx/commit/b59f19f6d9476561de26e41a04cc965da68cec8c))

- Update NamespaceContent to use endpointslices API (e2e)([9c5e0c3](https://github.com/forkline/ingress-nginx/commit/9c5e0c376e918c1b7befb2cab0dab524f88a5203))

- Resolve golangci-lint errors in e2e framework([4333140](https://github.com/forkline/ingress-nginx/commit/4333140fb20d80a964e59fd543755aa71fad8f14))

- Add perf to allowed commit types (commitlint)([ca4914d](https://github.com/forkline/ingress-nginx/commit/ca4914d3562486097483a15ab7a406a414f0ab72))

- Add commitlint and pre-commit to allowed scopes (commitlint)([d6ac263](https://github.com/forkline/ingress-nginx/commit/d6ac263e8368ef2b1fe21a797395ce7bb7ba12c9))

- Include NGINX base image in docker save and kind load (e2e)([ba68525](https://github.com/forkline/ingress-nginx/commit/ba68525400fd56fc9f3a95f42a0ef84430b20a3c))

- Fix yamllint line-length errors in workflow files (ci)([b1fa57c](https://github.com/forkline/ingress-nginx/commit/b1fa57c44a454e025a9fe1bdccc6181872d497fc))

- Fix multiline variable assignment in Helm workflow (ci)([62011a7](https://github.com/forkline/ingress-nginx/commit/62011a701a02fd4d35e6ee09d7f143794e8da8b4))

- Fix yamllint line-length in helm curl command (ci)([e1c166b](https://github.com/forkline/ingress-nginx/commit/e1c166b04570858aca56e3cb2e2f70d8b713eec7))

- Split SHA256 line to comply with yamllint 80-char limit (ci)([1775c37](https://github.com/forkline/ingress-nginx/commit/1775c378af465783bff83bcbb7d4f6a4897dd7f3))

- Use env directive for AH_SHA256 to fix shell and yamllint issues (ci)([ec630c7](https://github.com/forkline/ingress-nginx/commit/ec630c784664246d2521d629ee047ecd2bd7df9e))

- Move AH_SHA256 to run script to fix env directive parsing issue (ci)([84d7a7c](https://github.com/forkline/ingress-nginx/commit/84d7a7cf49bb7973999bb36c1cea52a6df0c3fe2))

- Use forkline e2e-test-runner image instead of upstream registry.k8s.io (ci)([7cfc09b](https://github.com/forkline/ingress-nginx/commit/7cfc09b9accee9ad7a6412ce22ed1e4b5490d769))

- Resolve CI failures in unify-versioning branch (ci)([46afbf5](https://github.com/forkline/ingress-nginx/commit/46afbf59ae55cc1c67f4d246638b44a7e91e4c83))

- Change test.sh shebang to /bin/sh for Alpine compatibility([285aa75](https://github.com/forkline/ingress-nginx/commit/285aa754ecc82b58e439a82ca02d1f71c36f7844))

- Build Go binaries directly in CI instead of run-in-docker.sh (e2e)([bb0b028](https://github.com/forkline/ingress-nginx/commit/bb0b028c14d32ed10e9863d001c4c42d2c240e11))

- Run unit tests directly instead of using Docker (ci)([6cf0f27](https://github.com/forkline/ingress-nginx/commit/6cf0f27f31c46338d7c25709a451d09716005302))

- Resolve yamllint line-length errors in e2e workflow (ci)([6abbe60](https://github.com/forkline/ingress-nginx/commit/6abbe60280f685e61e11c29af07deaa2a595fbc1))

- Output Go binaries to rootfs/bin/amd64 for Dockerfile (e2e)([18fab47](https://github.com/forkline/ingress-nginx/commit/18fab478d6babd06eb9f610a94e20e531dfc235c))

- Resolve CI test failures (tests)([9486785](https://github.com/forkline/ingress-nginx/commit/94867855463a99bfca45b384d1516fedf1a3b693))

- Add actions write permission for retry workflow (ci)([032425e](https://github.com/forkline/ingress-nginx/commit/032425e4d6919864a80aacac2de314a40fb7bc58))

- Resolve yamllint line-length error in e2e.yml (ci)([567fde2](https://github.com/forkline/ingress-nginx/commit/567fde22b01895110f3ea508cda71d5edaf9ab2f))

- Resolve lint errors in test files (test)([c89298e](https://github.com/forkline/ingress-nginx/commit/c89298e8c331e4eb9b88fcab432830b640a0eb7d))

- Resolve remaining lint issues (test)([394b920](https://github.com/forkline/ingress-nginx/commit/394b9202144f113c579aa1eaf5f5b036451127d3))

- Use temp directory for SSL tests and fix port conflicts (test)([3f7c72d](https://github.com/forkline/ingress-nginx/commit/3f7c72dd9dcb18c443e44e77f3f11e5e544d318f))

- Add constants for repeated string literals in flags tests (test)([f22272a](https://github.com/forkline/ingress-nginx/commit/f22272ae6a07ea05eb1bc36342ef16ddb0de6d3d))

- Resolve lint issues in flags_test.go and structure.go([57d13dc](https://github.com/forkline/ingress-nginx/commit/57d13dcd80b50678860849e9e2573b610056bfcc))

- Add ginkgo installation step to e2e workflow (ci)([efffc07](https://github.com/forkline/ingress-nginx/commit/efffc0717f426d007f1aadf12ec9255c1980e2e5))

- Update golangci-lint config for v2 compatibility (ci)([420139b](https://github.com/forkline/ingress-nginx/commit/420139bdc472cf5239bb46329b83cae3b6125c0a))

- Use temp directories for controller tests (test)([f5a9189](https://github.com/forkline/ingress-nginx/commit/f5a91896055c907c3377c17a4b59920e4cc385b1))

- Tag nginx base image with full registry path before docker save (e2e)([e0e717c](https://github.com/forkline/ingress-nginx/commit/e0e717c8f6c637df1b9881f47136605f3df0f3d1))

- Use temp directories for tests requiring /etc paths (test)([e1862ce](https://github.com/forkline/ingress-nginx/commit/e1862ce60af66cc3e7ba9120ee63b046636013e0))

- Fix test.sh PKG variable issue and update chart-testing version (ci)([47b5931](https://github.com/forkline/ingress-nginx/commit/47b5931f85696e453b5da0faf37eeedc6a281b67))

- Format maxmind.go to satisfy gofumpt (lint)([ef89374](https://github.com/forkline/ingress-nginx/commit/ef89374d7865f9575adc5a91d939508aeef54433))

- Allow capitalized commit types (commitlint)([e27db62](https://github.com/forkline/ingress-nginx/commit/e27db62c31fe44263571ea3fb662a50632ab1cca))

- Revert e2e-test-runner to upstream registry.k8s.io image (e2e)([04d2011](https://github.com/forkline/ingress-nginx/commit/04d201103d0960fbf0ffba084a8d1794eef1008e))

- Add timeouts to tests that hang in CI environments (test)([732e7d9](https://github.com/forkline/ingress-nginx/commit/732e7d984ba42312260749ba442eeba7bd64a736))

- Add lint to allowed scopes (commitlint)([9c930e9](https://github.com/forkline/ingress-nginx/commit/9c930e95db5127062706d05cf91eed82eb4aadae))

- Downgrade type/subject-empty rules to warnings (commitlint)([f8574a5](https://github.com/forkline/ingress-nginx/commit/f8574a5c276071021e8a2f8c0bfd56fe541edf2a))

- Include commitlint config in pre-commit cache key (ci)([309e7a9](https://github.com/forkline/ingress-nginx/commit/309e7a9c81a4e11556d1151ebc7c241c7181787a))

- Wait for cert-manager webhook certs before starting controller (helm)([9cfe009](https://github.com/forkline/ingress-nginx/commit/9cfe0098c0cf8788f2569e315daf96a35e9c560f))

- Use controller image for cert-manager wait initContainer (helm)([ece6d54](https://github.com/forkline/ingress-nginx/commit/ece6d54a92b7d6fa48b26c97e4512eacfb4dde26))

- Use POSIX-compliant [[:space:]] instead of \s in regex (ci)([5799506](https://github.com/forkline/ingress-nginx/commit/5799506b1ebe84a6ec566669697baac26372a014))

- Resolve yamllint line-length errors in pre-commit workflow([2c39a5f](https://github.com/forkline/ingress-nginx/commit/2c39a5f3be800b18ada5259c7865308ba988f1fc))

- Build static binaries with CGO_ENABLED=0 for Alpine compatibility (e2e)([7ec4c24](https://github.com/forkline/ingress-nginx/commit/7ec4c24528ecf2211fd92be1ad7bf2e7bbfae637))

- Build static e2e test binary with CGO_ENABLED=0 for Alpine compatibility (e2e)([62ee488](https://github.com/forkline/ingress-nginx/commit/62ee4881c2d34800f6d67d6c77145861a15e59b3))

- Resolve all pre-commit hook failures([282e743](https://github.com/forkline/ingress-nginx/commit/282e7433ba47dd36b05c279b2500d273fb7d9210))

- Exclude golden test data from trailing whitespace/EOF hooks([13f22c4](https://github.com/forkline/ingress-nginx/commit/13f22c48bc7ec66c3e39cb1b6f70297ba786103e))

- Exclude annotations-risk.md from end-of-file-fixer([78409cc](https://github.com/forkline/ingress-nginx/commit/78409cc08457f3b92ecaa99c123df1bff7376a7a))


### Chore

- Update pre-commit hook alessandrojcm/commitlint-pre-commit-hook to v9.25.0 (pre-commit)([27cabb8](https://github.com/forkline/ingress-nginx/commit/27cabb83a0f7d8fbb3eb65ff4649188e4bdb327b))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.141.2 (pre-commit)([ccf72e0](https://github.com/forkline/ingress-nginx/commit/ccf72e0075c2cfcb57012420ce3475f2ebf79417))


### CI

- Run required Go checks on all PRs with path-filtered skip([225e785](https://github.com/forkline/ingress-nginx/commit/225e78566da5816a444b7c750d23e4f73331a299))

- Add proactive scopes to prevent future timing races (commitlint)([6652dc0](https://github.com/forkline/ingress-nginx/commit/6652dc0e29726e66552f0b60ea705b3c3fb4d9b4))

- Publish latest images on every push to main([80524be](https://github.com/forkline/ingress-nginx/commit/80524be2ef69c2dd6c63e49c35f6f8d95ba5d441))


### Build

- Update module github.com/google/go-github/v48 to v84 (#15) (deps)([adc490c](https://github.com/forkline/ingress-nginx/commit/adc490cd6c0dfa281ddcf505e68ba7708e6c6c35))

- Update dependency kustomize to v5 (#11) (deps)([8592391](https://github.com/forkline/ingress-nginx/commit/8592391ef0a942a9fb7da59271289cdfe5686b2f))

- Update squidfunk/mkdocs-material Docker tag to v9.7.6 (#9) (deps)([9144ccc](https://github.com/forkline/ingress-nginx/commit/9144ccc11accd5dae129f396fb94fbb1f20518c3))

- Update dependency mkdocs-material to v9.7.6 (#5) (deps)([13f8874](https://github.com/forkline/ingress-nginx/commit/13f88742d9da33efa0ea1020a14ec32feb9e5405))

- Update go modules (#8) (deps)([45bcdaa](https://github.com/forkline/ingress-nginx/commit/45bcdaa890dbaee4ff9e9bfd2d5dcc4cd93f113c))

- Group all update types together in renovate config (deps)([3da319e](https://github.com/forkline/ingress-nginx/commit/3da319ebb6853786ec045a06a7b4154d6d64dbe8))

- Update google.golang.org/grpc/examples digest to 1c132b9 (#27) (deps)([aa1f09e](https://github.com/forkline/ingress-nginx/commit/aa1f09ed7da32782f2622a19bc99bce0140796a6))

- Update module helm.sh/helm/v4 to v4.1.4 [SECURITY] (#31) (deps)([3b1950f](https://github.com/forkline/ingress-nginx/commit/3b1950f45f7801c361a5c4b66529a9a3c28a014d))

- Update alpine Docker tag to v3.23.4 (#36) (deps)([3795290](https://github.com/forkline/ingress-nginx/commit/379529001d9f3a04391bcdda44ecd22ef85146fd))

- Update go modules (deps)([df66def](https://github.com/forkline/ingress-nginx/commit/df66def0a68ab5b33136b4b70257d56d24ef4911))

- Update go modules (deps)([10d7af6](https://github.com/forkline/ingress-nginx/commit/10d7af660cb6a2c5cfc53471364a7638f4d16480))

- Update module github.com/google/go-github/v84 to v85 (deps)([097aff4](https://github.com/forkline/ingress-nginx/commit/097aff4b3753a4b796e59be00b08fd19e43f0965))


### Performance

- Reuse published images from registry in e2e build job (ci)([011dc3f](https://github.com/forkline/ingress-nginx/commit/011dc3f1ed9f99d22fe25953c856f12894d57958))

- Use latest tag for e2e image reuse and push it on release (ci)([5f8cb19](https://github.com/forkline/ingress-nginx/commit/5f8cb196d44740baede5145e6411bc927f4ce0ed))


### Testing

- Add more regression tests for critical paths([9c78df8](https://github.com/forkline/ingress-nginx/commit/9c78df85ecdf82f3d617dbbaa36f02454156bcb4))

- Add annotation Validate, Equal, and error type tests to increase coverage([422af24](https://github.com/forkline/ingress-nginx/commit/422af24f07042aaffd007c8fc1c60891806ab88d))

- Add parser, task queue, util, and net tests for coverage([2bd8a68](https://github.com/forkline/ingress-nginx/commit/2bd8a6821faca653f1552ecd69a9082edb88e6df))

- Add inspector, store, and collector coverage tests([49298ce](https://github.com/forkline/ingress-nginx/commit/49298cea6503f990a396a114b14ac1b9ab222e45))

- Add dummy collector tests for metric package([b4409f9](https://github.com/forkline/ingress-nginx/commit/b4409f925340ff936d9fca08ce53318fbba34490))

- Add FilterIngresses test for store package([98bb16a](https://github.com/forkline/ingress-nginx/commit/98bb16a4e25baaa09c92fde5f0c28c9152f3fb94))

- Add createOpentelemetryCfg unit test([334b8e2](https://github.com/forkline/ingress-nginx/commit/334b8e2038e28aeef99f904620dd10846f6a6d6d))


### Pre-commit

- Update pre-commit hook norwoodj/helm-docs to v1.14.2 (#17)([82e9b9f](https://github.com/forkline/ingress-nginx/commit/82e9b9fecb09480e240c1e0a04313f32a1cfc180))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.104.1 (#10)([dc957c9](https://github.com/forkline/ingress-nginx/commit/dc957c9a6863c7d9dd9a32e836d285f7b5dadb4a))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.104.2 (#19)([43a136c](https://github.com/forkline/ingress-nginx/commit/43a136cf68268b6de42eb43eb81082ca84ddf524))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.104.3 (#22)([c496c02](https://github.com/forkline/ingress-nginx/commit/c496c024a8b41ccf2ccdede070e4f1dc6c5f2985))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.104.4 (#23)([0430476](https://github.com/forkline/ingress-nginx/commit/04304762cb4bd7ad11bf91d327378f5ee1616c05))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.104.6 (#24)([6ef6286](https://github.com/forkline/ingress-nginx/commit/6ef6286ad09a7f2c75e673ef3cd8b827b60fa064))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.104.7 (#25)([3a94390](https://github.com/forkline/ingress-nginx/commit/3a9439078d0dbcfd1c93fcbd68c1691b4ed2a50c))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.104.8 (#26)([a0da084](https://github.com/forkline/ingress-nginx/commit/a0da08463db829bda9544a87161617375fe4b533))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.104.10 (#28)([b72ab56](https://github.com/forkline/ingress-nginx/commit/b72ab5659fa1bbb738bbb7c97f744b4e7febd640))

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.140.0([5badf94](https://github.com/forkline/ingress-nginx/commit/5badf9482b5d81a234d7fd7ff92ec0a6c04d6854))


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

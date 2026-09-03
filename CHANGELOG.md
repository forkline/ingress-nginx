# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [v2026.9.3](https://github.com/forkline/ingress-nginx/tree/v2026.9.3) - 2026-09-03

### Chore

- Update pre-commit hook renovatebot/pre-commit-hooks (v44.39.1 → v44.59.3)

### Build

- Update google.golang.org/grpc/examples digest (7 updates)

- Update go modules (3 updates)

- Update module github.com/sirupsen/logrus to v1.10.2 (#361) (deps)([3eea43d](https://github.com/forkline/ingress-nginx/commit/3eea43dfa8c189d71c3e866957da4796cda5be2f))

- Update helm/kind-action action to v1.15.0 (#389) (deps)([1253e33](https://github.com/forkline/ingress-nginx/commit/1253e335ef9c8be4676c936d598e6316bce110e4))

- Update dependency nginx/nginx to v1.31.5 (#388) (deps)([a652212](https://github.com/forkline/ingress-nginx/commit/a652212694663c1095cd38c201ce8674d277d157))

## [v2026.8.21](https://github.com/forkline/ingress-nginx/tree/v2026.8.21) - 2026-08-21

### Fixed

- Update patches for NGINX 1.31.4 compatibility (nginx)([7803f4c](https://github.com/forkline/ingress-nginx/commit/7803f4c7e732c189192d6c3c9f33d8043af86751))

- Remove .orig artifacts from consolidated patch (nginx)([ab8fd0f](https://github.com/forkline/ingress-nginx/commit/ab8fd0f56e01dd79e6cd0a74d5234b10bf361419))

### Chore

- Update pre-commit hook renovatebot/pre-commit-hooks (v44.31.0 → v44.37.1)

- Upgrade nginx to 1.31.4 (nginx)([c203ac4](https://github.com/forkline/ingress-nginx/commit/c203ac432ab04eddad322dfac9b7d7e5c4c492ec))

### Build

- Update module github.com/stretchr/testify to v1.12.0 (#327) (deps)([8905228](https://github.com/forkline/ingress-nginx/commit/890522859b00652fd12b1c4a26566d7a92f36828))

- Update google.golang.org/grpc/examples digest (5 updates)

- Update module github.com/opencontainers/cgroups to v0.1.0 (#336) (deps)([5646dfb](https://github.com/forkline/ingress-nginx/commit/5646dfbf5b5572185334b1a3a3e2b82fba8a6816))

- Update go modules (3 updates)

- Update dependency nginx/nginx to v1.31.4 (#339) (deps)([6a4001f](https://github.com/forkline/ingress-nginx/commit/6a4001f3fa63a8ee29994ca909566a28894a74cc))

### Styling

- Fix trailing whitespace in consolidated patch([04d6188](https://github.com/forkline/ingress-nginx/commit/04d6188552336d4c1e97aaa894aa04afe508b0fe))

## [v2026.8.16](https://github.com/forkline/ingress-nginx/tree/v2026.8.16) - 2026-08-16

### Fixed

- Update go-github import paths from v89 to v90([c2fffcd](https://github.com/forkline/ingress-nginx/commit/c2fffcd0d23bfa37ab96dcdc7736c83b9191a11e))

### Documentation

- Replace upstream helm repo references with OCI chart in deploy guide([9dcf958](https://github.com/forkline/ingress-nginx/commit/9dcf958efe276d32e2e04105f77e4edba8593c2c))

- Replace upstream helm repo with OCI chart in chart README (chart)([ec8d4c8](https://github.com/forkline/ingress-nginx/commit/ec8d4c82f15d7359b9c06702d20d900014c122b0))

- Add installation section with OCI chart and image locations([e5d02bf](https://github.com/forkline/ingress-nginx/commit/e5d02bfe3363b1e894baaf8df85deb264eb59f62))

- Fix remaining helm repo references missed in OCI migration([cf3a8bf](https://github.com/forkline/ingress-nginx/commit/cf3a8bfbd208fdf1f813e814b00240b1789af0d7))

- Remove obsolete Get Involved section from README([0a7f3fb](https://github.com/forkline/ingress-nginx/commit/0a7f3fb3ad95d5886bb78798f77fb9f5ea535517))

### Chore

- Update pre-commit hook renovatebot/pre-commit-hooks (v43.272.6 → v44.30.4)

### Build

- Update module github.com/opencontainers/cgroups to v0.0.8 (#239) (deps)([8039345](https://github.com/forkline/ingress-nginx/commit/8039345df193c4812b470972b083c669c6ba5751))

- Update google.golang.org/grpc/examples digest (17 updates)

- Update module github.com/prometheus/common to v0.70.1 (#248) (deps)([f813063](https://github.com/forkline/ingress-nginx/commit/f813063c6028163b193f00951bfac287225cfdb1))

- Update go modules (5 updates)

- Update ossf/scorecard-action action to v2.4.4 (#259) (deps)([4a442bf](https://github.com/forkline/ingress-nginx/commit/4a442bf0f63673c017039d900e1c6b99bba7924e))

- Update module go.yaml.in/yaml/v3 to v3.0.5 (#271) (deps)([2401f26](https://github.com/forkline/ingress-nginx/commit/2401f266741c97d5b3774d52d499294fe9c2ea79))

- Update dependency helm to v4.2.4 (#317) (deps)([5f59268](https://github.com/forkline/ingress-nginx/commit/5f5926892626a6c74f3894888e9c9f2e128115d6))

- Update module github.com/google/go-github/v89 to v90 (deps)([4974b16](https://github.com/forkline/ingress-nginx/commit/4974b1685185173bbafa03083d8d6d02dbdb2b32))

## [v2026.7.20](https://github.com/forkline/ingress-nginx/tree/v2026.7.20) - 2026-07-20

### Fixed

- Update go-github imports from v88 to v89([be6f517](https://github.com/forkline/ingress-nginx/commit/be6f517aa6e084e1b71b81b3f406da5a2f829ae2))

- Replace nip.io with sslip.io for DNS resolution (e2e)([7ae2480](https://github.com/forkline/ingress-nginx/commit/7ae2480b8664497dfa91922a467b0cb7465bd9c5))

- Add HTTP readiness check for httpbun to reduce auth test flakiness([e6f9968](https://github.com/forkline/ingress-nginx/commit/e6f996852e30323ce44266450e0e30b98000cb7b))

- Add HTTP retry for auth E2E tests to reduce flakiness([f6ea709](https://github.com/forkline/ingress-nginx/commit/f6ea709be6b3b1f1406a73edf4abafb3bda2dd18))

- Add HTTP readiness check after endpoint wait to eliminate auth test flakiness (e2e)([d6ce5b5](https://github.com/forkline/ingress-nginx/commit/d6ce5b57f8890cdd66c99d5392b6e0aa56fa46aa))

- Rename gomodguard_v2 to gomodguard for golangci-lint v2 compatibility([cc53433](https://github.com/forkline/ingress-nginx/commit/cc53433f69b6937dffe4ef468b30fc9afd15aadf))

- Force nginx base image rebuild to include patch fixes([09aeb3e](https://github.com/forkline/ingress-nginx/commit/09aeb3e4a433da446d41c57521a81405ac6294c1))

- Fix crash with nginx 1.31.3: add NDK complex_value_end_code + set-misc not_found guard([1b0241c](https://github.com/forkline/ingress-nginx/commit/1b0241c9b860ff2a243f2adbdb7fc8ee9152de52))

### Chore

- Update pre-commit hook renovatebot/pre-commit-hooks (v43.246.0 → v43.272.3)

- Update NGINX to 1.31.3 (security release)([01e566c](https://github.com/forkline/ingress-nginx/commit/01e566c0b5276d9b5abf31bbe987168e1827c182))

### Build

- Update google.golang.org/grpc/examples digest (7 updates)

- Update go modules (4 updates)

- Update module github.com/pires/go-proxyproto to v0.14.0 (#163) (deps)([473a59a](https://github.com/forkline/ingress-nginx/commit/473a59a8686b80c8984ee2f8fcfeef8f22ae70df))

- Update dependency helm to v4.2.3 (#183) (deps)([d3b75d6](https://github.com/forkline/ingress-nginx/commit/d3b75d68a8638b6c19e5dda84808e60ce4679696))

- Update module github.com/google/go-github/v88 to v89 (deps)([eeecbfb](https://github.com/forkline/ingress-nginx/commit/eeecbfbc949c7bd6b2b2f5b3f5cf8bb504612723))

- Update actions/setup-go action to v7 (deps)([888f315](https://github.com/forkline/ingress-nginx/commit/888f3150cfbc9cd15c41cba9c95fa34dd1c626bb))

- Update dependency nginx/nginx to v1.31.3 (#209) (deps)([e6537b8](https://github.com/forkline/ingress-nginx/commit/e6537b8bb8095adc2b91ef07163f33f53f262719))

- Update dependency mkdocs-material to v9.7.7 (#221) (deps)([8f61d3f](https://github.com/forkline/ingress-nginx/commit/8f61d3fbad9d97f18de4038fdc2c49a508986083))

- Update squidfunk/mkdocs-material Docker tag to v9.7.7 (#222) (deps)([cccc592](https://github.com/forkline/ingress-nginx/commit/cccc59248e80054be122c2da7b326eb6b4e28463))

- Update actions/setup-python action to v7 (deps)([caa04b9](https://github.com/forkline/ingress-nginx/commit/caa04b9d1554f8629ef259e940d10f62a19579dd))

## [v2026.6.28](https://github.com/forkline/ingress-nginx/tree/v2026.6.28) - 2026-06-28

### Fixed

- Automerge minor updates, block nginx major automerge (renovate)([42151ca](https://github.com/forkline/ingress-nginx/commit/42151ca58e9ed0156916a3da8dee71752bd9710b))

- Prevent metrics endpoint exposure when disabled([e50c0a4](https://github.com/forkline/ingress-nginx/commit/e50c0a4b9ae90adc8d72673009feeaf05798215c))

### Chore

- Update pre-commit hook renovatebot/pre-commit-hooks (v43.227.0 → v43.245.0)

- Update pre-commit hook alessandrojcm/commitlint-pre-commit-hook to v9.26.0 (#136) (pre-commit)([8511621](https://github.com/forkline/ingress-nginx/commit/8511621fe14e512052b4f1e3738227be9d8d9cf3))

### Build

- Update github actions (3 updates)

- Update go modules (4 updates)

- Update google.golang.org/grpc/examples digest (2 updates)

- Update k8s.io/utils digest to be93311 (deps)([3c8aeb2](https://github.com/forkline/ingress-nginx/commit/3c8aeb2ebf7a9829691b6d289a2a8ed45aa50f62))

## [v2026.6.18](https://github.com/forkline/ingress-nginx/tree/v2026.6.18) - 2026-06-18

### Fixed

- Add timeouts to E2E workflow jobs (#79) (ci)([f4d2346](https://github.com/forkline/ingress-nginx/commit/f4d23469e849a05c41995faed298dac7dc21fe84))

### Chore

- Update pre-commit hook renovatebot/pre-commit-hooks (v43.201.1 → v43.222.1)

### Build

- Update google.golang.org/grpc/examples digest (5 updates)

- Update module github.com/prometheus/common to v0.68.0 (deps)([70f42ba](https://github.com/forkline/ingress-nginx/commit/70f42baa75e798751c554c3e859f2788d6e95fc3))

- Update go modules (6 updates)

- Update dependency helm (v4.2.1 → v4.2.2)

- Update github actions to v7 (deps)([f532336](https://github.com/forkline/ingress-nginx/commit/f53233660826cf925c87cfbc9fa5054a7a702f30))

- Update docker images (2 updates)

- Update dependency nginx/nginx to v1.31.2 (#116) (deps)([938b9a5](https://github.com/forkline/ingress-nginx/commit/938b9a56301d04c2bcd8eeebcb1e6a349e6cf676))

## [v2026.5.24](https://github.com/forkline/ingress-nginx/tree/v2026.5.24) - 2026-05-24

### Fixed

- Remove v prefix from container image tags in release notes([39dceba](https://github.com/forkline/ingress-nginx/commit/39dcebadfedfd296012f2a8ad73ecb5ddf356876))

- Update go-github import paths from v87 to v88 (deps)([4e52d9e](https://github.com/forkline/ingress-nginx/commit/4e52d9ed9cea60e034b7cd369c21e11c3107176a))

- Update NGINX_BASE to match NGINX_VERSION 1.31.0([95f504f](https://github.com/forkline/ingress-nginx/commit/95f504fbcea7ef00659fb7b45fdae4a7b8b98bf9))

- Remove Docker Hub tag from buildx push in e2e workflow([63e26f5](https://github.com/forkline/ingress-nginx/commit/63e26f5978bf06cfa54223d334c2bb561a5dbdfa))

### Build

- Update module github.com/google/go-github/v86 to v87 (#65) (deps)([9d440ec](https://github.com/forkline/ingress-nginx/commit/9d440ecfb312119e547e8fedb5fe9160a1de238b))

- Update google.golang.org/grpc/examples digest (4 updates)

- Enable automerge for google.golang.org/grpc/examples (deps)([2de964c](https://github.com/forkline/ingress-nginx/commit/2de964cb6262824d5321d95bb9c1802ba608ff30))

- Update go modules (deps)([c66bf96](https://github.com/forkline/ingress-nginx/commit/c66bf9652cd0ecaa5708fd6c1e2a987472505069))

- Update module github.com/google/go-github/v87 to v88 (deps)([64faa7b](https://github.com/forkline/ingress-nginx/commit/64faa7bea260919c5085f2881abe10bfa9c1f6de))

- Update dependency nginx/nginx (v1.31.0 → v1.31.1)

## [v2026.5.18](https://github.com/forkline/ingress-nginx/tree/v2026.5.18) - 2026-05-18

### Fixed

- Remove obsolete CVE-2025-23419 patch included in 1.30.1 (nginx)([362c07d](https://github.com/forkline/ingress-nginx/commit/362c07dc131440587d0e1130b80b8508f1095824))

- Nginx image caching by NGINX version and multi-platform builds (ci)([1f8b3b8](https://github.com/forkline/ingress-nginx/commit/1f8b3b8e72b7b25e7e991f465b1ccc450d56aa1e))

- Add yamllint config with 120 char line length (lint)([ffee14f](https://github.com/forkline/ingress-nginx/commit/ffee14f6550079df0477672aca75f82fff8d248f))

- Ensure E2E runs on push to main (ci)([86871aa](https://github.com/forkline/ingress-nginx/commit/86871aafff6a312989b369d33744cc3ea53d1bf3))

- Increase timeout to 360m in release images workflow (ci)([4a75adc](https://github.com/forkline/ingress-nginx/commit/4a75adcd9f0fbe9a1f3139afc067be0b36601cd4))

- Quote make variables to prevent target parsing errors (ci)([3018ef5](https://github.com/forkline/ingress-nginx/commit/3018ef59ec82a0fc00a4244002d95cf9fab2f1ad))

- Increase nginx build timeout to 360min for arm64 QEMU (ci)([db3de00](https://github.com/forkline/ingress-nginx/commit/db3de004016f6b8f6dd0e1d52c6b625c71e4a33a))

- Pass TAG to kube-webhook-certgen publish step (#59) (release)([0df4c3d](https://github.com/forkline/ingress-nginx/commit/0df4c3ddf3ce5071ecf7e6f177511784b2d9ecef))

- Strip v prefix from image TAG files([506fdd7](https://github.com/forkline/ingress-nginx/commit/506fdd7331d54f7688efcd79d42c752383229109))

### Build

- Update module github.com/google/go-github/v85 to v86 (#56) (deps)([18d8634](https://github.com/forkline/ingress-nginx/commit/18d8634ac4c2f6fbf7ac105fdb1ac56bc0c09d5e))

- Update module google.golang.org/grpc to v1.81.1 (#60) (deps)([6eff2ca](https://github.com/forkline/ingress-nginx/commit/6eff2ca90719fed84413f1121f14f580b4e65c37))

- Update google.golang.org/grpc/examples digest to 6602080 (#62) (deps)([671840a](https://github.com/forkline/ingress-nginx/commit/671840a6c2eb6b6bba8f957a3709b6dbc41c3dd9))

- Update go modules (#63) (deps)([729178d](https://github.com/forkline/ingress-nginx/commit/729178de09f5a77c1bdba62b8708c8cd82922eab))

## [v2026.5.14](https://github.com/forkline/ingress-nginx/tree/v2026.5.14) - 2026-05-14

### Added

- Update release script to maintain README Supported Versions table([b120e9d](https://github.com/forkline/ingress-nginx/commit/b120e9d6028de36b91674e8eb5f0d5d6e00624a6))

### Fixed

- Make gate jobs run on all PRs for required checks (ci)([ca995ff](https://github.com/forkline/ingress-nginx/commit/ca995ff9ce4b8d43c9f5120be59ab88e15be8aa8))

- Update to 1.30.1 stable and fix Renovate version tracking (#55) (nginx)([16be02e](https://github.com/forkline/ingress-nginx/commit/16be02e09826451ed54855ef7e8fb61ff96fabe0))

### Documentation

- Simplify README by removing duplicate information and consolidating notice([624d9ce](https://github.com/forkline/ingress-nginx/commit/624d9ceb2c0d97ef29f7c7d108cc9279bf5aae06))

### Chore

- Update pre-commit hook renovatebot/pre-commit-hooks to v43.150.0 (#45) (pre-commit)([248b673](https://github.com/forkline/ingress-nginx/commit/248b6735c4e47d38656d8ec6a8bcaebd55bb8c5d))

### Build

- Update github actions (#43) (deps)([3d7d717](https://github.com/forkline/ingress-nginx/commit/3d7d7171f53aab36b70b03b84d0eff93ac1cd58a))

- Update module google.golang.org/grpc to v1.81.0 (#44) (deps)([9dd98e5](https://github.com/forkline/ingress-nginx/commit/9dd98e50e2b9a6b2041055b875fa00a5fd71e3c3))

- Update module github.com/fsnotify/fsnotify to v1.10.1 (#46) (deps)([9648818](https://github.com/forkline/ingress-nginx/commit/9648818fbb271dafe703b12d72ae762b08d902c6))

- Update google.golang.org/grpc/examples digest (2 updates)

- Update go modules (2 updates)

- Update actions/dependency-review-action action to v5 (#50) (deps)([5265068](https://github.com/forkline/ingress-nginx/commit/5265068f6c41bd07155d843484bff7d8253c9308))

- Update dependency helm to v4.2.0 (#51) (deps)([430dc0b](https://github.com/forkline/ingress-nginx/commit/430dc0b927e3c91e3d174e8a1d8f02d744800acc))

- Enable Renovate tracking for upstream NGINX version (deps)([aa522d6](https://github.com/forkline/ingress-nginx/commit/aa522d6e8a3a7dbfd82fef4f1ac7932238ee240c))

## [v2026.5.3](https://github.com/forkline/ingress-nginx/tree/v2026.5.3) - 2026-05-02

### Fixed

- Correct sed pattern and regenerate chart docs for release([b3de357](https://github.com/forkline/ingress-nginx/commit/b3de357df0dca66805c3782e4886132330bbe048))

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

- Update go modules (3 updates)

- Group all update types together in renovate config (deps)([3da319e](https://github.com/forkline/ingress-nginx/commit/3da319ebb6853786ec045a06a7b4154d6d64dbe8))

- Update google.golang.org/grpc/examples digest to 1c132b9 (#27) (deps)([aa1f09e](https://github.com/forkline/ingress-nginx/commit/aa1f09ed7da32782f2622a19bc99bce0140796a6))

- Update module helm.sh/helm/v4 to v4.1.4 [SECURITY] (#31) (deps)([3b1950f](https://github.com/forkline/ingress-nginx/commit/3b1950f45f7801c361a5c4b66529a9a3c28a014d))

- Update alpine Docker tag to v3.23.4 (#36) (deps)([3795290](https://github.com/forkline/ingress-nginx/commit/379529001d9f3a04391bcdda44ecd22ef85146fd))

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

- Update pre-commit hook renovatebot/pre-commit-hooks (v43.104.1 → v43.140.0)

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

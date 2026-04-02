# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project
adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [unreleased]

### Chore

- renovate: Enable forkProcessing for fork support ([1c100e6](https://github.com/forkline/ingress-nginx/commit/1c100e6ec913eee676f3bfe704615acc0fee5d61))

## [v1.15.5](https://github.com/forkline/ingress-nginx/tree/v1.15.5) - 2026-04-02

### Fixed

- release: Point krew plugin to forkline artifacts ([71cb526](https://github.com/forkline/ingress-nginx/commit/71cb526e611748fb1540ce9921b93fca7e40d6ed))

## [v1.15.2](https://github.com/forkline/ingress-nginx/tree/v1.15.2) - 2026-04-02

### 404-server

- Graceful shutdown ([2f8e81e](https://github.com/forkline/ingress-nginx/commit/2f8e81e383d58b0b8d35b08c72b1e7d8982c6816))

### Added

- baremetal: Add kustomization.yaml ([7ddb734](https://github.com/forkline/ingress-nginx/commit/7ddb7343faa2286088409a2ea22bd47a7066e66a))
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
- Feat(#733)Support nginx bandwidth control ([890c57f](https://github.com/forkline/ingress-nginx/commit/890c57f2ca1c85d84e5bad821d9deae28959c069))
- Configurable proxy buffer number ([c92d29d](https://github.com/forkline/ingress-nginx/commit/c92d29d46211d0a6105f738b63fff5ca512127b7))
- Auth-req caching ([23504db](https://github.com/forkline/ingress-nginx/commit/23504db7708a5829ae918166d08dce8b25c881a3))
- Support to define trusted addresses for proxy protocol in stream block ([609e1b5](https://github.com/forkline/ingress-nginx/commit/609e1b5775099fb3477bbbc39a72372eaf3fdab8))
- Allow user to specify the maxmium number of retries in stream block. ([06f53bc](https://github.com/forkline/ingress-nginx/commit/06f53bcf058e9fb55b9232e29603e301be15b5a8))
- Add support for country databases ([1c6a1a0](https://github.com/forkline/ingress-nginx/commit/1c6a1a0e2360b60257ed7fe759d11176e837a042))
- Allow volume-type emptyDir in controller podsecuritypolicy ([39fea58](https://github.com/forkline/ingress-nginx/commit/39fea580851086c254637395ff03111edd7e8dcb))
- Feat(chart) Add volumes to default-backend deployment

Update changelog and Chart.yml

Signed-off-by: Pierre Péronnet <pierre.peronnet@ovhcloud.com> ([59f930d](https://github.com/forkline/ingress-nginx/commit/59f930dd25d4563d7903249b22e3568e402b4124))
- Multiple-cors-allow-origin support ([8a55801](https://github.com/forkline/ingress-nginx/commit/8a55801cc087e259493dc8ff787f3a37432c11a6))
- Add session-cookie-secure annotation ([f2e743f](https://github.com/forkline/ingress-nginx/commit/f2e743f561963f140e39f66bd2480d1ab2e68ab6))
- Added AdmissionController metrics ([a5bab6a](https://github.com/forkline/ingress-nginx/commit/a5bab6a7150f65dcf06588aeef31a41bad4b7da6))
- Always set auth cookie ([2c27e66](https://github.com/forkline/ingress-nginx/commit/2c27e66cc792f389186642a08406dc803c46ae4f))
- Support enbale nginx debug_connection ([2852e29](https://github.com/forkline/ingress-nginx/commit/2852e2998cbfb8c89f1b3d61de8ed03e0a1d0134))
- Migrate leaderelection lock to leases ([cf4dca8](https://github.com/forkline/ingress-nginx/commit/cf4dca8e43ad87e95d3a2fd4d708f3eb8c35de97))
- Update mimalloc to 1.7.6 ([0049796](https://github.com/forkline/ingress-nginx/commit/0049796682e509fc4c6a89931f093855fe54d014))
- Using LeaseLock for election ([730174f](https://github.com/forkline/ingress-nginx/commit/730174f73d8c1c581788f65bac8510c298d6ccd0))
- Switch from endpoints to endpointslices ([3579ed0](https://github.com/forkline/ingress-nginx/commit/3579ed04870c77979ec5bf18f4cd00c8763615a1))
- Add ovhcloud ([a3bed7a](https://github.com/forkline/ingress-nginx/commit/a3bed7ae4c02a2821e1aa811ed3ede71a35ddb5d))
- Support topology aware hints ([5b2a947](https://github.com/forkline/ingress-nginx/commit/5b2a9475dc2260dc15561fd117dc221b892464da))
- OpenTelemetry module integration ([c8cb916](https://github.com/forkline/ingress-nginx/commit/c8cb9167d3bfe6a2ed50796e179b3aec648a7e16))
- Add namespace overrides ([7ce6cc8](https://github.com/forkline/ingress-nginx/commit/7ce6cc88d8501628fd0fe1fb882579ad7d2a95fe))
- Add annotation to allow to add custom response headers ([1f4ee0e](https://github.com/forkline/ingress-nginx/commit/1f4ee0e235bfbc55d750b6a20cf4133381fa7dd7))
- Add grpc timeouts annotations ([d0e9934](https://github.com/forkline/ingress-nginx/commit/d0e9934789d5f07699788d5fa1f3819ebc0a8919))

### Fixed

- Chart: Mismatch between values.yml and README.md ([0f82342](https://github.com/forkline/ingress-nginx/commit/0f82342aa65f7339d57972452fbba5e562ff0100))
- ci: Allow helm unittest plugin install ([b378eed](https://github.com/forkline/ingress-nginx/commit/b378eedf4e292b068e974286f0a23e8cef8af0b1))
- ci: Stabilize go workflow on GitHub ([e07a658](https://github.com/forkline/ingress-nginx/commit/e07a6588222610abf9c166a649522f7c2ff27c18))
- ci: Run checks in project-compatible environments ([072e04d](https://github.com/forkline/ingress-nginx/commit/072e04d730bf4f1603c4d9e4b4fe03692ea8975d))
- ci: Install golangci-lint with current toolchain ([47d9d0e](https://github.com/forkline/ingress-nginx/commit/47d9d0e1174155226cde17f0ded11e1d4f23dfba))
- ci: Scope golangci-lint to new issues ([3aa95ce](https://github.com/forkline/ingress-nginx/commit/3aa95ce7c534093bb48142131e57a40a03c0a51d))
- ci: Fetch history for diff-based linting ([1a228b7](https://github.com/forkline/ingress-nginx/commit/1a228b7714a1f6612ad5c0d64b8b1f4e9ef5f0a9))
- ci: Limit image publishing to image changes ([ff620aa](https://github.com/forkline/ingress-nginx/commit/ff620aa39e6d3a01201f5511b4d07979f786c798))
- ci: Simplify image publishing to amd64 ([d9765fe](https://github.com/forkline/ingress-nginx/commit/d9765fe7a759442863ec55f94deaab4fc37489da))
- ci: Run image publishing on tags only ([d532994](https://github.com/forkline/ingress-nginx/commit/d532994200475836bedf6e0da447fd92e768dcee))
- controller: Typo in catch-all CheckIngress error message ([d5893d4](https://github.com/forkline/ingress-nginx/commit/d5893d4a2e0eb9cf16ad1a5c9e93d19c347e6110))
- cors: Ensure trailing comma treated as empty value to be ignored ([da51393](https://github.com/forkline/ingress-nginx/commit/da51393cac818545de796d3c3486c40127aa51f4))
- dashboard: Use regex for ingress ([6364e9d](https://github.com/forkline/ingress-nginx/commit/6364e9d3e3111a6c3d3fcfef0f6bfdd20d2ba957))
- dashboard: Remove unnecessary namespace variable in query ([4f58ef3](https://github.com/forkline/ingress-nginx/commit/4f58ef3e0f1563d8b10a917ff453ffdbca8f199a))
- docs: Describe MetalLB configuration via CRDs rather than configMap ([2843bb2](https://github.com/forkline/ingress-nginx/commit/2843bb264f5d31b2ed2514c300c65c33aca2557a))
- documentation: Fix some typos ([d1e955c](https://github.com/forkline/ingress-nginx/commit/d1e955ca0a2a770cdd464d260e37f6f578ea30c2))
- grafana-dashboard: Remove hardcoded namespace references ([101ab06](https://github.com/forkline/ingress-nginx/commit/101ab06010f80990009c212d7cb5c187e6502648))
- grafana-dashboard: Remove hardcoded namespace references ([45632e5](https://github.com/forkline/ingress-nginx/commit/45632e586397520afee83449737b85de822e8b24))
- hpa: Deprecated api version, bump to v2 ([27ffeeb](https://github.com/forkline/ingress-nginx/commit/27ffeeb18f73bbde1b66866e796714308cebf1f4))
- labels: Use complete labels variable on default-backend deployment ([707a5a0](https://github.com/forkline/ingress-nginx/commit/707a5a0bea93b241632cb5146ef74e30c35487c9))
- typo: Pluralize provider ([8a5eaa6](https://github.com/forkline/ingress-nginx/commit/8a5eaa63a93ee247b63b9cf6e3892abacd352d0c))
- Fixes #874

non-ascii character used. ([1a239ef](https://github.com/forkline/ingress-nginx/commit/1a239ef2ae38818db17a30d9506ab8388f9c9dcc))
- Fix typo in ingress/controllers/README.md ([54891ae](https://github.com/forkline/ingress-nginx/commit/54891aef04355b28a779c44b3f2ed2605ed2a05e))
- Fix typo in variable ProxyRealIPCIDR ([bd9ec42](https://github.com/forkline/ingress-nginx/commit/bd9ec42042468c6a83e582fa3757b5f0bc29ea08))
- Fix ref ([f96e964](https://github.com/forkline/ingress-nginx/commit/f96e964522c812052fd194eb70180266cb20c03f))
- Fix the wrong links to the examples and developer documentation ([fa608dd](https://github.com/forkline/ingress-nginx/commit/fa608dd317da923dd06c7c68c6d0e12dae0f85ee))
- Fix the wrong link to build/test/release ([2d01250](https://github.com/forkline/ingress-nginx/commit/2d012504f0cc08f7fbc23ba8a283733fa2dbc642))
- Fix typo for task.Queue's unit test case ([3cb68b4](https://github.com/forkline/ingress-nginx/commit/3cb68b421da4c9912d6b37defdb424dc41aa1067))
- Fix typo

Signed-off-by: fate-grand-order <chenjg@harmonycloud.cn> ([0cd3663](https://github.com/forkline/ingress-nginx/commit/0cd3663defb294d2003d675f229033458925e5b7))
- Fix wrong link(change titile) ([08149a7](https://github.com/forkline/ingress-nginx/commit/08149a7a216c98e877bd1b474ee6770947067d3e))
- Fix wrong link in the file of examples/README.md ([6c35adb](https://github.com/forkline/ingress-nginx/commit/6c35adbfd2e30403c84f10fe519f446ecbae7e14))
- Fix misspell "affinity" in main.go ([3d0e374](https://github.com/forkline/ingress-nginx/commit/3d0e374f9eb2b4dc96d581f0c5fd99ccf2e60bab))
- Fix some broken links
upgrade all nginx examples to latest version
moved some examples from contrib to this repo ([a2edde3](https://github.com/forkline/ingress-nginx/commit/a2edde35fc0cd0d59cceea4cafbe00d4f419c4ff))
- Fix nginx-udp-and-udp on same port ([23c4534](https://github.com/forkline/ingress-nginx/commit/23c45340be3f92b0fe9d616ad31b63c017b21f6d))
- Fix all go style mistakes about fmt.Errorf ([37bdb39](https://github.com/forkline/ingress-nginx/commit/37bdb3952e090e50f44207ad0e54b1f2a4ef1055))
- Fix vts readme ([be05d40](https://github.com/forkline/ingress-nginx/commit/be05d403ac473545fdfdaf656bc2f7ae856fb159))
- Fixed lua_package_path in nginx.tmpl

I did my own build of the nginx-ingress-controller and its docker image, but I had troubles with the `error_page.lua` module, which couldn't be loaded, there was an error in the log, module was not found.

I think the lua package path is wrong, here is a fix. ([beb17f3](https://github.com/forkline/ingress-nginx/commit/beb17f39aba12f21fce11d35bf29f5c7532d05ce))
- Fix header name ([9c56c72](https://github.com/forkline/ingress-nginx/commit/9c56c72464c973784e2d1e1ccc55ebceecca4157))
- Fix nginx reload flags '-c' ([9f32b74](https://github.com/forkline/ingress-nginx/commit/9f32b74feaaee03f386922353f914235ba1f6489))
- Fix nginx version to 1.13.3 to fix integer overflow in the range filter vulnerability ([53b8369](https://github.com/forkline/ingress-nginx/commit/53b83695d4056501ec655d4216579d6bb525885c))
- Fix the same udp port and tcp port, update nginx.conf error ([fa2c422](https://github.com/forkline/ingress-nginx/commit/fa2c422a6803e2483492a4267bfcd452f0e85c43))
- Fix typos in controllers/nginx/README.md ([e713e70](https://github.com/forkline/ingress-nginx/commit/e713e706cccaa4fabcff71965f99d374754ca23e))
- Fix typo ([6602b2f](https://github.com/forkline/ingress-nginx/commit/6602b2f0d883bfaec5a42a132b72bf0e2876a7bb))
- Fix typos ([9206a8b](https://github.com/forkline/ingress-nginx/commit/9206a8baf9139b0cede8fd52a1ca6b9c63edec59))
- Fix several titles ([34467ae](https://github.com/forkline/ingress-nginx/commit/34467aec825047b604b39cdd966bd4b05c300af0))
- Fix Type transform panic ([af6a7f6](https://github.com/forkline/ingress-nginx/commit/af6a7f6d1753afb71b8fa5e7ea2a6fb781d32066))
- Fix link to conformance suite ([fbd3778](https://github.com/forkline/ingress-nginx/commit/fbd3778fa4e5f46d8db6d88559a12e7ed068626e))
- Fix README of nginx-ingress-controller ([710eaff](https://github.com/forkline/ingress-nginx/commit/710eaffc69092f70c111d9777e258983ac673d76))
- Fix error when cert or key is nil ([003667f](https://github.com/forkline/ingress-nginx/commit/003667ff2ea6d7a8967685a291c9394b87002a33))
- Fix test ([15c6a11](https://github.com/forkline/ingress-nginx/commit/15c6a1175a0d71212c322653b797d6e67c6dd556))
- Fix link ([81f5170](https://github.com/forkline/ingress-nginx/commit/81f5170034537a11b2fc77b8f522974bb478fc0c))
- Fix link ([78639eb](https://github.com/forkline/ingress-nginx/commit/78639ebfbda60d57f2c2ef2ac22814ab821de86b))
- Fix ([b87c5ff](https://github.com/forkline/ingress-nginx/commit/b87c5ff39d22890ffe002a537f10c968c32c7574))
- Fix broken GCE-GKE service descriptor

fixes #1546 ([489b851](https://github.com/forkline/ingress-nginx/commit/489b851835142e66ce67ceed656abfa24cc765b7))
- Fixed https port forwarding

443 was redirected to named port 'http' instead of 'https' (points to https://github.com/kubernetes/ingress-nginx/blob/master/deploy/provider/patch-service-without-rbac.yaml#L35)
In the AWS/GCP examples it was fine ('https') so I fixed it for azure as well. ([f5ddf6a](https://github.com/forkline/ingress-nginx/commit/f5ddf6af7685bb880619c2a0c171996b0ff9e44b))
- Fix typo in user-guide/annotations.md ([a7a76bd](https://github.com/forkline/ingress-nginx/commit/a7a76bd69292b3451ce8d6fb7dfe2c2b95ff3247))
- Core() is deprecated use CoreV1() instead. ([b3cec74](https://github.com/forkline/ingress-nginx/commit/b3cec74e79e69dd8d1155876d3ed6f26d6df44de))
- Replace deprecated methods. ([a3136aa](https://github.com/forkline/ingress-nginx/commit/a3136aa0490e4964a6c1fd583a9a0de1712cd6c2))
- Fix typos in docs. ([e2ce52a](https://github.com/forkline/ingress-nginx/commit/e2ce52a55ed0b596ecaf89dd3754f4e0864f1721))
- Some typo. ([9cf0b11](https://github.com/forkline/ingress-nginx/commit/9cf0b11fc7c14ff32cb5d10f3c779997106b4988))
- Fix var checked ([8688953](https://github.com/forkline/ingress-nginx/commit/86889532aa0701afed8e8f73c6c7d865b0189c16))
- Fix typo error for server name _

Signed-off-by: qiupeng-huacloud <qiupeng@chinacloud.com.cn> ([c2cbcf7](https://github.com/forkline/ingress-nginx/commit/c2cbcf7a58c47165933215b9c149ea7abb9b3fc5))
- Fix broken links in static-ip readme

Signed-off-by: LinWengang <linwengang@chinacloud.com.cn> ([20675cc](https://github.com/forkline/ingress-nginx/commit/20675cccd02c0dd70213d4a4f080dc8a79596484))
- Fix typo stickyness to stickiness ([f5d764c](https://github.com/forkline/ingress-nginx/commit/f5d764cc9c1c334fde254b813a94cf76e8f8da93))
- Fix wrong annotation ([36acee2](https://github.com/forkline/ingress-nginx/commit/36acee29ff26ee6474a06c3eac951f11b213ae57))
- Fix spell error reslover -> resolver ([7ede78f](https://github.com/forkline/ingress-nginx/commit/7ede78f00480562fe08cab909b8175e28ade2abb))
- Fix limit-req-status-code doc ([309a794](https://github.com/forkline/ingress-nginx/commit/309a79483f3cc3d9eb70cb826d4c77dd639a70ed))
- Fix wrong json tag

json tags are case sensitive when encode, change omitEmpty to omitempty ([ebcdfad](https://github.com/forkline/ingress-nginx/commit/ebcdfade8eb7b80b6cb190fca108872e1cccc3b3))
- Fix grammer mistake

fix grammer mistake
```release-note
None
``` ([d27a132](https://github.com/forkline/ingress-nginx/commit/d27a13223fa2a3c4065ed645caf1b1e4cffb5558))
- Fix go test TestSkipEnqueue error, move queue.Run ([f4caa13](https://github.com/forkline/ingress-nginx/commit/f4caa13b2879c192da045ea9385a0f050248835d))
- Fix wrong config generation when upstream-hash-by is set ([df50487](https://github.com/forkline/ingress-nginx/commit/df50487a35459f5b0ac0d7a2d9f621114b754307))
- Fix-link ([65fde75](https://github.com/forkline/ingress-nginx/commit/65fde75a7f5483d336906849d197c65832656356))
- Cannot set $service_name if use rewrite ([0b0a274](https://github.com/forkline/ingress-nginx/commit/0b0a274a9a6a35dd90d5ff8211c40b5a3144aa7b))
- Empty ingress path ([1f93a1c](https://github.com/forkline/ingress-nginx/commit/1f93a1ccadbf91711dc5f29b1ab33c00c242cce1))
- Fix nil pointer when ssl with ca.crt ([17f6996](https://github.com/forkline/ingress-nginx/commit/17f6996941234a7a5bcc3fbf45daf7199a5899c0))
- Fix make verify-all failures ([9198e2c](https://github.com/forkline/ingress-nginx/commit/9198e2c14b616414d047ab733d4a2f29a464a02b))
- Fill missing patch yaml config. ([bde3394](https://github.com/forkline/ingress-nginx/commit/bde3394fb290f01bf9167a33a856a476a21fa2e4))
- Fix the default cookie name in doc ([4237f29](https://github.com/forkline/ingress-nginx/commit/4237f290ed737e26771c9ad1d6eed25d017d22dd))
- Fix flaky dynamic configuration test ([b2084c0](https://github.com/forkline/ingress-nginx/commit/b2084c057d142a45ed115ce18d2606ad5a51a622))
- Fix flaky test ([2eb0286](https://github.com/forkline/ingress-nginx/commit/2eb0286c8aca36fc56ecaf464499ebe89d7ef096))
- Fix bug with lua sticky session implementation and refactor balancer ([7ac4e1d](https://github.com/forkline/ingress-nginx/commit/7ac4e1db3013432f72d18c6bd97d6ef44c2caba5))
- Fix ewma.balance and add unit tests for it ([04b7356](https://github.com/forkline/ingress-nginx/commit/04b7356190ce5d7e7bb9aec8afbe3038dce9e0c4))
- Fix changelog link in README.md ([f7ee676](https://github.com/forkline/ingress-nginx/commit/f7ee676c130c2f04ee5239b251a4a87a74b860fb))
- Fix nginx conf test error when not found active service endpoints ([aeab703](https://github.com/forkline/ingress-nginx/commit/aeab7035f8d8e716478f16eedb0c48fac04af402))
- Fix for #1930, make sessions sticky, for ingress with multiple rules and backends

* for an ingress with session affinity cookie, set the location as path on the cookie when unique
* the previous behaviour ( cookie path=/ ) is preserved for ingresses with multiple rules for the same backend (locations not unique)

added e2e tests for session affinity, setting path on sticky config

added tests:
* it should set the path to /something on the generated cookie
* it should set the path to / on the generated cookie if there's more than one rule referring to the same backend ([1a320ae](https://github.com/forkline/ingress-nginx/commit/1a320ae289048ece4585076949bbf5f65c81402b))
- Use the correct opentracing plugin for Jaeger ([85d1742](https://github.com/forkline/ingress-nginx/commit/85d17422839841eccec37e6accd93f766e13de3d))
- Fix custom-error-pages functionality in dynamic mode ([ed19dc3](https://github.com/forkline/ingress-nginx/commit/ed19dc3bc680361581bbaf372947bf07ed12ad15))
- Fix issues introduced in #2804 ([a2692ce](https://github.com/forkline/ingress-nginx/commit/a2692ce94679d72e78d459547151a7923fa33eda))
- Fix bug with lua e2e test suite ([fa74877](https://github.com/forkline/ingress-nginx/commit/fa74877256cb916b4a3ed6be4c3656de5912339a))
- Fix the bug #2799, add prefix (?i) in rewrite statement and add new e2e
test. ([72a2aa1](https://github.com/forkline/ingress-nginx/commit/72a2aa171a575a34b3d686f79b3f84f657a565bd))
- Sort TCP/UDP upstream order ([6d9772c](https://github.com/forkline/ingress-nginx/commit/6d9772ce007195c640ade310385b8b4088abfedd))
- Fixed rewrites for paths not ending in / ([e428095](https://github.com/forkline/ingress-nginx/commit/e428095e3ce2e8f15140d44496e1d2a2fc80d360))
- Fix variable parsing when key is number ([27cd1af](https://github.com/forkline/ingress-nginx/commit/27cd1af4a7e7476ae078aaf3762095266064c445))
- Fix flaky luarestywaf test ([aa3e06b](https://github.com/forkline/ingress-nginx/commit/aa3e06b1890b7e7c434cfdc687086b72fd712d3b))
- Fixed jsonpath command in examples ([cbdd12f](https://github.com/forkline/ingress-nginx/commit/cbdd12f898095976caf0c8907297c14a674e4d14))
- Don't try and find local certs when secretName is not specified ([6648620](https://github.com/forkline/ingress-nginx/commit/66486203dbb95dddba69a88dc33a350e0b6bc6f1))
- Fix some typos

Signed-off-by: Lei Gong <lgong@alauda.io> ([e73510d](https://github.com/forkline/ingress-nginx/commit/e73510d8189969b9f88aefbfd29cf5cb7e430cca))
- Fix missing datasource value ([8298d27](https://github.com/forkline/ingress-nginx/commit/8298d27899310ffff31528b2935de1eeca73819c))
- Fix typos ([1d0e752](https://github.com/forkline/ingress-nginx/commit/1d0e7523391d97dffecb197c23bbaa08f5954de0))
- Fix newlines location denied ([6454608](https://github.com/forkline/ingress-nginx/commit/6454608c6ca03fc32d71659415038882a2afa03f))
- Fix two bugs with backend-protocol annotation ([cdb244e](https://github.com/forkline/ingress-nginx/commit/cdb244e579ff41e4bd76b2d539bc3511a33f3084))
- Fix typos ([51fffc6](https://github.com/forkline/ingress-nginx/commit/51fffc653d8b0acc71818058fd643d903b068147))
- Fix bug with balancer.lua configuration ([5cc116f](https://github.com/forkline/ingress-nginx/commit/5cc116fa100e30b882a162bf5a1829d3f93ebb39))
- Fix logging calls ([9d227ab](https://github.com/forkline/ingress-nginx/commit/9d227ab62db6068d1bbec84a1ad2e04ce5ad947c))
- Fix sticky session implementation ([9e639f9](https://github.com/forkline/ingress-nginx/commit/9e639f97888ead927ab918210f410e9b787cc270))
- Fix baremetal.md link ([c4834e5](https://github.com/forkline/ingress-nginx/commit/c4834e5063b16186b3d46186a674422ee0ea5684))
- Fix typo ([f1b2a54](https://github.com/forkline/ingress-nginx/commit/f1b2a540fd8e58a2f2503c6d51de0f6fb8c03d7f))
- Fix typos ([76b5a7b](https://github.com/forkline/ingress-nginx/commit/76b5a7b45e87cb1d1e869a6b76d4200dfac54f9a))
- Fix the typos ([76aae20](https://github.com/forkline/ingress-nginx/commit/76aae20b641a62849f0ca16813e2f6d61c4ae553))
- Fix logging calls ([2850fb5](https://github.com/forkline/ingress-nginx/commit/2850fb538a787362693bf28c3f94114515bb57ad))
- Fix Status key conflic, fixes https://github.com/kubernetes/ingress-nginx/issues/3451 ([068d633](https://github.com/forkline/ingress-nginx/commit/068d633e81ebc36dac74fb5048fbfe492f3db84e))
- Fix an ewma unit test ([a4bad90](https://github.com/forkline/ingress-nginx/commit/a4bad90f1f418f735e25814475dad39af5a46d50))
- Fix status updated: make sure ingress.status is copied ([7a00d52](https://github.com/forkline/ingress-nginx/commit/7a00d521417aa54fc84360dc3e4140b13a3e0ee2))
- Fix dind-cluster executable path ([17d4b2c](https://github.com/forkline/ingress-nginx/commit/17d4b2c64c9d5c81f6f1ed9e31406a762312a5c6))
- Fix issue with failing e2e tests ([57440c9](https://github.com/forkline/ingress-nginx/commit/57440c946404470dd03dfb23cce9d26505b307b1))
- Fixed test case for math.randomseed. ([c782f22](https://github.com/forkline/ingress-nginx/commit/c782f22c5d8c7e6c1ded62f704dfd523f9e5e2e1))
- Fix no config change bug with custom-http-errors ([94a9a47](https://github.com/forkline/ingress-nginx/commit/94a9a47225c8f1647fe599999a7ab7c645ec33a2))
- Rename proxy-buffer-number to proxy-buffers-number ([dc63e5d](https://github.com/forkline/ingress-nginx/commit/dc63e5d185a5a9df0786e0c0cb73cf4006b4547a))
- Run gofmt ([6305e1d](https://github.com/forkline/ingress-nginx/commit/6305e1d1520490a30aad68c2990f5c8a587a53f8))
- Fix function comment ([c934509](https://github.com/forkline/ingress-nginx/commit/c9345093091fd0ba8a3fd517b1e6a979cabc1d3d))
- Fix typo: delete '`'

fix typo: delete '`' ([fd1f200](https://github.com/forkline/ingress-nginx/commit/fd1f200eb481967775dc692aeb53f2f71709b08a))
- Fix custom default backend test title ([107bec6](https://github.com/forkline/ingress-nginx/commit/107bec676af394d1e2865a82fe6b68516020b64c))
- Fix dynamic cert bug ([fc6e7c9](https://github.com/forkline/ingress-nginx/commit/fc6e7c9be033fac2994de75328d218173d272274))
- Fix luacheck warning ([93f00b2](https://github.com/forkline/ingress-nginx/commit/93f00b2143d8057f01a42e0065047c1d9a7c28d5))
- Fix typo: deployement->deployment ([8ea40bb](https://github.com/forkline/ingress-nginx/commit/8ea40bbcb98922c854443e2a14519a4b2baa5354))
- Fix e2e-test make target

- explicitly wait for api token
- only use posix shell conditionals ([731c2d8](https://github.com/forkline/ingress-nginx/commit/731c2d8e4c42250e2538be97e1e12dd2349de913))
- Fix source file mods ([c4ced9d](https://github.com/forkline/ingress-nginx/commit/c4ced9d6944709d246650a653a26ab4ff49ae839))
- Fix monitor test after move to openresty ([2b46c3a](https://github.com/forkline/ingress-nginx/commit/2b46c3a056abc576b5ff0e6846831cd7fe9aa1a8))
- Fix lua lints ([97d3a0d](https://github.com/forkline/ingress-nginx/commit/97d3a0ddab6badb6e86bb337bd241572f9129575))
- 　fix image link.

Signed-off-by: Jintao Zhang <zhangjintao9020@gmail.com> ([5f1ebb4](https://github.com/forkline/ingress-nginx/commit/5f1ebb4c8409d47a41bceef9506bbae8c37eaf32))
- Fix dev-env script ([c898ad6](https://github.com/forkline/ingress-nginx/commit/c898ad6974cc26c9aa5d0dfb02d6a83bba1a71b0))
- Fix test by setting default luashareddicts ([94052b1](https://github.com/forkline/ingress-nginx/commit/94052b1bfc498545ca2efb9542046ce1bd3b23f2))
- Fix lua certificate handling tests ([57db904](https://github.com/forkline/ingress-nginx/commit/57db904c929ca73744707ea124288d73f7bde0da))
- Fix typo ([2ba1a9e](https://github.com/forkline/ingress-nginx/commit/2ba1a9e71aa34cf5b04edc4650b3f7fc194ff2f7))
- Fix bug with new and running configuration comparison ([d8a3d61](https://github.com/forkline/ingress-nginx/commit/d8a3d616b4a355967e27c8a2422bb14374f38acc))
- Fix ingress name in get example ([efc64c8](https://github.com/forkline/ingress-nginx/commit/efc64c85a4a6706e23cf7dd3b54b31d7ea122e23))
- Fix duplicate hsts bug ([54918c0](https://github.com/forkline/ingress-nginx/commit/54918c0ff29aad7c3d10df72e08a6805e838ce6d))
- Fix formatting ([c9f4ed9](https://github.com/forkline/ingress-nginx/commit/c9f4ed928312771125e01ee679017302e2939f05))
- Fix typo ([ad8a01f](https://github.com/forkline/ingress-nginx/commit/ad8a01f945e2ba0d40b5b00f6042a5de9180666f))
- Remove unnecessary if statement when redirect annotation is defined ([eefb32c](https://github.com/forkline/ingress-nginx/commit/eefb32c667afe7ec9c3791ec325c10707d3d4e17))
- Fix e2e rests related to the removed string ([39c01c6](https://github.com/forkline/ingress-nginx/commit/39c01c6f968aa30014313118b3bc34506f71b674))
- Fix typo in changelog ([093f354](https://github.com/forkline/ingress-nginx/commit/093f354325180c251761965152565c57930b05d8))
- Fix oauth2-proxy image repository ([e51556f](https://github.com/forkline/ingress-nginx/commit/e51556f692be5623e30d011a3b86bb93bbfd57fc))
- Fix undefined variable $auth_cookie  error when location is denied

(add) isLocationAllowed check before setting the cookie ([7767230](https://github.com/forkline/ingress-nginx/commit/7767230e6a6cd8372eca98c08c59a09441fd1b40))
- Fix for #5666 ([4ff9830](https://github.com/forkline/ingress-nginx/commit/4ff9830056106b096c301d46a41b5f62f367a7cc))
- Fix json tag for SSLPreferServerCiphers

related https://github.com/kubernetes/ingress-nginx/pull/5534 ([8557677](https://github.com/forkline/ingress-nginx/commit/8557677a5ee5d396df7ad73e741b837eec9b03d7))
- Fixed some typos

Signed-off-by: Frank Gadban <frankgad@outlook.de> ([e9059ee](https://github.com/forkline/ingress-nginx/commit/e9059eef0131ef2a2008c3ed8938c634addb8860))
- Fix for 5590 ([6239a66](https://github.com/forkline/ingress-nginx/commit/6239a66e5ed3a820a1e3bae1cb32b977063e6a33))
- Fix variable ordering in error message ([af29ec1](https://github.com/forkline/ingress-nginx/commit/af29ec11376b606f2f2fdcf10e5ec884daa6ee44))
- Log warning if empty ingress class is monitored. Improve docs related to --ingress-class ([903e511](https://github.com/forkline/ingress-nginx/commit/903e511b9d35d0d62350e04f470ab822f20ef0e0))
- Fix podAnnotations quotes for #6315

bumped chart version, daemonset podannotations

missing end on podannotations

ci values files

new lines at the end of files ([3ae837b](https://github.com/forkline/ingress-nginx/commit/3ae837b4b0af05546a590bc733a1c1189bcc000f))
- Fix for 6219 ([f7372d6](https://github.com/forkline/ingress-nginx/commit/f7372d603e458f32cabc41da791701d0282b0f4c))
- OWASP CoreRuleSet rules for NodeJS and Java ([0cf475a](https://github.com/forkline/ingress-nginx/commit/0cf475ad2d0505cceea29796064113d7c8e4c296))
- Fix controller service annotations ([9872e37](https://github.com/forkline/ingress-nginx/commit/9872e37b0dcb1f485ad9a244b9e7c61edd50a61c))
- Empty IngressClassName, Error handling ([e5fa90d](https://github.com/forkline/ingress-nginx/commit/e5fa90db9b65a8855ebcb22b72e800a3a0ff0c5c))
- Fix log-format-upstream sample

- Align column names to snake case.
- Align the space. ([1182569](https://github.com/forkline/ingress-nginx/commit/11825698ff59f6b882dd162fbd29c0f85b00919e))
- Fixed misspell

Update rootfs/etc/nginx/lua/plugins/README.md ([1ad89c8](https://github.com/forkline/ingress-nginx/commit/1ad89c8bb21e47ad6e345c8ebdf82bd03ddee3db))
- Fix for 6564

review comments ([57f8106](https://github.com/forkline/ingress-nginx/commit/57f81068a4663e32d2be0810524191e07a59a368))
- Move maxmindLicenseKey to controller.maxmindLicenseKey ([64a6e87](https://github.com/forkline/ingress-nginx/commit/64a6e87c1be5b3d39db91b0f953508c3fb3408a5))
- Fix flaky lua tests ([1e9650a](https://github.com/forkline/ingress-nginx/commit/1e9650a0f91450b1e938d734c596f559fc6660c5))
- Fix typo ([8852266](https://github.com/forkline/ingress-nginx/commit/885226618fa3f9fc635b9f1373b109d88e215904))
- Fix generated code for the new year ([bc6a271](https://github.com/forkline/ingress-nginx/commit/bc6a2718d2c98719f51348dbb348b4efae2d415c))
- Fix ipmatcher installation ([dfed436](https://github.com/forkline/ingress-nginx/commit/dfed436b9a98f8a870b3ff07ffb7d39eda6888cf))
- Fix link in annotation docs ([8c193a2](https://github.com/forkline/ingress-nginx/commit/8c193a229749cfcab5e50f8e61dce5e75bf206d7))
- Fix the documentation for the proxy-ssl-secret and the auth-tls-secret annotations ([15eff82](https://github.com/forkline/ingress-nginx/commit/15eff8220a2dccd022c48dd43ca9193c22aca2c5))
- Fix scaleTargetRef definition for KEDA v.2 ([d70652a](https://github.com/forkline/ingress-nginx/commit/d70652a0eb1a51642d529e201afee6d97ba9ddf6))
- Fix crl not reload when crl got updated in the ca secret ([4ddb0c7](https://github.com/forkline/ingress-nginx/commit/4ddb0c724a889dd73da7b9f004ee80cf07dc281b))
- Fix typos ([7cdc819](https://github.com/forkline/ingress-nginx/commit/7cdc819bb33e47628aa626912255249952eb18c1))
- Fix version with gen script ([6d53dd1](https://github.com/forkline/ingress-nginx/commit/6d53dd1430f15af4d75bc4df90b89c64c9d65d24))
- Fixed markdown typo

Signed-off-by: Mathis Van Eetvelde <mathis.vaneetvelde@protonmail.com> ([727020b](https://github.com/forkline/ingress-nginx/commit/727020b3acaaa336df86b6d1102c4563030d97a5))
- Use exponential backoff mechanism to listen on nginx.StatusPort ([a088870](https://github.com/forkline/ingress-nginx/commit/a08887040b8a4e4057f3e58166c08b351b5115a3))
- Fixing wording for #7094 ([e2f421b](https://github.com/forkline/ingress-nginx/commit/e2f421b9f442d1e31392cea9156be06f36f606b4))
- Fix for #7197 & #7285 ([9e274dd](https://github.com/forkline/ingress-nginx/commit/9e274dd41cc214219be873680184912407437f48))
- Discover mounted geoip db files ([c9d5b21](https://github.com/forkline/ingress-nginx/commit/c9d5b21a657c5267b396893f8e2064d9ec01e60e))
- Fix ingress-nginx panic when the certificate format is wrong.

* fix ingress-nginx panic when the certificate format is wrong.

Signed-off-by: wang_wenhu <976400757@qq.com>

* Add unit test.

Signed-off-by: wang_wenhu <976400757@qq.com>

* Update controller_test.go ([6593cb2](https://github.com/forkline/ingress-nginx/commit/6593cb244b21f6312c7a438708568807a9135a44))
- Fix 7591 ([9f9122c](https://github.com/forkline/ingress-nginx/commit/9f9122c38141542aac7efb2f09f610e4ece57605))
- Fix charts README.md to give additional detail on prometheus metrics …

* fix charts README.md to give additional detail on prometheus metrics configuration

* charts readme.md remove blank line ([f182b31](https://github.com/forkline/ingress-nginx/commit/f182b317ba7d69d1dfc4abcf5e39b22e705beea9))
- Fix cli flag typo in faq ([e779194](https://github.com/forkline/ingress-nginx/commit/e7791941ba6fd72ec343862ed5ae1de1f058a5a6))
- Fix typos. ([557a765](https://github.com/forkline/ingress-nginx/commit/557a765754f83da8f71917aba4a6c75bc5118bb7))
- Upgrade lua-resty-balancer to v0.04 ([0606ef8](https://github.com/forkline/ingress-nginx/commit/0606ef8282bdb46d9163dbee5be416f8f8c26217))
- Fix overlap check when ingress is configured as canary ([e8d9144](https://github.com/forkline/ingress-nginx/commit/e8d914475910e6de6e2dc1977d22db0161eb96f4))
- Fix reference to contributing.md in issue-triage.md ([068eccb](https://github.com/forkline/ingress-nginx/commit/068eccb6073ebbb8c8113dc1be5665adac17bcfc))
- Fixed issue 7807 ([0b24ade](https://github.com/forkline/ingress-nginx/commit/0b24ade145de1166d203c49ddac9229c911eb03b))
- Fix thread synchronization issue #6245 ([880ea6d](https://github.com/forkline/ingress-nginx/commit/880ea6dba887b7b2f88feabbdf95cff7203ee7a2))
- Fix ingress syntax. ([ed34f6c](https://github.com/forkline/ingress-nginx/commit/ed34f6c93d27fa69253822cf0d66f542d369a1c8))
- Fix missing `\-` in regex expression for CORS wildcard domain ([100057d](https://github.com/forkline/ingress-nginx/commit/100057d0c5daec02331bb3412bd5d72e203e899e))
- Fix compling kubectl-ingress_nginx error and add some descriptions for build command. ([ab4fa42](https://github.com/forkline/ingress-nginx/commit/ab4fa4246797a33bc7b45b45097a91086fe9ad48))
- Go-grpc Dockerfile ([e621c6e](https://github.com/forkline/ingress-nginx/commit/e621c6e973ffd12ea0e1fbfdac47c0dd217f10de))
- Fix to really execute plugins in order ([6163231](https://github.com/forkline/ingress-nginx/commit/6163231ef604664ccb0728367245527f13ec9fe7))
- Fix nginx compilation flags

* use '-O2' instead of '-Og'
  '-O2' produce production optimized binary while '-Og' is used mostly
  for debugging
* use '-mtune=generic' instead of '-mtune=native'
  '-mtune=native' produce optimal code for builder host system, but it
  can be sub-optimal for execution host system ([8ccec84](https://github.com/forkline/ingress-nginx/commit/8ccec8449652b08d65d0a29838e66b5ea8624b29))
- Fix custom-error-pages file not exist ([8e9bf7f](https://github.com/forkline/ingress-nginx/commit/8e9bf7f62c454bed1009a9a99ded721f1ed61ca0))
- Fix helmdoc push shell script ([2db580a](https://github.com/forkline/ingress-nginx/commit/2db580a51385c26a356ace9b511d0878955c5444))
- Fix inconsistent-label-cardinality for prometheus metrics: nginx_ingress_controller_requests

* fix inconsistent-label-cardinality

for prometheus metrics: nginx_ingress_controller_requests

* add host to collectorLabels only if metricsPerHost is true ([86964b1](https://github.com/forkline/ingress-nginx/commit/86964b15a889c37b41f520b028643147213d0086))
- Deny locations with invalid auth-url annotation ([1e2ce80](https://github.com/forkline/ingress-nginx/commit/1e2ce80846c91f3c034ea62915bfe6d756c9cea3))
- Fix the cloud build

Signed-off-by: James Strong <strong.james.e@gmail.com> ([5c47803](https://github.com/forkline/ingress-nginx/commit/5c47803d0f0d37cd83dc942a7e1c3ccd8565681a))
- Fix indent on env

* fix indent on env

* revert tag to      - TAG=$_GIT_TAG ([e51c151](https://github.com/forkline/ingress-nginx/commit/e51c15160e9b4d641fc9e289601676e1adc70e19))
- Fix code highlight ([08fcd94](https://github.com/forkline/ingress-nginx/commit/08fcd942c7fa141d806cf689a57e0e2b65f235a7))
- Fix build and upgrade otel to latest version ([79aa435](https://github.com/forkline/ingress-nginx/commit/79aa43540c67037ceae494a2f655adb00671ead3))
- Fix change log changes list ([59c6c05](https://github.com/forkline/ingress-nginx/commit/59c6c058fe7fc9c3ea8ebdc29c8712f3ef363dd3))
- Fix typo ([67e430c](https://github.com/forkline/ingress-nginx/commit/67e430cb3abf387a2be80c00abe46808e007e254))
- Fix opentelemetry-cpp-contrib sha256 ([2d2ec55](https://github.com/forkline/ingress-nginx/commit/2d2ec558ae0445f5aa826f1224587d6bab752ac9))
- Fix the gosec test and a make target for it ([f5d3ab4](https://github.com/forkline/ingress-nginx/commit/f5d3ab450538a02d26436f7c7ce508324b88db87))
- Fix bullet md format ([653f1e8](https://github.com/forkline/ingress-nginx/commit/653f1e8a9bf7351e7a1b9c0c243b88317abfa387))
- Add MAC_OS variable for static-check ([60b7143](https://github.com/forkline/ingress-nginx/commit/60b714336e81ee8089a3e989b0df40d1730b87e6))
- Test-runner Makefile ([e55e848](https://github.com/forkline/ingress-nginx/commit/e55e8488daa27634484ed295d4de021a023fb0a0))
- Test-runner prow build ([423008b](https://github.com/forkline/ingress-nginx/commit/423008b75282616413b743321d60d3a9557a570e))
- Make use of sed portable for BSD and GNU ([f9dcc13](https://github.com/forkline/ingress-nginx/commit/f9dcc13a0e0aa81fbf85da26cfb3104287257ac4))
- Test-runner prow build ([92f81e7](https://github.com/forkline/ingress-nginx/commit/92f81e7449fa5bce5777b2aba924373984fb8530))
- Change cloudbuild configuration ([fe116d6](https://github.com/forkline/ingress-nginx/commit/fe116d62cb4018d5d85847c45050b7d46286d4bc))
- Change all cloudbuild jobs configuration ([f0490cb](https://github.com/forkline/ingress-nginx/commit/f0490cbfbf29a7a05caaac29998dde56173ac2bb))
- Fix permissions

Signed-off-by: James Strong <strong.james.e@gmail.com> ([9162fe0](https://github.com/forkline/ingress-nginx/commit/9162fe0aa82996ca0e4988833c0d3d5626123e03))
- Bump k8s dependencies to fix go-restful CVE ([f6d04d3](https://github.com/forkline/ingress-nginx/commit/f6d04d3e3bb3db11641644962c1a5317cacf6cd8))
- Fixed deprecated ginkgo flags ([c6b70ec](https://github.com/forkline/ingress-nginx/commit/c6b70ec349cc68a627cdd1f71dd6038ce2a0580f))
- Fix LD_LIBRARY_PATH ([9a42ded](https://github.com/forkline/ingress-nginx/commit/9a42ded8bac9630a55e20e29d47db4af3ce904e6))
- Fix otel init_module ([981ce38](https://github.com/forkline/ingress-nginx/commit/981ce38a7f0455bee35d692af237089ce5a42220))
- Do not apply job-patch psp on Kubernetes 1.25 and newer ([67f7d3d](https://github.com/forkline/ingress-nginx/commit/67f7d3da63ba723c2d5f8f2d1dacd40a72a248c2))
- Fix chroot module mount path ([1a078af](https://github.com/forkline/ingress-nginx/commit/1a078af307d3cb52e3ad4c17e6a26bf4ef82581b))
- Fixed broken helm version comparision ([5fb3b97](https://github.com/forkline/ingress-nginx/commit/5fb3b974afb1cf12fd6e67c6db4c23499087d270))
- Fix wrong tag ([c211aa8](https://github.com/forkline/ingress-nginx/commit/c211aa83d3fee9d1d9e3b631cbecb7035b8fefa8))
- Fix e2e resource leak when ginkgo exit before clear resource ([c9faac2](https://github.com/forkline/ingress-nginx/commit/c9faac2222b5d510d040d08fca20854035eb0585))
- Handle 401 and 403 by external auth ([6aac006](https://github.com/forkline/ingress-nginx/commit/6aac00648b6353d25e455594c525d07cf6d80960))
- Fix ports ([499dbf5](https://github.com/forkline/ingress-nginx/commit/499dbf57af5d0d7eb2fdb298cc89b5f4bf179854))
- Fix typo in docs. ([a383cfc](https://github.com/forkline/ingress-nginx/commit/a383cfc55100aae3894629bea224472c578909c0))
- Fix svc long name

Signed-off-by: tombokombo <tombo@sysart.tech>

Signed-off-by: tombokombo <tombo@sysart.tech> ([490ecff](https://github.com/forkline/ingress-nginx/commit/490ecffc5242be8c54f8735d4af4d4ccad39acb7))
- Fix CVE-2022-27664 ([8949379](https://github.com/forkline/ingress-nginx/commit/894937993107e1dbe6e1b28bd916b06790a07d3c))
- Fix broken annotation yaml ([69a811d](https://github.com/forkline/ingress-nginx/commit/69a811dde9368f0464c1e050648a979a67fad3e8))
- Fixed multiple ginkgo versions ([1e08519](https://github.com/forkline/ingress-nginx/commit/1e08519a733b50134eae37f300f96b58bdb616e5))
- Missing CORS headers when auth fails ([3aa53aa](https://github.com/forkline/ingress-nginx/commit/3aa53aaf5b210dd937598928e172ef1478e90e69))
- Fix change images

Signed-off-by: James Strong <james.strong@chainguard.dev>

Signed-off-by: James Strong <james.strong@chainguard.dev> ([8b5a25f](https://github.com/forkline/ingress-nginx/commit/8b5a25fa141dff2d7f6bb8eecc07ee66a1f1a210))
- Disable auth access logs ([424cc86](https://github.com/forkline/ingress-nginx/commit/424cc8671b1c6df1db08283179d113fbaa5b2213))
- Fixed boiler plate lint

Signed-off-by: James Strong <strong.james.e@gmail.com> ([a4fd7c8](https://github.com/forkline/ingress-nginx/commit/a4fd7c80960270c5567e62427d01e6bafdcda640))
- Fixed boiler plate lint

Signed-off-by: James Strong <strong.james.e@gmail.com> ([22d9c35](https://github.com/forkline/ingress-nginx/commit/22d9c35edec8faeca4cf4f8b72e8f08da20c1618))
- Fix some comments

Signed-off-by: cui fliter <imcusg@gmail.com> ([82e836f](https://github.com/forkline/ingress-nginx/commit/82e836fbe89815ba57dcc1c8f9e370773d6bf291))
- Controller psp's volume config ([e3c9382](https://github.com/forkline/ingress-nginx/commit/e3c9382fc927cfbdfaa7a006b5ef2b0c332242d1))
- Fix controller tag in release

Signed-off-by: James Strong <james.strong@chainguard.dev> ([c3a22a2](https://github.com/forkline/ingress-nginx/commit/c3a22a219ac2826901bd6426e3fdf25519e53a64))
- Avoid builds and tests for changes to markdown ([0cb3dcf](https://github.com/forkline/ingress-nginx/commit/0cb3dcfd5c313305bf36924c300572dbd7f07fc1))
- Fix broken kubernetes.io/user-guide/ docs links ([4c00085](https://github.com/forkline/ingress-nginx/commit/4c00085c173566519f0509cd8c1b0a4aa0f81aa0))
- Add canary to sidebar in examples ([dd4a703](https://github.com/forkline/ingress-nginx/commit/dd4a703637f723c34081cd67d51dc5451a3a7771))
- Obsolete warnings ([30de599](https://github.com/forkline/ingress-nginx/commit/30de5999c15fb0dc1581950685f1e027957515d5))
- Fix gcloud builds

Signed-off-by: James Strong <strong.james.e@gmail.com> ([e8097d8](https://github.com/forkline/ingress-nginx/commit/e8097d8b8ff4aa4573a8d25e4bb98093323e6302))
- Fix deps sha

Signed-off-by: Jintao Zhang <zhangjintao9020@gmail.com> ([c83422f](https://github.com/forkline/ingress-nginx/commit/c83422fd6592b7b6aea2f58b5688929a5170875a))
- Add /etc/mime.types #10309 ([6b05e9b](https://github.com/forkline/ingress-nginx/commit/6b05e9b06e3a61062a33c6c8338f2e253c881aec))
- Update action file to auto release plugin #10197 ([8d0b00d](https://github.com/forkline/ingress-nginx/commit/8d0b00dd263b32f179460ab881095c3e95efbcbc))
- Path with sepecial characters warning #10281 #10308 ([c3a28ab](https://github.com/forkline/ingress-nginx/commit/c3a28ab45da2360ac5b590b640f6b5e6a20055b0))
- Remove curl on base container #9716 ([4664b74](https://github.com/forkline/ingress-nginx/commit/4664b741ff789817ed32923cf0391cb0a4bc77df))
- Fix path to faq.md in examples/rewrite/README.md ([fcda809](https://github.com/forkline/ingress-nginx/commit/fcda809ab034fffd14d8ff9b39441b2d12da420e))
- Fix brotli build issues

Signed-off-by: James Strong <strong.james.e@gmail.com> ([3a28016](https://github.com/forkline/ingress-nginx/commit/3a28016a6461d029253d9c6bb1c013f2d59efd5d))
- Validate x-forwarded-prefix annotation with RegexPathWithCapture ([9cdd51d](https://github.com/forkline/ingress-nginx/commit/9cdd51d5dcbf01fe6f06a8129027b07bf63fcb9b))
- Adjust unfulfillable validation check for session-cookie-samesite annotation ([13d95d0](https://github.com/forkline/ingress-nginx/commit/13d95d026ad8e2d7ccfb39252e801cc4cc99f479))
- Remove tcpproxy copy error handling ([44e550e](https://github.com/forkline/ingress-nginx/commit/44e550ea72f673fadeae0559a773feb9cbf3eec6))
- Disable cluster wide controller role permissions ([6c876bb](https://github.com/forkline/ingress-nginx/commit/6c876bba9ad90414c3ace6473c1ad0912db5d0d5))
- Fix OpenTelemtry image build

* fix OpenTelemtry image build

* use fpic ([8227888](https://github.com/forkline/ingress-nginx/commit/8227888ba0a9b09afb4514284ef89f331d4c60b9))
- Fix run command in dockerfile of test-runner-image ([760bf8e](https://github.com/forkline/ingress-nginx/commit/760bf8eb0c8033e71398a8c95f817e8ead5e54c6))
- Live-docs script ([c660f9e](https://github.com/forkline/ingress-nginx/commit/c660f9e3eb85cde56c0eddebb09852fcaceb42fe))
- Fix datasource, $exported_namespace variable in grafana nginx dashboard

* grafana/dashboards/nginx.json: re-add exported_namespace as a variable (was deleted entirely in #9523)

* dashboards/nginx.json: switch around ingress and namespace selectors, and rename "Exported Namespace" to "Ingress Namespace"

authored by tghartland at https://gist.github.com/tghartland/9147d88f991a95d4bab0fa7278c237eb

* dashboards/nginx.json: make "Ingress Request Volume" and "Ingress Success Rate" panels look at selected Ingress Namespaces only, and rename two panel titels to use the renamed variable

as suggested by tghartland in https://github.com/kubernetes/ingress-nginx/pull/9092#issuecomment-1285840900

* dashboards/nginx.json: apply Ingress Namespace selection to "Ingress Percentile Response Times and Transfer Rates" as well

this is from https://github.com/kubernetes/ingress-nginx/pull/9092#issuecomment-1287114743 also by tghartland ([1bc20da](https://github.com/forkline/ingress-nginx/commit/1bc20da92f02daba83725d06fe71582deca12d62))
- Fix geoip2 configuration docs ([c25b80c](https://github.com/forkline/ingress-nginx/commit/c25b80ca0067f0d4adf9dd879d2477d4830bff0f))
- Fixes brotli build issue ([25d2758](https://github.com/forkline/ingress-nginx/commit/25d2758e94268559f971e8b28ddcd980f5b6448a))
- Update kube version requirement to 1.21 ([95554dc](https://github.com/forkline/ingress-nginx/commit/95554dccd25a5b6debac9a850c0b84ec2f815fe8))
- Fix path in file changed detected message

* fix path in file changed detected message

Signed-off-by: Tom Hayward <thayward@infoblox.com>

* fix typo in log message

* explain code per review comments

---------

Signed-off-by: Tom Hayward <thayward@infoblox.com> ([48fbdfe](https://github.com/forkline/ingress-nginx/commit/48fbdfe3ba0c0e258890c970e2561caecea532dd))
- Fix ref error

Signed-off-by: James Strong <strong.james.e@gmail.com> ([b330f96](https://github.com/forkline/ingress-nginx/commit/b330f96482df33a837120b6151ee248abadad6bb))
- Fix for docs issue 11432 ([ec29659](https://github.com/forkline/ingress-nginx/commit/ec296594c788659b8c1427dba0956cef71768efa))
- Fixed fastcgi userguide ([8ca27e7](https://github.com/forkline/ingress-nginx/commit/8ca27e7ee902ec6b77c7cfb1749254cdbab36c24))
- Ensure changes in MatchCN annotation are detected ([bcb98c0](https://github.com/forkline/ingress-nginx/commit/bcb98c0c8dec15f796a67496b7f283508f957bbd))
- Fix formatting ([fddf4e0](https://github.com/forkline/ingress-nginx/commit/fddf4e034c0d766a4815247b414cbd71a1869456))
- Fix e2e tests for cgroups ([8621dfc](https://github.com/forkline/ingress-nginx/commit/8621dfc66dbe9c84b4e0bcae8916f44ef4d2620a))
- Fix lint errors ([ac0f6fc](https://github.com/forkline/ingress-nginx/commit/ac0f6fcd398065bd2ce54c655480a7ed32526127))
- Fix tests ([b2d67ff](https://github.com/forkline/ingress-nginx/commit/b2d67ff92bad02c041be12b7f02d66fc417a2e6a))
- Fix v1 test ([e9f3717](https://github.com/forkline/ingress-nginx/commit/e9f371787e82e11b6009b77d740e88edd8691bf8))
- Fix DNS issues with unresolvable backends with ExternalName

Co-authored-by: Pierre Ozoux <pierre@ozoux.net> ([12eecbe](https://github.com/forkline/ingress-nginx/commit/12eecbe471b77df177cb9cfc574404c87ee0db45))

### Documentation

- annotations: Explicit redirect status code ([f6a4307](https://github.com/forkline/ingress-nginx/commit/f6a430775c79b2d7a994ec630ab04a66007aa4cd))
- charts: Using helm-docs for chart ([71de8e1](https://github.com/forkline/ingress-nginx/commit/71de8e1a23fac0f3fc4c6cf87d9e454573604a69))
- deploy: Fix helm install command for helm v3 ([dbb0970](https://github.com/forkline/ingress-nginx/commit/dbb0970393183b4b112d83549b25847b9adf4c45))
- helm: Fix value key in readme for enabling certManager ([12da492](https://github.com/forkline/ingress-nginx/commit/12da492f0112a5833f6437e00a485c23562856b2))
- index: Fix typo helm value example ([b219555](https://github.com/forkline/ingress-nginx/commit/b219555e75684d6df91a03415507d3b310b831a5))
- tls: Add warning for not supporting TLSv1, TLSv1.1 ([a4de4de](https://github.com/forkline/ingress-nginx/commit/a4de4debecfd6fc96c564b286a6dee04f5398e3a))
- Use generic instead of OSS controller ([8c9e880](https://github.com/forkline/ingress-nginx/commit/8c9e88058cc64547fec198a68e7b1b574bdaed58))
- Add list of known annotations ([d871f44](https://github.com/forkline/ingress-nginx/commit/d871f4409bce5c0c10b838036b4f39575239a128))
- Fix app-root and clean up haproxy ([d4d369d](https://github.com/forkline/ingress-nginx/commit/d4d369d8c149ae4415a6774d9cf2aa5da2cbda8c))
- Update dead link ([a910661](https://github.com/forkline/ingress-nginx/commit/a910661aaf721762c8b7b612d36242d77c7c22b4))
- Fix typos and clarify wording ([854bf4b](https://github.com/forkline/ingress-nginx/commit/854bf4bb961d5be5290529c55837a5441541fd88))
- Remove duplicated section ([8bdb5e4](https://github.com/forkline/ingress-nginx/commit/8bdb5e42f21ce6d213ea20341d3f7a4c480db64e))
- Precisations on the usage of the InfluxDB module ([5c680ba](https://github.com/forkline/ingress-nginx/commit/5c680ba62973bc97ffe1a26e7329cbccd43cef89))
- Bare-metal considerations ([775b835](https://github.com/forkline/ingress-nginx/commit/775b8358c61628c374c7ea568f3c6c1b159c0a71))
- Add MetalLB and externalIPs to bare-metal deployment page ([82aad99](https://github.com/forkline/ingress-nginx/commit/82aad99da20d23373c6ef9965b4c087c74d187f2))
- Add docs for proxy-buffer-number ([81e4440](https://github.com/forkline/ingress-nginx/commit/81e4440bdb9b82ad8e6d39c9c82d82eeb29c1b8f))
- Reference buildx as a requirement for docker builds ([d131353](https://github.com/forkline/ingress-nginx/commit/d13135329aeaf1a4f83b44029f124f68270fa2b5))
- Fix use-gzip wrong markdown style ([12dddcc](https://github.com/forkline/ingress-nginx/commit/12dddcca179e786d57b7d6f2494675b919c8f7e3))
- Update development.md ([6545b1f](https://github.com/forkline/ingress-nginx/commit/6545b1fe69a835bf7b7d2664ed6cc24f138486b4))
- Update development.md to use ingress-nginx-* ([13b6532](https://github.com/forkline/ingress-nginx/commit/13b65323d5bce36dad26f391b9fe14e5f561f09a))
- Fix grammar and inconsistencies ([f41dbe1](https://github.com/forkline/ingress-nginx/commit/f41dbe14bc78eeb7e3ee3a88af20bd218695b6be))
- Move Azure deploy note to right item ([270f52a](https://github.com/forkline/ingress-nginx/commit/270f52ab21bcd3b4e4161f52ac6bb590c0e78bbd))
- Proxy-real-ip-cidr ([cdaf1bd](https://github.com/forkline/ingress-nginx/commit/cdaf1bdd84921cd46d5aded97c3b087350250983))
- Docs：update troubleshooting.md

* Update troubleshooting.md

Made the troubleshooting steps a bit more fluid IMHO.

* Update troubleshooting.md

Fixed introduced troubleshooting workflow change.

* Update troubleshooting.md

Fixed token path in new proposed workflow.

* Update troubleshooting.md

Fixed terminology (pod vs. container)

* Changed verb to get CLA refresh.

* Updating PR with requested changes.

Signed-off-by: Robert Jackson <robert@aztek.io> ([c4cc9a5](https://github.com/forkline/ingress-nginx/commit/c4cc9a504acbc17eded62911b56f6da44f0e94b9))
- Docs for migration to apiVersion networking.k8s.io/v1 ([1510c06](https://github.com/forkline/ingress-nginx/commit/1510c06045ece4e199ebec85e7ec90cf15e19747))
- Clarify default-backend behavior ([f84006d](https://github.com/forkline/ingress-nginx/commit/f84006d62f796621532948e66c51ababe6bc093c))
- Remove extra symbol ([0963a45](https://github.com/forkline/ingress-nginx/commit/0963a459220e08be1d0ec956d4ef908441958649))
- Docs_multiple_instances_one_cluster_ticket_7543 ([2ff5af0](https://github.com/forkline/ingress-nginx/commit/2ff5af08d425d9b20744c333f4d68a7781c0b86b))
- Fix typo'd executible name ([d6284d1](https://github.com/forkline/ingress-nginx/commit/d6284d16728063c870aa602e0585999607acbcaa))
- Correct typo ([feba7e1](https://github.com/forkline/ingress-nginx/commit/feba7e1ffc7f2047f1cae5cb5f093f5b34c58abd))
- Fix inconsistent controller annotation ([6eecefd](https://github.com/forkline/ingress-nginx/commit/6eecefd3dab9e1784be609e3e83af0418a216539))
- Fix changelog formatting ([3bd3231](https://github.com/forkline/ingress-nginx/commit/3bd32316babefb3f04827154dcdda593803516b8))
- Comment out TODO heading on the deployment page ([b0d1982](https://github.com/forkline/ingress-nginx/commit/b0d198252fb74127eecc54f09100ad0a9ad3adce))
- Updated the content of deploy/rbac.md ([c3ea3b8](https://github.com/forkline/ingress-nginx/commit/c3ea3b861e04db08a9edb67d25fc96502b32a657))
- Canary weighted deployments example ([adbad99](https://github.com/forkline/ingress-nginx/commit/adbad99a71234f1f570133c36c1961c7dd383e94))
- Add lua testing documentation ([388987c](https://github.com/forkline/ingress-nginx/commit/388987c4e7500478ce76ff0c845d43256b603a00))
- Add netlify configuration ([f1e3f2f](https://github.com/forkline/ingress-nginx/commit/f1e3f2fa3c38a8a5473bd253175c3f6525bc3f53))
- Change Dockefile url ref main ([686aeac](https://github.com/forkline/ingress-nginx/commit/686aeac5961f37eaf1ddfa2fa320df4ccf0cf005))
- Swap explanation to match example ([d9baff9](https://github.com/forkline/ingress-nginx/commit/d9baff90d7d9fa451331882acb6af3a6c0cde14a))
- Update configmap docs for enable-global-auth option ([cd3e5d3](https://github.com/forkline/ingress-nginx/commit/cd3e5d323d026b810e922a5f0e79679f56d370ce))
- Add index for global-auth-always-set-cookie ([7f723c5](https://github.com/forkline/ingress-nginx/commit/7f723c59855e82614582ff7b2efd1783b1afc2ee))
- Update annotations docs with missing session-cookie section ([c295cd1](https://github.com/forkline/ingress-nginx/commit/c295cd1c4bdb385f0a97371a4f4fad6470b39751))
- Add vouch-proxy OAuth example ([7d75abb](https://github.com/forkline/ingress-nginx/commit/7d75abb0ffab8f8d10445ef38ecd202477ed6267))
- Update the 404 link to FAQ ([a302cc5](https://github.com/forkline/ingress-nginx/commit/a302cc5ccaa62ee6c47a965240dbba98879ebf96))
- Update Ingress-NGINX v1.10.1 compatibility with Kubernetes v1.30 ([8691884](https://github.com/forkline/ingress-nginx/commit/8691884033b60030b39b806269adcddc3a90716a))
- Update OpenSSL Roadmap link ([e4986a7](https://github.com/forkline/ingress-nginx/commit/e4986a74cd14119b2e16262e74f9080722b151e6))
- Add deployment for AWS NLB Proxy. ([b795512](https://github.com/forkline/ingress-nginx/commit/b79551287e42ec15306fecc31f3d8dee23b7f840))

### Add

- Documentation for proxy-ssl-location-only ([bc79fe1](https://github.com/forkline/ingress-nginx/commit/bc79fe1532c39f318d373ee9d5c548278ad4ec19))

### Added

- Support for http header passing from external authentication service response ([302fa5f](https://github.com/forkline/ingress-nginx/commit/302fa5f4bb53ba048695e7e5fed545f014419db4))
- Example of external service response headers propagation (requested by @aledbf) ([5ee1eed](https://github.com/forkline/ingress-nginx/commit/5ee1eed434292a5b4f1f38f4bcfb77a15e6d1271))

### Annotations

- Allow commas in URLs. ([d4c4911](https://github.com/forkline/ingress-nginx/commit/d4c49112a4a5abdcced93bf34c589147b594da34))
- Reload on custom header changes. ([29d1e20](https://github.com/forkline/ingress-nginx/commit/29d1e2014bffb1cbb0f1e4e045f3ad285f8a0c5f))
- Deny newlines. ([698c3c0](https://github.com/forkline/ingress-nginx/commit/698c3c0df104e5358fa7ff6871ef246fdd6ff44e))
- Allow ciphers with underscores. ([8c1ecd7](https://github.com/forkline/ingress-nginx/commit/8c1ecd7655bd052a26e64d3361dede3096cd80c6))
- Quote auth proxy headers. ([3d90678](https://github.com/forkline/ingress-nginx/commit/3d90678bfe2c78c188c26e4ac5b0cf6779f8b39f))
- Fix log format. ([57e6899](https://github.com/forkline/ingress-nginx/commit/57e6899d86751519e9526eb9f68b296bbc29d9b0))
- Respect changes to `auth-proxy-set-headers`. ([ea86daa](https://github.com/forkline/ingress-nginx/commit/ea86daa7dfa3a7fa5f895212f6291103468197ab))
- Add `^` and `$` to auth method regex. ([a24822a](https://github.com/forkline/ingress-nginx/commit/a24822a8f2a2f2d894f63e809047c042c78b3377))
- Use dedicated regular expression for `proxy-cookie-domain`. ([ded3b60](https://github.com/forkline/ingress-nginx/commit/ded3b6071484147e20adfb900ee1e30be34f2924))
- Consider aliases in risk evaluation. ([3a9889b](https://github.com/forkline/ingress-nginx/commit/3a9889ba5bdf598c90446be415ca8d3a5ad81790))

### Annotations/AuthTLS

- Allow named redirects. ([2fb7fd2](https://github.com/forkline/ingress-nginx/commit/2fb7fd2e5c72b509b5c2c3a8f8f8a035b2016874))

### Bugfix

- Fix incomplete log ([b65ceee](https://github.com/forkline/ingress-nginx/commit/b65ceee1a865dfacb89c00f8fd7f561c45127b42))
- Non-host canary ingress use default server name as host to merge ([b6dc384](https://github.com/forkline/ingress-nginx/commit/b6dc384afb14a9eeca4abfc41918325909643d8f))

### Build

- Build luajit with debug symbols ([1a84719](https://github.com/forkline/ingress-nginx/commit/1a84719e0b42782db96af9a07e40c16aa1c7be1f))
- Build lua-bridge-tracer ([f4b3174](https://github.com/forkline/ingress-nginx/commit/f4b31743c07504b37745d8493910859310f02ca4))
- Remove unnecessary tag line in e2e ([5c65cf4](https://github.com/forkline/ingress-nginx/commit/5c65cf498e19d021a2832ecdf9c94a9bf7c06ad5))
- Remove docker version check ([714637b](https://github.com/forkline/ingress-nginx/commit/714637bec598202ae54c218eef3529b4bf239734))
- Build yaml-cpp lib in image builder ([5794a93](https://github.com/forkline/ingress-nginx/commit/5794a9360ac1c7e435edfb33379858ee09f79578))
- Build release 1.6.1 image

Signed-off-by: James Strong <james.strong@chainguard.dev> ([90c857d](https://github.com/forkline/ingress-nginx/commit/90c857d8e59124fc600be89ac80ee3fc6e9d930a))
- Build 1.6.2 to fix #9569

Signed-off-by: James Strong <james.strong@chainguard.dev> ([3348a60](https://github.com/forkline/ingress-nginx/commit/3348a6038edb72e19681ce273148d26ba54abea9))
- Build release 1.6.1 image

Signed-off-by: James Strong <james.strong@chainguard.dev> ([fb2223b](https://github.com/forkline/ingress-nginx/commit/fb2223b12c70cdf746349393a7f44b753576b49c))
- Build 1.6.2 to fix ([7782c70](https://github.com/forkline/ingress-nginx/commit/7782c70030bc5cbe9f981fc4126c31dd715f9427))
- Always use local `tmp` dir on macOS. ([59a0da7](https://github.com/forkline/ingress-nginx/commit/59a0da769de6f8a88852060321136c41dd571ef9))

### CI

- ci: Replace upstream automation with fork-owned workflows ([e24441b](https://github.com/forkline/ingress-nginx/commit/e24441b70bafbe61a0836ef588d682ec7f2ff054))
- helm: Fix Helm Chart release action 422 error ([868df87](https://github.com/forkline/ingress-nginx/commit/868df87bb31478fe239e3cd426eb6c1d4bb3ae04))
- Update kubernetes versions ([45d2452](https://github.com/forkline/ingress-nginx/commit/45d245237e83c90dff83b6bf3410730bf39309f6))
- Remove setup-kind step ([8736b3b](https://github.com/forkline/ingress-nginx/commit/8736b3b9a76ef9d499f5cd69886aa85b4c927248))
- Remove setup-helm step ([4f528fc](https://github.com/forkline/ingress-nginx/commit/4f528fc70a3bfcc2aff1614234891b27aebc1608))
- Replace `chart-testing` image by `e2e-test-runner`. ([648cb8b](https://github.com/forkline/ingress-nginx/commit/648cb8bb0a3934312250435cb31f9cc9a10956ea))
- Bump forgotten Ginkgo versions. ([9ca96df](https://github.com/forkline/ingress-nginx/commit/9ca96df6af0ca1897416a26d27f8ca5386ef5027))
- Grant checks write permissions to E2E Test Report. ([1c0a3dd](https://github.com/forkline/ingress-nginx/commit/1c0a3ddf0369df411c05696186ca769d84e178d9))
- Fix chart testing. ([6608eb2](https://github.com/forkline/ingress-nginx/commit/6608eb23b022b8840f65ffcc3576c7c68719489f))
- Update KIND images. ([76f90ec](https://github.com/forkline/ingress-nginx/commit/76f90ec8cf67f6cc6650865a1916aa2581ea205f))
- Update KIND images. ([62b97c7](https://github.com/forkline/ingress-nginx/commit/62b97c7b994761e5effdcff79b0e5d57ff426531))
- Update `kubectl` to v1.31.5. ([8e58582](https://github.com/forkline/ingress-nginx/commit/8e58582ddc04d11ecba557b0dace2ab0d56d5453))
- Update Artifact Hub to v1.20.0. ([90eb6aa](https://github.com/forkline/ingress-nginx/commit/90eb6aac4f3fb6cd5edd84999bce4c96e8eaff58))
- Update `kubectl` to v1.32.2. ([7ed3578](https://github.com/forkline/ingress-nginx/commit/7ed3578b6166a5c91fef93b34fa759975d9e2a2e))
- Update KIND images. ([bbbf1d4](https://github.com/forkline/ingress-nginx/commit/bbbf1d41469c5370be25bc9617e129f0aae169ce))
- Update Kubernetes to v1.32.3. ([6e1858c](https://github.com/forkline/ingress-nginx/commit/6e1858c95c33a40904992d4cc07d62c94eeaba44))
- Update KIND to v1.32.3. ([c7fb821](https://github.com/forkline/ingress-nginx/commit/c7fb8213e31ee3d4bacd1acd3f09cb02830c83fa))
- Do not fail fast. ([ad9200e](https://github.com/forkline/ingress-nginx/commit/ad9200eae3d55875ce62caf955c6dce8b817639f))
- Update Kubernetes to v1.32.4. ([166b197](https://github.com/forkline/ingress-nginx/commit/166b197b1d64474a764cb7a3db019ba325872bf6))
- Update Kubernetes. ([1d71617](https://github.com/forkline/ingress-nginx/commit/1d7161791dcb8f1ec7c9e4f909719981b6f2ca5f))
- Update Kubernetes to v1.33.2. ([8da05a5](https://github.com/forkline/ingress-nginx/commit/8da05a5a107a3596778ad46f4ea209fcb99ec062))
- Update Kubernetes to v1.33.3. ([abd6138](https://github.com/forkline/ingress-nginx/commit/abd6138f96f7ab327d918525e8708f0018d44336))
- Update KIND to v1.33.2. ([ee00f72](https://github.com/forkline/ingress-nginx/commit/ee00f72bc027ae9d679299ee37d1209c01c22305))
- Fix typo. ([15ee736](https://github.com/forkline/ingress-nginx/commit/15ee7365970f7b2e932b39d4be5f23d9e09d9c67))
- Update Kubernetes to v1.33.4. ([d6aa282](https://github.com/forkline/ingress-nginx/commit/d6aa282926979ab5041d567acc4ae9e9a88d55ca))
- Update KIND to v1.34.0. ([8b87b02](https://github.com/forkline/ingress-nginx/commit/8b87b02ba4340ae1d022399a436aae8a537dad91))
- Update Kubernetes to v1.34.0. ([d8fbe21](https://github.com/forkline/ingress-nginx/commit/d8fbe217ce9db15a33b63da8d8e927a32ae1cb98))
- Update Helm to v3.18.6. ([47dac2d](https://github.com/forkline/ingress-nginx/commit/47dac2d918e4630d4166c3f6c201e9b0f8ac040a))
- Update Kubernetes to v1.34.1. ([d10c78b](https://github.com/forkline/ingress-nginx/commit/d10c78bc59b374a0bef007dd4097540dd13b4ffe))
- Update Helm to v3.19.0. ([84637cf](https://github.com/forkline/ingress-nginx/commit/84637cfd3eaaed9e528b168009be976ff43f4c3f))
- Update Helm to v3.19.1. ([fe1f1c5](https://github.com/forkline/ingress-nginx/commit/fe1f1c5ff31ac7c8f19688f24510b09dbedd812c))
- Update Kubernetes to v1.34.2. ([c6c5602](https://github.com/forkline/ingress-nginx/commit/c6c5602c48eaf7046930bf9ec0ce00e016d9dc20))
- Update Helm to v3.19.2. ([bcc3a2d](https://github.com/forkline/ingress-nginx/commit/bcc3a2d1a7bd9613637015574bccd9a5d49529d5))
- Pin Helm version. ([3e9820e](https://github.com/forkline/ingress-nginx/commit/3e9820eece14a23674ff43604317390d67cee750))
- Update Helm to v4.0.1. ([1f50c2e](https://github.com/forkline/ingress-nginx/commit/1f50c2e63e4184dc12cb90d826f5cf95406303f7))
- Disable verification for Helm Unit Test. ([8c30e82](https://github.com/forkline/ingress-nginx/commit/8c30e82edcee1caddde70eb6d564519da3ff94ab))
- Update Kubernetes to v1.34.3. ([59c0bbe](https://github.com/forkline/ingress-nginx/commit/59c0bbecc1ef991ba572e18a4f529b51c3c311f0))
- Update Helm to v4.0.2. ([06f26b6](https://github.com/forkline/ingress-nginx/commit/06f26b6b331d33e8da68cefd600fc6e68968b702))
- Update KIND to v1.34.2. ([cfe539a](https://github.com/forkline/ingress-nginx/commit/cfe539a907d1985cd538e418d73a27225cf6e11d))
- Update Helm to v4.0.4. ([2533b10](https://github.com/forkline/ingress-nginx/commit/2533b1007e6ee88d535fbebbcc60a31c1b1e3d15))
- Update KIND to v1.34.3. ([6dfad97](https://github.com/forkline/ingress-nginx/commit/6dfad976969dfd6f45fe089108e09d498538c3d6))
- Update Kubernetes to v1.35.0. ([3126e6f](https://github.com/forkline/ingress-nginx/commit/3126e6f89b7be5992ac3fe811072f815603dc378))
- Update Helm to v4.0.5. ([0966d0f](https://github.com/forkline/ingress-nginx/commit/0966d0f2f4f513d46c54d26e7beedfefc3041671))
- Update Helm to v4.1.0. ([3980fc6](https://github.com/forkline/ingress-nginx/commit/3980fc664469595d5fbd5389ceaaa3fefbea2301))
- Update Helm to v4.1.1. ([72d99b3](https://github.com/forkline/ingress-nginx/commit/72d99b3719f1736563ee11ce98d8ee87254c44f3))
- Update Kubernetes to v1.35.1. ([296ce3b](https://github.com/forkline/ingress-nginx/commit/296ce3bec2f6b6c01815c651af0a5a0f36fd9f07))
- Update KIND to v1.35.1. ([349785f](https://github.com/forkline/ingress-nginx/commit/349785f92b96537d2752f06fc2a209deb9729d48))
- Update Kubernetes to v1.35.2. ([ab6c151](https://github.com/forkline/ingress-nginx/commit/ab6c151c36414ac434bed41f0a075be91f32aeb8))
- Update Helm to v4.1.3. ([6b74b11](https://github.com/forkline/ingress-nginx/commit/6b74b11670ec6db29b077f5a0e357b56b4f40697))
- Update Kubernetes to v1.35.3. ([52f3f41](https://github.com/forkline/ingress-nginx/commit/52f3f412859e35d5a881be9b216f8fa24dfb5a2d))

### Change

- ToYaml to range ([826da96](https://github.com/forkline/ingress-nginx/commit/826da966fcc211823e51de506ba23d57c18a74b7))

### Changelog.md

- Update references to sigs.k8s.io/promo-tools ([c2fe736](https://github.com/forkline/ingress-nginx/commit/c2fe736d48a9733ba766615dbd8e64f19dcf1223))

### Chart

- Drop `controller.headers`, rework DH param secret. ([58e5a2c](https://github.com/forkline/ingress-nginx/commit/58e5a2c01f72bd1e4d9e924d5059b29350577d76))
- Improve `README.md`. ([7fcafff](https://github.com/forkline/ingress-nginx/commit/7fcafff04617528619950b1cef190a3b424e4f2b))
- Rework network policies. ([0b0ce03](https://github.com/forkline/ingress-nginx/commit/0b0ce031ac11b2738f85f42d38fdd1fe10fde120))
- Improve #10539. ([9cb3919](https://github.com/forkline/ingress-nginx/commit/9cb3919e848134d1ece70b6cd38542fd6b8a9b46))
- Fix pod selectors in `NOTES.txt`. ([6499a6b](https://github.com/forkline/ingress-nginx/commit/6499a6bd0431ee590234f205ca482a4ec7a84828))
- Tighten `securityContext`s and Pod Security Policies. ([8b026f4](https://github.com/forkline/ingress-nginx/commit/8b026f42d5db4e1da910a4c0e49e9511d239ba4f))
- Promote myself to approver & reviewer. ([e6d3bbb](https://github.com/forkline/ingress-nginx/commit/e6d3bbb520c2d902372a5215c9f1d2c68ba44c08))
- Put me in alphabetical order. ([f3f0ee5](https://github.com/forkline/ingress-nginx/commit/f3f0ee539df34b57bf9d88f0019f5bb1ec5e927a))
- Rename `changelog.md.gotmpl` into `changelog/helm-chart.md.gotmpl`. ([559c03d](https://github.com/forkline/ingress-nginx/commit/559c03d1d379542e88b09997300081c5299356ef))
- Improve `changelog/helm-chart.md.gotmpl`. ([84ced1e](https://github.com/forkline/ingress-nginx/commit/84ced1ed1c4018eba38a96066e9f2506a9cc31bf))
- Rename `changelog/Changelog-*.md` into `changelog/helm-chart-*.md`. ([b8e4e3c](https://github.com/forkline/ingress-nginx/commit/b8e4e3cebabd84a08c418691a596c1d30b28bc02))
- Split `CHANGELOG.md` into `changelog/helm-chart-*.md`. ([7b9e356](https://github.com/forkline/ingress-nginx/commit/7b9e3566f79f9519977958e8cd1afe8f913515b2))
- Simplify image templating. ([815a1c5](https://github.com/forkline/ingress-nginx/commit/815a1c56a99e011549a5ba308dda89a925e547d0))
- Improve #10673. ([2f7f4d7](https://github.com/forkline/ingress-nginx/commit/2f7f4d70eb29e71ff28c6465f213cf6893137d24))
- Revert verion `4.8.4`. ([0e12525](https://github.com/forkline/ingress-nginx/commit/0e12525bdd74e9145e861aafcd3ce6fead59b94b))
- Add Gacko to maintainers. ([9de651a](https://github.com/forkline/ingress-nginx/commit/9de651aa7d393e4fc797d99f7cfc81fc6d20c649))
- Remove useless `default` from `_params.tpl`. ([48b9831](https://github.com/forkline/ingress-nginx/commit/48b98311229913803063ed34a738fa1f8fd7c012))
- Set `--enable-metrics` depending on `controller.metrics.enabled`. ([3e740fe](https://github.com/forkline/ingress-nginx/commit/3e740fe8e747eb37c35abb64e1d6f5db42e5b15c))
- Deploy `PodDisruptionBudget` with KEDA. ([b5c4476](https://github.com/forkline/ingress-nginx/commit/b5c447612c48e9d517e5c6dd4ae0f0a176203cde))
- Improve IngressClass documentation. ([2894b8a](https://github.com/forkline/ingress-nginx/commit/2894b8a060fd50a3b991f9f084e0cbc5bae4133c))
- Add Gacko to maintainers. Again. ([aa5deed](https://github.com/forkline/ingress-nginx/commit/aa5deedae3dc9d55fd12f482d92c1cb86337cc0a))
- Align HPA & KEDA conditions. ([9480cde](https://github.com/forkline/ingress-nginx/commit/9480cde724d167c3246563586b5aaae86d26de00))
- Render `controller.ingressClassResource.parameters` natively. ([112b9bb](https://github.com/forkline/ingress-nginx/commit/112b9bb028da76b81c98387bacb4841a78c67e32))
- Add IngressClass aliases. ([56a0968](https://github.com/forkline/ingress-nginx/commit/56a0968675d8ed2599bf9ae96c82b0a64f4dbb3d))
- Make `controller.config` templatable. ([ad274ab](https://github.com/forkline/ingress-nginx/commit/ad274ab2c62604a8d6b950793536df97f214a18f))
- Add unit tests for default backend & topology spread constraints. ([531b007](https://github.com/forkline/ingress-nginx/commit/531b007b604afce3fee8fa23a96dc5626401f762))
- Remove `controller.enableWorkerSerialReloads`. ([987039c](https://github.com/forkline/ingress-nginx/commit/987039c0146cd1a90b10dc0fb5b2489ab6015b2c))
- Make admission webhook patch job RBAC configurable. ([0c17748](https://github.com/forkline/ingress-nginx/commit/0c17748c4412b5f6ac8532586ed41accb0471c74))
- Fix `IngressClass` annotations. ([90ef458](https://github.com/forkline/ingress-nginx/commit/90ef45852ca4bf0db53c508c098ff265f2f9fb03))
- Make pod affinity templatable. ([af9e524](https://github.com/forkline/ingress-nginx/commit/af9e5246ad09a304ef68d11f1467185dbd12fe72))
- Explicitly set `runAsGroup`. ([36df47f](https://github.com/forkline/ingress-nginx/commit/36df47fcc4be503a5db028dd2d8e6d9475353ad2))
- Remove `isControllerTagValid`. ([e972a35](https://github.com/forkline/ingress-nginx/commit/e972a35e98d476171746629dec9ddc3ea4d3be05))
- Bump Kube Webhook CertGen & OpenTelemetry. ([593f05e](https://github.com/forkline/ingress-nginx/commit/593f05ed571fdd69e06504d13c4f15308643ad66))
- Use generic values for `ConfigMap` test. ([f6595f5](https://github.com/forkline/ingress-nginx/commit/f6595f554abb6692f15296d75cc1f764e08679b8))
- Add tests for `PrometheusRule` & `ServiceMonitor`. ([5d457c7](https://github.com/forkline/ingress-nginx/commit/5d457c7daa377b404d3d5bc2dbcfca477525ccfa))
- Add `controller.metrics.prometheusRule.annotations`. ([3cde777](https://github.com/forkline/ingress-nginx/commit/3cde7770dd5bdc582d3bfc7e1307a04b6a80c094))
- Implement `controller.admissionWebhooks.service.servicePort`. ([a647bc1](https://github.com/forkline/ingress-nginx/commit/a647bc1b7aca33171809b143e96b35da1c989b79))
- Improve default backend service account. ([61f56cb](https://github.com/forkline/ingress-nginx/commit/61f56cb490d6cdc9c8e89ef3040e7eb609e88c2a))
- Remove Pod Security Policy. ([0276039](https://github.com/forkline/ingress-nginx/commit/027603927b46f086c1b77dc1ab76f0667343718b))
- Align default backend `PodDisruptionBudget`. ([435d536](https://github.com/forkline/ingress-nginx/commit/435d5365b4c37ba0f9e11273573e18fdbe6f93cb))
- Test `controller.minAvailable` & `controller.maxUnavailable`. ([b2bc961](https://github.com/forkline/ingress-nginx/commit/b2bc9618d3aaae3234fe30d36a937b2f7488357a))
- Add `defaultBackend.maxUnavailable`. ([43a7d8d](https://github.com/forkline/ingress-nginx/commit/43a7d8d5fe90ecbccadb956f6592928691eb0881))
- Implement `unhealthyPodEvictionPolicy`. ([17209eb](https://github.com/forkline/ingress-nginx/commit/17209eb3a93bf2d3c08c639125f050070e53acd2))
- Add `controller.progressDeadlineSeconds`. ([7b8d293](https://github.com/forkline/ingress-nginx/commit/7b8d293d9bb7d1e78908c9a762460dbf64c01157))
- Extend image tests. ([24a9f97](https://github.com/forkline/ingress-nginx/commit/24a9f972ff3fe3a9a1c05736a4b97641327302bb))
- Improve CI. ([f369ffb](https://github.com/forkline/ingress-nginx/commit/f369ffb0734cc63fdf85959c4e41be0606ca8f1e))
- Add `global.image.registry`. ([45fc886](https://github.com/forkline/ingress-nginx/commit/45fc8860cfc1012c1025dbb2c474c34f3aa9b0f9))
- Add `controller.metrics.service.enabled`. ([f3bfa56](https://github.com/forkline/ingress-nginx/commit/f3bfa56c61446e1de7423038cc551c3e8eaa237b))
- Bump Kube Webhook CertGen. ([657393e](https://github.com/forkline/ingress-nginx/commit/657393e7b3b94e41f1b66081d9a5323e7eb097dc))
- Suggest `matchLabelKeys` in Topology Spread Constraints. ([0edf16f](https://github.com/forkline/ingress-nginx/commit/0edf16ff6bff89bd61750c38558b3bf801ec5ced))
- Add ServiceAccount tests. ([bd76cf8](https://github.com/forkline/ingress-nginx/commit/bd76cf8f05bb48e36e76eee5a333691499379be4))
- Set `automountServiceAccountToken` in workloads. ([e07f0f6](https://github.com/forkline/ingress-nginx/commit/e07f0f6890b79c0eb9dc7f541b4ce05a868b4c74))
- Rework ServiceMonitor. ([d0a0430](https://github.com/forkline/ingress-nginx/commit/d0a04308c87d46aa0bb28a85af888da06ae8cf88))
- Implement ServiceMonitor limits. ([260976b](https://github.com/forkline/ingress-nginx/commit/260976b8d8f1497acbee9eb8ae410daf326064b7))
- Add service cluster IPs. ([be8abe7](https://github.com/forkline/ingress-nginx/commit/be8abe7a5cf701b8e00502e39166cf652076e378))
- Bump Kube Webhook CertGen. ([29513e8](https://github.com/forkline/ingress-nginx/commit/29513e8564a9a7efd0bf2dfdc86fd2eece190185))
- Add `controller.service.trafficDistribution`. ([506ded7](https://github.com/forkline/ingress-nginx/commit/506ded73b1c69b62757a176d606c65ba9f69d711))
- Bump Kube Webhook CertGen. ([f246d43](https://github.com/forkline/ingress-nginx/commit/f246d43e4c280793b406448641e96ab2e01524b6))
- Add `controller.service.external.labels` & `controller.service.internal.labels`. ([fe91e8e](https://github.com/forkline/ingress-nginx/commit/fe91e8e4220a7efe43dbb20990e3dfd817cdb8f3))
- Add `controller.admissionWebhooks.certManager.*.revisionHistoryLimit`. ([f333b0b](https://github.com/forkline/ingress-nginx/commit/f333b0bdfcdf9b2c284aa51cca36967c0ae8a37c))
- Bump Kube Webhook CertGen. ([3a31ad2](https://github.com/forkline/ingress-nginx/commit/3a31ad2b026967fd4c1707fb76708236eb2b22ed))
- Bump Kube Webhook CertGen. ([b6fc9d4](https://github.com/forkline/ingress-nginx/commit/b6fc9d4490539c96ef1941831303b24b05bd95f0))
- Implement `runtimeClassName`. ([551c0c4](https://github.com/forkline/ingress-nginx/commit/551c0c42012c1dd4efab96c1d9c26b3a35143422))
- Remove validation for removed API. ([30be84a](https://github.com/forkline/ingress-nginx/commit/30be84a20af4c3bf289322607142730bfabb01f4))
- Bump Kube Webhook CertGen. ([666f7d2](https://github.com/forkline/ingress-nginx/commit/666f7d2aecb38e8503171b328afae83b668a6ad9))
- Add `activeDeadlineSeconds`. ([6338a3a](https://github.com/forkline/ingress-nginx/commit/6338a3ac71f5b5e38971d81dd80a49ca51e563d4))
- Bump Kube Webhook CertGen. ([2a467c5](https://github.com/forkline/ingress-nginx/commit/2a467c5419e6094abf34bdb94b821877787897bf))
- Remove trailing whitespace. ([8105f08](https://github.com/forkline/ingress-nginx/commit/8105f0861fbd69558e58c52f93540c81edad3a4f))
- Template default backend extra volumes. ([f659cdf](https://github.com/forkline/ingress-nginx/commit/f659cdf892d77b40790f2566abe281ab37f9bd96))
- Push to OCI registry. ([81c3267](https://github.com/forkline/ingress-nginx/commit/81c3267b32852c5a1edfacdc4877533b17252859))
- Bump Kube Webhook CertGen. ([37f344b](https://github.com/forkline/ingress-nginx/commit/37f344bed3e6ee7606e638f69ea8e8bc2a26f39a))
- Add volumes for webhook patch job. ([203b97a](https://github.com/forkline/ingress-nginx/commit/203b97ae13118f9b61ca74ec3ed50b8df5dfbe04))
- Bump Kube Webhook CertGen. ([02adf6a](https://github.com/forkline/ingress-nginx/commit/02adf6a380137281dd707dc8a657d8b1c4a43516))
- Add resize policy. ([a031a08](https://github.com/forkline/ingress-nginx/commit/a031a0893bcb777400a90cf189647f48b90bf6e0))
- Bump Kube Webhook CertGen. ([7678cac](https://github.com/forkline/ingress-nginx/commit/7678cacd7ade167c778326be49f5c2001875d228))
- Add `controller.metrics.serviceMonitor.scrapeTimeout`. ([85806f9](https://github.com/forkline/ingress-nginx/commit/85806f957c8a4c4d845163f356ed56be9bdedda8))
- Make extra init containers templatable. ([cd3c6fb](https://github.com/forkline/ingress-nginx/commit/cd3c6fb941a58eec9f388a4acbfa3f13b28f2260))

### Chore

- build: Fix Run make dev-env syntax error ([46d87d3](https://github.com/forkline/ingress-nginx/commit/46d87d3462014d7e188bff0082f0c4de044d9bc0))
- dep: Upgrade github.com/emicklei/go-restful/v3 to 3.10 ([e97e928](https://github.com/forkline/ingress-nginx/commit/e97e9285428bf78b3ddc334c1f670636fd7c2427))
- dep: Change lua-resty-cookie's repo ([0a054d1](https://github.com/forkline/ingress-nginx/commit/0a054d1f586d030b5d7d2b4f585a0725962277f6))
- deps: Upgrade headers-more module to 0.37 ([e78af97](https://github.com/forkline/ingress-nginx/commit/e78af97ecdf10187b43786df216478df6ea440d1))
- deps: Group update k8s.io packages to v0.30.0 ([0308931](https://github.com/forkline/ingress-nginx/commit/030893148be3a2ea655b9fc014ca24c7756d233d))
- Add test to internal ingress resolver pkg ([7a533f0](https://github.com/forkline/ingress-nginx/commit/7a533f035b0888763e5f104050298c94152b737f))
- Add Artifact Hub lint ([40c69a1](https://github.com/forkline/ingress-nginx/commit/40c69a1ef706a76c6b47d317003965ba9186cde3))
- V1.2.0-beta.0 release ([e86e7ee](https://github.com/forkline/ingress-nginx/commit/e86e7eebf551d85b8ac1d0b91a87d4810d89fcf6))
- Release v1.2.1 ([c32f9a4](https://github.com/forkline/ingress-nginx/commit/c32f9a43279425920c41ba2e54dfcb1a54c0daf7))
- Remove stable.txt from release process ([fb3b288](https://github.com/forkline/ingress-nginx/commit/fb3b2882c25caa96384c49bdb8e501e7364ddd99))
- Start v1.3.0 release process ([caac91c](https://github.com/forkline/ingress-nginx/commit/caac91ce662f0fac188bc466874172c037eeabcb))
- Bump NGINX version v1.21.4 ([bf8362c](https://github.com/forkline/ingress-nginx/commit/bf8362cb500b413d04d04b1312df1914b4b193f8))
- Update NGINX to 1.21.6 ([0b5e068](https://github.com/forkline/ingress-nginx/commit/0b5e0685112e4537ee20a0bdbba451e9f6158aa3))
- Create httpbun image ([6d91c2a](https://github.com/forkline/ingress-nginx/commit/6d91c2a54c6e54c987078f6e01868385e3308ee6))
- Update httpbin to httpbun ([0bdb643](https://github.com/forkline/ingress-nginx/commit/0bdb64373c48660192dc87fe17989ab574422932))
- Start v1.8.0 release process ([0e94cc1](https://github.com/forkline/ingress-nginx/commit/0e94cc1b234c62fbc3a37ba1b991c2accb5e6959))
- Pkg imported more than once ([114ae77](https://github.com/forkline/ingress-nginx/commit/114ae77fb7ca30cb97326f41a18a1cb75ecc5665))
- Move httpbun to be part of framework ([60bf6ba](https://github.com/forkline/ingress-nginx/commit/60bf6ba6429feb4a5d8b485d1bcee87187fbc368))
- Remove echo friom canary tests ([1eeabe9](https://github.com/forkline/ingress-nginx/commit/1eeabe97b5cf8ad04d25f6b435579dce388da4bd))
- Remove echo from snippet tests ([f8bf5a3](https://github.com/forkline/ingress-nginx/commit/f8bf5a3086fc52114040a60335e2f483e819a14c))
- Bump OpenResty to v1.21.4.2 ([6416ed8](https://github.com/forkline/ingress-nginx/commit/6416ed821d726d78d8988ed1ae189bc243903471))
- Fix function names in comment ([e5b6636](https://github.com/forkline/ingress-nginx/commit/e5b663690300180d0c3b0b349afaa89fbe5dd19d))

### Chores

- Remove recently added whitespaces. ([365d886](https://github.com/forkline/ingress-nginx/commit/365d886c1d8db4c6e4cc543f5ee08adc5480d0d1))
- Align security contacts & chart maintainers to actual owners. ([e084ad0](https://github.com/forkline/ingress-nginx/commit/e084ad0a5eefe7ccfa80a0f4eac69f923efd0374))
- Migrate deprecated `wait.Poll*` to context-aware equivalents. ([02b3870](https://github.com/forkline/ingress-nginx/commit/02b3870f0715e73d0b8690697daac55151617553))

### Config

- Fix panic on invalid `lua-shared-dict`. ([ac23d40](https://github.com/forkline/ingress-nginx/commit/ac23d4069befefdb5260d3502608063459fb9b92))
- Remove notes about future defaults. ([683c203](https://github.com/forkline/ingress-nginx/commit/683c203df4db69f11426a9472a6e8c5abc8ea676))
- Use stronger ciphers first. ([66d8d6c](https://github.com/forkline/ingress-nginx/commit/66d8d6c61c53eb70757dda2efd965db20c9f81bd))

### Config/Annotations

- Add `relative-redirects`. ([698960e](https://github.com/forkline/ingress-nginx/commit/698960e9b77eb7f5bc8100333800dfa580619fbb))
- Add `proxy-busy-buffers-size`. ([d1dc3e8](https://github.com/forkline/ingress-nginx/commit/d1dc3e827f818ee23a08af09e9a7be0b12af1736))
- Fix `proxy-busy-buffers-size`. ([75a5907](https://github.com/forkline/ingress-nginx/commit/75a590772c763e5083344ba63d06d7c6982591f5))
- Remove `proxy-busy-buffers-size` default value. ([da24841](https://github.com/forkline/ingress-nginx/commit/da24841bc9bd48dbd6fa6e83d5c4a975f3fe250a))

### Controller

- Make Leader Election TTL configurable. ([7c8af49](https://github.com/forkline/ingress-nginx/commit/7c8af4928b9bbc02264d5c7fad95308f1744a143))
- Fix panic in alternative backend merging. ([e5c29d1](https://github.com/forkline/ingress-nginx/commit/e5c29d1ce437446f750875094de9fbe9e05c8250))
- Fail annotation parsing fast and report errors. ([a5cd15d](https://github.com/forkline/ingress-nginx/commit/a5cd15d74a16dc3510283dc7224082f83b40f690))
- Several security fixes. ([cfe3923](https://github.com/forkline/ingress-nginx/commit/cfe3923bd657a82226eb58d3307204a8a8802db4))
- Add traffic distribution support. ([df48ec7](https://github.com/forkline/ingress-nginx/commit/df48ec7a210d5e995c94b5a98af82e11be71003f))
- Fix SSL session ticket path. ([b74cfa8](https://github.com/forkline/ingress-nginx/commit/b74cfa882ce32a6acdd9767266366e289fe967ec))
- Fix nil pointer in path validation. ([ae7a44c](https://github.com/forkline/ingress-nginx/commit/ae7a44cc9f20c016fc58a078a245ec0ba9769452))
- Fix `limit_req_zone` sorting. ([2455bfd](https://github.com/forkline/ingress-nginx/commit/2455bfd7496c8e4765da278a8150805f6d6a25ce))
- Fix host/path overlap detection for multiple rules. ([6393cdd](https://github.com/forkline/ingress-nginx/commit/6393cdda38fc11a034a206867a99b565628687e5))
- Fix sync for when host clock jumps to future. ([8523207](https://github.com/forkline/ingress-nginx/commit/85232076f4ca78acf50e4c30d318c16006e2f9a7))
- Verify UIDs. ([d5cd861](https://github.com/forkline/ingress-nginx/commit/d5cd861e011a7ccc2368a74300863cabf028765d))
- Use 4KiB buffers for PROXY protocol parsing in TLS passthrough. ([4cb48ea](https://github.com/forkline/ingress-nginx/commit/4cb48eaac113b3bad534e7d4db722813a70ca445))
- Enable SSL Passthrough when requested on before HTTP-only hosts. ([efcd54f](https://github.com/forkline/ingress-nginx/commit/efcd54f92b28f798d4458a899541449aee317b2c))

### DOCS

- Add clarification regarding ssl passthrough ([54f1568](https://github.com/forkline/ingress-nginx/commit/54f1568e11f297171583c1f6a51ade7dd78abd22))
- Correct ssl-passthrough annotation description. ([d004fca](https://github.com/forkline/ingress-nginx/commit/d004fcac05a024a4950d2cae2786c1ea69cc08f7))

### DaemonSet

- Implement OpenTelemetry resources. ([8f54b53](https://github.com/forkline/ingress-nginx/commit/8f54b538d9b32649411d9e507ca38fbee55ce51c))

### Dashboard

- Remove `ingress_upstream_latency_seconds`. ([e6851d9](https://github.com/forkline/ingress-nginx/commit/e6851d91df309e7251796873ac4c4d23da827b8e))

### Deploy

- Use LoadBalancer for KIND. ([4f62e98](https://github.com/forkline/ingress-nginx/commit/4f62e980bec084513fdc05a03f91e2481218d1d9))

### Deployment/DaemonSet

- Label pods using `ingress-nginx.labels`. ([47eb3a1](https://github.com/forkline/ingress-nginx/commit/47eb3a17fd95e9205dbe0c6bdb474ac4603f498d))
- Template `topologySpreadConstraints`. ([4869c8b](https://github.com/forkline/ingress-nginx/commit/4869c8b462e1c02597ce978051a694992c8417ae))
- Fix templating & value. ([2d03da6](https://github.com/forkline/ingress-nginx/commit/2d03da63344b63544f8f8f108c54b326fecb7c36))
- Remove `distroless` from `extraModules` templating. ([97d4a83](https://github.com/forkline/ingress-nginx/commit/97d4a83e75ff6ff41243edb5fededbc251aa56b5))

### Development

- Bump Kubernetes to v1.31.4. ([eb62c2a](https://github.com/forkline/ingress-nginx/commit/eb62c2a465939bba354aa00758d8d19b2c6467b5))
- Bump Kubernetes to v1.31.4. ([bdc5017](https://github.com/forkline/ingress-nginx/commit/bdc5017cb0f23527d21d8c9e1677750b5a805167))
- Update Kubernetes to v1.32.0. ([ca5fed8](https://github.com/forkline/ingress-nginx/commit/ca5fed8b070a90f5e0db9da37ee6b119dbd91675))
- Update KIND images. ([d5ab319](https://github.com/forkline/ingress-nginx/commit/d5ab31965708263f90bc21eca87cb2431c865958))

### Doc

- Add `remote_addr` into default values in configmap for TCP logging format ([ee24bf1](https://github.com/forkline/ingress-nginx/commit/ee24bf1bbcfc85ed41c629fcbd0b005e3301b6fe))
- Adding initial hardening guide ([b56258c](https://github.com/forkline/ingress-nginx/commit/b56258c06833748c435c795ef067064e43ffe1f5))
- Updating hardening guide after PR feedback ([98fb569](https://github.com/forkline/ingress-nginx/commit/98fb56912c59d2065d569735ca54991fae7cf58a))

### Docs

- Examples: fix link ([c0e6dae](https://github.com/forkline/ingress-nginx/commit/c0e6daebef4a8996ce47c189f2c8de674816a3a8))
- Configmap: use-gzip ([f5b0905](https://github.com/forkline/ingress-nginx/commit/f5b090518dfb4c396a4e0638e8af8f273ceabd3a))
- Remove redundant --election-id arg from Multiple Ingresses ([93cd78a](https://github.com/forkline/ingress-nginx/commit/93cd78aa456d896e6d1f473f40184fc13c38e037))
- Add documentation about default ingress helm value, corrections to only ingress section ([39e721d](https://github.com/forkline/ingress-nginx/commit/39e721de731e95bf87359378d11af503c31928b7))
- Keep title in navbar in upcase ([6ff70f0](https://github.com/forkline/ingress-nginx/commit/6ff70f015dd8b12d5e96adc6ef3c16e1c9060944))
- Remove opentracing and zipkin from docs ([20d9a60](https://github.com/forkline/ingress-nginx/commit/20d9a609b57934fe2d939ec4b34e2b8634a6d73f))
- Improve default certificate usage. ([0abc9ea](https://github.com/forkline/ingress-nginx/commit/0abc9eaff1791e2960cd8b8d612f8461abb23a1d))
- Specify `ingressClass` for multi-controller setup. ([f0787c3](https://github.com/forkline/ingress-nginx/commit/f0787c3027ca7f2f8caa113d301e4bad3ba2a613))
- Add information about HTTP/3 support. ([6a111a9](https://github.com/forkline/ingress-nginx/commit/6a111a974b1c235f1aa18177c3b26c61318e6eb7))
- Format NGINX configuration table. ([879747a](https://github.com/forkline/ingress-nginx/commit/879747a92fb780b06a0b40f2faf7509727fa345c))
- Clarify `from-to-www` redirect direction. ([e1d81b7](https://github.com/forkline/ingress-nginx/commit/e1d81b78188a6e439b877bc4f6c493e281a00587))
- Fix `from-to-www` redirect description. ([c6e86c8](https://github.com/forkline/ingress-nginx/commit/c6e86c86dc0a007e3bda34ee6eee9cabed34eb2c))
- Fix typo in AWS LB Controller reference ([2e3c2c1](https://github.com/forkline/ingress-nginx/commit/2e3c2c121d10a64f46cd49e5f2a32977cc886cb6))
- Add note about `--watch-namespace`. ([0b98b17](https://github.com/forkline/ingress-nginx/commit/0b98b1783e4b50de4ba82c340dc7a37392720837))
- Convert `opentelemetry.md` from CRLF to LF. ([883c09f](https://github.com/forkline/ingress-nginx/commit/883c09fb57e41a4f31409088d5418ad801f59413))
- Add health check annotations for AWS. ([8d6435b](https://github.com/forkline/ingress-nginx/commit/8d6435b8a0a73e01b198c1f7fd0dd10e76e32482))
- Add a multi-tenant warning. ([114421f](https://github.com/forkline/ingress-nginx/commit/114421f94c8f92d84bc2aeca62e163b7feec3cb5))
- Clarify external & service port in TCP/UDP services explanation. ([162e393](https://github.com/forkline/ingress-nginx/commit/162e3932a28fc2ef583f549e5f9cca0acd3aacdb))
- Add Pod Security Admission. ([1c0f4fa](https://github.com/forkline/ingress-nginx/commit/1c0f4fa8b25f75d1214349b54a48e2547e83afc3))
- Fix limit-rate-after references ([440575e](https://github.com/forkline/ingress-nginx/commit/440575e151b24ecac5cfdeb4cc1dd1483948778b))
- Add CPU usage note for `--metrics-per-undefined-host`. ([0909a61](https://github.com/forkline/ingress-nginx/commit/0909a61ea375b202eb40a8c4744d2cc44b90df7b))
- Add guide on how to set a Maintenance Page. ([94e39e3](https://github.com/forkline/ingress-nginx/commit/94e39e32cfe48c2da2ec259a138bb349ca73399c))
- Clarify rate limits are per ingress controller replica. ([0374af9](https://github.com/forkline/ingress-nginx/commit/0374af94ef9cbf2b50aa57ae8ae68c7dbafec290))
- Improve bare-metal setup. ([36f2d40](https://github.com/forkline/ingress-nginx/commit/36f2d40db984bb1a30e132633a4e74cbc7422646))
- Fix character format. ([c2bab5a](https://github.com/forkline/ingress-nginx/commit/c2bab5a2c90ae1c3edbb4e05bb98b880184be51e))
- Enable code copy button. ([97bbec4](https://github.com/forkline/ingress-nginx/commit/97bbec446b6fc24129846c9b602681e239ea2503))
- Migrate to AR. ([460aa90](https://github.com/forkline/ingress-nginx/commit/460aa90e713b8060ed31f5a36132ec9cd0d9dc87))
- Update link to `values.yaml`. ([4a94726](https://github.com/forkline/ingress-nginx/commit/4a94726c74306a7287a37eaf8c9e8601e9897d0c))
- Use `enable-global-auth` annotation instead of non-existing ConfigMap option. ([aa8cf9a](https://github.com/forkline/ingress-nginx/commit/aa8cf9a795ea7059f36624db41fd0cb2fafc132a))
- Fix OpenTelemetry listing. ([8aca34b](https://github.com/forkline/ingress-nginx/commit/8aca34bfaa3589f9ae5ca300df4d8d825424eb28))
- Fix link in installation instructions. ([3477745](https://github.com/forkline/ingress-nginx/commit/3477745198b713cf7420fed52156f023e675edd0))
- Enable metrics in manifest-based deployments. ([c5eaeb1](https://github.com/forkline/ingress-nginx/commit/c5eaeb1a38f288a26ce04f9173e1ea961b7e0cd5))
- Improve formatting in `monitoring.md`. ([fee0175](https://github.com/forkline/ingress-nginx/commit/fee0175680a2e171b3f6b2c1bbc68ecc89e82c4d))
- How to modify NLB TCP timeout. ([387b459](https://github.com/forkline/ingress-nginx/commit/387b459cfb7babbc6fbc4c64720add1bb6c4df04))
- Add OpenTelemetry defaults. ([a71de6b](https://github.com/forkline/ingress-nginx/commit/a71de6b0f8223f40adae4658013275ea76d4bd49))
- Fix function names in comments. ([cbf5f27](https://github.com/forkline/ingress-nginx/commit/cbf5f275082282441f24c56e38f8b65140b5907f))
- Improve `opentelemetry-trust-incoming-span`. ([23a4c20](https://github.com/forkline/ingress-nginx/commit/23a4c209b4de7cc6ea2cce89a586a8825704cfb7))
- Update prerequisites in `getting-started.md`. ([d7c33ea](https://github.com/forkline/ingress-nginx/commit/d7c33ea648e295bb3c041813fdf0d039089c73b2))
- Fix links and formatting in user guide. ([801ea32](https://github.com/forkline/ingress-nginx/commit/801ea32ff64fcaf9881174cc3304e476e5ce7c8b))
- Use HTTPS for NGINX links. ([4e1b438](https://github.com/forkline/ingress-nginx/commit/4e1b438988768f99f8f3fd186d7f5e11150219b1))
- Remove `X-XSS-Protection` header from hardening guide. ([94acc25](https://github.com/forkline/ingress-nginx/commit/94acc258b7e412855e08e13772e5627092cb1e90))
- Fix default config values and links. ([1e1824b](https://github.com/forkline/ingress-nginx/commit/1e1824b34c9c0a65bcd89fee6a8d9c063dbc44fe))
- Bump mkdocs to v9.6.16, fix links. ([e11414f](https://github.com/forkline/ingress-nginx/commit/e11414fae120932e4695e50038105bab4195000d))
- Replace no-break spaces (U+A0). ([b485724](https://github.com/forkline/ingress-nginx/commit/b4857248801c120dcab4bf550254f4112d2c8c95))
- Remove `datadog` ConfigMap options. ([a5af4bf](https://github.com/forkline/ingress-nginx/commit/a5af4bf2d1d65f43cae2027b3b0bc02eeb382c21))
- Update link to Kubernetes controller documentation. ([fb745c1](https://github.com/forkline/ingress-nginx/commit/fb745c1e6d0b11091abc23e66793e09e58a91f81))
- Fix typo. ([6c58f84](https://github.com/forkline/ingress-nginx/commit/6c58f84b9e5379192d1eea3e66d49e834aa365f0))
- Clarify regex docs. ([0c5a029](https://github.com/forkline/ingress-nginx/commit/0c5a029884dcda6615e800b08dad9d726db1ce89))
- Add retirement blog post. ([f7d0aa0](https://github.com/forkline/ingress-nginx/commit/f7d0aa0b8b181d4d2184eadcfde6431732520570))
- Remove duplicate in log format. ([76009e2](https://github.com/forkline/ingress-nginx/commit/76009e20316a76e0608b3eeaf659fd9cd797e36c))
- Fix typos. ([19b635d](https://github.com/forkline/ingress-nginx/commit/19b635df2d758e15a56ba1cbe1012eeaa52ec52e))
- Clarify valid values for `proxy-request-buffering`. ([6f1daa5](https://github.com/forkline/ingress-nginx/commit/6f1daa552764dfa4b2aa908406ecc5000976079a))
- Add retirement notice to website. ([26549c1](https://github.com/forkline/ingress-nginx/commit/26549c1492b6bb4cdfde57e381dec50b04242bcd))
- Clarify PROXY protocol is not supported on GKE default load balancer. ([50faa99](https://github.com/forkline/ingress-nginx/commit/50faa99bad82aa1f579f9125012e9b978363e639))

### Exoscale

- Use HTTP healthcheck mode ([f5b9d9c](https://github.com/forkline/ingress-nginx/commit/f5b9d9c51267ec1672e717628371512d05e00ef4))

### FIX

- Ingress was not creating the endpoint when target port is string ([14ae787](https://github.com/forkline/ingress-nginx/commit/14ae787b404fe11f7af93cb19e9d3d4eda5f7720))

### Feat

- Canary supports using specific match strategy to match header value. ([0b33650](https://github.com/forkline/ingress-nginx/commit/0b33650bb80bee8c750bbf020abf0c440d6126f6))

### Fix

- Update config map name ([2387c38](https://github.com/forkline/ingress-nginx/commit/2387c386241626e04b3b287e12f479e752dfb84f))
- Fillout missing health check timeout on health check. ([b28577a](https://github.com/forkline/ingress-nginx/commit/b28577a4bf5d47565d86a3feb518d4946ab4f18d))

### Fixed

- Error parsing with-rbac.yaml: error converting YAML to JSON: yaml: line 36: did not find expected key ([2be35d0](https://github.com/forkline/ingress-nginx/commit/2be35d024beddb8dc62f70421acfafcb31c599ce))

### GCE

- Don't update URL Map if unchanged ([0950910](https://github.com/forkline/ingress-nginx/commit/0950910e29d6c0d7fba7fda77c6b3d8ba7631e7a))
- Revert node handlers removal ([11c0306](https://github.com/forkline/ingress-nginx/commit/11c0306937cf20707672d06bae6bd9c07df2a500))

### GitHub

- Improve Dependabot. ([cb1dcb3](https://github.com/forkline/ingress-nginx/commit/cb1dcb3e550d95563798ba06beb4281962d27933))
- Fix `exec` in issue template. ([6ceccbd](https://github.com/forkline/ingress-nginx/commit/6ceccbd67b140b7626670ad17f926f121a9e5563))
- Remove 'Stale Issues and PRs' workflow. ([ea2a37f](https://github.com/forkline/ingress-nginx/commit/ea2a37fac092930a30547bf566ed6a30019ee9f5))
- Bump Chart Testing action. ([b8b5c27](https://github.com/forkline/ingress-nginx/commit/b8b5c27f5ac94aac72d41076c9143aed146ce0c0))

### Go

- Bump to v1.22.5. ([2603677](https://github.com/forkline/ingress-nginx/commit/26036777c9f81cf217e323cbd70c5ed8374911f6))
- Bump to v1.22.6. ([86e8137](https://github.com/forkline/ingress-nginx/commit/86e81373ea0f2c1abc76fc0710070e168ee646b5))
- Bump to v1.22.7. ([0111961](https://github.com/forkline/ingress-nginx/commit/0111961e7d79b386c7717fa392d727e8fdadfc2a))
- Bump to v1.22.8. ([bf287e4](https://github.com/forkline/ingress-nginx/commit/bf287e4331f7de96d06ab9c4a740d873b493b029))
- Bump to v1.23.3. ([4a44778](https://github.com/forkline/ingress-nginx/commit/4a447782ac80c8219f45dd390d09a6056863fddb))
- Bump to v1.23.4. ([53ca6e2](https://github.com/forkline/ingress-nginx/commit/53ca6e292180ade2eeb5507a498a76c647005b5d))
- Clean `go.work.sum`. ([453160d](https://github.com/forkline/ingress-nginx/commit/453160de3d84bb7f5f31167cb53320228b4a5ea0))
- Stop using workspace. ([8111b07](https://github.com/forkline/ingress-nginx/commit/8111b07adbe4ade4aba96bd52457b05fc737628f))
- Bump to v1.23.5. ([06c990f](https://github.com/forkline/ingress-nginx/commit/06c990f80a1d88c06d6b2a42fd5360966cdf8214))
- Replace `golang.org/x/exp/slices` with `slices`. ([68a35a8](https://github.com/forkline/ingress-nginx/commit/68a35a8a120b9f343d8df84be81049359e3b09f4))
- Bump to v1.23.6. ([b63cc4d](https://github.com/forkline/ingress-nginx/commit/b63cc4dc5046bc161edc17652f0a92d565de2969))
- Bump to v1.24.1. ([07a1133](https://github.com/forkline/ingress-nginx/commit/07a1133585b5200609b5523673aa443a20ac7a1f))
- Update dependencies. ([3a9ebee](https://github.com/forkline/ingress-nginx/commit/3a9ebee8cd463f35722286d5e1d7c36f0d1d41d6))
- Update dependencies. ([fa1e9b8](https://github.com/forkline/ingress-nginx/commit/fa1e9b87374ae94683b37f4cceb46bc2dd27c2db))
- Fix Mage. ([1d7abc1](https://github.com/forkline/ingress-nginx/commit/1d7abc12ef727bca69a9e35b88d1167d6d502e9f))
- Bump to v1.24.2. ([23a640f](https://github.com/forkline/ingress-nginx/commit/23a640ff6c5464ec48ab2beafd419e050b886230))
- Update dependencies. ([8f2593b](https://github.com/forkline/ingress-nginx/commit/8f2593bb872ae7fbfa75c203d7fdda00fe90d1d4))
- Update dependencies. ([add38db](https://github.com/forkline/ingress-nginx/commit/add38dbb6ff13fc6d09a3d3caf8e824a9b74c252))
- Update dependencies. ([4470615](https://github.com/forkline/ingress-nginx/commit/4470615d0d4a3aaf1679e910ed7d99df3ee04485))
- Update dependencies. ([e06a171](https://github.com/forkline/ingress-nginx/commit/e06a171dc90503dac702862199df36a553c77eba))
- Update dependencies. ([06e18f5](https://github.com/forkline/ingress-nginx/commit/06e18f5b70c30992e3cbe60664ef781c922abdc6))
- Update dependencies. ([28004d3](https://github.com/forkline/ingress-nginx/commit/28004d3dfc5fe9101bfb548c7989b53d239cc9f5))
- Update dependencies. ([2604f08](https://github.com/forkline/ingress-nginx/commit/2604f08bfc90b77cbe6fe9918bc0843f3c5b2220))
- Bump to v1.24.3. ([e1feec9](https://github.com/forkline/ingress-nginx/commit/e1feec90ffdad31e8887f80b40558d4395e239e3))
- Update dependencies. ([bc490b0](https://github.com/forkline/ingress-nginx/commit/bc490b09c4ccbda27ff36cafc2b42bd4b05e7708))
- Update dependencies. ([262bb77](https://github.com/forkline/ingress-nginx/commit/262bb7714f3da0b0d2192820d56a8e6853e81eea))
- Bump to v1.24.4. ([d432b10](https://github.com/forkline/ingress-nginx/commit/d432b108a91a32b11f6df0b8e53c965aa836a99d))
- Update dependencies. ([311a208](https://github.com/forkline/ingress-nginx/commit/311a2082c5622c7bdaff24273110e6f2a0f3b5da))
- Update dependencies. ([e2b7261](https://github.com/forkline/ingress-nginx/commit/e2b7261b77e66e24ec0e8276cfe94d31406397da))
- Update dependencies. ([b299406](https://github.com/forkline/ingress-nginx/commit/b299406898b1ded8e4a8682df774dc64d25f797f))
- Update dependencies. ([a8e6de7](https://github.com/forkline/ingress-nginx/commit/a8e6de7488dd27d1ed09465cc51e3ad7c950eabd))
- Bump to v1.24.5. ([266905d](https://github.com/forkline/ingress-nginx/commit/266905dcc53a9d05264c85c44f9259ee8676e745))
- Update dependencies. ([2b54cbf](https://github.com/forkline/ingress-nginx/commit/2b54cbf01c61a8b5274c18680956196b4248d4be))
- Update dependencies. ([9bf4d1a](https://github.com/forkline/ingress-nginx/commit/9bf4d1a8dc653bd409dac5c972ad4dd8aba561fb))
- Update dependencies. ([60f7eea](https://github.com/forkline/ingress-nginx/commit/60f7eeab6f4da73336460366adcf3e120552a2ac))
- Bump to v1.24.6. ([99070f9](https://github.com/forkline/ingress-nginx/commit/99070f9827c3bafc0777bb7a1d70d09638ac81b0))
- Update dependencies. ([c88c397](https://github.com/forkline/ingress-nginx/commit/c88c3972dc633999a8cccd4b2b71e150e448e8ad))
- Update dependencies. ([6b95221](https://github.com/forkline/ingress-nginx/commit/6b95221a22d889e38f51ade1a3c6feacf81fed49))
- Update dependencies. ([0e09802](https://github.com/forkline/ingress-nginx/commit/0e09802bc770680c82ab56402c39d78b07512b88))
- Bump to v1.25.0. ([602ce8f](https://github.com/forkline/ingress-nginx/commit/602ce8f918e59a44fe730c30931979477b3f2cf9))
- Update dependencies. ([041f28e](https://github.com/forkline/ingress-nginx/commit/041f28ee6104b26a44dbcb8f35e7024be4572786))
- Bump to v1.25.1. ([0aed652](https://github.com/forkline/ingress-nginx/commit/0aed6521068546f9164891a91f2499dcd9125ff9))
- Update dependencies. ([a78d971](https://github.com/forkline/ingress-nginx/commit/a78d971822c0d39fe32cb4864430c8b7d24c2e05))
- Update dependencies. ([095ff42](https://github.com/forkline/ingress-nginx/commit/095ff424dc51ee351997327dd8116f4f46f9a1e3))
- Update dependencies. ([f1699e2](https://github.com/forkline/ingress-nginx/commit/f1699e29f9ece7f72aefca5373eb74434683d0ed))
- Update dependencies. ([278bb7f](https://github.com/forkline/ingress-nginx/commit/278bb7f812aa0cf43658e2421177402533241e06))
- Update dependencies. ([dd50cf3](https://github.com/forkline/ingress-nginx/commit/dd50cf3cf22e54249203a15ea6e15351b7b905d0))
- Update dependencies. ([26f78fd](https://github.com/forkline/ingress-nginx/commit/26f78fd87103f79d4dfd5ed9a10c253b4dcd7db0))
- Bump to v1.25.2. ([70e01ff](https://github.com/forkline/ingress-nginx/commit/70e01ff4ba2408f6311fad82d20db36b55756195))
- Update dependencies. ([e15ad59](https://github.com/forkline/ingress-nginx/commit/e15ad593e28e74ba87bbea9d0c29ec0eb3d1ae4b))
- Bump to v1.25.3. ([cf998c7](https://github.com/forkline/ingress-nginx/commit/cf998c707ba72d9fea108c5aa7191fc2a09fd365))
- Bump to v1.25.4. ([3ca0307](https://github.com/forkline/ingress-nginx/commit/3ca0307fa76ac3c47f70723850c0ad58fa8670bd))
- Update dependencies. ([d3949ca](https://github.com/forkline/ingress-nginx/commit/d3949ca83a1348ac348fe3174016e34d78d7df82))
- Update dependencies. ([110485d](https://github.com/forkline/ingress-nginx/commit/110485d8af21941df23a6278173aa158ff22d9c5))
- Update dependencies. ([6a374fe](https://github.com/forkline/ingress-nginx/commit/6a374fe3c64799c1f2806d18312785e7047d782d))
- Bump to v1.25.5. ([a83f6dd](https://github.com/forkline/ingress-nginx/commit/a83f6dd4e83c02f8929164b4b5232d495819291a))
- Update dependencies. ([cc19bbf](https://github.com/forkline/ingress-nginx/commit/cc19bbfa3720be026c590a5448d549835d870962))
- Update dependencies. ([24e6a52](https://github.com/forkline/ingress-nginx/commit/24e6a52f4f6d9f4b104d3264ffc69fed8dafe0ba))
- Update dependencies. ([0167cdd](https://github.com/forkline/ingress-nginx/commit/0167cdda685280d3a9389b01eff9c39c7fbcd910))
- Update dependencies. ([2fcf0fd](https://github.com/forkline/ingress-nginx/commit/2fcf0fd6e17892891dcdb40edeef207f736ef0cd))
- Bump to v1.25.6. ([9c95a75](https://github.com/forkline/ingress-nginx/commit/9c95a7564140ed758f3d0ed929011f963617fcbb))
- Update dependencies. ([f8bdc24](https://github.com/forkline/ingress-nginx/commit/f8bdc2458b451a61a76241f565acb1ac75740355))
- Update dependencies. ([cf40d49](https://github.com/forkline/ingress-nginx/commit/cf40d49e22e96fd86d3ebdf558c2f032208333c1))
- Update dependencies. ([322b8a4](https://github.com/forkline/ingress-nginx/commit/322b8a417be70160b6d172f7d037d314dffbbbca))
- Bump to v1.25.7. ([488ae18](https://github.com/forkline/ingress-nginx/commit/488ae181aaf2613d8f3cc87829a264cbae28b343))
- Update dependencies. ([3830ec7](https://github.com/forkline/ingress-nginx/commit/3830ec7303598b363c2e1bc485bb3d5eef61139a))
- Update dependencies. ([f24c031](https://github.com/forkline/ingress-nginx/commit/f24c0312b55dfc88ee00901f2be7fdcf8f4a9f9c))
- Bump to v1.26.1. ([9189dc5](https://github.com/forkline/ingress-nginx/commit/9189dc5ca8b88549e5ba143966016e92e4b0c81b))
- Update dependencies. ([9f34639](https://github.com/forkline/ingress-nginx/commit/9f34639ae89759655adf27494532466a1acae20c))
- Update dependencies. ([5efc236](https://github.com/forkline/ingress-nginx/commit/5efc2362e1e191953ff562e3597be2965e85814e))
- Update dependencies. ([2a1acab](https://github.com/forkline/ingress-nginx/commit/2a1acab51d8cc010f377dd6245269dac5063831f))
- Update dependencies. ([61872ff](https://github.com/forkline/ingress-nginx/commit/61872ff4300229f9751376ad53c2633b0355d18c))

### Godeps,vendor

- Add k8s.io/client-go/tools/cache/testing ([bb09963](https://github.com/forkline/ingress-nginx/commit/bb09963b661f374d943dafad9f8f1c06c98c82ba))

### HPA

- Add `controller.autoscaling.annotations` to `values.yaml`. ([6ed6a76](https://github.com/forkline/ingress-nginx/commit/6ed6a76200595a7c6843322c60491fef2c86a864))
- Autoscaling/v2beta1 deprecated, bump apiVersion to v2 for defaultBackend ([e2d076c](https://github.com/forkline/ingress-nginx/commit/e2d076c4fc52f6c9cdcf9af93d9ae41821303faa))
- Use capabilites & align manifests. ([06612e6](https://github.com/forkline/ingress-nginx/commit/06612e6ffd044171e33b5e8315b9605eb8730c55))

### Hack

- Bump `golangci-lint` to v2.3.0. ([9481de4](https://github.com/forkline/ingress-nginx/commit/9481de4988e48498e2007dad535c10a30046ecd9))

### Hotfix

- Influxdb module disable, see: https://github.com/influxdata/nginx-influxdb-module/pull/9 ([36b5ec8](https://github.com/forkline/ingress-nginx/commit/36b5ec86e099126c33a49ee08ec148e284dde57e))

### Images

- Bump `kube-webhook-certgen`. ([e564e83](https://github.com/forkline/ingress-nginx/commit/e564e831c5f3e1b60766c0e20cbb4fc522dbffc3))
- Bump `NGINX_BASE` to v0.0.9. ([a86ddb5](https://github.com/forkline/ingress-nginx/commit/a86ddb5f032d623713b0f1ade65c00115b101dc2))
- Bump `test-runner`. ([cee3fb3](https://github.com/forkline/ingress-nginx/commit/cee3fb3b131b1bfdf6f0d0ed54cd36a190f4e7a0))
- Trigger NGINX build. ([290de76](https://github.com/forkline/ingress-nginx/commit/290de76a1b3c862bc44ee42ad7e2d017ce5251a5))
- Bump `NGINX_BASE` to v0.0.10. ([2bdca3c](https://github.com/forkline/ingress-nginx/commit/2bdca3ccc7ba60779290a2ae00687143e0dc7d3c))
- Trigger `test-runner` build. ([d6f2b86](https://github.com/forkline/ingress-nginx/commit/d6f2b86508e2e2b89912ad4954d001c6b6a44f1a))
- Re-run `test-runner` build. ([151fca0](https://github.com/forkline/ingress-nginx/commit/151fca0c9cd48a8d2ff77dd4f076d91accf1de8d))
- Trigger NGINX build. ([fd7e02b](https://github.com/forkline/ingress-nginx/commit/fd7e02b97617d7869f583ff0182a893d5ac61d7f))
- Bump `NGINX_BASE` to v0.0.11. ([1e6e2e1](https://github.com/forkline/ingress-nginx/commit/1e6e2e1b376404332c0ae4ae705c1b446d992e66))
- Trigger NGINX build. ([8d96714](https://github.com/forkline/ingress-nginx/commit/8d96714c4e88d53c855c8c7e9f6cb093f1d40827))
- Bump `NGINX_BASE` to v0.0.12. ([ffbbb44](https://github.com/forkline/ingress-nginx/commit/ffbbb449f79580431d789678efc53aa17688faba))
- Trigger `test-runner` build. ([3f0129a](https://github.com/forkline/ingress-nginx/commit/3f0129aa8cf192680d6b356b682eaa61a3873c01))
- Trigger other builds. ([b933310](https://github.com/forkline/ingress-nginx/commit/b933310da5914adbf3f243b5cf902c245d5a0cef))
- Trigger failed builds. ([0dd1bf5](https://github.com/forkline/ingress-nginx/commit/0dd1bf5fb93101fc62a43c24d26ab26d0b027249))
- Trigger `test-runner` build. ([2c42176](https://github.com/forkline/ingress-nginx/commit/2c4217629c37f5e72afad0ba135ea970115dee2a))
- Use latest Alpine 3.20 everywhere. ([8d0e2ef](https://github.com/forkline/ingress-nginx/commit/8d0e2ef9f473d49a0766a818401006d155774e66))
- Bump OpenTelemetry C++ Contrib. ([ee61440](https://github.com/forkline/ingress-nginx/commit/ee61440780fd62f30e7328aae061894c4d387c7f))
- Remove OpenTelemetry. ([3f6e6ae](https://github.com/forkline/ingress-nginx/commit/3f6e6aef78259370b95213d094cf5d74df64aba2))
- Remove NGINX v1.21. ([e33ca05](https://github.com/forkline/ingress-nginx/commit/e33ca05c7b54f81c983565338e0eff828ed49193))
- Trigger NGINX build. ([551c9ab](https://github.com/forkline/ingress-nginx/commit/551c9ab8273752297dd8ef639dcfc8da98077ed1))
- Bump `NGINX_BASE` to v1.0.0. ([9f49f80](https://github.com/forkline/ingress-nginx/commit/9f49f80f63671f135e536d5de2db09fab70c5640))
- Trigger `test-runner` build. ([114a6ab](https://github.com/forkline/ingress-nginx/commit/114a6abbf5e70d1e14dfec2e8340d80e18ef79c7))
- Trigger other builds. ([90259d6](https://github.com/forkline/ingress-nginx/commit/90259d65c5b8cba0ef2d14b6a1df07f2b5463a8e))
- Build `s390x` controller. ([deb01b9](https://github.com/forkline/ingress-nginx/commit/deb01b9f2c5812c6477f9750610ee60660388758))
- Drop `s390x`. ([fdfc97a](https://github.com/forkline/ingress-nginx/commit/fdfc97a7fb090afa3715811169bde75fa662a9e8))
- Trigger `e2e-test-echo` build. ([293b4fe](https://github.com/forkline/ingress-nginx/commit/293b4fef24d67d6539134c15d21d5c8221b118fc))
- Trigger `test-runner` build. ([02a3933](https://github.com/forkline/ingress-nginx/commit/02a3933ea9a8a7e761f75d734023fdbf6fe4fccd))
- Bump `gcb-docker-gcloud` to v20241110-72bb0b1665. ([79c684f](https://github.com/forkline/ingress-nginx/commit/79c684f9cec106271811576202dc4ef10f24bc70))
- Bump Alpine to v3.21. ([c160bff](https://github.com/forkline/ingress-nginx/commit/c160bfff69423a9b734e3e4514bd298adfadd540))
- Bump `gcb-docker-gcloud` to v20241217-ff46a068cd. ([1b596fb](https://github.com/forkline/ingress-nginx/commit/1b596fbdb899d1f29d3807b539c3757d1826f27a))
- Trigger NGINX build. ([ba4166f](https://github.com/forkline/ingress-nginx/commit/ba4166fe4a55454bde4e37e5917bd45d167b223d))
- Bump `NGINX_BASE` to v1.1.0. ([0ef18ba](https://github.com/forkline/ingress-nginx/commit/0ef18ba7fb7ffe5491bbabbb510eee0d17e3ae2a))
- Trigger `test-runner` build. ([68ed4e7](https://github.com/forkline/ingress-nginx/commit/68ed4e7b656cbb15b9d751825be61a82667508d3))
- Trigger other builds (1/2). ([30e1eee](https://github.com/forkline/ingress-nginx/commit/30e1eee243955709993e25c2142f0df5312b05e4))
- Trigger other builds (2/2). ([227de50](https://github.com/forkline/ingress-nginx/commit/227de501177ae39de10a50689d4e9d9b780fa141))
- Trigger NGINX build. ([9026c03](https://github.com/forkline/ingress-nginx/commit/9026c03fdf435424807f04212ccad1fb37096190))
- Bump `NGINX_BASE` to v2.0.0. ([9756893](https://github.com/forkline/ingress-nginx/commit/97568931850c2dda8fa5341ab01412f216fe3eed))
- Trigger Test Runner build. ([a188f4e](https://github.com/forkline/ingress-nginx/commit/a188f4eb1997c9281c970fc92ac048fdce5d5356))
- Trigger other builds (1/2). ([b932ac0](https://github.com/forkline/ingress-nginx/commit/b932ac06676d1e5bb04973a9793c392a1e1922ed))
- Trigger other builds (2/2). ([3e8586b](https://github.com/forkline/ingress-nginx/commit/3e8586b9b2adced5d7bd4fd437bf19eac5c31731))
- Bump `gcb-docker-gcloud` to v20250116-2a05ea7e3d. ([133b028](https://github.com/forkline/ingress-nginx/commit/133b02834cf4227d0a65d8b111563934307ddd56))
- Update `kubectl` to v1.31.5. ([240c249](https://github.com/forkline/ingress-nginx/commit/240c249f7b63a01067c7ec6540b37f7e1968ce4d))
- Migrate to AR. (1/2) ([b916cf5](https://github.com/forkline/ingress-nginx/commit/b916cf50792a1912bfd09904aec54905665242e8))
- Migrate to AR. (2/2) ([2153cab](https://github.com/forkline/ingress-nginx/commit/2153cab0bb5ad509ac1b6cd0798f94fe9794a8cd))
- Update `kubectl` to v1.32.2. ([6a889f4](https://github.com/forkline/ingress-nginx/commit/6a889f4d4078fd676f4dab757b4f6e8fe9c67159))
- Rework. (1/3) ([b6e5ca2](https://github.com/forkline/ingress-nginx/commit/b6e5ca2a68449c74a6fc572d9a2aa750dc320fb9))
- Rework. (2/3) ([5e2d76c](https://github.com/forkline/ingress-nginx/commit/5e2d76c0f096836c3a67fdf987a3c606effc5319))
- Rework. (3/3) ([8164d4e](https://github.com/forkline/ingress-nginx/commit/8164d4efd2e5ede95aab5c6e7913ed90eda8f4b4))
- Trigger NGINX build. ([013890d](https://github.com/forkline/ingress-nginx/commit/013890dfcb080818ca7c6b7ad8f7da4346512e04))
- Bump `NGINX_BASE` to v2.0.1. ([f8ec75b](https://github.com/forkline/ingress-nginx/commit/f8ec75b41dad869aeef76b146539b8c424b2837a))
- Trigger Test Runner build. ([dc7ad95](https://github.com/forkline/ingress-nginx/commit/dc7ad9533aed923da123270181d972ed1958ac24))
- Trigger other builds (1/2). ([b8791c0](https://github.com/forkline/ingress-nginx/commit/b8791c04bee6d3f204143c84c053f8bb8a2696fc))
- Trigger other builds (2/2). ([870550c](https://github.com/forkline/ingress-nginx/commit/870550cb261dd9f3e34cdf09c2592fcb7fd21caa))
- Extract modules. ([46f0fd4](https://github.com/forkline/ingress-nginx/commit/46f0fd4e9f5755bdb2a5b887e7cd6674b158649b))
- Fix FromAsCasing. ([35abf6b](https://github.com/forkline/ingress-nginx/commit/35abf6b969ed1165521a408ca811b61ab2d8cb2b))
- Trigger NGINX build. ([8eb2152](https://github.com/forkline/ingress-nginx/commit/8eb2152d288994da60fd3f8852f17bf2dcd843ce))
- Bump `NGINX_BASE` to v2.0.2. ([ce1f130](https://github.com/forkline/ingress-nginx/commit/ce1f130ca95ad55dd980a2d878e45317306a4897))
- Trigger Test Runner build. ([46ee137](https://github.com/forkline/ingress-nginx/commit/46ee1377e632d751bb7421604e68e59cc62fbf97))
- Trigger NGINX build. ([1a4d476](https://github.com/forkline/ingress-nginx/commit/1a4d4761a0a8d0b8fc5b74ed2f71bfed21d1ffad))
- Bump NGINX to v2.0.3. ([9105640](https://github.com/forkline/ingress-nginx/commit/91056403059b1bee8c4056286272dd3bd41abcfa))
- Trigger Test Runner build. ([ba56882](https://github.com/forkline/ingress-nginx/commit/ba56882c944eca65a2d227674cb93cad53238193))
- Trigger other builds (1/2). ([5bf90a1](https://github.com/forkline/ingress-nginx/commit/5bf90a1f49b4f36da4f87fcea26bbaf47d088531))
- Trigger other builds (2/2). ([35a7798](https://github.com/forkline/ingress-nginx/commit/35a77984cfa5b18d29cc3f75afa9ccd93539bc50))
- Trigger NGINX build. ([4422d4f](https://github.com/forkline/ingress-nginx/commit/4422d4f3ea38b083f3d82b1f7c6c37b4767180fa))
- Bump NGINX to v2.1.0. ([d62898f](https://github.com/forkline/ingress-nginx/commit/d62898fc14930914b8f00206141e6ac3ab48deea))
- Trigger Test Runner build. ([1408b74](https://github.com/forkline/ingress-nginx/commit/1408b7447682e42778ea2e446336d4047f8f2634))
- Bump GCB Docker GCloud to v20250513-9264efb079. ([b0160e6](https://github.com/forkline/ingress-nginx/commit/b0160e6456ce3b40572c07caf52af68abe754ff4))
- Build Go gRPC Greeter Server from scratch. ([3f256fb](https://github.com/forkline/ingress-nginx/commit/3f256fb82dcbe9c19606bad1757cd82ec44d7e65))
- Trigger NGINX build. ([974ca67](https://github.com/forkline/ingress-nginx/commit/974ca67915475496764a72682971319fe0bf805f))
- Bump NGINX to v2.1.1. ([66c248a](https://github.com/forkline/ingress-nginx/commit/66c248a2b2380f69db6fd5cc662702d7bda62f9d))
- Trigger Test Runner build. ([9569a76](https://github.com/forkline/ingress-nginx/commit/9569a76edc3a26378539d27a97724c01c5ca8dcb))
- Trigger other builds (1/2). ([ea45520](https://github.com/forkline/ingress-nginx/commit/ea455201e6f5604b446d897b768bd69fe6b0be15))
- Trigger other builds (2/2). ([f0f2db5](https://github.com/forkline/ingress-nginx/commit/f0f2db512fc88e6e558d9672e9947a5ddd6e530e))
- Trigger controller build. ([e4b964b](https://github.com/forkline/ingress-nginx/commit/e4b964bc970d5d16b7a1bce897a98c623fe4b3cd))
- Fix LuaRocks. ([4bae139](https://github.com/forkline/ingress-nginx/commit/4bae1397532d159c01cadfe070ba21cacbbc98a4))
- Update LuaRocks to v3.12.0. ([2363343](https://github.com/forkline/ingress-nginx/commit/2363343c7ae7aa96b01f1f2bf17e2db411991414))
- Bump Alpine to v3.22. ([44329a8](https://github.com/forkline/ingress-nginx/commit/44329a804efbdcec74fd91d62b988dc6f621ca8a))
- Trigger NGINX build. ([1f3f2bc](https://github.com/forkline/ingress-nginx/commit/1f3f2bcb62364211aa60a4ef9d35a79f8ccbc908))
- Bump NGINX to v2.2.0. ([457e398](https://github.com/forkline/ingress-nginx/commit/457e398de12ab0e8ce81ad54edbf0d7e778ae083))
- Trigger Test Runner build. ([60095a6](https://github.com/forkline/ingress-nginx/commit/60095a694885a0f5fc01d9f4bb7332e57e76e346))
- Trigger other builds (1/2). ([1d64e7c](https://github.com/forkline/ingress-nginx/commit/1d64e7c43a5a129dd8dceae9617ab2fdc9ff0090))
- Trigger other builds (2/2). ([5a4f379](https://github.com/forkline/ingress-nginx/commit/5a4f379ee8e47b5a01e9dbd4c6bcb75a7f14a1bb))
- Trigger controller build. ([4cbb78a](https://github.com/forkline/ingress-nginx/commit/4cbb78a9dc4f1888af802b70ddf980272e01268b))
- Trigger NGINX build. ([617b310](https://github.com/forkline/ingress-nginx/commit/617b31084bf0e4cd3c8daae694b409b3b5c4d777))
- Bump NGINX to v2.2.1. ([e64d6a5](https://github.com/forkline/ingress-nginx/commit/e64d6a53610e52d85c2526b34621f75b5c047ef5))
- Trigger Test Runner build. ([de10fe7](https://github.com/forkline/ingress-nginx/commit/de10fe7930b9ecd1b33f347c97901142a4376702))
- Trigger other builds (1/2). ([68161c7](https://github.com/forkline/ingress-nginx/commit/68161c733d5a948c4f660afc47c43067638a1df4))
- Trigger other builds (2/2). ([0f18595](https://github.com/forkline/ingress-nginx/commit/0f185955aedcddb3291afd8899cee2b1a0d99750))
- Remove redundant ModSecurity-nginx patch. ([7419b7a](https://github.com/forkline/ingress-nginx/commit/7419b7a15a22952b24ab19ff56e8686a90e560d4))
- Trigger NGINX build. ([eeed081](https://github.com/forkline/ingress-nginx/commit/eeed0814bdd2ae202d7de66f45a9ff26ea7b5036))
- Bump NGINX to v2.2.2. ([95ec994](https://github.com/forkline/ingress-nginx/commit/95ec9946f8233b92028607cecdaa34c28908bc99))
- Trigger Test Runner build. ([782c332](https://github.com/forkline/ingress-nginx/commit/782c3320d039243fd090ca0ee8fef9fa1d6fc877))
- Trigger other builds (1/2). ([0d5409c](https://github.com/forkline/ingress-nginx/commit/0d5409c936963ddb706c562b5c87bbfb9857af20))
- Trigger other builds (2/2). ([bcff4e2](https://github.com/forkline/ingress-nginx/commit/bcff4e20b2cd867f435f194ad93cdfa7a53f2adb))
- Use Alpine v3.22.1. ([75217c0](https://github.com/forkline/ingress-nginx/commit/75217c0e2b6dfe5c57b310ae474387fa46a5c56b))
- Trigger NGINX build. ([fee63bd](https://github.com/forkline/ingress-nginx/commit/fee63bd05ffdd41dc630ce86ff08aad132f3a4c7))
- Bump NGINX to v2.2.3. ([15d4732](https://github.com/forkline/ingress-nginx/commit/15d47329a5a06fbd63bc9ac558f768754eaced1a))
- Trigger Test Runner build. ([c66aa89](https://github.com/forkline/ingress-nginx/commit/c66aa896a6732b676aaa59ea0dc65b454af1ae1b))
- Trigger other builds (1/2). ([0dd597e](https://github.com/forkline/ingress-nginx/commit/0dd597e46bd0d1c47763bbd2f4cc827485e1e545))
- Trigger other builds (2/2). ([98e761a](https://github.com/forkline/ingress-nginx/commit/98e761a1810acfcf5938c6d5ecdb0fb31f9c99bc))
- Bump Alpine to v3.22.2. ([5121c2e](https://github.com/forkline/ingress-nginx/commit/5121c2e3f59d4e846feb70870b806fc3f29bf431))
- Trigger NGINX build. ([5c4cedb](https://github.com/forkline/ingress-nginx/commit/5c4cedbb53e2125391584ad98b510a6d708bf931))
- Bump NGINX to v2.2.4. ([fd4c8ac](https://github.com/forkline/ingress-nginx/commit/fd4c8acc0c2589a059b748e04ae972923da29c9d))
- Trigger Test Runner build. ([ab51ade](https://github.com/forkline/ingress-nginx/commit/ab51adeab2976fbcdd5dfee63c3c9474392fabb1))
- Trigger other builds (1/2). ([6673cab](https://github.com/forkline/ingress-nginx/commit/6673cabc23aab16554111b40500492f38ce860c2))
- Trigger other builds (2/2). ([4f895e1](https://github.com/forkline/ingress-nginx/commit/4f895e1a1804f8b9b0573eb11d45b7182ad402b6))
- Bump other images. ([5119318](https://github.com/forkline/ingress-nginx/commit/511931881b9829e2eb7fe8968fa1e13294646361))
- Trigger controller build. ([52c0a83](https://github.com/forkline/ingress-nginx/commit/52c0a83ac9bc72e9ce1b9fe4f2d6dcc8854516a8))
- Bump GCB Docker GCloud to v20251110-7ccd542560. ([1f4e7a3](https://github.com/forkline/ingress-nginx/commit/1f4e7a37914e58df2f44dfe617eeb46c77ea246c))
- Trigger NGINX build. ([901e259](https://github.com/forkline/ingress-nginx/commit/901e259c3afcfd9c5fa531ffbadddce3ef3d0a0f))
- Bump NGINX to v2.2.5. ([08a97c2](https://github.com/forkline/ingress-nginx/commit/08a97c2ac33fbead97de24c80933f63e8af21415))
- Trigger Test Runner build. ([78cc1a1](https://github.com/forkline/ingress-nginx/commit/78cc1a1cd29e8e57ba83c55dbeb1b3a81c37b928))
- Trigger other builds (1/2). ([1808cd9](https://github.com/forkline/ingress-nginx/commit/1808cd9f2befa011bec94d46b369c8ebc30b29fc))
- Trigger other builds (2/2). ([cc7f992](https://github.com/forkline/ingress-nginx/commit/cc7f992c570568668fe59ec6ba69db058e5baf2b))
- Update LuaRocks to v3.12.2. ([ad4b653](https://github.com/forkline/ingress-nginx/commit/ad4b6535153105ca669167bece3da0b86efcdbed))
- Bump other images. ([4a73100](https://github.com/forkline/ingress-nginx/commit/4a731008939bcc7c232952c53240310bac741d43))
- Bump Alpine to v3.23.0. ([b33b024](https://github.com/forkline/ingress-nginx/commit/b33b02493b7420dc5e91fb1b713966cd4515c6e1))
- Bump GCB Docker GCloud to v20251110-7ccd542560. ([3d5559a](https://github.com/forkline/ingress-nginx/commit/3d5559af2b4cf6de1045b5579bdbca0871d384e3))
- Bump GCB Docker GCloud to v20251211-4c812d4cd8. ([d499ed9](https://github.com/forkline/ingress-nginx/commit/d499ed9b8f9a132f62672031afee73ba95551113))
- Bump Alpine to v3.23.2. ([b779068](https://github.com/forkline/ingress-nginx/commit/b7790688eeae97e6d67e115809357246707ae2a4))
- Bump GCB Docker GCloud to v20251222-9ed298b43e. ([c1fc4bf](https://github.com/forkline/ingress-nginx/commit/c1fc4bf16ebe4a22ae23165b87688f64a6a8ab1f))
- Trigger NGINX build. ([f9d782b](https://github.com/forkline/ingress-nginx/commit/f9d782bc96b3ee95685a24bcea8fda983bd2930c))
- Bump GCB Docker GCloud to v20260108-7f313c340e. ([96e2edd](https://github.com/forkline/ingress-nginx/commit/96e2edd48cd63dda214b0879306850bf979e6f14))
- Bump NGINX to v2.2.6. ([5f1c3a7](https://github.com/forkline/ingress-nginx/commit/5f1c3a7e5d45f88115a55ad59abf7656448b781e))
- Trigger Test Runner build. ([a2f3d94](https://github.com/forkline/ingress-nginx/commit/a2f3d949588897e2f6a2938f54599f42e15e2ee7))
- Trigger other builds (1/2). ([c26980a](https://github.com/forkline/ingress-nginx/commit/c26980a0c9a87be6f6cc981ed77e046f1ea4d278))
- Trigger other builds (2/2). ([1c8edf0](https://github.com/forkline/ingress-nginx/commit/1c8edf0af3f90275b36c725f7aca2bec428724f1))
- Bump other images. ([624ead9](https://github.com/forkline/ingress-nginx/commit/624ead9e8ba73f28bc5f0402afc66fc3229ad31b))
- Bump GCB Docker GCloud to v20260127-c1affcc8de. ([f4caf3b](https://github.com/forkline/ingress-nginx/commit/f4caf3b4e89bfcd2f6f7a9af726e9c37491f9a64))
- Trigger NGINX build. ([a997c44](https://github.com/forkline/ingress-nginx/commit/a997c44fa78b08fdcd4ab8895648a3ea18f993a4))
- Bump NGINX to v2.2.7. ([8f58b24](https://github.com/forkline/ingress-nginx/commit/8f58b24e2004511c1e667db789b47a6cd1edf617))
- Trigger Test Runner build. ([cd37f45](https://github.com/forkline/ingress-nginx/commit/cd37f45d8ace8d844470543e6ccee3ac466a4d9e))
- Trigger other builds (1/2). ([2f57f15](https://github.com/forkline/ingress-nginx/commit/2f57f152f115f24adc14ae45c0995c5e260ab2f7))
- Trigger other builds (2/2). ([2b7afb6](https://github.com/forkline/ingress-nginx/commit/2b7afb689f59ce02b4cd5a571416580d0687c9b6))
- Bump other images. ([5ed6ced](https://github.com/forkline/ingress-nginx/commit/5ed6cedb9c2af6bf5e5dc1b024fb5be7e4ef4596))
- Bump Alpine to v3.23.3. ([9055a5c](https://github.com/forkline/ingress-nginx/commit/9055a5c7909dd4230c563638471a46a160ece5df))
- Trigger NGINX build. ([ca46007](https://github.com/forkline/ingress-nginx/commit/ca46007e558c28ec8b8652f41f48d2e8ccda485a))
- Bump NGINX to v2.2.8. ([d27738d](https://github.com/forkline/ingress-nginx/commit/d27738d5c9fd128c6345e1820731c6725c7aac20))
- Trigger Test Runner build. ([91faeab](https://github.com/forkline/ingress-nginx/commit/91faeab1930813525eecfac25689788295d522bd))
- Trigger other builds (1/2). ([cdc4e3e](https://github.com/forkline/ingress-nginx/commit/cdc4e3e0e451b679dd586341e2593416dc0f9b20))
- Trigger other builds (2/2). ([008a082](https://github.com/forkline/ingress-nginx/commit/008a08242a9f400f55f38ed50f3b436c99ffcbcf))
- Bump other images. ([f701b56](https://github.com/forkline/ingress-nginx/commit/f701b561f66cbbf73269b78693b57b850f6045bc))
- Trigger controller build. ([00f349e](https://github.com/forkline/ingress-nginx/commit/00f349eae47a4a8b194e53866867107697bd5b41))
- Trigger NGINX build. ([c153522](https://github.com/forkline/ingress-nginx/commit/c1535222b36f77e8aeb44a8d64d2bfa0cd71273f))
- Trigger Test Runner build. ([e3d7fba](https://github.com/forkline/ingress-nginx/commit/e3d7fba51f59fbeb0d6fe1de0218b8980ada28f1))
- Trigger other builds (1/2). ([e944ce6](https://github.com/forkline/ingress-nginx/commit/e944ce62f0de5504015205e5563020cd9dbcb292))
- Trigger other builds (2/2). ([b704853](https://github.com/forkline/ingress-nginx/commit/b70485329a3008128affa906895e526be18e61ac))
- Bump other images. ([e077c3f](https://github.com/forkline/ingress-nginx/commit/e077c3f6aa7c173c20e2360ceec783edb682c407))

### Ingresses

- Allow `.` in `Exact` and `Prefix` paths. ([618aae1](https://github.com/forkline/ingress-nginx/commit/618aae18515213bcf3fb820e6f8c234703d844b2))

### KEP

- Availability zone aware routing ([d9de505](https://github.com/forkline/ingress-nginx/commit/d9de5053412908f8f8b654966f966785f4a26d00))

### Lua

- Remove plugins from `.luacheckrc` & E2E docs. ([bde6a6b](https://github.com/forkline/ingress-nginx/commit/bde6a6bc3e5751c1223369462db365cfd1f1a268))
- Extract external auth into file. ([7356c4f](https://github.com/forkline/ingress-nginx/commit/7356c4f40f49bb2898d08f7bc272cfd04e0c40db))
- Fix `ExternalName` services without endpoints. ([76e2f69](https://github.com/forkline/ingress-nginx/commit/76e2f6944982e2dafaac984b8c12e29c86b17608))
- Fix type mismatch. ([85999aa](https://github.com/forkline/ingress-nginx/commit/85999aadbefa10534edb6998e903439c8833d030))

### Mage

- Implement static check recommendations. ([dbe4994](https://github.com/forkline/ingress-nginx/commit/dbe499437e179ca6160ccb06eca5ebbf3e99efba))
- Stop mutating release notes. ([2d67ec2](https://github.com/forkline/ingress-nginx/commit/2d67ec293526470a9f91bdc3794834670f5f44c7))
- Rewrite `updateChartValue` to obsolete outdated libraries. ([3d24186](https://github.com/forkline/ingress-nginx/commit/3d2418625a00c483bc4b827f1bac35504ae635a7))

### Make

- Add `helm-test` target. ([01d3d80](https://github.com/forkline/ingress-nginx/commit/01d3d80cfdcfa2a9f8f4b86457cbac852256c998))

### Metrics

- Remove `ingress_upstream_latency_seconds`. ([eee2760](https://github.com/forkline/ingress-nginx/commit/eee2760907f029b808d95cd90abf290a9f0f09dd))
- Add `--metrics-per-undefined-host` argument. ([034c3cc](https://github.com/forkline/ingress-nginx/commit/034c3ccad41374822a9180dceb3e5b999dc43f2f))
- Fix namespace in `nginx_ingress_controller_ssl_expire_time_seconds`. ([9e6c406](https://github.com/forkline/ingress-nginx/commit/9e6c40664fe81c210a6e07e35bee0592842b9f05))
- Disable by default. ([75c77e5](https://github.com/forkline/ingress-nginx/commit/75c77e5dc3cdbf93725cd586afe7288b2b52eefd))
- Fix `nginx_ingress_controller_config_last_reload_successful`. ([59dbc01](https://github.com/forkline/ingress-nginx/commit/59dbc01fea0d4669cd2c899df9ce0796936e42e6))

### NGINX

- Remove inline Lua from template. ([6510535](https://github.com/forkline/ingress-nginx/commit/6510535ae0164e48d60d0c3df79bffd35031b3f3))
- Remove unused substitutions module. ([c8ab89c](https://github.com/forkline/ingress-nginx/commit/c8ab89c0211abba8bcf13a4db061c613fb37de3a))
- Bump OpenTelemetry. ([5b142ed](https://github.com/forkline/ingress-nginx/commit/5b142ed7c44327a949037833531b7cf11057a1ad))
- Bump ModSecurity. ([69fd353](https://github.com/forkline/ingress-nginx/commit/69fd353086305fbbcb2b64bcc0724d0b960de3e0))
- Bump to OpenResty v1.27.1.1. ([1ece0dd](https://github.com/forkline/ingress-nginx/commit/1ece0ddbc110c207213709b5ffcecfa7f0ac46be))
- Align quotes. ([cc34197](https://github.com/forkline/ingress-nginx/commit/cc341973b022b807f6cf9eafdcdbbf3baa151b44))
- Update ModSecurity. ([cb04b22](https://github.com/forkline/ingress-nginx/commit/cb04b22b805d94cb0d0c0468e162c045f9019091))
- Add NJS. ([3710e62](https://github.com/forkline/ingress-nginx/commit/3710e6254184c722638ab0f281843cae8b31cf36))
- Add X-Original-Forwarded-Host header. ([da54ac6](https://github.com/forkline/ingress-nginx/commit/da54ac6b25239eed173e5b81a703290398ed6ba7))
- Correctly determine client IP. ([cf0a441](https://github.com/forkline/ingress-nginx/commit/cf0a44191c496fbabb8d12edbdd9f8b6c987b5ba))
- Bump to OpenResty v1.27.1.2. ([cd5c23c](https://github.com/forkline/ingress-nginx/commit/cd5c23cc65c22a5bf4169ef8877df461152ce2cb))
- Disable mimalloc's architecture specific optimizations. ([19d2bd8](https://github.com/forkline/ingress-nginx/commit/19d2bd86cb76c4db434e2cf2e67b0122963521f7))
- Update OWASP CRS to v4.22.0. ([d31735a](https://github.com/forkline/ingress-nginx/commit/d31735af5f4e1bb8946009272e7a1eac66da4340))

### NIT

- Correct comment re default of server-tokens=false ([aecc5ba](https://github.com/forkline/ingress-nginx/commit/aecc5bac2151d8ef5305de6b91e3adeb4b59a86f))

### Network

- Rework IPv6 check. ([ba22a30](https://github.com/forkline/ingress-nginx/commit/ba22a30c660c81ba2695e2072ecf5e0c5a636fc2))

### Owners

- Promote Gacko to `ingress-nginx-maintainers` & `ingress-nginx-reviewers`. ([bf3fa53](https://github.com/forkline/ingress-nginx/commit/bf3fa531676ce2b61effb55f5a94345dbc382d09))
- Promote Gacko to admin. ([68b59db](https://github.com/forkline/ingress-nginx/commit/68b59db3e932db354287547c8137fe1d696d3013))

### PDB

- Add `maxUnavailable`. ([170af7b](https://github.com/forkline/ingress-nginx/commit/170af7be88ccbcd3a13f67811ab728d056377a26))

### POC

- Setting upstream vhost for nginx. ([2120aab](https://github.com/forkline/ingress-nginx/commit/2120aab66cc62af94024475e46835bb0e14ef8b4))

### Performance

- Json encoding share to eatch request ([68ec350](https://github.com/forkline/ingress-nginx/commit/68ec350388b2b52edb74cca13e0715bb35211343))
- Avoid unnecessary byte/string conversion ([d02ba28](https://github.com/forkline/ingress-nginx/commit/d02ba28b9609c106ea4de8e0f3dbf5abc3d06fdc))

### Plugin

- Bump `goreleaser` to v2. ([5ae018e](https://github.com/forkline/ingress-nginx/commit/5ae018e5df4ac54c5280a96b43808090a4c24ec5))
- Improve error handling. ([6841107](https://github.com/forkline/ingress-nginx/commit/6841107b8b1a443438b68fec2b3d728f542d6c44))
- Change `rewriteTargetWithoutCaptureGroup` lint to include any numbered capture group. ([2eba780](https://github.com/forkline/ingress-nginx/commit/2eba7801e74cb80e8acacc6c898984ba486eca81))

### Proposal

- E2e tests for regex patterns ([e44cab7](https://github.com/forkline/ingress-nginx/commit/e44cab7245ad0e0c31cb6d0c947fa1aa51cff1ba))

### README

- Update `external-dns` link. ([91a89bc](https://github.com/forkline/ingress-nginx/commit/91a89bcc0c1f63045de37e5dd8a4f7b6a3a85788))
- Fix support matrix. ([988ebd9](https://github.com/forkline/ingress-nginx/commit/988ebd9a0fbe122c4409b3ba30f0d1dcb3c423d6))

### Refactor

- Refactor rate limit whitelist ([f045fa6](https://github.com/forkline/ingress-nginx/commit/f045fa6d881809d0c4fad738c3bb640712b4e24d))
- Refactor controllers.go ([a61017a](https://github.com/forkline/ingress-nginx/commit/a61017ae4e9debd41913c4c63d18a66db4a2f66a))
- Refactor balancer into more testable and extensible interface ([e9dc275](https://github.com/forkline/ingress-nginx/commit/e9dc275b81dc4ab17479463ea15a1d71efdd6ea0))
- Refactor some lua code ([cb47558](https://github.com/forkline/ingress-nginx/commit/cb4755835e7ba4fd0a5c056a8141653bb85bfb47))
- Remove unnecessary libthrift libraries, as Jaeger is statically linked ([0e26cff](https://github.com/forkline/ingress-nginx/commit/0e26cff42f705df64c3cf1261b4c80f4a23502d1))
- Refactor lua balancer and fix ipv6 issue ([4b07e73](https://github.com/forkline/ingress-nginx/commit/4b07e73e5d7589134719ded0aa1093c19ea6241f))
- Refactor GetFakeSSLCert ([13a7e2c](https://github.com/forkline/ingress-nginx/commit/13a7e2c5d089950f2ad92f683999756ffb341bfe))
- Refactor force ssl redirect logic ([8c64b12](https://github.com/forkline/ingress-nginx/commit/8c64b12a967670934cf50b2757c0f827e62d313d))
- Refactor ssl handling in preperation of OCSP stapling ([4bb9106](https://github.com/forkline/ingress-nginx/commit/4bb9106be267334315c301891b55de17a49203bf))
- Use more specific var name ([eb112ea](https://github.com/forkline/ingress-nginx/commit/eb112ea06cf1952e2b082ddab1e81f08c325bb53))
- Update DaemonSet and Deployment command params to use templates ([5a52d99](https://github.com/forkline/ingress-nginx/commit/5a52d99ae85cfe5ef9535291b8326b0006e75066))
- Refactor helm ci tests part I

* refactor helm ci tests part I

Signed-off-by: cpanato <ctadeu@gmail.com>

* update indentation

Signed-off-by: cpanato <ctadeu@gmail.com>

* fix path

Signed-off-by: cpanato <ctadeu@gmail.com>

* more updates

Signed-off-by: cpanato <ctadeu@gmail.com>

* add helm-lint job

Signed-off-by: cpanato <ctadeu@gmail.com>

---------

Signed-off-by: cpanato <ctadeu@gmail.com> ([c0767cc](https://github.com/forkline/ingress-nginx/commit/c0767ccc6164882aa83861ae0de68301f2e47505))

### Repository

- Rename `Changelog.md.gotmpl` into `changelog/controller.md.gotmpl`. ([6cd7331](https://github.com/forkline/ingress-nginx/commit/6cd7331bd59798bfaa97b4a4e55cf16927d0baf1))
- Improve `changelog/controller.md.gotmpl`. ([7e34a67](https://github.com/forkline/ingress-nginx/commit/7e34a676b975b0795ff181ccf75d992ea7c6364d))
- Rename `changelog/Changelog-*.md` into `changelog/controller-*.md`. ([84bdad5](https://github.com/forkline/ingress-nginx/commit/84bdad5341e8da721557176049321e3af54dc566))
- Align `changelog/controller-*.md` to `changelog/controller.md.gotmpl`. ([433781c](https://github.com/forkline/ingress-nginx/commit/433781c9186609ff6a31c76bbf8e1b5075e1384b))
- Add changelogs from `release-v1.10`. ([cb2cdde](https://github.com/forkline/ingress-nginx/commit/cb2cdde10e81f7f425d804e681f347847706c7a6))
- Update owners. ([bd3ee3e](https://github.com/forkline/ingress-nginx/commit/bd3ee3ed3f08d0e204bea93c5b35aa1306f7245d))
- Remove `netlify.toml`. ([7877ead](https://github.com/forkline/ingress-nginx/commit/7877ead6dbf7f039165cad6d6325810b39164f61))

### Security

- Follow-up on recent changes. ([e9f6c8e](https://github.com/forkline/ingress-nginx/commit/e9f6c8e8f2d630e7394b9d2788a4d3feee5aa919))
- Harden socket creation and validate error code input. ([cca7690](https://github.com/forkline/ingress-nginx/commit/cca7690f315d1c545f1cfdcb24d3bafc61cde10e))

### Status

- Add support for multiple Node IP addresses. ([6917c48](https://github.com/forkline/ingress-nginx/commit/6917c4869c09cec3a6fac330ac6ae8a0ea808b31))

### Store

- Handle panics in service deletion handler. ([371cb38](https://github.com/forkline/ingress-nginx/commit/371cb3891cec23052ad93e565ef7f837889b687c))

### TLS.md

- Move the TLS secret misc bit to the TLS document ([451a01b](https://github.com/forkline/ingress-nginx/commit/451a01bb0a2a470bf9ea8b44312769aa4f3bd51e))
- Clarify how to set --default-ssl-certificate ([f65c8f0](https://github.com/forkline/ingress-nginx/commit/f65c8f0aaaa3cbcdd8cc685055969c8bb9e1cc30))
- Remove the frankly useless curl output in the default certificate section ([aca5097](https://github.com/forkline/ingress-nginx/commit/aca5097a5621d833c9c0c281f666d2af92a65266))
- Reformat and grammar check ([ec56200](https://github.com/forkline/ingress-nginx/commit/ec56200ee021fca15ee8fbd891de6a687b4e758a))
- Remove useless manual TOC ([ed48199](https://github.com/forkline/ingress-nginx/commit/ed48199b3034cdd19409b826275c82900297d040))

### Template

- Bypass custom error pages when handling auth URL requests. ([c0a37c1](https://github.com/forkline/ingress-nginx/commit/c0a37c15a1843b3a503e16d79697bd4dff8aa2cc))
- Quote all `location` and `server_name` directives, and escape quotes and backslashes. ([605dea8](https://github.com/forkline/ingress-nginx/commit/605dea88408f5fbea84880641473fc9e3e50c06c))
- Use `RawURLEncoding` instead of `URLEncoding` with padding removal. ([328c38e](https://github.com/forkline/ingress-nginx/commit/328c38e08e549a760ee083899f6d7aeafa19f58c))
- Quote `proxy_pass`. ([f356bde](https://github.com/forkline/ingress-nginx/commit/f356bde11ce8b043301284e00f864121db2e9ec2))
- Remove path from comment. ([b86d9f0](https://github.com/forkline/ingress-nginx/commit/b86d9f09e6aaf3b4276c8c07f19eb4fa32a9b4d3))

### Test

- Remove gRPC Fortune Teller. ([f5babef](https://github.com/forkline/ingress-nginx/commit/f5babefc88776a5e853bd74b87f3617eceee6ae6))

### Testing

- gzip: Reach ingress ([0e3e32d](https://github.com/forkline/ingress-nginx/commit/0e3e32d0aef58ec870495268de70058c901fbc19))
- Test fix ([d5ede33](https://github.com/forkline/ingress-nginx/commit/d5ede33f88239cf3a5397fd8ba8edc1cd227ec07))
- Test to assert nameservers are passed to lua ([7d927a3](https://github.com/forkline/ingress-nginx/commit/7d927a3f416b2e47554e0ab4a6f732ae8283e050))
- Test for ewma:after_balance function ([c03ac37](https://github.com/forkline/ingress-nginx/commit/c03ac375efeda744859fe90a3dba54b7fa2478f3))
- Testing that a secure cookie gets set when being in ssl mode

Signed-off-by: Fabian Topfstedt <topfstedt@schneevonmorgen.com> ([f03c8a8](https://github.com/forkline/ingress-nginx/commit/f03c8a85443b6429d18d717bb232278f19a397e3))
- Test to make sure dynamic cert works trailing dot in domains ([f771e72](https://github.com/forkline/ingress-nginx/commit/f771e7247a308bbdf6827872a0a415648e4efe5e))
- Test modsecurity-snippet ([1ee081c](https://github.com/forkline/ingress-nginx/commit/1ee081ccc805a0e12705223efc6fcf4269ccbe6c))
- Testing output of sarif file

Signed-off-by: James Strong <strong.james.e@gmail.com> ([e55a84e](https://github.com/forkline/ingress-nginx/commit/e55a84e8a031fe665e2aecac57ac4668b7debf21))
- Test the new e2e test images

Signed-off-by: James Strong <james.strong@chainguard.dev>

Signed-off-by: James Strong <james.strong@chainguard.dev> ([30d6f7e](https://github.com/forkline/ingress-nginx/commit/30d6f7e14043dafba40f078e9f2d8cb4e4f97c51))
- Testing auto change

Signed-off-by: James Strong <strong.james.e@gmail.com> ([ef5bf06](https://github.com/forkline/ingress-nginx/commit/ef5bf06c61b12e78e0e3168a3bc7fd4222fb5db0))
- Testing auto change

Signed-off-by: James Strong <strong.james.e@gmail.com> ([c015c62](https://github.com/forkline/ingress-nginx/commit/c015c628b5e4e07078a04a5f8c3ced8441c255a3))
- Test kind updates

Signed-off-by: James Strong <strong.james.e@gmail.com> ([d712dd9](https://github.com/forkline/ingress-nginx/commit/d712dd9d92f1f46eab5f55d4643f3a45a7a97f90))

### Tests

- Replace deprecated `grpc.Dial` by `grpc.NewClient`. ([0718c89](https://github.com/forkline/ingress-nginx/commit/0718c89203df192e797ca38271073770a3cd4312))
- Bump `test-runner` to v20240717-1fe74b5f. ([ebee23e](https://github.com/forkline/ingress-nginx/commit/ebee23ec2581adabf15f7c54a07674b30aa4d497))
- Bump `e2e-test-runner` to v20240729-04899b27. ([b0f8182](https://github.com/forkline/ingress-nginx/commit/b0f81825fe1969fe68db18e4bd5711f06ac6f154))
- Bump `e2e-test-runner` to v20240812-3f0129aa. ([f19e926](https://github.com/forkline/ingress-nginx/commit/f19e9265b0ca266c7f2bc5e4d2ac137479e8b842))
- Bump `e2e-test-runner` to v20240829-2c421762. ([6ca67b5](https://github.com/forkline/ingress-nginx/commit/6ca67b5296fb0e43652692ff43eba08b246cf294))
- Bump `e2e-test-runner` to v20241004-114a6abb. ([23c2552](https://github.com/forkline/ingress-nginx/commit/23c2552113935efff0b9c9484ca554b72959459e))
- Bump `e2e-test-runner` to v20241104-02a3933e. ([b3742aa](https://github.com/forkline/ingress-nginx/commit/b3742aa5de49e8a3a578ae0e3f07d6c54fee322a))
- Bump `e2e-test-runner` to v20241224-68ed4e7b. ([efa41b7](https://github.com/forkline/ingress-nginx/commit/efa41b7aaf60e386bc53f4b1343f37af7948d1d4))
- Bump Test Runner to v20250112-a188f4eb. ([5c7b74c](https://github.com/forkline/ingress-nginx/commit/5c7b74c5db46a2c675bbec4b1a3f8820804ef102))
- Bump Test Runner to v2.0.1. ([0109163](https://github.com/forkline/ingress-nginx/commit/010916360e1388f75ca69d889966cfab82be4a1a))
- Fallback to `yq`. ([4666f35](https://github.com/forkline/ingress-nginx/commit/4666f35c92a4de2ba823cfad7605cc2336ac9cd1))
- Bump Test Runner to v2.0.2. ([84db18a](https://github.com/forkline/ingress-nginx/commit/84db18ad79e87a8157b18cb4d283f8b89095f2c8))
- Bump Test Runner to v2.0.3. ([d0fc1e3](https://github.com/forkline/ingress-nginx/commit/d0fc1e3d387f2489913d8b6f2d392ee278537cd0))
- Bump Test Runner to v2.1.0. ([5803538](https://github.com/forkline/ingress-nginx/commit/580353804ab0af70399cb2aea0c8c687a1b74ea7))
- Bump Test Runner to v2.1.1. ([e0bf254](https://github.com/forkline/ingress-nginx/commit/e0bf2544b0c37e28dce85bf3831a2ea3b1ae8092))
- Bump Test Runner to v2.2.0. ([9e770b8](https://github.com/forkline/ingress-nginx/commit/9e770b849abd2c7017d20092e8877150c632b942))
- Bump Test Runner to v2.2.1. ([f598b64](https://github.com/forkline/ingress-nginx/commit/f598b64ade33367c43eec8f2a1debab2060fe69f))
- Add `ssl-session-*` config values tests. ([9baa28b](https://github.com/forkline/ingress-nginx/commit/9baa28b2a4a59db2f37a83736b0caffae86f8def))
- Enhance SSL Proxy. ([4c87d58](https://github.com/forkline/ingress-nginx/commit/4c87d58a2d7e5228b3a79fea4d9b37cb3233767e))
- Enable default backend access logging tests. ([4d05806](https://github.com/forkline/ingress-nginx/commit/4d05806dd61d6fe1570d21166ecbf99e04d62e7a))
- Bump Ginkgo to v2.24.0. ([09c9c57](https://github.com/forkline/ingress-nginx/commit/09c9c5714eb0300cb5044282b525af3b0b4efec6))
- Bump Ginkgo to v2.25.0. ([2bb9ebc](https://github.com/forkline/ingress-nginx/commit/2bb9ebc3e82082d298ad5357a01813053b276ad0))
- Bump Ginkgo to v2.25.1. ([679ab41](https://github.com/forkline/ingress-nginx/commit/679ab41d795ca655ca0b98ced98560ae853f44e9))
- Bump Test Runner to v2.2.2. ([44a8e37](https://github.com/forkline/ingress-nginx/commit/44a8e37d1aed27452584901ec7d28e710aa0f3d3))
- Bump Ginkgo to v2.25.2. ([8da30f4](https://github.com/forkline/ingress-nginx/commit/8da30f4d357161a20a4c751f1fb7df9c28dc7ea7))
- Bump Ginkgo to v2.25.3. ([5304ce4](https://github.com/forkline/ingress-nginx/commit/5304ce429a910a108f679973934e1029851b556b))
- Bump Test Runner to v2.2.3. ([f5315ff](https://github.com/forkline/ingress-nginx/commit/f5315ff26bd9afd31021d508d34f572bcfbdbfcd))
- Bump Test Runner to v2.2.4. ([64492d8](https://github.com/forkline/ingress-nginx/commit/64492d82da5a30aa5bdd5056afee9ccde80ae2ca))
- Bump Ginkgo to v2.27.2. ([dcfe691](https://github.com/forkline/ingress-nginx/commit/dcfe691576e0f1d5db2cd13a611c72a73e175039))
- Bump Test Runner to v2.2.5. ([0ad5060](https://github.com/forkline/ingress-nginx/commit/0ad50603c175e66fcf4aa4576645a989c923fafb))
- Bump Ginkgo to v2.27.3. ([29ff442](https://github.com/forkline/ingress-nginx/commit/29ff442cd49b4980afae756074afe870f9f92571))
- Bump Ginkgo to v2.27.5. ([7794b1c](https://github.com/forkline/ingress-nginx/commit/7794b1c2f2494303dbd51fdcaa6c24a9ce8be25c))
- Bump Test Runner to v2.2.6. ([d8ab0e8](https://github.com/forkline/ingress-nginx/commit/d8ab0e88145a7ff5438333ef76dd9706d9d59338))
- Bump Test Runner to v2.2.7. ([9ccfa50](https://github.com/forkline/ingress-nginx/commit/9ccfa50bc98e6236377994b811f23dbbb563dac4))
- Bump Ginkgo to v2.28.1. ([b4e67be](https://github.com/forkline/ingress-nginx/commit/b4e67beced864d80cda24b07f9074f308749a7ca))
- Bump Test Runner to v2.2.8. ([7450c78](https://github.com/forkline/ingress-nginx/commit/7450c78fdd3a8312185e998e691fef7b49f3fe3c))
- Bump Test Runner to v2.2.9. ([f402411](https://github.com/forkline/ingress-nginx/commit/f402411e6290437fc16ee0a37d9def73850c4dc0))

### Typo

- Docs/examples/rewrite/README.md ([ccacef6](https://github.com/forkline/ingress-nginx/commit/ccacef6a8a4462e7fbb57211b5ca1c9b99b5158e))

### UPT

- Annotation enhancement for resty-lua-waf ([04a89ce](https://github.com/forkline/ingress-nginx/commit/04a89ce2343be1ce91e80d4c5fa597014961afad))
- Align waf options ([bab521e](https://github.com/forkline/ingress-nginx/commit/bab521e81abc03252d4cdde6118aec7bf9e103c6))
- Updated e2e testing title for lua test ([3c2c0d0](https://github.com/forkline/ingress-nginx/commit/3c2c0d085872aac1e36168fb5f8a7b50643ab113))
- Updated e2e test and default true for process-multipart-body annotation ([bf03046](https://github.com/forkline/ingress-nginx/commit/bf03046a80e42f968f61e3e7727f4979971826ea))
- Add variable to define custom sampler host and port, add commituser ([31ffad8](https://github.com/forkline/ingress-nginx/commit/31ffad8fa161ea7535a2c76658a1f529aa778431))
- Modify configmap to include jaeger sampler host and jaeger sampler port ([d468cd5](https://github.com/forkline/ingress-nginx/commit/d468cd5ec5a6bbd10e7348a343481d4656565fbe))
- Opentracing configmap  documentation ([616b1e2](https://github.com/forkline/ingress-nginx/commit/616b1e239a5e8b204258c12c0431c1b6040c2d1f))

### Util

- Fix panic for empty `cpu.max` file. ([9597db7](https://github.com/forkline/ingress-nginx/commit/9597db7a80c6da4dd940147acab78c2ae6ff12a7))

### Values

- Add missing `controller.metrics.service.labels`. ([a069617](https://github.com/forkline/ingress-nginx/commit/a069617ef88a9a33b8cbd2f4d92a4c2574dac818))
- Fix indention of commented values. ([5806b58](https://github.com/forkline/ingress-nginx/commit/5806b5800328c2b8f45808cc4fb32f4f4913c13d))

### WIP

- Avoid reloads implementing Equals in structs ([75a4a61](https://github.com/forkline/ingress-nginx/commit/75a4a61254f2b0d7671f0a7dc58f5827f365b985))

### [docs]

- Fix Prerequisite section ([9bb18c9](https://github.com/forkline/ingress-nginx/commit/9bb18c9a17070d421504a15b571a5ed4ec1be718))

### \core\pkg\ingress\errors

- Delete unuseful variable ([c05b7a0](https://github.com/forkline/ingress-nginx/commit/c05b7a0094a26cc218787bef39f644c62f7b25ff))

### Add

- (admission-webhooks) ability to set securityContext for job-containers createSecret and patchWebhook ([ac1a336](https://github.com/forkline/ingress-nginx/commit/ac1a3363bde234252e750845d4a9547dad997e84))

### Admission

- Improved log messages for ingress name ([0122aba](https://github.com/forkline/ingress-nginx/commit/0122aba44dfeb72847c660641839ae2dc53d1a93))

### Bugfix

- Set canary attributes when initializing balancer ([41c925f](https://github.com/forkline/ingress-nginx/commit/41c925f390e4b572ea633e5f39cd6fee81a3a2bf))
- Do not merge catch-all canary backends with itself ([ec28539](https://github.com/forkline/ingress-nginx/commit/ec28539e43b681f17b3dfa1aaaf0838f97ee4dc5))
- Fixed duplicated seeds. ([0110629](https://github.com/forkline/ingress-nginx/commit/011062967ab28b8c1cd9903570d67e04f67e4a45))
- When secret includes ca.crt store it on disk even in dynamic cert mode ([5667ea5](https://github.com/forkline/ingress-nginx/commit/5667ea5d67a9ce518b3047bd6142f74a03be498b))
- Check all previously failing upstreams, not just the last one ([e2c6202](https://github.com/forkline/ingress-nginx/commit/e2c620232416676e4abed29be5188b7e59b87a43))
- Always update trafficShapingPolicy when using ewma as load-balance even if endpoints not change, otherwise update trafficShapingPolicy will not working ([8ca5450](https://github.com/forkline/ingress-nginx/commit/8ca5450e22e17095cadf3e87b1bd9520241457a2))

### Catalog

- Add traefik ([83fc1d8](https://github.com/forkline/ingress-nginx/commit/83fc1d8a0991d105917be68d5c44d2e8da64a828))
- Add alb-ingress-controller ([e220fb8](https://github.com/forkline/ingress-nginx/commit/e220fb84f9d1aede2c134fe86bbe049efa59c931))

### Chart

- Allow setting allocateLoadBalancerNodePorts ([d6a0f46](https://github.com/forkline/ingress-nginx/commit/d6a0f46c320a97dbbe4594ccdc9c2d4ac37a622e))

### Chart/ghaction

- Set the correct permission to have access to push a release ([4dda149](https://github.com/forkline/ingress-nginx/commit/4dda149ed02dad03cc6dcc862be1b8c93dc9824a))

### Cleanup

- Fix typos in framework.go ([3c05cc4](https://github.com/forkline/ingress-nginx/commit/3c05cc42258498059b97eddc01bef76237f6bca7))
- Remove ioutil for new go version ([787ea74](https://github.com/forkline/ingress-nginx/commit/787ea74b6beffeb14f20cbf3a0d95c315ee3d8e3))

### Configmap

- Option to not trust incoming tracing spans ([7d5452d](https://github.com/forkline/ingress-nginx/commit/7d5452d00b7584195c4f9239f8938447a2a1b3f5))

### Configmap.md

- Convert hyphens in name column to non-breaking-hyphens ([519f72e](https://github.com/forkline/ingress-nginx/commit/519f72e2f95bc774309608447b5005e1192b3aef))

### Controller

- Don't panic when ready condition in a endpointslice is missing ([4aef45c](https://github.com/forkline/ingress-nginx/commit/4aef45c17734d7dcba32a2032f97ac604c59d7bd))

### Core

- Allow disabling node-lister via flag ([adc2a7d](https://github.com/forkline/ingress-nginx/commit/adc2a7d74cb76e7774d97181e6873cc3e943ddf3))

### Deploy

- Add protocol to all Container/ServicePorts ([5b918e2](https://github.com/forkline/ingress-nginx/commit/5b918e2d95db0308e293ec3d2a0c54ccb3fd378d))

### Doc

- Update docs and fixed typos ([3402d07](https://github.com/forkline/ingress-nginx/commit/3402d07ff0cc1bc4c228eb668353c1dc0256a17f))
- Improvement ([452515c](https://github.com/forkline/ingress-nginx/commit/452515ca2f4961aef7e1624c596f4e75865f71f8))
- Fix deployment manifest example ([230d8e6](https://github.com/forkline/ingress-nginx/commit/230d8e67921c5747cc29baf230ba09431bcbfc4c))
- Update NEW_CONTRIBUTOR.md ([59d80f0](https://github.com/forkline/ingress-nginx/commit/59d80f05bc6a363576c2b3c4707d768f1b7b8f66))

### Echoheaders

- Nginx-slim:0.21 ([64f19e0](https://github.com/forkline/ingress-nginx/commit/64f19e0ec362fdc63b9f8df14017b964f37c674d))
- Release echoserver:1.7 ([03056aa](https://github.com/forkline/ingress-nginx/commit/03056aa563b805685d1b7d113625254fe1d7b5ab))

### Examples/nginx/rbac

- Give access to own namespace ([8cd18bc](https://github.com/forkline/ingress-nginx/commit/8cd18bc205f534acc93d77b5863a8b336a915418))

### Faq

- Gce: fixup some spelling errors ([8edf306](https://github.com/forkline/ingress-nginx/commit/8edf306391eb7502b8ab0c3fc336a958b5358940))

### Glbc

- Watch backend service ([a94d31e](https://github.com/forkline/ingress-nginx/commit/a94d31e87db7c8e71f3e788575a39c95bf463766))

### Helm

- Add new ingressClass resource ([cec3c0a](https://github.com/forkline/ingress-nginx/commit/cec3c0af3d713319749641a4b714688d610c8b9e))
- ServiceMonitor: sane default namespaceSelector ([a665a40](https://github.com/forkline/ingress-nginx/commit/a665a409da87028896dbb3d8dfc78cf8a154e275))
- Fix opentelemetry module installation for daemonset ([8c7981b](https://github.com/forkline/ingress-nginx/commit/8c7981bfa21c6dab551ed1f4c48db217b5d852ca))
- Use .Release.Namespace as default for ServiceMonitor namespace ([e17927b](https://github.com/forkline/ingress-nginx/commit/e17927ba52e852c3a604f257838ad1c00a2a8e86))
- Add resources to opentelemetry init container ([06c64bf](https://github.com/forkline/ingress-nginx/commit/06c64bf5672ed7167f0c030b750ec9062bc86c83))
- Opentelemetry addon allow configuration of registry with setting tag ([7e31f81](https://github.com/forkline/ingress-nginx/commit/7e31f818ff5871dcf54b55e656c052d455df6901))

### Images

- Use k8s-staging-test-infra/gcb-docker-gcloud ([a9029d2](https://github.com/forkline/ingress-nginx/commit/a9029d2bc7ae8626152e5fc010829374fc950950))
- Upgrade to Alpine 3.18 ([c2e1f34](https://github.com/forkline/ingress-nginx/commit/c2e1f34cbefe95014540b3fb833fbe9580339cc7))
- Upgrade to Alpine 3.18.4 ([362ec37](https://github.com/forkline/ingress-nginx/commit/362ec37778146520374e82c594d3aeae82c7f617))
- Upgrade to Alpine 3.18.5 ([6152695](https://github.com/forkline/ingress-nginx/commit/6152695c78075ef6ea17c97923d46e27e606e51c))

### Images/kube-webhook-certgen/rootfs

- Improvements ([260910c](https://github.com/forkline/ingress-nginx/commit/260910c0a0ad74639174044aaa10108af1911135))
- Add support for patching APIService objects ([9acf62d](https://github.com/forkline/ingress-nginx/commit/9acf62d867145a071ad715020292bbb887289ba0))
- Add missing tests and fix regression ([5452364](https://github.com/forkline/ingress-nginx/commit/54523641a89a2b026180eb1e779152b8e939b11a))

### Ingress

- Nginx controller watches referenced tls secrets ([00b2180](https://github.com/forkline/ingress-nginx/commit/00b2180a8fd3834ddb25b54380e614941002de2a))
- Adds configurable SSL redirect nginx controller ([3ae80fd](https://github.com/forkline/ingress-nginx/commit/3ae80fd3cc89f2190891e7f016043dd92d9d3a65))
- Use POD_NAMESPACE as a namespace in cli parameters ([e4de1e6](https://github.com/forkline/ingress-nginx/commit/e4de1e62b825ad2c7cbde9440b33eecc29cd6c81))
- Removed unnecessary whitespaces ([d8fbe2f](https://github.com/forkline/ingress-nginx/commit/d8fbe2f5824501c0527920f3b28b51fdb241b88a))

### Ingress-path-matching

- Doc typo ([a904527](https://github.com/forkline/ingress-nginx/commit/a90452774ac35a5059157a7b3090890ff23a626a))

### Ingress/controllers/README.md

- Fix a link ([9630813](https://github.com/forkline/ingress-nginx/commit/963081375beb55af23c04a2284de9794925cb11c))

### Ingress/controllers/nginx

- WebSocket documentation ([0373ce6](https://github.com/forkline/ingress-nginx/commit/0373ce6f31a09bbd9210d1ad2ad4423a2ad1fcc3))

### Issue

- 8739 fix doc issue ([6a83cdd](https://github.com/forkline/ingress-nginx/commit/6a83cddb0367a70754d96336f4dd7d96398fb25d))

### Minor

- Formatting ([d3b9525](https://github.com/forkline/ingress-nginx/commit/d3b952552a8152e84905c04054387c38dd8ac943))

### Misc

- Improve build scripts ([162ecb9](https://github.com/forkline/ingress-nginx/commit/162ecb97e98d3e118caaec4b1b034484e598a602))

### Multiple-ingress.md

- Rework page for clarity and less repetition ([572aac4](https://github.com/forkline/ingress-nginx/commit/572aac442170a147711a4e215a9896763fdcc7b0))

### Netlify

- Only trigger preview when there are changes in docs. ([a2f3036](https://github.com/forkline/ingress-nginx/commit/a2f3036e20bcd500daee9d9eeebf807257c75c7b))

### Nginx

- Also listen on ivp6 ([8fe1efe](https://github.com/forkline/ingress-nginx/commit/8fe1efe3965c0125577db8321092f9ae76340e58))

### Nginx/README.md

- Clarify app-root and fix example hyperlink ([6d3e966](https://github.com/forkline/ingress-nginx/commit/6d3e9666bad3799463e06cd3518bc48df41f9237))

### Nginx/pkg/config

- Delete unuseful variable ([f4da971](https://github.com/forkline/ingress-nginx/commit/f4da971b86f379b857cfad77f70781418c2f5eb6))

### Nginx/proxy

- Allow specifying next upstream behaviour ([5503e8d](https://github.com/forkline/ingress-nginx/commit/5503e8d0e91c31108aa025c4f69d562c962c92b2))

### Oauth-external-auth

- README.md: Link to oauth2-proxy, dashboard-ingress.yaml ([dab11bd](https://github.com/forkline/ingress-nginx/commit/dab11bdf4eea72c7cb60060bd9dbc59f75f523a3))

### Rbac-nginx

- ResourceNames cannot filter create verb ([4618fd2](https://github.com/forkline/ingress-nginx/commit/4618fd2f64a904b1949f1d0a9a76ebe6ab8cb719))

### Tcpproxy

- Increase buffer size to 16K ([5628f76](https://github.com/forkline/ingress-nginx/commit/5628f765fe883dd8c13ccd3084e9003ffd3e28d5))
- Increase buffer size to 16K ([909a818](https://github.com/forkline/ingress-nginx/commit/909a8185920c45f99292cce705dba07c770e6f74))

### Tracing

- Upgrade to dd-opentracing-cpp v1.3.7 ([05e5956](https://github.com/forkline/ingress-nginx/commit/05e5956545afa74ce332791075db14ac7ffa54ff))

### Types.go

- Fix typo in godoc ([e247fdb](https://github.com/forkline/ingress-nginx/commit/e247fdb7b6f2a712e251fb7fe46a217eeeadc1d2))

### Webhook

- Remove useless code. ([53a232f](https://github.com/forkline/ingress-nginx/commit/53a232f829220abdc40e42dc9ab44f9975b7e319))



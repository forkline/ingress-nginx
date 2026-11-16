#!/usr/bin/env bash
set -euo pipefail

script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
script="${script_dir}/prepare-nginx-base-image.sh"
tmp=$(mktemp -d)
trap 'rm -rf "${tmp}"' EXIT

cat > "${tmp}/docker" <<'MOCK_DOCKER'
#!/usr/bin/env bash
set -euo pipefail
printf '%s\n' "$*" >> "${DOCKER_LOG}"
if [[ "${1:-}" == "pull" && "${MOCK_FAIL_PULL_ONCE:-false}" == "true" \
      && ! -e "${MOCK_PULL_STATE}" ]]; then
  touch "${MOCK_PULL_STATE}"
  exit 1
fi
if [[ "${1:-}" == "buildx" && "${2:-}" == "build" \
      && "${MOCK_BUILD_FAIL:-false}" == "true" ]]; then
  exit 1
fi
MOCK_DOCKER
chmod +x "${tmp}/docker"

image=ghcr.io/forkline/ingress-nginx/nginx:1.31.8
local_image=ingress-nginx/nginx:2026.11.6

assert_has() {
  if ! grep -Fxq -- "$1" "${DOCKER_LOG}"; then
    printf 'Expected command missing: %s\n' "$1" >&2
    cat "${DOCKER_LOG}" >&2
    exit 1
  fi
}
assert_lacks() {
  if grep -Fq -- "$1" "${DOCKER_LOG}"; then
    printf 'Unexpected command found: %s\n' "$1" >&2
    cat "${DOCKER_LOG}" >&2
    exit 1
  fi
}

run_case() {
  local label=$1
  local expected_status=$2
  shift 2
  export DOCKER_LOG="${tmp}/${label}.log"
  export MOCK_PULL_STATE="${tmp}/${label}.pull-state"
  : > "${DOCKER_LOG}"
  if env PATH="${tmp}:${PATH}" NGINX_IMAGE="${image}" TAGNGINX=2026.11.6 \
    "$@" bash "${script}" > "${tmp}/${label}.stdout" 2>&1; then
    status=0
  else
    status=1
  fi
  if [[ "${status}" != "${expected_status}" ]]; then
    printf 'Case %s returned %s, expected %s\n' \
      "${label}" "${status}" "${expected_status}" >&2
    cat "${tmp}/${label}.stdout" >&2
    exit 1
  fi
}

# Registry hit: one amd64 pull and a tag, no multi-arch rebuild.
run_case registry-hit 0 NGINX_CHANGED=false
assert_has "pull --platform linux/amd64 ${image}"
assert_has "tag ${image} ${local_image}"
assert_lacks 'buildx'
assert_lacks 'rmi'
[[ $(wc -l < "${DOCKER_LOG}") -eq 2 ]]

# Changed build context: publish both architectures and re-pull amd64.
run_case forced-rebuild 0 NGINX_CHANGED=true
assert_has "buildx build --builder ingress-nginx --platform linux/amd64,linux/arm64 rootfs --tag ${image} --no-cache --push"
assert_has "pull --platform linux/amd64 ${image}"
assert_has "tag ${image} ${local_image}"
assert_lacks 'rmi'

# Registry miss: fall back to the cached multi-platform build, then pull amd64.
run_case registry-miss 0 NGINX_CHANGED=false MOCK_FAIL_PULL_ONCE=true
assert_has "buildx build --builder ingress-nginx --platform linux/amd64,linux/arm64 rootfs --tag ${image} --cache-from type=gha --cache-to type=gha,mode=max --push"
assert_has "tag ${image} ${local_image}"
assert_lacks 'rmi'
[[ $(grep -Fc "pull --platform linux/amd64 ${image}" "${DOCKER_LOG}") -eq 2 ]]

# A failed build must fail the step, without tagging a stale image.
run_case failed-build 1 NGINX_CHANGED=true MOCK_BUILD_FAIL=true
if grep -q '^tag ' "${DOCKER_LOG}"; then
  echo 'Must not tag an image after failure' >&2
  exit 1
fi

# A failed post-build pull must fail instead of tagging a stale image.
run_case failed-post-build-pull 1 NGINX_CHANGED=true MOCK_FAIL_PULL_ONCE=true
if grep -q '^tag ' "${DOCKER_LOG}"; then
  echo 'Must not tag an image after failure' >&2
  exit 1
fi

printf 'All five nginx image handoff scenarios passed.\n'

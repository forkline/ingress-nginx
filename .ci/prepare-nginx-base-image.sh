#!/usr/bin/env bash
set -euo pipefail

: "${NGINX_IMAGE:?NGINX_IMAGE must be set}"
: "${TAGNGINX:?TAGNGINX must be set}"

# The controller and E2E jobs consume only amd64 images. Always pull that
# platform, whether we build the multi-platform image or reuse it from GHCR.
if [[ "${NGINX_CHANGED:-false}" != "true" ]] && \
  docker pull --platform linux/amd64 "${NGINX_IMAGE}"; then
  echo "Reused nginx base image (${NGINX_IMAGE}) from registry"
else
  echo "Building nginx base image (${NGINX_IMAGE})..."
  docker buildx create --name ingress-nginx --bootstrap || true
  if [[ "${NGINX_CHANGED:-false}" == "true" ]]; then
    buildx_args=(--no-cache)
  else
    buildx_args=(--cache-from type=gha --cache-to type=gha,mode=max)
  fi
  docker buildx build \
    --builder ingress-nginx \
    --platform linux/amd64,linux/arm64 \
    rootfs \
    --tag "${NGINX_IMAGE}" \
    "${buildx_args[@]}" \
    --push
  echo "Built and pushed nginx base image (${NGINX_IMAGE})"
  docker pull --platform linux/amd64 "${NGINX_IMAGE}"
fi

# Never remove the source image before tagging it: on containerd-backed
# Docker engines, docker inspect can return an OCI index ID that docker rmi
# deletes along with the reference.
docker tag "${NGINX_IMAGE}" "ingress-nginx/nginx:${TAGNGINX}"

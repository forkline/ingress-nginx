#!/usr/bin/env bash
#
# Load a single-platform Docker image into kind.
# Docker's containerd-backed image store can retain a multi-platform index
# without the other platforms' blobs, causing "ctr: content digest ... not found"
# when kind load docker-image uses "ctr images import --all-platforms".
# See https://kind.sigs.k8s.io/docs/user/known-issues/#unable-to-kind-load-docker-images
#
set -euo pipefail

IMAGE="${1:?usage: load-kind-image.sh IMAGE CLUSTER_NAME NODE1[,NODE2...]}"
CLUSTER_NAME="${2:?missing kind cluster name}"
NODES="${3:?missing kind node list}"

FIRST_NODE="${NODES%%,*}"
NODE_ARCH="$(kubectl get node "${FIRST_NODE}" -o jsonpath='{.status.nodeInfo.architecture}')"
if [[ -z "${NODE_ARCH}" ]]; then
  echo "Could not determine Kubernetes node architecture for ${FIRST_NODE}" >&2
  exit 1
fi

ARCHIVE_DIR="$(mktemp -d)"
trap 'rm -rf "${ARCHIVE_DIR}"' EXIT
ARCHIVE="${ARCHIVE_DIR}/image.tar"

echo "[kind] exporting ${IMAGE} for linux/${NODE_ARCH}..."
docker image save --platform "linux/${NODE_ARCH}" --output "${ARCHIVE}" "${IMAGE}"
kind load image-archive --name="${CLUSTER_NAME}" --nodes="${NODES}" "${ARCHIVE}"

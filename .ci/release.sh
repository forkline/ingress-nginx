#!/bin/bash
set -e

check_dependencies() {
    local missing=()
    
    if ! command -v git-cliff &>/dev/null; then
        missing+=("git-cliff")
    fi
    
    if ! command -v helm-docs &>/dev/null; then
        if [ -x "$(go env GOPATH 2>/dev/null)/bin/helm-docs" ]; then
            export PATH="$PATH:$(go env GOPATH)/bin"
        else
            missing+=("helm-docs")
        fi
    fi
    
    if [ ${#missing[@]} -gt 0 ]; then
        echo "Error: Missing required dependencies:"
        for dep in "${missing[@]}"; do
            echo "  - $dep"
        done
        echo ""
        echo "Install with:"
        echo "  cargo install git-cliff"
        echo "  go install github.com/norwoodj/helm-docs/cmd/helm-docs@latest"
        exit 1
    fi
}

ensure_full_history() {
    if git rev-parse --is-shallow-repository 2>/dev/null | grep -q "true"; then
        echo "Shallow clone detected. Fetching full history and tags..."
        git fetch --unshallow --quiet
    fi
    if [ -z "$(git tag -l)" ]; then
        echo "No tags found. Fetching tags..."
        git fetch --tags --quiet
    fi
}

check_dependencies
ensure_full_history

if ! [ "$(git rev-list --count origin/main..HEAD 2>/dev/null || echo 1)" -eq 0 ]; then
    echo "There are commits in this branch. Please merge them first."
    echo "CHANGELOG template needs main commit ID."
    exit 1
fi

CURRENT_VERSION=$(cat TAG)
BASE_VERSION="v$(date +%-Y.%-m.%-d)"

if git rev-parse "$BASE_VERSION" >/dev/null 2>&1; then
    HIGHEST_SUFFIX=$(git tag -l "${BASE_VERSION}-*" \
        | sed "s/^${BASE_VERSION}-//" \
        | sort -n | tail -1)
    if [ -z "$HIGHEST_SUFFIX" ]; then
        NEW_VERSION="${BASE_VERSION}-1"
    else
        NEW_VERSION="${BASE_VERSION}-$((HIGHEST_SUFFIX + 1))"
    fi
else
    NEW_VERSION="$BASE_VERSION"
fi

echo "Current version: $CURRENT_VERSION"
echo "New version:     $NEW_VERSION"
echo ""

if [ "$NEW_VERSION" = "$CURRENT_VERSION" ]; then
    echo "Version unchanged. Nothing to do."
    exit 0
fi

echo "$NEW_VERSION" > TAG

TAG_VERSION="${NEW_VERSION#v}"
for img in nginx kube-webhook-certgen test-runner cfssl custom-error-pages e2e-test-echo fastcgi-helloserver go-grpc-greeter-server httpbun ext-auth-example-authsvc; do
    echo "$TAG_VERSION" > "images/$img/TAG"
done

NGINX_VERSION=$(grep 'export NGINX_VERSION=' images/nginx/rootfs/build.sh | sed "s/.*NGINX_VERSION=//")
sed -i "s|ghcr.io/forkline/ingress-nginx/nginx:.*|ghcr.io/forkline/ingress-nginx/nginx:$NGINX_VERSION|" NGINX_BASE

make update-version

make update-changelog

echo "Running helm-docs to update chart README..."
helm-docs --chart-search-root charts

echo "Updating README Supported Versions table..."
CHART_VERSION="${NEW_VERSION#v}"
NGINX_VERSION=$(grep 'export NGINX_VERSION=' images/nginx/rootfs/build.sh | sed "s/.*NGINX_VERSION=//")
ALPINE_VERSION=$(grep '^FROM alpine:' images/nginx/rootfs/Dockerfile | head -1 | sed 's/.*alpine://')
K8S_VERSIONS=$(grep '|    ✅' README.md | head -1 | cut -d'|' -f4 | xargs)

if [ -z "$K8S_VERSIONS" ]; then
    K8S_VERSIONS="1.35, 1.34, 1.33, 1.32, 1.31"
fi

NEW_ROW="|    ✅     | **$NEW_VERSION**         | $K8S_VERSIONS | $ALPINE_VERSION         | $NGINX_VERSION        | $CHART_VERSION           |"
sed -i "s/|    ✅/|    🔄/g" README.md
sed -i "/^| :-------:/a\\$NEW_ROW" README.md

git add .
git commit -m "release: prepare $NEW_VERSION"

echo ""
echo "✅ Release commit created for $NEW_VERSION"
echo "After pushing to main, tag and release are automatically done"

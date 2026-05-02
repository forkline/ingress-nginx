#!/bin/bash
set -e

if ! [ "$(git rev-list --count origin/main..HEAD)" -eq 0 ]; then
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

for img in nginx kube-webhook-certgen test-runner cfssl custom-error-pages e2e-test-echo fastcgi-helloserver go-grpc-greeter-server httpbun ext-auth-example-authsvc; do
    echo "$NEW_VERSION" > "images/$img/TAG"
done

sed -i "s|ghcr.io/forkline/ingress-nginx/nginx:.*|ghcr.io/forkline/ingress-nginx/nginx:$NEW_VERSION|" NGINX_BASE

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

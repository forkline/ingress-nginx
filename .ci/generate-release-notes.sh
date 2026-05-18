#!/bin/bash
set -e

VERSION="${1:?Usage: $0 <version> [prev_version]}"
PREV_VERSION="${2:-}"
CHANGELOG_ENTRY="${3:-/tmp/changelog.txt}"
NOTES_FILE="/tmp/release-notes.md"
REGISTRY="ghcr.io/forkline/ingress-nginx"

if [ -z "$PREV_VERSION" ]; then
    if git rev-parse --is-shallow-repository 2>/dev/null | grep -q "true"; then
        git fetch --unshallow --quiet 2>/dev/null || true
    fi
    if [ -z "$(git tag -l)" ]; then
        git fetch --tags --quiet
    fi
    PREV_VERSION=$(git tag --sort=-creatordate | grep -v "^${VERSION}$" | head -1)
fi

> "$NOTES_FILE"

{
    echo "## What's Changed"
    echo ""

    if [ -f "$CHANGELOG_ENTRY" ] && [ -s "$CHANGELOG_ENTRY" ]; then
        sed '/^---$/,$d' "$CHANGELOG_ENTRY" | grep -v '^## \['
        echo ""
    fi

    echo "### Container Images"
    echo ""
    echo "- \`${REGISTRY}/controller:${VERSION}\`"

    if [ -f images/nginx/TAG ]; then
        echo "- \`${REGISTRY}/nginx:$(cat images/nginx/TAG)\`"
    fi

    if [ -f images/kube-webhook-certgen/TAG ]; then
        echo "- \`${REGISTRY}/kube-webhook-certgen:$(cat images/kube-webhook-certgen/TAG)\`"
    fi

    echo ""
    echo "### Helm Chart"
    echo ""
    echo "- appVersion: \`${VERSION#v}\`"
    echo ""
    echo "### kubectl plugin"
    echo ""
    echo "Binaries available on the [GitHub Release](https://github.com/forkline/ingress-nginx/releases/tag/${VERSION}) page."

    if [ -n "$PREV_VERSION" ]; then
        echo ""
        echo "**Full Changelog**: https://github.com/forkline/ingress-nginx/compare/${PREV_VERSION}...${VERSION}"
    fi

} >> "$NOTES_FILE"

echo "Release notes written to ${NOTES_FILE}"

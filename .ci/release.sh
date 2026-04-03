#!/bin/bash
set -e

if ! [ "$(git rev-list --count origin/main..HEAD)" -eq 0 ]; then
    echo "There are commits in this branch. Please merge them first."
    echo "CHANGELOG template needs main commit ID."
    exit 1
fi

CURRENT_VERSION=$(cat TAG)
BASE_VERSION="v$(date +%Y.%m.%d)"

if git rev-parse "$BASE_VERSION" >/dev/null 2>&1; then
    HIGHEST_SUFFIX=$(git tag -l "${BASE_VERSION}.*" \
        | sed "s/^${BASE_VERSION}\.//" \
        | sort -n | tail -1)
    if [ -z "$HIGHEST_SUFFIX" ]; then
        NEW_VERSION="${BASE_VERSION}.1"
    else
        NEW_VERSION="${BASE_VERSION}.$((HIGHEST_SUFFIX + 1))"
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

make update-version

make update-changelog

git add .
git commit -m "release: prepare $NEW_VERSION"

echo ""
echo "✅ Release commit created for $NEW_VERSION"
echo "After pushing to main, tag and release are automatically done"

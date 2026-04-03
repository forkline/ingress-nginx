#!/bin/bash
set -e

if ! [ "$(git rev-list --count origin/main..HEAD)" -eq 0 ]; then
    echo "There are commits in this branch. Please merge them first."
    echo "CHANGELOG template needs main commit ID."
    exit 1
fi

echo "Current version: $(cat TAG)"
echo ""
echo "Enter new date-based version (format: vYYYY.MM.DD or vYYYY.MM.DD.N for multiple releases per day)"
echo "Example: v2026.04.03 or v2026.04.03.1"
read -p "New version: " NEW_VERSION

if [[ ! $NEW_VERSION =~ ^v[0-9]{4}\.[0-9]{2}\.[0-9]{2}(\.[0-9]+)?$ ]]; then
    echo "ERROR: Invalid version format. Must be vYYYY.MM.DD or vYYYY.MM.DD.N"
    exit 1
fi

echo "$NEW_VERSION" > TAG

make update-version

make update-changelog

git add .
git commit -m "release: prepare $NEW_VERSION"

echo ""
echo "✅ Release commit created for $NEW_VERSION"
echo "After pushing to main, tag and release are automatically done"

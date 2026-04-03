# Release Process

This document describes the release process for ingress-nginx.

## Overview

The release process is fully automated using git-cliff for changelog generation and GitHub Actions for artifact building and publishing.

## Date-based Versioning

This project uses **date-based versioning** (format: `vYYYY.MM.DD` or `vYYYY.MM.DD-N` for multiple releases per day, using semver pre-release notation for Helm compatibility).

Examples:
- `v2026.04.03` - First release on April 3, 2026
- `v2026.04.03-1` - Second release on April 3, 2026 (semver pre-release)

### Why Date-based Versioning?

Since this is a **maintenance-only project** that:
- Does NOT add new features
- Does NOT make breaking changes
- Only updates dependencies and fixes bugs

Semantic versioning is meaningless here. Date versions clearly answer: **"When was this last maintained for security?"**

## Creating a Release

### Prerequisites

- You must be on the `main` branch with no uncommitted changes
- All changes should be merged to `main` first
- Ensure you have `git-cliff` installed

### Automated Release Script

Run the release script:

```bash
.ci/release.sh
```

This script will:
1. Verify you're on `main` branch with no unmerged commits
2. Prompt for the new version (date-based format)
3. Update the TAG file
4. Update Chart.yaml (appVersion and version)
5. Generate CHANGELOG with git-cliff
6. Create a release commit

### Manual Release Process

If you need more control:

```bash
# 1. Update version in TAG file
echo "v2026.04.04" > TAG

# 2. Update version in all files
make update-version

# 3. Update CHANGELOG
make update-changelog

# 4. Commit the changes
git add .
git commit -m "release: prepare v2026.04.04"
```

### What Happens Next

After pushing to `main`:

1. **Auto-tag workflow** automatically creates a GPG-signed tag
2. **Release workflow** builds and publishes:
   - Docker images (controller, nginx, kube-webhook-certgen)
   - Helm chart (OCI to ghcr.io)
   - kubectl plugin binaries
   - GitHub release with changelog

## Version Updates

The `update-version` Makefile target updates:
- `TAG` - Main version file
- `charts/ingress-nginx/Chart.yaml` - Helm chart version and appVersion

## Changelog Generation

Changelogs are generated using [git-cliff](https://github.com/orhun/git-cliff) with conventional commits:

- `cliff.toml` - Main CHANGELOG generation
- `.ci/cliff-chart.toml` - Helm chart Artifact Hub annotations

### Chart Annotations

For Helm chart releases, chart-specific changes can be generated:

```bash
# Generate Artifact Hub changes annotation
git-cliff --config .ci/cliff-chart.toml --strip all v2026.04.03..HEAD
```

## Release Checklist

Before creating a release:

- [ ] All dependency updates merged
- [ ] Tests passing (`make test`)
- [ ] E2E tests passing (if applicable)
- [ ] CHANGELOG reviewed
- [ ] Version follows date-based format
- [ ] No uncommitted changes

## Troubleshooting

### Tag Already Exists

If you see "Tag already exists":
- The version in CHANGELOG.md already has a tag
- Update CHANGELOG.md with a new version

### Wrong Version Format

Version must match: `vYYYY.MM.DD` or `vYYYY.MM.DD-N`

Valid examples:
- `v2026.04.03`
- `v2026.12.31`
- `v2026.04.03-1`

Invalid examples:
- `v1.15.8` (semantic versioning)
- `2026.04.03` (missing 'v' prefix)
- `v2026.4.3` (must use zero-padding)

---
name: release
description: Prepare and publish a new release. Use when the user asks to release, cut a release, or publish a new version.
---

## Purpose

Release a new version of ingress-nginx using the automated release script and CI pipeline.

## When to use

Use this skill when:
- The user asks to release a new version
- The user asks to cut a release or publish
- The user asks to tag a new version

## Prerequisites

Before releasing, verify:

1. **Working tree is clean** — no uncommitted changes
2. **Dependencies installed** — `git-cliff` and `helm-docs` are required
3. **There are new commits** since the last tag worth releasing

Check with:
```bash
git status --short
command -v git-cliff && command -v helm-docs || echo "Missing deps"
git tag --sort=-creatordate | head -3
git log --oneline <latest_tag>..HEAD
```

### Required Dependencies

The release script now checks for these upfront and fails fast if missing:

```bash
# Install git-cliff (changelog generator from conventional commits)
cargo install git-cliff

# Install helm-docs (generates Helm chart README)
go install github.com/norwoodj/helm-docs/cmd/helm-docs@latest
```

**Note**: If `helm-docs` is installed via `go install`, it may be in `$(go env GOPATH)/bin` which isn't always in PATH. The release script auto-detects this, but you can also add it manually:

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Release Process

### Step 1: Run the release script

```bash
.ci/release.sh
```

This script **automatically**:
1. Detects and fixes shallow clones (fetches full history + tags)
2. Verifies you're on main with no unmerged commits
3. Computes the new version (`vYYYY.M.D`, auto-incrementing suffix if same-day)
4. Updates `TAG` file and all image `TAG` files
5. Updates `NGINX_BASE` reference
6. Runs `make update-version` (Chart.yaml, values.yaml)
7. Runs `make update-changelog` (git-cliff from conventional commits)
8. Squashes repetitive changelog entries (pre-commit hooks, digest updates, go modules, etc.)
9. Updates helm-docs and README Supported Versions table
10. Creates commit: `release: prepare <version>`

### Step 2: Push to main

```bash
git push origin main
```

### Step 3: Verify automation triggers

After pushing, the auto-tag workflow (`auto-tag.yml`) will:
1. Read the new version from CHANGELOG.md
2. Create a signed git tag
3. Push the tag

The tag then triggers build/release workflows for:
- Container images (`ghcr.io/forkline/ingress-nginx/controller`) — published for `linux/amd64,linux/arm64`
- Helm charts (`ghcr.io/forkline/helm-charts/ingress-nginx`)
- kubectl plugin binaries (GitHub release)

Monitor with:
```bash
gh run list --limit 5
```

## CI Pipeline Details

### Nginx base image caching

The nginx base image is tagged with the actual NGINX version (e.g. `1.30.1`) in addition to the release tag. This enables caching across workflow runs:

- **E2E workflow**: Pulls by NGINX version. If not in registry, builds multi-platform and pushes. Pulls amd64 variant back locally for controller/e2e builds.
- **Release workflow**: Checks if NGINX version manifest exists. Skips build if already published by E2E.

This means NGINX version changes are validated by E2E tests before they reach the release workflow.

### Published platforms

All images are published for `linux/amd64` and `linux/arm64`.

## What NOT to do

| Mistake | Why it's wrong | Fix |
|---------|---------------|-----|
| Manually editing CHANGELOG.md | git-cliff generates it from conventional commits automatically | Use `.ci/release.sh` |
| Manually editing TAG, Chart.yaml, values.yaml | `make update-version` handles all of them | Use `.ci/release.sh` |
| Creating git tags manually | `auto-tag.yml` creates signed tags automatically | Just push to main |
| Running from a feature branch | CHANGELOG generation needs main commit IDs | Checkout main first |
| Releasing with dirty working tree | Script will fail or produce incomplete release | Commit or stash changes first |
| Pushing directly to main when protected | Branch protection requires PR workflow | Create release branch and PR |
| Worrying about shallow clones | Script auto-detects and fetches full history | Just run `.ci/release.sh` |
| Running without git-cliff/helm-docs | Script will fail mid-process, leaving inconsistent state | Install deps first (script now checks upfront) |

## Protected Branch Workflow

If `main` is protected (requires PR):

1. **Create release branch** from current HEAD:
   ```bash
   git checkout -b release/v2026.X.Y
   ```

2. **Run release script on the branch**:
   ```bash
   .ci/release.sh
   ```

3. **Push and create PR**:
   ```bash
   git push origin release/v2026.X.Y
   git-api pr create --title "release: prepare v2026.X.Y" --body "..."
   ```

4. **After merge**, the auto-tag workflow triggers on `main`

**Why PR workflow?** The release script needs to compute versions from `origin/main` commit history. Running on a branch that tracks `main` works correctly because:
- `git rev-list --count origin/main..HEAD` returns 0 before the release commit
- git-cliff correctly generates CHANGELOG from main's commit IDs

## Troubleshooting

### "There are commits in this branch. Please merge them first."
You're on a branch with unmerged commits. Switch to main and merge first.

### "Error: Missing required dependencies: git-cliff / helm-docs"
The release script now checks for dependencies before making any changes. Install them:
```bash
cargo install git-cliff
go install github.com/norwoodj/helm-docs/cmd/helm-docs@latest
export PATH="$PATH:$(go env GOPATH)/bin"  # if helm-docs not found
```

### "Shallow clone detected. Fetching full history and tags..."
The release script detected a shallow clone (e.g. CI with `fetch-depth: 1`). It automatically fetches full history and tags so git-cliff can generate a correct CHANGELOG. No action needed.

### "Version unchanged. Nothing to do."
Already released this version today. If you need a same-day re-release, the script auto-increments the suffix (`v2026.5.14-1`, `v2026.5.14-2`, etc.) only if the base tag already exists. Check `cat TAG` to see current version.

### "Protected branch update failed for refs/heads/main"
The main branch requires PR workflow. Create a release branch instead:
```bash
git checkout -b release/v2026.X.Y
.ci/release.sh
git push origin release/v2026.X.Y
git-api pr create --title "release: prepare v2026.X.Y" --body "..."
```

### Release script failed mid-process (before dependency check fix)
Old versions of the script would update TAG/Chart.yaml then fail at git-cliff. Fix by:
1. Resetting changes: `git checkout -- .`
2. Installing missing dependencies
3. Re-running `.ci/release.sh`

**Current script**: Checks dependencies FIRST, fails fast before any changes.

### Auto-tag workflow didn't trigger
Ensure:
- The CHANGELOG.md has a new version entry as the first `## [v...]` heading
- The tag doesn't already exist: `git tag -l | grep <version>`
- The `PAT` and `GPG_PRIVATE_KEY` secrets are configured

## Key Files

| File | Role |
|------|------|
| `.ci/release.sh` | Main release script — runs everything |
| `.ci/squash-changelog.py` | Post-processes CHANGELOG to squash repetitive entries |
| `TAG` | Current version file |
| `images/*/TAG` | Per-image version tags |
| `NGINX_BASE` | Base nginx image reference |
| `cliff.toml` | git-cliff configuration (conventional commit parsing) |
| `CHANGELOG.md` | Generated changelog (auto-tag reads version from here) |
| `.github/workflows/auto-tag.yml` | Creates signed tag on push to main |
| `.github/workflows/e2e.yml` | E2E tests — builds nginx if NGINX version changed |
| `.github/workflows/docker_images.yml` | Release images — skips nginx build if already published |
| `Makefile` | `update-version`, `update-changelog`, and `release` targets |

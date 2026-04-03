# AGENTS.md - Critical Information for Maintaining This Repository

This document provides essential context for AI coding agents and contributors working on this repository.

## Repository Status: Maintenance Mode

**This repository is in strict maintenance mode.** We do NOT add new features or make significant behavior changes. The sole purpose is to keep the existing functionality working correctly while updating dependencies.

### Primary Maintenance Activities

1. **Dependency Updates** (automated via Renovate):
   - NGINX base image updates (`images/nginx/TAG`, `NGINX_BASE`)
   - Go modules (`go.mod`, `go.sum`, `GOLANG_VERSION`)
   - Kubernetes client modules
   - Security vulnerability fixes

2. **Bug Fixes**: Only critical bugs that break existing functionality

3. **Build/Release Reliability**: CI/CD workflow improvements that don't change behavior

## Critical Rule: No Regressions Allowed

**Every change must preserve existing behavior exactly.** This is non-negotiable.

### Why This Matters
- Users depend on predictable behavior
- No new features means no behavior changes
- Dependency updates can silently introduce regressions
- Test coverage must prevent regressions

### How to Ensure No Regressions

1. **Test Coverage Baseline** (Issue #4 - HIGH PRIORITY):
   - Coverage must NEVER decrease from baseline
   - All PRs must pass coverage checks
   - Renovate PRs validated automatically
   - Critical update lanes must have documented coverage

2. **Before Any Change**:
   - Run full test suite: `make test`
   - Run E2E tests for NGINX/template changes
   - Check that all existing tests still pass
   - Verify behavior unchanged in affected areas

3. **Dependency Update Protocol**:
   - NGINX updates: Verify template rendering, Lua integration
   - Go/K8s updates: Verify API client behavior, controller logic
   - Security fixes: Test edge cases thoroughly

## Maintenance Lanes (From MAINTENANCE.md)

### NGINX Lane
**Files**: `images/nginx/TAG`, `NGINX_BASE`, `images/nginx/rootfs/`

**What to check**:
- NGINX template generation (`internal/ingress/controller/template/`)
- Lua script integration (`rootfs/etc/nginx/lua/`)
- NGINX configuration rendering
- E2E tests must pass

**Update process**:
1. Update `images/nginx/TAG`
2. Update image build inputs
3. Update `NGINX_BASE`
4. Run: `make image`, smoke checks, `make kind-e2e-test`
5. Verify coverage maintained

### Go and Kubernetes Lane
**Files**: `GOLANG_VERSION`, `go.mod`, `go.sum`

**What to check**:
- Unit tests: `make test`
- Lint: `make golint-check`
- Build: `make build`
- E2E tests: `make kind-e2e-test`
- Coverage: must not decrease

**Update process**:
1. Update `GOLANG_VERSION`
2. Run `go get -u` or accept Renovate PR
3. Run: `make test`, `make build`, `make kind-e2e-test`
4. Verify coverage report shows no decrease

### Security Lane
**Triggers**: Renovate PRs, Trivy alerts, `govulncheck` output

**What to check**:
- Security fixes don't break edge cases
- Behavior unchanged for affected packages
- Test coverage for security-critical code paths

**Update process**:
1. Review security alert details
2. Accept Renovate PR or manual update
3. Run full test suite
4. Verify affected package coverage maintained
5. Check E2E for integration scenarios

## Testing Requirements

### Current Test Infrastructure
- **Unit tests**: 151 test files (`make test`)
- **E2E tests**: 89 test files (`make kind-e2e-test`)
- **Lua tests**: (`make lua-test`)
- **Helm tests**: (`make helm-test`)

### Test Coverage Implementation (Issue #4)
**STATUS**: Not yet implemented - HIGH PRIORITY

**Required implementation**:
1. Add `-coverprofile=coverage.out -covermode=atomic` to `test/test.sh`
2. Create `.github/codecov.yml` with threshold: **0% decrease allowed**
3. Add Codecov integration to `.github/workflows/go.yml`
4. Enforce coverage checks on ALL PRs
5. Track coverage for maintenance lanes separately

**Coverage policy**:
- **Project threshold**: 0% decrease from baseline (baseline becomes floor)
- **Patch threshold**: 80% for new code
- **Coverage must never decrease** - this is the golden rule
- Renovate PRs automatically checked
- Coverage badge in README

### Critical Test Areas (Must Have Good Coverage)
1. `internal/nginx/` - NGINX template rendering
2. `internal/ingress/controller/` - Core controller logic
3. `internal/ingress/annotations/` - All annotation handlers (extensive tests exist)
4. `internal/k8s/` - Kubernetes API client
5. `internal/admission/` - Admission webhook
6. Template generation logic
7. Lua integration points

## Branch Model & Release Process

### Branch Model
- **Single branch**: `main` only (no develop, no release branches)
- **Release tags**: Date-based tags (e.g., `v2026.04.03`) on `main`
- **No long-lived branches** except `main`

### Versioning Strategy
**Date-based versioning** (`vYYYY.MM.DD` or `vYYYY.MM.DD.N` for multiple releases per day)

**Why date-based instead of semantic versioning?**
- We don't add features (no MINOR version meaning)
- We don't make breaking changes (no MAJOR version meaning)
- Only dependency updates and bug fixes (PATCH-only)
- Date versions answer: "When was this last maintained for security?"
- More honest about the nature of maintenance-only software

### Release Process
1. Update CHANGELOG.md with new date-based version
2. Auto-tag workflow creates tag `vYYYY.MM.DD` on CHANGELOG update
3. GitHub Actions build and publish:
   - Container images to `ghcr.io/forkline/ingress-nginx/controller`
   - Helm charts to `ghcr.io/forkline/helm-charts/ingress-nginx`
   - kubectl plugin binaries to GitHub release
4. No manual release steps (fully automated)

### Pull Request Requirements
**All PRs must**:
- Pass pre-commit checks (`pre-commit run --all-files`)
- Pass unit tests (`make test`)
- Pass E2E tests (for behavior changes)
- Pass coverage checks (once Issue #4 implemented)
- Have no coverage decrease
- Preserve existing behavior exactly

## Renovate Integration

**Renovate automates dependency updates** but every PR must be validated.

### Renovate PR Validation Checklist
For every Renovate PR:

1. **Check PR type**:
   - NGINX base image → Run E2E tests, verify template coverage
   - Go module → Run unit tests + E2E, verify affected package coverage
   - Kubernetes module → Run full test suite, verify k8s/ coverage
   - Security fix → Run tests, verify edge case coverage

2. **Run tests**:
   ```bash
   make test
   make kind-e2e-test  # if behavior could change
   ```

3. **Check coverage** (once implemented):
   - Coverage report shows no decrease
   - Affected packages maintain coverage
   - Codecov PR check passes

4. **Manual verification**:
   - Review diff for unexpected changes
   - Check changelog for breaking changes
   - Verify no behavior modifications

5. **Approve if**:
   - All tests pass
   - Coverage maintained
   - Behavior unchanged
   - No breaking changes in dependency

## Workflow Files (CI/CD)

### Active Workflows
- `.github/workflows/pre-commit.yml` - Linting, formatting checks
- `.github/workflows/go.yml` - Go tests, lint, build (Phase 2 of MAINTENANCE.md)
- `.github/workflows/e2e.yml` - E2E test suite
- `.github/workflows/helm.yml` - Helm chart tests and publishing
- `.github/workflows/docker_images.yml` - Container image building
- `.github/workflows/security.yml` - Trivy, govulncheck, dependency review
- `.github/workflows/pages.yml` - Documentation publishing

### Coverage Workflow (To Be Added)
**File**: `.github/workflows/go.yml` (modification)

**Add**:
```yaml
- name: Run tests with coverage
  run: make test

- name: Upload coverage to Codecov
  uses: codecov/codecov-action@v4
  with:
    files: ./coverage.out
    flags: unittests
    fail_ci_if_error: true
```

## Code Style & Conventions

### Go Code
- Standard Go conventions
- Use existing libraries and patterns
- Check neighboring files for style
- Run `make golint-check` before committing

### Testing Style
- Use `testing` package (existing pattern)
- Testify for assertions (already in use)
- Mock external dependencies
- Focus on behavior preservation

### No Comments Policy
**IMPORTANT**: Do NOT add comments to code unless explicitly requested by user.
Follow the pattern from existing code.

## What NOT to Do

### Prohibited Actions
❌ Add new features
❌ Change existing behavior
❌ Add new configuration options
❌ Modify NGINX templates for new capabilities
❌ Add new annotation types
❌ Refactor for "cleaner code" (unless fixing bugs)
❌ Change API contracts
❌ Add unnecessary abstractions
❌ Commit code with decreased test coverage

### Allowed Actions
✅ Update dependencies (with testing)
✅ Fix critical bugs
✅ Improve CI/CD reliability
✅ Fix security vulnerabilities
✅ Update documentation for accuracy
✅ Add tests to increase coverage
✅ Refactor to fix bugs (minimal changes)

## Decision Framework for Changes

**Before making any change, ask**:

1. **Is this a dependency update?** → Follow maintenance lane protocol
2. **Is this a bug fix?** → Verify it doesn't change expected behavior
3. **Is this a CI improvement?** → Must not change code behavior
4. **Is this a new feature?** → **STOP - Not allowed**
5. **Does this decrease coverage?** → **STOP - Not allowed**
6. **Does this change behavior?** → **STOP - Not allowed unless fixing bug**

## Key Files Reference

### Configuration Files
- `MAINTENANCE.md` - Maintenance plan and phases
- `CONTRIBUTING.md` - Contribution guidelines
- `renovate.json` - Renovate configuration
- `.github/codecov.yml` - (To be created for Issue #4)
- `.golangci.yml` - Linting configuration

### Critical Code Areas
- `internal/ingress/controller/` - Core controller
- `internal/nginx/` - NGINX integration
- `internal/ingress/annotations/` - Annotation handlers
- `internal/k8s/` - Kubernetes client
- `internal/admission/` - Admission webhook
- `rootfs/etc/nginx/lua/` - Lua scripts
- `test/e2e/` - E2E test suite

### Test Files
- `test/test.sh` - Unit test runner (add coverage flags here)
- `test/e2e/run-kind-e2e.sh` - E2E test runner
- `Makefile` - Test targets (`make test`, `make kind-e2e-test`)

## Issue #4: Test Coverage Implementation

**This is the highest priority maintenance task.**

### Why Critical
- Dependency updates happen constantly (Renovate)
- Each update risks regression
- No coverage = no regression detection
- Maintenance mode depends on preserving behavior

### Implementation Priority
**Phase 1** (immediate):
- Add coverage to `test/test.sh`
- Generate baseline report
- Upload to Codecov

**Phase 2** (next):
- Integrate into CI workflow
- Add PR checks
- Fail on coverage decrease

**Phase 3** (enforcement):
- Configure thresholds (0% decrease)
- Document requirements
- Add badge to README

### Success Metrics
- Baseline established: ✅
- CI integration: Pending
- Coverage never decreases: Pending
- Renovate PRs validated: Pending
- Coverage badge: Pending

## Summary

**Remember**:
1. Maintenance mode = preserve behavior, update dependencies
2. Test coverage = regression prevention foundation
3. Coverage must NEVER decrease (Issue #4)
4. All changes validated by tests + coverage
5. Renovate automates updates, we validate them
6. Single branch (`main`), simple tags (`vX.Y.Z`)
7. No new features, no behavior changes

**Before every change**:
- Run tests
- Check coverage (once implemented)
- Verify behavior unchanged
- Ensure no regressions

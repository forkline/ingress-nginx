# Test Coverage Guide

This document explains the test coverage tracking implementation for the ingress-nginx fork.

## Overview

This repository enforces strict test coverage policies to prevent regressions during dependency updates. Since the repository is in maintenance mode with no new features, preserving existing behavior is critical.

## Coverage Requirements

### Mandatory Rules

1. **Coverage must never decrease** from baseline
2. **All PRs must pass coverage checks**
3. **New code must have at least 80% coverage** (patch threshold)
4. **Renovate PRs are automatically validated** for coverage

### Why Coverage Matters

- **Maintenance Mode**: We preserve existing behavior, not add features
- **Dependency Updates**: Updates can silently introduce regressions
- **Safety Net**: Coverage detects unintended behavior changes
- **Critical Paths**: Ensures important code remains tested across updates

## Running Tests with Coverage

### Local Development

```bash
# Run tests with coverage (generates coverage.out)
make test

# View coverage summary
go tool cover -func=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html

# Open in browser
open coverage.html
```

### Coverage Report Location

After running `make test`, the coverage file is generated at:
- `coverage.out` - Coverage data in Go cover format

## CI/CD Integration

### GitHub Actions Workflow

The `.github/workflows/go.yml` workflow automatically:

1. Runs unit tests with coverage enabled
2. Uploads coverage to Codecov
3. Fails the CI if coverage upload fails

### Codecov Configuration

The `.github/codecov.yml` file configures:

- **Project threshold**: 0% decrease allowed (baseline becomes floor)
- **Patch threshold**: 80% for new code
- **Exclusions**: Test files, generated code, vendor, docs, etc.

### Coverage Badge

The README.md includes a Codecov badge showing current coverage status.

## Critical Coverage Areas

These directories are especially important for coverage:

### `internal/nginx/`
- NGINX template generation
- Lua integration
- Configuration rendering

### `internal/ingress/controller/`
- Core controller logic
- Ingress processing
- Configuration management

### `internal/ingress/annotations/`
- All annotation handlers
- Annotation parsing and application

### `internal/k8s/`
- Kubernetes API client interactions
- Resource watching
- Event handling

### `internal/admission/`
- Admission webhook logic
- Validation and mutation

## Coverage by Update Lane

Different dependency update lanes have specific coverage requirements:

### NGINX Base Image Updates
- Verify `internal/nginx/` coverage maintained
- Run E2E tests for template changes
- Check Lua integration coverage

### Go Module Updates
- Verify affected packages maintain coverage
- Check for API compatibility changes
- Run full test suite

### Kubernetes Module Updates
- Verify `internal/k8s/` and `internal/ingress/controller/` coverage
- Check client behavior changes
- Validate API compatibility

### Security Fixes
- Run full test suite
- Verify edge case coverage
- Check affected packages

## Understanding Coverage Reports

### Overall Coverage

The total percentage of code covered by tests across the entire project.

### Package Coverage

Coverage broken down by Go package. Helps identify areas needing attention.

### File Coverage

Line-by-line coverage within files. Shows exactly which code paths are tested.

## Troubleshooting

### Coverage Upload Fails

1. Check that `coverage.out` was generated
2. Verify Codecov token is configured (for private repos)
3. Check GitHub Actions logs for details

### Coverage Decreases

1. Review which files/lines lost coverage
2. Add tests for uncovered code paths
3. Ensure tests still pass with changes
4. Check for refactoring that removed tested code

### Cannot Generate Coverage

1. Ensure Go 1.20+ is installed
2. Run `go mod download` first
3. Check for build errors in test compilation
4. Verify PKG environment variable is set

## Best Practices

### Writing Testable Code

- Keep functions focused and small
- Avoid global state where possible
- Use dependency injection
- Mock external dependencies

### Increasing Coverage

- Focus on critical code paths first
- Test error conditions and edge cases
- Add integration tests for complex logic
- Don't sacrifice test quality for coverage numbers

### Maintaining Coverage

- Run tests locally before pushing
- Review Codecov PR comments
- Address coverage decreases immediately
- Update tests when refactoring

## Renovate Integration

Renovate automatically creates PRs for dependency updates. Each Renovate PR:

1. Triggers the full test suite with coverage
2. Uploads coverage to Codecov
3. Must pass coverage checks before merge
4. Is reviewed for coverage changes

### Renovate PR Validation Checklist

- [ ] All tests pass
- [ ] Coverage maintained or increased
- [ ] No unexpected coverage decreases
- [ ] Affected packages have adequate coverage

## Metrics and Reporting

### Codecov Dashboard

Visit the Codecov dashboard for:
- Coverage trends over time
- Package-level breakdowns
- PR coverage diffs
- Historical comparisons

### Coverage Metrics

- **Line Coverage**: Percentage of code lines executed
- **Function Coverage**: Percentage of functions called
- **Branch Coverage**: Percentage of code branches taken

## Future Improvements

Phase 4 and 5 of Issue #4 will add:

1. **Per-package coverage targets**: Specific thresholds for critical areas
2. **E2E coverage integration**: Behavior-level coverage tracking
3. **Coverage trend alerts**: Notifications for coverage erosion
4. **Renovate-specific checks**: Lane-specific coverage validation

## Questions?

- Open an issue for coverage-related problems
- Check existing issues before creating new ones
- Reference this document when discussing coverage

## References

- [Go Cover Documentation](https://pkg.go.dev/cmd/cover)
- [Codecov Documentation](https://docs.codecov.com)
- [Testing Best Practices](https://golang.org/doc/effective_go#testing)
- Issue #4: Test Coverage Implementation
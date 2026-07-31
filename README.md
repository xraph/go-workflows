# xraph/go-workflows

Reusable GitHub Actions workflows for xraph Go repositories.

Consumers pin to the moving major tag:

```yaml
jobs:
  ci:
    uses: xraph/go-workflows/.github/workflows/go-ci.yml@v1
    secrets: inherit
```

## Workflows

| workflow | purpose |
|---|---|
| `go-ci.yml` | test matrix, lint, verify, security |
| `go-release.yml` | tag-push and manual-dispatch releases |
| `go-binary-release.yml` | GoReleaser cross-platform binaries |
| `codeql.yml` | CodeQL analysis for Go |

## Pinned tool versions

| tool | version |
|---|---|
| golangci-lint | v2.12.2 |
| gosec | v2.28.0 |
| govulncheck | v1.6.0 |
| goreleaser | v2.17.1 |
| actionlint | v1.7.12 |

Bumps are deliberate commits to this repository. See `CHANGELOG.md`.

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

## Recovering a failed dispatch release

On the `workflow_dispatch` path the tag is pushed before the GitHub release is
created. Tests and lint run first, so the common failures happen before anything
is mutated — but if release creation itself fails, the tag is already public.

To retry:

```bash
git push --delete origin v1.2.3
git fetch --prune --prune-tags
```

Then re-run the dispatch. A GitHub release requires its tag to exist, so
release-then-tag is not possible; this is a known trade-off.

## Caller permissions

`go-release.yml` needs `contents: write` **declared by the caller**. A reusable
workflow cannot grant itself permissions. Omitting it produces a failure at
release-creation time, not at parse time.

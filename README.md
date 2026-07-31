# xraph/workflows

Reusable GitHub Actions workflows for xraph Go repositories.

Consumers pin to the moving major tag:

```yaml
jobs:
  ci:
    uses: xraph/workflows/.github/workflows/go-ci.yml@v1
    secrets: inherit
```

## Workflows

| workflow | purpose |
|---|---|
| `go-ci.yml` | test matrix, lint, verify, security |
| `go-release.yml` | tag-push and manual-dispatch releases |
| `go-binary-release.yml` | GoReleaser cross-platform binaries |
| `codeql.yml` | CodeQL analysis for Go |
| `semantic-release.yml` | conventional-commit versioning and GitHub releases |

## Pinned tool versions

| tool | version |
|---|---|
| golangci-lint | v2.12.2 |
| gosec | v2.28.0 |
| govulncheck | v1.6.0 |
| goreleaser | v2.17.1 |
| actionlint | v1.7.12 |

Bumps are deliberate commits to this repository. See `CHANGELOG.md`.

## Module path for /v2+

`module-path` defaults to `github.com/<owner>/<repo>`, with no major-version
suffix. Repositories that have gone through a `/v2`, `/v3`, etc. major-version
bump must set `module-path` explicitly (e.g. `github.com/xraph/foo/v2`) — the
default will not add the suffix for you.

## Submodules input

`submodules` (on `go-release.yml`) is for **packages within the root module** —
directories that share the root `go.mod` but get their own install line and
proxy-warm call in the release notes, like `go-utils`'s `errs` and `log`. It is
**not** for repositories containing genuinely nested Go modules (their own
`go.mod`, e.g. `discovery/go.mod` in `farp`). Go resolves versions for a nested
module from a `<submodule>/vX.Y.Z` tag, and this workflow only ever creates the
bare `$VERSION` tag on the root module — so a nested module listed here gets an
install line and proxy warm that fail. Repositories with real nested modules
need their own per-submodule tagging step; this workflow does not implement
that.

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

A called (`workflow_call`) workflow can only **narrow** the permissions the
caller's token already has — it can never widen them. Job-level `permissions`
inside the reusable workflow are therefore a ceiling, not a grant: every
caller must declare, at the top level of its own workflow file, whatever
permissions its jobs actually need.

- `go-release.yml` needs `contents: write` **declared by the caller**.
  Omitting it produces a failure at release-creation time, not at parse time.
- `codeql.yml` needs `actions: read`, `contents: read`, and
  `security-events: write` **declared by the caller**. Omitting it produces a
  `startup_failure` immediately, since the default repo token is read-only.

## Binary track

`go-binary-release.yml` wraps GoReleaser. The consumer repository owns its
`.goreleaser.yml` — the config is genuinely per-project (build matrix, ldflags,
Homebrew and Scoop blocks) and is not something this repository should dictate.

Secrets, all optional: `HOMEBREW_TAP_TOKEN`, `SCOOP_BUCKET_TOKEN`,
`GPG_PRIVATE_KEY`. Pass them explicitly rather than with `secrets: inherit` so
the blast radius of a token stays visible in the caller.

No repository consumes this yet. `forge`, `forgeui` and `smart-form` already
carry a `.goreleaser.yml` and are the intended first consumers.

## semantic-release

`semantic-release.yml` is language-agnostic. It reads conventional commits,
computes the next version, writes the changelog, and creates the GitHub release.
Anything ecosystem-specific belongs in the consumer's own `.releaserc.json`,
via `@semantic-release/exec`.

The pinned set is Node 22 with `semantic-release@25`. The current plugin
generation declares `engines.node: ^22.14.0 || >=24.10.0`, so Node 20 fails at
install. `spindle` and `farp` run `@23` on Node 20; they are unaffected.

### Caller permissions

```yaml
permissions:
  contents: write         # tag, and commit CHANGELOG.md via @semantic-release/git
  issues: write           # comment on resolved issues
  pull-requests: write    # and on the PRs that closed them
```

Omitting `issues:` or `pull-requests:` only logs a warning and skips the
comment. Omitting `contents: write` fails the release.

### Branch protection

`@semantic-release/git` commits `CHANGELOG.md` back to `main`. If `main` is
protected without an allowance for `github-actions[bot]`, that push fails and
the release aborts partway. Either exempt the bot, or drop
`@semantic-release/git` and let the changelog live only in the GitHub release.

## Security reporting

`go-ci.yml` runs gosec with `-no-fail -fmt sarif` and uploads the result, so
findings appear in the repository's Security tab with history, triage and
dismissal — not only in job logs. It then fails the build when
`gosec-fail-on-findings` (default `true`) is set. Pass `false` for
advisory-only scanning.

SARIF upload is skipped on pull requests from forks, which have no write token.
Those runs degrade to log-only output.

`codeql.yml` defaults to the `security-extended,security-and-quality` query
packs. Pass `queries: ''` for CodeQL's standard set.

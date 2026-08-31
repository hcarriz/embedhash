# 0002: Automate Go tool releases with Release Please

Status: Accepted
Date: 2026-08-31

The release-credential portion of this decision is superseded by
[ADR 0003](0003-use-github-token-for-releases.md).

## Context

`embedhash` needs reproducible semantic-version releases without requiring a
maintainer to calculate each next version. The repository uses Conventional
Commits and contains both an importable library and a command under
`cmd/embedhash`.

The command is intended for module-scoped use through Go's `tool` directive and
`go tool embedhash`. It is not distributed as standalone binary archives.

## Decision

Use Release Please on GitHub. Conventional Commits determine the proposed
semantic version and generated changelog entries. Release Please opens and
maintains a release pull request; merging that pull request is the maintainer's
publication approval and causes the gated workflow to create a `v`-prefixed Go
module tag and GitHub Release.

Before Release Please runs, GitHub Actions must test the declared Go 1.24 floor
and pass the stable-toolchain release gate: formatting and whitespace, vet,
race and shuffled tests, builds, tidy module metadata, absence of root
`replace` directives or `go.work`, Staticcheck, and govulncheck. Actions and
analysis tools are pinned to reviewed versions. Staticcheck and govulncheck are
Go 1.27 tool directives in a separate `tools/go.mod`, keeping them out of the
public module graph.

The workflow publishes source, the Go module tag, changelog, and GitHub Release
notes only. It does not use GoReleaser or attach compiled binaries.

## Consequences

- Maintainers review a calculated version and release notes in an ordinary pull
  request rather than calculating a version manually.
- Consumers can pin the command with `go get -tool ...@vX.Y.Z` and run it with
  `go tool embedhash` without a globally installed binary.
- A repository-scoped release credential is required so Release Please pull
  requests trigger validation and the release identity can create protected
  tags.
- Conventional Commit accuracy directly affects version calculation and
  changelog quality.
- The repository does not provide standalone downloads for users without a Go
  toolchain.

## Alternatives considered

- Accept a version through `workflow_dispatch`: rejected because maintainers do
  not want to calculate routine release versions manually.
- Publish cross-platform binaries with GoReleaser: rejected because the command
  is intended to run as a module-scoped Go tool.
- Release on every qualifying commit without a release pull request: rejected
  because the pull request provides a reviewable publication checkpoint.

## Follow-up

Configure the repository settings in [`docs/github-settings.md`](../github-settings.md),
merge the first Release Please pull request, and verify `go get -tool` plus
`go tool embedhash` from a separate module.

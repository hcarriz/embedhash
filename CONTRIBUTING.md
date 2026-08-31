# Contributing to embedhash

Thank you for helping improve `embedhash`. Keep changes focused and preserve
the library's compatibility, deterministic output, and security guidance.

## Before starting

Read [AGENTS.md](AGENTS.md) for repository-specific engineering rules. Discuss
changes to exported APIs or observable behavior, dependencies, the minimum Go
version, hash defaults, security policy, CI, and releases before implementing
them.

For non-documentation work, resolve material questions about scope, behavior,
compatibility, security, and acceptance criteria before editing. Record the
answers in the pull request. Do not report a vulnerability publicly; follow
[SECURITY.md](SECURITY.md).

## Development environment

Use Go 1.24 or newer. A currently supported, fully patched Go release is
recommended, especially for security-related work.

Run the repository checks from the module root:

```sh
go fmt ./...
go vet ./...
go test -count=1 ./...
go test -race -count=1 ./...
go test -shuffle=on -count=1 ./...
go mod tidy
git diff --check
```

After `go mod tidy`, verify that any `go.mod` or `go.sum` change is intentional.
The repository currently has no fuzz targets or benchmarks. Add focused ones
when a change introduces attacker-influenced parsing or performance-sensitive
behavior.

## Design, documentation, and tests

- Prefer the standard library and justify every new dependency.
- Keep the public API small and the package layout flat unless a new package has
  a distinct, independently useful responsibility.
- Add behavioral tests for observable changes and update user documentation in
  the same change.
- Cover relevant errors, boundary cases, nil values, untrusted inputs, and
  concurrency.
- Keep tests deterministic; use explicit synchronization instead of timing
  sleeps.
- Preserve deterministic generated output and the path, filename, and hash
  behavior documented in [README.md](README.md).
- Add useful Go doc comments to every exported identifier.

## Decision records

Durable decisions belong in numbered architecture decision records under
[`docs/decisions/`](docs/decisions/). Add or supersede a record when a choice
affects exported API or behavior, compatibility, the minimum Go version,
dependencies, security boundaries, concurrency guarantees, generated format,
supported environments, CI, or releases.

Accepted records are historical. When a decision changes, write a later record
and link the two instead of silently rewriting the old decision.

## Commits and pull requests

Use [Conventional Commits 1.0.0](https://www.conventionalcommits.org/en/v1.0.0/)
for authored commits:

```text
<type>[optional scope][!]: <description>
```

Use `feat` for user-visible additions, `fix` for bug fixes, and `!` with a
`BREAKING CHANGE:` footer for incompatible changes. Common non-release types
include `docs`, `test`, `refactor`, `perf`, `build`, `ci`, and `chore`.

A pull request should explain:

- the problem and intended outcome;
- observable behavior, compatibility, and security effects;
- material questions and accepted answers;
- tests and analysis performed;
- user-facing documentation changed; and
- maintainer decisions still required and any ADR added or superseded.

## Sensitive changes

Get explicit maintainer approval before changing an exported API, the minimum
Go version, dependencies, hash or output defaults, security guarantees,
licensing, ownership, CI, or release policy. Security fixes and reports must use
the private process in [SECURITY.md](SECURITY.md).

## Releases

No automated release workflow is currently committed. Adding release
automation or publishing a release requires maintainer approval. Releases use
semantic, `v`-prefixed Go module tags and accompanying GitHub release notes.
Published versions are immutable: fix a bad release by rolling forward, and use
a later `retract` directive if a module version must no longer be selected.

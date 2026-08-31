# Repository guide

## Purpose and scope

`embedhash` is a Go library and command that discovers files selected by
`//go:embed` and generates a deterministic Go map of their content hashes.

These instructions apply to the entire repository. Keep the library focused on
embedded-file hash discovery and source generation. This repository does not
currently contain a `.go-library-standard.yml` adoption marker, so do not claim
formal conformance to the Go Library Standard without a maintainer-approved
audit and adoption change.

Before changing CI, security policy, release automation, dependency tooling, or
the minimum Go version, consult the maintained
[Go Library Standard](https://github.com/hcarriz/library-agents) and its GitHub
platform guidance.

## Required questions before changes

Before making any non-documentation change, inspect the relevant code and
documentation, ask targeted questions that materially affect implementation,
and wait for the answers. This includes library code, tests, executable
examples, generated-code inputs, configuration, dependencies, build files, CI,
and release automation, even when a change appears mechanical.

Cover only unresolved choices relevant to the task: scope and non-goals, public
API and observable behavior, errors, input trust and resource limits,
concurrency, compatibility, dependencies and Go support, security consequences,
and acceptance criteria. Include a recommendation and tradeoff when the answer
is not obvious. Reading, diagnosis, planning, and prose-only documentation
edits do not require this checkpoint.

Summarize accepted answers before implementation. Do not repeat settled
questions unless new evidence creates a conflict.

## Layout

- `embedhash.go` contains the public library API, package discovery, hashing,
  and Go source rendering.
- `embedhash_test.go` contains black-box behavioral tests for the public API.
- `cmd/embedhash/` contains the command-line entry point and filesystem writes.
- `internal/valid/` is the embedded-file fixture used by tests.
- `docs/decisions/` contains durable architecture decision records.
- `.github/pull_request_template.md` records change decisions and verification.

Keep the root library package flat. Add a package only when it has a distinct,
independently useful responsibility.

## Approval-sensitive changes

Get explicit maintainer approval before changing:

- an exported identifier, signature, or observable error or behavior;
- the minimum Go version or supported release policy;
- runtime, test, or development dependencies;
- default hash algorithm, generated filename, variable name, source format, or
  path normalization;
- command overwrite behavior or filesystem scope;
- security guarantees or the vulnerability-reporting channel;
- CI, releases, licensing, ownership, or repository automation policy; or
- formal adoption of a library standard or an exception to one.

## API and security invariants

- `New` returns one result for each matching package that has embedded files and
  does not itself write files.
- Library defaults are MD5, `HashesForEmbedded`, and `embedhashes.go`. The CLI
  deliberately overrides the filename default with `embedhash.go`.
- Generated keys are relative to the package directory and slash-normalized.
- `Result.Save` sorts entries by key and formats valid Go source with the
  standard formatter. Its generated-code header and deterministic ordering are
  observable behavior.
- `OutputFileName` keeps only `filepath.Base`; output stays in the package
  directory.
- MD5 is not collision resistant and must not be presented as security-grade
  integrity or authentication. An unkeyed SHA-256 digest also does not
  authenticate content.
- Package loading and file hashing have no built-in resource limits. The CLI is
  not a sandbox and overwrites destination files non-atomically.
- A caller-supplied `hash.Hash` must not be used concurrently while `New` is
  using it. Generated maps are mutable and require caller synchronization when
  writes may race with reads.

Keep [README.md](README.md), [SECURITY.md](SECURITY.md), tests, and Go doc
comments consistent with these invariants.

## Go and implementation style

The declared source-compatibility floor is Go 1.24 because that is the explicit
floor in `go.mod`. Treat a change to it as a consumer compatibility decision;
do not infer support for an older version solely from current syntax.

- Use only language and standard-library features available at the declared
  floor.
- Prefer direct control flow, concrete types, useful zero values, and small
  consumer-defined interfaces.
- Avoid package-level mutable state. Every goroutine must have a clear stop
  condition.
- Wrap errors only when adding useful operation context and preserve causes
  with `%w` or `errors.Join`.
- Add useful Go doc comments to every exported identifier.
- Preserve unrelated user changes.

## Decision records

Record task-local choices in the pull request and final handoff. Add or
supersede an ADR under `docs/decisions/` when a choice durably affects exported
API or behavior, compatibility, the minimum Go version, dependencies, security
boundaries, concurrency guarantees, generated formats, supported environments,
CI, or releases.

Accepted ADRs explain why a constraint exists. Keep current instructions here
and current user guidance in topic documentation. Never silently rewrite an
accepted ADR; supersede it with a later record and link both directions.

## Testing and verification

Add or update behavioral tests for every observable change. Prefer table-driven
black-box tests and deterministic synchronization. Add native fuzzing when
attacker-influenced parsing gains a stable invariant worth exercising.

Run:

```sh
go fmt ./...
go vet ./...
go test -count=1 ./...
go test -race -count=1 ./...
go test -shuffle=on -count=1 ./...
go mod tidy
git diff --check
```

The repository currently has no fuzz targets, benchmarks, pinned Staticcheck,
or pinned govulncheck tool. Do not claim those checks ran unless the repository
adds and runs them deliberately. Documentation-only changes still require a
rendered-structure review and `git diff --check`.

## Commits and releases

Every author-controlled commit uses Conventional Commits 1.0.0. The default
branch is `master`; documentation and automation must agree with it.

No release workflow is currently committed. Get maintainer approval before
adding one or publishing a release. Releases use semantic `v`-prefixed Go
module tags and GitHub release records. Published versions are immutable; roll
fixes forward and retract an unsuitable module version in a later release.

## Handoff

For implementation work, report questions resolved, decisions made, checks
run, documentation updated, ADRs added or superseded, and remaining
compatibility or security concerns. For an explicit standard audit, also report
requirement identifiers, approved exceptions, unverified host settings, and
remaining conformance gaps.

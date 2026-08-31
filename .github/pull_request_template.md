## Problem and outcome

<!-- What problem does this change solve, and what result should reviewers expect? -->

## Behavior and compatibility

<!-- Describe observable behavior and any exported API, error, minimum-Go, dependency, generated-format, or performance effect. -->

## Questions and decisions

<!-- Summarize material questions, accepted answers, and implementation choices. Link any ADR added or superseded, or explain why no ADR is required. -->

## Security

<!-- Describe effects on trust boundaries, file writes, untrusted input, resource limits, and hash guarantees. Write "No security effect" with a short reason when applicable. -->

## Verification

- [ ] `go fmt ./...`
- [ ] `go vet ./...`
- [ ] `go test -count=1 ./...`
- [ ] `go test -race -count=1 ./...`
- [ ] `go test -shuffle=on -count=1 ./...`
- [ ] `go mod tidy` leaves intentional module metadata.
- [ ] Pinned Staticcheck and govulncheck pass.
- [ ] `git diff --check`
- [ ] Relevant fuzz targets or benchmarks ran, or are not applicable.

List exact commands and noteworthy results:

## Documentation and approval

- [ ] User-facing documentation matches observable behavior.
- [ ] Exported identifiers have useful Go doc comments.
- [ ] The proposed commit or squash title follows Conventional Commits.
- [ ] Maintainer approval-sensitive decisions are identified below.
- [ ] A durable ADR was added or superseded when required.

Required maintainer decisions or remaining concerns:

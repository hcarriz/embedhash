# Architecture decisions

Use this directory for decisions that establish durable library or repository
constraints. ADRs explain why a choice exists. Keep current operational
instructions in [`AGENTS.md`](../../AGENTS.md) and current user guidance in
[`README.md`](../../README.md).

Write an ADR when a choice materially affects:

- exported API or observable behavior;
- compatibility or migration;
- the minimum Go version or dependency policy;
- security boundaries or resource limits;
- concurrency guarantees;
- generated source format, filenames, or path normalization;
- supported environments; or
- CI, releases, or repository policy.

Do not create an ADR for an implementation detail local to one change. Record
that choice in the pull request instead.

## Format

Copy [`0000-template.md`](0000-template.md) to the next four-digit number and a
short kebab-case title:

```text
0001-stable-generated-format.md
```

Each ADR contains:

```text
# NNNN: Decision title

Status: Proposed | Accepted | Superseded by ADR NNNN
Date: YYYY-MM-DD

## Context
## Decision
## Consequences
## Alternatives considered
## Follow-up
```

An accepted ADR is a historical record. Correct small factual or link errors
without changing its decision. When the choice changes, write a later ADR, mark
the old record superseded, and link both directions.

## Accepted decisions

None yet.

## Superseded decisions

None yet.

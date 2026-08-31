# 0001: Exclude generated output from embedded hashes

Status: Accepted
Date: 2026-08-31

## Context

An embed pattern such as `//go:embed *` can match the configured generated
output file. After the first generation creates that file, the next generation
hashes the previous output and adds its digest to the new output. The generated
file then changes its own input on every run, so an unchanged source tree does
not produce stable output.

A generated file cannot reliably contain its own content digest because adding
the digest changes the content being digested.

## Decision

For every matching package, `New` excludes the configured output path from the
embedded files it hashes. This applies to both the default output filename and
filenames selected with `OutputFileName`.

The exclusion is limited to the exact output path in that package. Other files
matched by the embed directives remain in the generated map.

## Consequences

Generating twice from an unchanged source tree produces byte-for-byte identical
output, including when a wildcard embed pattern matches the generated file.

The generated map cannot contain a digest for its own source file. A caller
that needs that digest must calculate it separately after generation and must
not attempt to store it inside the same generated file.

## Alternatives considered

- Include the output and document repeated changes. Rejected because it defeats
  deterministic generation and creates perpetual diffs.
- Search for a self-consistent digest. Rejected because embedding the digest
  changes the bytes being hashed and does not provide a practical stable
  contract.
- Require every caller to exclude the output in its embed pattern. Rejected
  because wildcard patterns make the failure easy to introduce and the library
  already knows the exact output path.

## Follow-up

None.

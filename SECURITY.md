# Security policy

## Supported versions

Until the first tagged release, security fixes are applied to the `master`
branch. After tagged releases begin, only the latest release on the current
major module line receives security fixes unless a release note states
otherwise.

Use a currently supported, fully patched Go toolchain even though the module's
source-compatibility floor is Go 1.24.

## Reporting a vulnerability

Do not open a public issue, discussion, or pull request for a suspected
vulnerability. Use GitHub's
[private vulnerability report](https://github.com/hcarriz/embedhash/security/advisories/new)
to contact the maintainers without public disclosure.

Include:

- affected versions, commit, Go version, and environment;
- impact and a plausible attack path;
- minimal reproduction details;
- known workarounds; and
- preferred attribution, if any.

Do not send live credentials, unrelated personal data, or destructive proofs
of concept.

## Security expectations

`embedhash` is a local build-time library and command. It loads Go package
metadata, reads every file reported for embedding, and the command writes a
generated file into each matching package. It is not a sandbox: run it only on
source trees you trust, review the package pattern, and use `-dry` before
writing when destinations are uncertain.

The command replaces an existing destination file and does not write
atomically. The project does not impose limits on file count, file size, or
total hashing work. Callers running it over untrusted or unexpectedly large
trees must provide their own isolation and resource limits.

MD5 is the default for compatibility and non-adversarial change detection. MD5
must not be used for collision-resistant integrity checks, authentication,
signatures, or other security decisions. SHA-256 is available through the
command's `-hash sha256` flag and the library's `Hasher` option, but an unkeyed
digest alone does not authenticate content.

## Maintainer process

Maintainers will investigate actionable reports privately, coordinate fixes
and disclosure, add regression coverage, and publish release notes identifying
affected and fixed versions. Published Go module versions will not be replaced;
fixes roll forward, and unsuitable versions are retracted in a later release.

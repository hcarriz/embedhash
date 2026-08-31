# 0003: Use GitHub's workflow token for releases

Status: Accepted
Date: 2026-08-31
Supersedes: the release-credential portion of
[ADR 0002](0002-automate-go-tool-releases.md)

## Context

The initial release workflow required a repository secret named
`RELEASE_TOKEN`. When that secret was not configured, Release Please received
an empty required `token` input and failed before it could propose a release.

GitHub supplies a short-lived `GITHUB_TOKEN` to each workflow run. Using it
avoids maintaining a personal access token or GitHub App credential, but
GitHub may require approval for pull-request workflows caused by that token.

## Decision

Pass `secrets.GITHUB_TOKEN` to Release Please and explicitly target `master`.
Grant write access to Contents, Issues, and Pull requests only to the Release
Please job; keep the workflow's default permission read-only.

Enable the repository setting that allows GitHub Actions to create and approve
pull requests. Maintainers must approve any approval-required `Minimum Go` and
`Release Gate` runs on Release Please pull requests before merging them.

## Consequences

- The workflow no longer depends on a separately provisioned release secret.
- Release operations use a short-lived, repository-scoped workflow token.
- Release Please pull-request checks may require explicit workflow approval.
- Repository and tag rules must allow the release workflow identity to create
  release pull requests, `v*` tags, and GitHub Releases.

## Alternatives considered

- Keep `RELEASE_TOKEN`: rejected because a missing external secret caused the
  workflow failure and the repository does not currently need unattended
  downstream workflow triggering.
- Use a fine-grained personal access token or GitHub App token: deferred. This
  remains an option if release pull-request checks must run without approval.

## Follow-up

Enable the required GitHub Actions repository setting, rerun the failed
workflow, and confirm that Release Please opens or updates its pull request.
Approve that pull request's validation workflows if GitHub requests it.

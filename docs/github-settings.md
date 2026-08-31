# GitHub release settings

The repository workflow validates changes and uses Release Please to propose
and publish releases. The following GitHub settings must be configured after
the workflow reaches `master`; they cannot be enforced by committed files.

## Release credential

Create a repository secret named `RELEASE_TOKEN`. Use a narrowly scoped
fine-grained personal access token or GitHub App token whose identity can:

- read repository metadata and contents;
- create and update release pull requests and their labels;
- create release tags and GitHub Releases; and
- trigger workflows for its release pull requests.

Grant repository permissions for Contents, Issues, and Pull requests only as
needed for those operations. Do not use an organization-wide token when a
repository-scoped credential is available. Rotate the credential according to
the maintainer's credential policy.

Release Please uses this token instead of the default `GITHUB_TOKEN` so the
release pull request triggers the normal pull-request checks.

## Branch rules

Protect `master` against force pushes and deletion. Require pull requests and
the stable checks named `Minimum Go` and `Release Gate` when the collaboration
model permits it. Release Please's pull request must pass the same checks as any
other change.

The repository uses Conventional Commits. Configure squash merging so the pull
request title becomes the resulting commit subject, or otherwise ensure every
authored commit that reaches `master` follows the convention.

## Tag and release rules

Protect tags matching `v*` against updates and deletion. Allow the identity
behind `RELEASE_TOKEN` to create new release tags after the workflow gate
passes. Enable immutable GitHub Releases when available.

Never move, replace, or silently delete a published tag. Roll a correction
forward and retract an unsuitable Go module version in a later release.

## Release lifecycle

Each push to `master` runs the release gate. Release Please reads Conventional
Commits and opens or updates a release pull request containing the calculated
version, changelog, and manifest update. Merging that pull request causes the
next gated run to create the `v`-prefixed tag and GitHub Release.

No binary archives are published. Consumers install a tagged release as a Go
tool:

```sh
go get -tool github.com/hcarriz/embedhash/cmd/embedhash@vX.Y.Z
go tool embedhash -help
```

After the first release, verify the tag points to the tested `master` commit,
the GitHub Release contains the expected notes, and the two commands above
succeed from a separate module.

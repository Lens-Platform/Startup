# Releasing

This document will cover how to properly release Blackspace.

Blackspace adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html) and the release
process is dependent on the version change.

<!-- MarkdownTOC autolink="true" style="ordered" indent="   " -->

- [Releasing](#releasing)
  - [Quick Start](#quick-start)
    - [Patch Releases](#patch-releases)
    - [Major/Minor Releases](#majorminor-releases)
  - [Fixing Up a Release](#fixing-up-a-release)

<!-- /MarkdownTOC -->
## Quick Start

### Patch Releases

1. Create a new branch from the latest `vMAJOR.MINOR.PATCH` tag. Ex: `git checkout -b v1.2.3 v1.2.2`
2. Make the appropriate changes/fixes.
3. Update the [`/RELEASE_NOTES.md`] file.
4. Update the [`/CHANGELOG.md`] and [`/src/services/*/CHANGELOG.md`] headers to reflect the new version `vMAJOR.MINOR.PATCH - YYYY-MM-DD`
5. Commit the changes above with message "Release vMAJOR.MINOR.PATCH"
6. Create a new tag named `vMAJOR.MINOR.PATCH`
7. Push the new tag
8. Delete the temporary branch you created.
9. [All done](https://i.giphy.com/media/3ohzdIvnUKKjiAZTSU/giphy.webp)

### Major/Minor Releases
 
1. Switch to the `master` branch, this should be reflective of the new version's changes.
2. Update the [`/RELEASE_NOTES.md`] file.
3. Update the [`/CHANGELOG.md`] and [`/src/services/*/CHANGELOG.md`] headers to reflect the new version `vMAJOR.MINOR.PATCH - YYYY-MM-DD`
4. Commit the changes above with message `"Release vMAJOR.MINOR.PATCH"`
5. Create a new tag named `vMAJOR.MINOR.PATCH`
6. Push the new tag.
7. Update the [`/CHANGELOG.md`] header to reflect the new upcoming version `vNEW_MAJOR.NEW_MINOR-dev`
8. Commit the changes above with message `"Start vNEW_MAJOR.NEW_MINOR+1"`
9. [All done](https://i.giphy.com/media/3ohzdIvnUKKjiAZTSU/giphy.webp)

## Fixing Up a Release

If you tried to cut a release and the CI failed for some unexpected reason, you can follow these steps to recover:

1. Switch to the `v$VERSION` branch.
1. Delete the `v$VERSION` tag.
1. Cherry pick the commits directly to the branch.
1. Run `make release` while still on that branch.
1. Commit and tag accordingly.
1. Cherry pick that commit back to master so that the release is carried over.

[All done]: https://i.giphy.com/media/3ohzdIvnUKKjiAZTSU/giphy.webp

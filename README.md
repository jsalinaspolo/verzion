[![Actions Status](https://github.com/jsalinaspolo/verzion/workflows/Test/badge.svg)](https://github.com/jsalinaspolo/verzion/actions)
[![codecov](https://codecov.io/gh/jsalinaspolo/verzion/branch/main/graph/badge.svg)](https://codecov.io/gh/jsalinaspolo/verzion)


# `verzion`
Verzion prints out the next version to release.
The main documentation is embedded in the binary. `verzion -h`

## Help
```
$ verzion -h
* Verzion prints the next version to release as a string.

    Your Verzion (current directory):
      - From tags: 4.1.4-deps-git-fix
      - From packed tags: 4.1.0
      - From VERSION file: 4.2.0
      - Verzion will output: 4.2.0

* It's mostly so that you don't have to update your VERSION file each release.

* It looks at the local git tags and VERSION file, compares them,
  and prints out a sensible semantic version (https://semver.org).

* Verzions are printed in the following format:
  [Major].[Minor].[Patch]-[Branch]-[Misc]-[Sha]

* Your VERSION file should be in the format [Major].[Minor]
  Patch numbers in VERSION files are ignored.

* By default, running zersion increments the patch number, e.g. 1.1.1 -> 1.1.2
  To print the current version instead, use 'verzion -c'.

* By default, prints next release version incrementing patch number.

  -c	print without incrementing the patch version. Overrides all other flags.
  -v	print the version of Verzion itself.
```

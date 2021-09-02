[![Actions Status](https://github.com/jsalinaspolo/verzion/workflows/Test/badge.svg)](https://github.com/jsalinaspolo/verzion/actions)
[![codecov](https://codecov.io/gh/jsalinaspolo/verzion/branch/main/graph/badge.svg)](https://codecov.io/gh/jsalinaspolo/verzion)

# `Verzion`

Versioning should be simple, but most of the times it is solved in different ways. Some times it is managed  manually, based on build numbers, based on tags, based on date times and with many other strategies. 
Most of the times, versioning between services and application in an organisation are not consistent because the solution is custom per team, department or repository.

Verzion solves how versioning is managed making it deterministic and consistent following [Semantic Versioning](https://semver.org).

## Install

> Linux
```bash
 curl -sSL https://github.com/jsalinaspolo/verzion/releases/download/v0.6.0/verzion_0.6.0_linux_amd64.tar.gz | tar -xz
 chmod +x ./verzion
./verzion -v
```

> MacOS
```bash
curl -sSL https://github.com/jsalinaspolo/verzion/releases/download/v0.6.0/verzion_0.6.0_darwin_amd64.tar.gz | tar -xz
chmod +x ./verzion
./verzion -v
```

> Windows

Unzip https://github.com/jsalinaspolo/verzion/releases/download/v0.6.0/verzion_0.6.0_windows_amd64.zip

Add the folder to your PATH

Run `verzion -v`  

## How

Verzion CLI prints out the next version to release as a string. It uses a VERSION file in the repository, but avoid having to update the file in each release. 

Verzion uses Semantic Versioning `MAJOR`. `MINOR`. `PATCH` format, increasing by default the `MINOR` version but allowing `PATCH` versioning by parameter.

By default, it looks if the current commit sha exists in a tag. If it exists, will return the tag version. Hence, any builds with the same codebase (same commit hash), will return the same tag version to be deterministic.

If the commit has not been tagged, it will look at the tags and get the highest then will increment the `MINOR` version.

Verzion allows `PATCH` versions by parameter `--patch [version]`. It will look for the tags containing the `version` parameter and get the highest, then increment the `PATCH` version.

Verzion allows metadata by adding a parameter --metadata [metadata] or specific ones 
like the commit sha using the flag `--sha` or branch name using the flag `--branch`

The main documentation is the binary. `verzion -h`

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

* Versions are printed in the following format:
  [Major].[Minor].[Patch]+[Sha]

* Your VERSION file should be in the format Major[.Minor]
  Minor and Patch numbers in VERSION files are ignored.

* By default, running `verzion` increments the minor number, e.g. 1.1.1 -> 1.2.0
  To print the current version instead, use 'verzion -c'.

* By default, prints next release version incrementing patch number.

  -c	        print without incrementing the patch version. Overrides all other flags.
  -p [version]  increment patch version
  -s            append the current commit sha.
  -m            append a miscellaneous string (32 char limit, [0-9A-Za-z-.+] only).
  -v	        print the version of Verzion itself.
```

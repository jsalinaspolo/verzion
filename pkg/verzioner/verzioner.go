package verzioner

import (
	"github.com/jsalinaspolo/verzion/internal/git"
	"github.com/jsalinaspolo/verzion/internal/verzion"
	"sort"
)

type RepositoryPath struct {
	Path string `default:"."`
}

// findVersion encapsulates the logic of verzion.
// This function igores errors. For our use case, we always want to print a version,
func FindVersion(current bool, repoPath RepositoryPath) string {
	fileTagVersion, _ := git.FromFileTags(repoPath.Path)

	// Only check packed refs if there's no file tags.
	tagVersion := fileTagVersion
	if tagVersion.Equal(verzion.Zero) {
		packedVersion, _ := git.FromPackedRefs(repoPath.Path)
		tagVersion = packedVersion
	}

	// Increment the patch version of our last tag (unless `-c` is set).
	if !current {
		tagVersion.Patch++
	}

	// Parse a version from the VERSION file.
	fileVersion, _ := verzion.FromFile(repoPath.Path+"VERSION")

	// Ignore any patch number in the VERSION file.
	fileVersion.Patch = 0

	// Sort the two versions and take the latest.
	versions := verzion.Slice{fileVersion, tagVersion}
	sort.Stable(versions)
	latestVersion := versions[1]

	// If `-c` is on, return the latest tagged version.
	// If there are no tagged versions, return the VERSION file content or 0.0.0.
	if current {
		if tagVersion.Equal(verzion.Zero) {
			return latestVersion.String()
		}

		return tagVersion.String()
	}

	return latestVersion.String()
}

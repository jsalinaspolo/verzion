package verzioner

import (
	"sort"
	"strings"

	"github.com/jsalinaspolo/verzion/internal/git"
	"github.com/jsalinaspolo/verzion/internal/verzion"
)

type RepositoryPath struct {
	Path string
}

// FindVersion encapsulates the logic of verzion.
func FindVersion(current bool, sha bool, branch bool, repoPath RepositoryPath) (string, error) {
	commitHash, _ := git.FindLatestCommit(repoPath.Path)
	tagVersion, err := git.FindTagByHash(repoPath.Path, commitHash)

	// Current commit sha has not been tagged
	if err != nil {
		v, err := git.FromPatchBranch(repoPath.Path)
		if err == nil { // Patch Version
			// Increment patch version (unless `-c` is set).
			if !current {
				v.Patch++
			}
		} else { // Minor Version
			fileTagVersion, _ := git.FromFileTags(repoPath.Path)
			// Only check packed refs if there's no file tags.
			v = fileTagVersion
			if v.Equal(verzion.Zero) {
				packedVersion, _ := git.FromPackedRefs(repoPath.Path)
				v = packedVersion
			}

			// Increment the minor version of our last tag (unless `-c` is set).
			if !current {
				v.Minor++
				v.Patch = 0
			}

			// Parse a version from the VERSION file.
			fileVersion, _ := verzion.FromVersionFile(repoPath.Path + "/VERSION")
			// Sort the two versions and take the latest.
			versions := verzion.Slice{fileVersion, v}
			sort.Stable(versions)
			v = versions[1]
		}

		// TODO should change the logic and name variables
		tagVersion = v
	}

	// TODO this does not make sense
	latestVersion := tagVersion

		// If `-c` is on, return the latest tagged version.
	// If there are no tagged versions, return the VERSION file content or 0.0.0.
	if current {
		return tagVersion.String(), nil
	}

	var metadata []string

	// Add branch flag.
	if branch {
		b, _ := git.Branch(repoPath.Path)
		trimmedBranch := strings.TrimSpace(b)
		if len(trimmedBranch) > 0 && trimmedBranch != "master" {
			metadata = append(metadata, trimmedBranch)
		}
	}

	// Add sha flag.
	if sha {
		sha, err := git.FindShortCommitSha(repoPath.Path)
		if err != nil {
			return "", err
		}
		metadata = append(metadata, sha)
	}

	latestVersion.AddMetadata(metadata)
	latestVersion.Metadata = strings.Join(metadata, ".")
	return latestVersion.String(), nil
}

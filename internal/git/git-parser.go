package git

import (
	"fmt"
	verzion2 "github.com/jsalinaspolo/verzion/internal/verzion"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

// FromFileTags parses local file tags and returns the greatest Verzion.
func FromFileTags(repoPath string) (verzion2.Verzion, error) {
	tagVersions := verzion2.Slice{}

	// Parse tags from local files.
	if fileTags, err := ioutil.ReadDir(filepath.Join(repoPath, ".git", "refs", "tags")); err == nil {
		for _, file := range fileTags {
			v, err := verzion2.FromString(file.Name())
			if err != nil {
				continue
			}

			tagVersions = append(tagVersions, v)
		}
	}

	if len(tagVersions) == 0 {
		return verzion2.Verzion{}, fmt.Errorf("could not parse any tag files out of `%s`", repoPath)
	}

	sort.Stable(tagVersions)
	return tagVersions[len(tagVersions)-1], nil
}

// FromPackedRefs returns the last parsable packed ref.
func FromPackedRefs(repoPath string) (verzion2.Verzion, error) {
	content, err := ioutil.ReadFile(filepath.Join(repoPath, ".git", "packed-refs"))
	if err != nil {
		return verzion2.Verzion{}, err
	}

	refs := strings.Split(string(content), "\n")

	for i := len(refs) - 1; i >= 0; i-- {
		refLine := strings.Fields(refs[i])
		if len(refLine) != 2 {
			continue
		}

		v, err := verzion2.FromString(strings.TrimPrefix(refLine[1], "refs/tags/"))
		if err != nil {
			continue
		}

		return v, nil
	}

	return verzion2.Verzion{}, fmt.Errorf("could not parse any packed refs out of `%s`", repoPath)
}

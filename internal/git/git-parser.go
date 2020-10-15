package git

import (
	"fmt"
	"github.com/jsalinaspolo/verzion/internal/verzion"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strings"
)

type Tag struct {
	Hash    string
	Version verzion.Verzion
}

func (t Tag) Less(cmp Tag) bool {
	return t.Version.Less(cmp.Version)
}

// FromFileTags parses local file tags and returns the greatest Verzion.
func FromFileTags(repoPath string) (verzion.Verzion, error) {
	// Parse tags from local files.
	tagVersions, err := parseRefsTags(repoPath)
	if err != nil {
		return verzion.Verzion{}, err
	}

	sort.Stable(tagVersions)
	return tagVersions[len(tagVersions)-1], nil
}

// FromPackedRefs returns the last parsable packed ref.
func FromPackedRefs(repoPath string) (verzion.Verzion, error) {
	content, err := ioutil.ReadFile(filepath.Join(repoPath, ".git", "packed-refs"))
	if err != nil {
		return verzion.Verzion{}, err
	}

	refs := strings.Split(string(content), "\n")

	for i := len(refs) - 1; i >= 0; i-- {
		refLine := strings.Fields(refs[i])
		if len(refLine) != 2 {
			continue
		}

		v, err := verzion.FromString(strings.TrimPrefix(refLine[1], "refs/tags/"))
		if err != nil {
			continue
		}

		return v, nil
	}

	return verzion.Verzion{}, fmt.Errorf("could not parse any packed refs out of `%s`", repoPath)
}

func parseRefsTags(repoPath string) (verzion.Slice, error) {
	tagVersions := verzion.Slice{}
	if fileTags, err := ioutil.ReadDir(filepath.Join(repoPath, ".git", "refs", "tags")); err == nil {
		for _, file := range fileTags {
			v, err := verzion.FromString(file.Name())
			if err != nil {
				continue
			}

			tagVersions = append(tagVersions, v)
		}
	}

	if len(tagVersions) == 0 {
		return nil, fmt.Errorf("could not parse any tag files out of `%s`", repoPath)
	}

	return tagVersions, nil
}

// FindTagByHsh returns the tag version if the commit hash has a tag
func FindTagByHash(repoPath string, hash string) (verzion.Verzion, error) {
	tagVersions := Slice{}
	gitRefsTagPath := filepath.Join(repoPath, ".git", "refs", "tags")

	if fileTags, err := ioutil.ReadDir(gitRefsTagPath); err == nil {
		for _, file := range fileTags {
			content, err := ioutil.ReadFile(filepath.Join(gitRefsTagPath, file.Name()))
			if err != nil {
				continue
			}

			v, err := verzion.FromString(file.Name())
			if err != nil {
				continue
			}

			tag := Tag{Hash: string(content), Version: v}
			if hash == tag.Hash {
				tagVersions = append(tagVersions, tag)
			}
		}

		if len(tagVersions) == 0 {
			return verzion.Verzion{}, fmt.Errorf("could not find any tag with the hash `%s`", hash)
		}

		sort.Stable(tagVersions)
	}
	return tagVersions[len(tagVersions)-1].Version, nil
}

// FindLatestCommit determine the latest commit sha
func FindLatestCommit(repoPath string) (string, error) {
	c, err := ioutil.ReadFile(filepath.Join(repoPath, ".git", "HEAD"))
	if err != nil {
		return "", err
	}

	content := string(c)

	// If is not detach, extract reference
	if strings.HasPrefix(content, "ref:") {
		ref := strings.TrimPrefix(content, "ref: ")
		c, err := ioutil.ReadFile(filepath.Join(repoPath, ".git", ref))
		if err != nil {
			return "", err
		}
		return string(c), nil
	}

	return content, nil
}

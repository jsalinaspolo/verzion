package verzioner

import (
	"testing"

	"github.com/jsalinaspolo/verzion/internal/git"
	"github.com/jsalinaspolo/verzion/internal/verzion"
	"github.com/stretchr/testify/require"
)

const latestCommit = `7a9d0ca3e6e684ca2f35197511e0290496d94215`

func before(t *testing.T) string {
	tempDir := t.TempDir()
	git.StubHead(t, tempDir, []byte(latestCommit))
	return tempDir
}

func TestFindVersion(t *testing.T) {
	t.Run("should increase minor for zero verzion when empty repository", func(t *testing.T) {
		tempDir := before(t)
		v, err := FindVersion(false, false, false, "", RepositoryPath{Path: tempDir})

		require.NoError(t, err)
		require.Equal(t, "0.1.0", v)
	})

	t.Run("should use VERSION file when empty repository", func(t *testing.T) {
		tempDir := before(t)
		git.StubVersion(t, tempDir, "2.0")

		v, err := FindVersion(false, false, false, "", RepositoryPath{Path: tempDir})

		require.NoError(t, err)
		require.Equal(t, "2.0.0", v)
	})

	t.Run("should get current version based on latest tags ", func(t *testing.T) {
		tempDir := before(t)
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})
		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(true, false, false, "", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "1.3.0", v)
	})

	t.Run("should get current zero verzion when empty repository", func(t *testing.T) {
		tempDir := before(t)
		v, err := FindVersion(true, false, false, "", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "0.0.0", v)
	})

	t.Run("should get next version based on latest  tags", func(t *testing.T) {
		tempDir := before(t)
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})
		tags = append(tags, git.Tag{Hash: "444", Version: verzion.Verzion{Major: 1, Minor: 3, Patch: 1}})
		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(false, false, false, "", RepositoryPath{Path: tempDir})

		require.NoError(t, err)
		require.Equal(t, "1.4.0", v)
	})

	t.Run("should use tag version if commit matches the tag", func(t *testing.T) {
		tempDir := before(t)
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 4}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: latestCommit, Version: verzion.Verzion{Major: 1, Minor: 2, Patch: 3}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})

		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(false, false, false, "", RepositoryPath{Path: tempDir})

		require.NoError(t, err)
		require.Equal(t, "1.2.3", v)
	})

	t.Run("should use tagged version if commit matches the tag and VERSION is bigger", func(t *testing.T) {
		tempDir := before(t)
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 4}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: latestCommit, Version: verzion.Verzion{Major: 1, Minor: 2, Patch: 3}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})

		git.StubRefsTags(t, tempDir, tags)
		git.StubVersion(t, tempDir, "2.0")


		v, err := FindVersion(false, false, false, "", RepositoryPath{Path: tempDir})

		require.NoError(t, err)
		require.Equal(t, "1.2.3", v)
	})

	t.Run("should use VERSION if bigger than tag", func(t *testing.T) {
		tempDir := before(t)
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})
		git.StubRefsTags(t, tempDir, tags)
		git.StubVersion(t, tempDir, "2.0")

		v, err := FindVersion(false, false, false, "", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "2.0.0", v)
	})

	t.Run("should add commit sha to version ", func(t *testing.T) {
		tempDir := before(t)
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})
		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(false, true, false, "", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "1.4.0+7a9d0ca", v)
	})

	t.Run("should add branch name to version ", func(t *testing.T) {
		tempDir := before(t)
		git.StubHead(t, tempDir, []byte("ref: refs/heads/branch-name"))
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})
		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(false, false, true, "", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "1.4.0+branch-name", v)
	})

	t.Run("should increase patch if is a patch branch", func(t *testing.T) {
		tempDir := before(t)
		git.StubHead(t, tempDir, []byte("ref: refs/heads/patch-v1.1"))
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 0}})
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})
		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(false, false, false, "v1.1", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "1.1.1", v)
	})

	t.Run("should increase patch if is a patch branch with previous patch", func(t *testing.T) {
		tempDir := before(t)
		git.StubHead(t, tempDir, []byte("ref: refs/heads/patch-v1.1"))
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 0}})
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 1, Patch: 1}})
		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(false, false, false, "v1.1", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "1.1.2", v)
	})

	t.Run("should increase patch if is a patch branch having highest VERSION file ", func(t *testing.T) {
		tempDir := before(t)
		git.StubHead(t, tempDir, []byte("ref: refs/heads/patch-v1.1"))
		git.StubVersion(t, tempDir, "2.0")
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 0}})
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})
		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(false, false, false, "v1.1", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "1.1.1", v)
	})

	t.Run("should get current patch version", func(t *testing.T) {
		tempDir := before(t)
		git.StubHead(t, tempDir, []byte("ref: refs/heads/patch-v1.1"))
		var tags []git.Tag
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 0}})
		tags = append(tags, git.Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, git.Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, git.Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})
		git.StubRefsTags(t, tempDir, tags)

		v, err := FindVersion(true, false, false, "v1.1", RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "1.1.0", v)
	})
}

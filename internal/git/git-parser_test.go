package git

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jsalinaspolo/verzion/internal/verzion"
	"github.com/stretchr/testify/require"
)

func TestFromFileTags(t *testing.T) {
	t.Run("get latest tag version", func(t *testing.T) {
		tempDir := t.TempDir()
		StubRefsTags(t, tempDir, nil)

		v, err := FromFileTags(tempDir)
		require.NoError(t, err)
		require.Equal(t, verzion.Verzion{Major: 1, Minor: 9, Patch: 0}, v)
	})

	t.Run("fail if there is no tag files", func(t *testing.T) {
		_, err := FromFileTags(t.TempDir())
		require.Error(t, err)
	})
}

func TestFromPackedRefs(t *testing.T) {
	t.Run("parse packed refs ", func(t *testing.T) {
		tempDir := t.TempDir()
		StubPackedRefs(t, tempDir)

		v, err := FromPackedRefs(tempDir)
		require.NoError(t, err)
		require.Equal(t, verzion.Verzion{Major: 1, Minor: 0, Patch: 9}, v)
	})

	t.Run("parse only tags references ", func(t *testing.T) {
		input := []byte(`
# pack-refs with: peeled fully-peeled sorted
2fe51dbfd9bd6fc66d818cd4ee110ffc9e951d42 refs/remotes/origin/v9.9.9
b12b761f6332ea6de8b98f69921061b90a39379d refs/tags/v1.0.4
^8cc58fe566e5d0ef5cadaf456653ab764f1327fb`)
		tempDir := t.TempDir()
		folder := filepath.Join(tempDir, ".git")
		os.MkdirAll(folder, os.ModePerm)
		tmpFile := filepath.Join(folder, "packed-refs")

		err := ioutil.WriteFile(tmpFile, input, 0666)

		v, err := FromPackedRefs(tempDir)
		require.NoError(t, err)
		require.Equal(t, verzion.Verzion{Major: 1, Minor: 0, Patch: 4}, v)
	})

	t.Run("parse only tags references reverse ", func(t *testing.T) {
		input := []byte(`
# pack-refs with: peeled fully-peeled sorted
b12b761f6332ea6de8b98f69921061b90a39379d refs/tags/v1.0.4
^8cc58fe566e5d0ef5cadaf456653ab764f1327fb
2fe51dbfd9bd6fc66d818cd4ee110ffc9e951d42 refs/remotes/origin/v9.9.9`)

		tempDir := t.TempDir()
		folder := filepath.Join(tempDir, ".git")
		os.MkdirAll(folder, os.ModePerm)
		tmpFile := filepath.Join(folder, "packed-refs")

		err := ioutil.WriteFile(tmpFile, input, 0666)

		v, err := FromPackedRefs(tempDir)
		require.NoError(t, err)
		require.Equal(t, verzion.Verzion{Major: 1, Minor: 0, Patch: 4}, v)
	})

	t.Run("fail if packed refs does not exist ", func(t *testing.T) {
		_, err := FromPackedRefs(t.TempDir())
		require.Error(t, err)
	})
}

func TestFindTagByHash(t *testing.T) {
	t.Run("should find version from the tags", func(t *testing.T) {
		hash := "222"

		var tags []Tag
		tags = append(tags, Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, Tag{Hash: "333", Version: verzion.Verzion{Major: 1, Minor: 3}})

		tempDir := t.TempDir()
		StubRefsTags(t, tempDir, tags)
		v, err := FindTagByHash(tempDir, hash)

		require.NoError(t, err)
		require.Equal(t, verzion.Verzion{Major: 1, Minor: 2, Patch: 0}, v)
	})

	t.Run("should find version from highest tag", func(t *testing.T) {
		hash := "222"

		var tags []Tag
		tags = append(tags, Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 4}})
		tags = append(tags, Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})
		tags = append(tags, Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 3}})

		tempDir := t.TempDir()
		StubRefsTags(t, tempDir, tags)
		v, err := FindTagByHash(tempDir, hash)

		require.NoError(t, err)
		require.Equal(t, verzion.Verzion{Major: 1, Minor: 4, Patch: 0}, v)
	})

	t.Run("should fail when not found", func(t *testing.T) {
		hash := "000"

		var tags []Tag
		tags = append(tags, Tag{Hash: "111", Version: verzion.Verzion{Major: 1, Minor: 1}})
		tags = append(tags, Tag{Hash: "222", Version: verzion.Verzion{Major: 1, Minor: 2}})

		tempDir := t.TempDir()
		StubRefsTags(t, tempDir, tags)
		_, err := FindTagByHash(tempDir, hash)

		require.EqualError(t, err, fmt.Sprintf("could not find any tag with the hash `%s`", hash))
	})
}

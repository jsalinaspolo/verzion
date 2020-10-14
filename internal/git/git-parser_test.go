package git

import (
	"fmt"
	"github.com/jsalinaspolo/verzion/internal/verzion"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFromFileTags(t *testing.T) {
	t.Run("get latest tag version", func(t *testing.T) {
		input := []byte(`078174542934ec4907a66cf334ed4c4eee744fa9`)
		tempDir := t.TempDir()

		folder := filepath.Join(tempDir, ".git", "refs", "tags")
		os.MkdirAll(folder, os.ModePerm)

		for i := 1; i < 10; i++ {
			tmpFile := filepath.Join(folder, fmt.Sprintf("v1.%d.0", i))
			err := ioutil.WriteFile(tmpFile, input, 0666)
			require.NoError(t, err)
		}

		v, err := FromFileTags(tempDir)
		require.NoError(t, err)
		assert.Equal(t, v, verzion.Verzion{Major: 1, Minor: 9, Patch: 0})
	})

	t.Run("fail if there is no tag files", func(t *testing.T) {
		_, err := FromFileTags(t.TempDir())
		require.Error(t, err)
	})
}

func TestFromPackedRefs(t *testing.T) {
	t.Run("parse packed refs ", func(t *testing.T) {
		input := []byte(`
# pack-refs with: peeled fully-peeled sorted
b12b761f6332ea6de8b98f69921061b90a39379d refs/tags/v1.0.4
^8cc58fe566e5d0ef5cadaf456653ab764f1327fb
7af46b154eee287dac21f3558f3c4a61e60beebd refs/tags/v1.0.5
^ea9a3fa9b265f68b5c42c1a3a8a94345dfbb594f
dc72cb9bdb8169df4d30b7d34a36c2fcfbd9bbe9 refs/tags/v1.0.6
^cc2c4cecce961f6697f71589266523743041aa10
f57b8f4e053ad955041747c6a43e9d73e36bac4e refs/tags/v1.0.7
^fca1dc37471033693c12cfcf336eb54a653c2e34
1f8ff7c624026d397f39c252aabdfcc7287eb9c4 refs/tags/v1.0.8
^4bc7f007daeb5da4c7fb230fb5f3cebfedc02a95
bf2b15cd5b04a142f51adb73ac4601251375bb88 refs/tags/v1.0.9
^3e8d1a24ed8d7dc827330be704254ccf5ac95a55`)

		tempDir := t.TempDir()
		folder := filepath.Join(tempDir, ".git")
		os.MkdirAll(folder, os.ModePerm)

		tmpFile := filepath.Join(folder, "packed-refs")

		err := ioutil.WriteFile(tmpFile, input, 0666)
		require.NoError(t, err)

		v, err := FromPackedRefs(tempDir)
		require.NoError(t, err)
		assert.Equal(t, v, verzion.Verzion{Major: 1, Minor: 0, Patch: 9})
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
		assert.Equal(t, v, verzion.Verzion{Major: 1, Minor: 0, Patch: 4})
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
		assert.Equal(t, v, verzion.Verzion{Major: 1, Minor: 0, Patch: 4})
	})

	t.Run("fail if packed refs does not exist ", func(t *testing.T) {
		_, err := FromPackedRefs(t.TempDir())
		require.Error(t, err)
	})
}

package verzion

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestFindVersion(t *testing.T) {
	t.Run("should increase path for zero verzion when empty repository", func(t *testing.T) {
		v := FindVersion(false, RepositoryPath{Path: os.TempDir()})
		require.Equal(t, v, "0.0.1")
	})

	t.Run("should get current version based on latest tags ", func(t *testing.T) {
		input := []byte(`078174542934ec4907a66cf334ed4c4eee744fa9`)

		folder := filepath.Join(os.TempDir(), ".git", "refs", "tags")
		os.MkdirAll(folder, os.ModePerm)
		defer os.RemoveAll(folder)

		for i := 1; i < 10; i++ {
			tmpFile := filepath.Join(folder, fmt.Sprintf("v1.%d.0", i))
			err := ioutil.WriteFile(tmpFile, input, 0666)
			require.NoError(t, err)
		}

		v := FindVersion(false, RepositoryPath{Path: os.TempDir()})
		assert.Equal(t, v, "1.9.1")
	})

	t.Run("should get current zero verzion when empty repository", func(t *testing.T) {
		v := FindVersion(true, RepositoryPath{Path: os.TempDir()})
		require.Equal(t, v, "0.0.0")
	})

	t.Run("should get next version based on latest  tags", func(t *testing.T) {
		input := []byte(`078174542934ec4907a66cf334ed4c4eee744fa9`)

		folder := filepath.Join(os.TempDir(), ".git", "refs", "tags")
		os.MkdirAll(folder, os.ModePerm)
		defer os.RemoveAll(folder)

		for i := 1; i < 10; i++ {
			tmpFile := filepath.Join(folder, fmt.Sprintf("v1.%d.2", i))
			err := ioutil.WriteFile(tmpFile, input, 0666)
			require.NoError(t, err)
		}

		v := FindVersion(true, RepositoryPath{Path: os.TempDir()})
		require.Equal(t, v, "1.9.2")
	})

	t.Run("should use VERSION if bigger than tag", func(t *testing.T) {
		input := []byte(`078174542934ec4907a66cf334ed4c4eee744fa9`)

		folder := filepath.Join(os.TempDir(), ".git", "refs", "tags")
		os.MkdirAll(folder, os.ModePerm)
		defer os.RemoveAll(folder)

		versionFile := filepath.Join(os.TempDir(), "VERSION")

		err := ioutil.WriteFile(versionFile, []byte(`2.1`), 0666)
		defer os.Remove(versionFile)
		require.NoError(t, err)

		for i := 1; i < 10; i++ {
			tmpFile := filepath.Join(folder, fmt.Sprintf("v1.%d.2", i))
			err := ioutil.WriteFile(tmpFile, input, 0666)
			require.NoError(t, err)
		}

		v := FindVersion(false, RepositoryPath{Path: os.TempDir()})
		require.Equal(t, "2.1.0", v)
	})
}

package verzioner

import (
	"fmt"
	"github.com/jsalinaspolo/verzion/internal/git"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindVersion(t *testing.T) {
	tempDir := t.TempDir()
	git.StubHead(t, tempDir, []byte(`7a9d0ca3e6e684ca2f35197511e0290496d94215`))

	t.Run("should increase path for zero verzion when empty repository", func(t *testing.T) {
		v, err := FindVersion(false, RepositoryPath{Path: t.TempDir()})

		require.NoError(t, err)
		require.Equal(t, "0.0.1", v)
	})

	t.Run("should get current version based on latest tags ", func(t *testing.T) {
		input := []byte(`078174542934ec4907a66cf334ed4c4eee744fa9`)
		tempDir := t.TempDir()

		folder := filepath.Join(tempDir, ".git", "refs", "tags")
		os.MkdirAll(folder, os.ModePerm)

		for i := 1; i < 10; i++ {
			tmpFile := filepath.Join(folder, fmt.Sprintf("v1.%d.0", i))
			err := ioutil.WriteFile(tmpFile, input, 0666)
			require.NoError(t, err)
		}

		v, err := FindVersion(false, RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		assert.Equal(t, "1.9.1", v)
	})

	t.Run("should get current zero verzion when empty repository", func(t *testing.T) {
		v, err := FindVersion(true, RepositoryPath{Path: t.TempDir()})
		require.NoError(t, err)
		require.Equal(t, "0.0.0", v)
	})

	t.Run("should get next version based on latest  tags", func(t *testing.T) {
		input := []byte(`078174542934ec4907a66cf334ed4c4eee744fa9`)
		tempDir := t.TempDir()

		folder := filepath.Join(tempDir, ".git", "refs", "tags")
		os.MkdirAll(folder, os.ModePerm)

		for i := 1; i < 10; i++ {
			tmpFile := filepath.Join(folder, fmt.Sprintf("v1.%d.2", i))
			err := ioutil.WriteFile(tmpFile, input, 0666)
			require.NoError(t, err)
		}

		v, err := FindVersion(true, RepositoryPath{Path: tempDir})

		require.NoError(t, err)
		require.Equal(t, "1.9.2", v)
	})

	t.Run("should use VERSION if bigger than tag", func(t *testing.T) {
		input := []byte(`078174542934ec4907a66cf334ed4c4eee744fa9`)
		tempDir := t.TempDir()

		folder := filepath.Join(tempDir, ".git", "refs", "tags")
		os.MkdirAll(folder, os.ModePerm)

		versionFile := filepath.Join(tempDir, "VERSION")

		err := ioutil.WriteFile(versionFile, []byte(`2.1`), 0666)
		require.NoError(t, err)

		for i := 1; i < 10; i++ {
			tmpFile := filepath.Join(folder, fmt.Sprintf("v1.%d.2", i))
			err := ioutil.WriteFile(tmpFile, input, 0666)
			require.NoError(t, err)
		}

		v, err := FindVersion(false, RepositoryPath{Path: tempDir})
		require.NoError(t, err)
		require.Equal(t, "2.1.0", v)
	})
}

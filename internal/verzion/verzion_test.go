package verzion

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestFromString(t *testing.T) {

	t.Run("should parse a string ", func(t *testing.T) {
		var tests = []struct {
			input    string
			expected Verzion
		}{
			{"1.0", Verzion{1, 0, 0, ""}},
			{"v1.0", Verzion{1, 0, 0, ""}},
			{"1.1", Verzion{1, 1, 0, ""}},
			{"1.0.0", Verzion{1, 0, 0, ""}},
			{"v1.0.0", Verzion{1, 0, 0, ""}},
			{"v1.2.3", Verzion{1, 2, 3, ""}},
			{"v1.2.3+c4b0a06", Verzion{1, 2, 3, "c4b0a06"}},
			{"v1.0.3-c4b0a06", Verzion{1, 0, 3, "c4b0a06"}},
		}

		for _, test := range tests {
			v, err := FromString(test.input)
			require.NoError(t, err)
			assert.Equal(t, v, test.expected)
		}
	})

	t.Run("should fail if the string is not valid ", func(t *testing.T) {
		_, err := FromString("1")
		require.Error(t, err)
	})
}

func TestVerzion_Less(t *testing.T) {
	t.Run("should compare less than", func(t *testing.T) {
		v0 := Verzion{1, 0, 0, ""}
		v1 := Verzion{1, 1, 0, ""}

		require.True(t, v0.Less(v1))
		require.False(t, v1.Less(v0))

	})
	t.Run("same version should be less than itself", func(t *testing.T) {
		v0 := Verzion{1, 0, 0, ""}
		v1 := Verzion{1, 0, 0, ""}

		require.True(t, v0.Less(v1))
		require.True(t, v1.Less(v0))

	})
}

func TestVerzion_Equal(t *testing.T) {
	t.Run("should be true when equals ignoring suffix", func(t *testing.T) {
		v0 := Verzion{1, 2, 3, "a"}
		v1 := Verzion{1, 2, 3, "b"}

		require.True(t, v0.Equal(v1))
		require.True(t, v1.Equal(v0))
	})

	t.Run("should be false when are not equal", func(t *testing.T) {
		var tests = []struct {
			v0 Verzion
			v1 Verzion
		}{
			{
				v0: Verzion{0, 0, 0, ""},
				v1: Verzion{0, 0, 1, ""},
			},
			{
				v0: Verzion{1, 0, 0, ""},
				v1: Verzion{1, 0, 1, ""},
			},
			{
				v0: Verzion{1, 0, 0, ""},
				v1: Verzion{1, 1, 0, ""},
			},
			{
				v0: Verzion{1, 1, 0, ""},
				v1: Verzion{1, 0, 0, ""},
			},
			{
				v0: Verzion{1, 1, 1, ""},
				v1: Verzion{1, 1, 0, ""},
			},
		}

		for _, test := range tests {
			t.Run(fmt.Sprintf("compare %v %v", test.v0, test.v1), func(t *testing.T) {
				require.False(t, test.v0.Equal(test.v1))
				require.False(t, test.v1.Equal(test.v0))
			})
		}
	})
}

func TestVerzion_String(t *testing.T) {
	v := Verzion{1, 2, 3, "d43d0dc"}
	require.Equal(t, v.String(), "1.2.3-d43d0dc")
}

func TestFromFile(t *testing.T) {
	input := []byte(`1.2.3`)

	tmpFile := filepath.Join(os.TempDir(), "VERSION")

	err := ioutil.WriteFile(tmpFile, input, 0666)
	defer os.Remove(tmpFile)

	require.NoError(t, err)
	v, err := FromFile(tmpFile)
	require.NoError(t, err)
	require.Equal(t, v, Verzion{Major: 1, Minor: 2, Patch: 3})
}

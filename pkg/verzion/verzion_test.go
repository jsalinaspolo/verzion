package verzion

import (
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

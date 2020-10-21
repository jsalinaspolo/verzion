package git

import (
	"testing"

	"github.com/jsalinaspolo/verzion/internal/verzion"
	"github.com/stretchr/testify/require"
)

func TestSlice_Swap(t *testing.T) {
	t.Run("should swap values", func(t *testing.T) {
		tag1 := Tag{Hash: "hash1", Version: verzion.Verzion{Major: 1, Minor: 1}}
		tag2 := Tag{Hash: "hash2", Version: verzion.Verzion{Major: 2, Minor: 2}}
		s := Slice{
			tag1,
			tag2,
		}

		require.Equal(t, s[0], tag1)
		s.Swap(0, 1)
		require.Equal(t, s[0], tag2)
	})

	t.Run("should panic when swap of index out of range", func(t *testing.T) {
		tag1 := Tag{Hash: "hash1", Version: verzion.Verzion{Major: 1, Minor: 1}}
		s := Slice{
			tag1,
		}

		require.Panics(t, func() { s.Swap(0, 1) }, "should throw panic")
	})
}

func TestSlice_Len(t *testing.T) {
	t.Run("should get Len", func(t *testing.T) {
		tag1 := Tag{Hash: "hash1", Version: verzion.Verzion{Major: 1, Minor: 1}}
		tag2 := Tag{Hash: "hash2", Version: verzion.Verzion{Major: 2, Minor: 2}}
		s := Slice{
			tag1,
			tag2,
		}

		require.Equal(t, s.Len(), 2)
	})

	t.Run("should get Len of an empty slice", func(t *testing.T) {
		s := Slice{}
		require.Equal(t, s.Len(), 0)
	})
}

func TestSlice_Less(t *testing.T) {
	t.Run("should compare Tags version using less", func(t *testing.T) {
		tag1 := Tag{Hash: "hash1", Version: verzion.Verzion{Major: 1, Minor: 1}}
		tag2 := Tag{Hash: "hash2", Version: verzion.Verzion{Major: 2, Minor: 2}}
		s := Slice{
			tag1,
			tag2,
		}

		require.True(t, s.Less(0, 1))
		require.False(t, s.Less(1, 0))
	})
}

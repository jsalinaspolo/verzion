package verzion

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSlice_Swap(t *testing.T) {
	t.Run("should swap values", func(t *testing.T) {
		v1 := Verzion{Major: 1, Minor: 1}
		v2 := Verzion{Major: 2, Minor: 2}
		s := Slice{
			v1,
			v2,
		}

		require.Equal(t, s[0], v1)
		s.Swap(0, 1)
		require.Equal(t, s[0], v2)
	})

	t.Run("should panic when swap of index out of range", func(t *testing.T) {
		v1 := Verzion{Major: 1, Minor: 1}
		s := Slice{
			v1,
		}

		require.Panics(t, func() { s.Swap(0, 1) }, "should throw panic")
	})
}

func TestSlice_Len(t *testing.T) {
	t.Run("should get Len", func(t *testing.T) {
		v1 := Verzion{Major: 1, Minor: 1}
		v2 := Verzion{Major: 2, Minor: 2}
		s := Slice{
			v1,
			v2,
		}

		require.Equal(t, s.Len(), 2)
	})

	t.Run("should get Len of an empty slice", func(t *testing.T) {
		s := Slice{}
		require.Equal(t, s.Len(), 0)
	})
}

func TestSlice_Less(t *testing.T) {
	t.Run("should compare Verzions using less", func(t *testing.T) {
		v1 := Verzion{Major: 1, Minor: 1}
		v2 := Verzion{Major: 2, Minor: 2}
		s := Slice{
			v1,
			v2,
		}

		require.True(t, s.Less(0, 1))
		require.False(t, s.Less(1, 0))
	})
}

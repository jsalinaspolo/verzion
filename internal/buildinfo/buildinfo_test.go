package buildinfo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPrint(t *testing.T) {
	// set version
	Version = "0.1"

	require.Equal(t, Version, Print())
}

package buildinfo

import (
	"fmt"
)

// Build information. Populated at build-time.
var (
	Version string
)

// Print returns formatted build info
func Print() string {
	return fmt.Sprintf("%v", Version)
}

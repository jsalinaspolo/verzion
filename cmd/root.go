package cmd

import (
	"fmt"
	"os"
)

// Execute runs the main command logic.
func Execute() {
	if err := versionCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}

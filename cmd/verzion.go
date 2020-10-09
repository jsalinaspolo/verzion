package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "verzion",
	Short: "Prints the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("verzion hi:")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"
	"github.com/jsalinaspolo/verzion/pkg/verzion"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func versionCmd() *cobra.Command {
	const currentFlag = "current"

	cmd := &cobra.Command{
		Use:   "verzion",
		Short: "Prints the CLI version",
		PreRun: func(cmd *cobra.Command, args []string) {
			_ = viper.BindPFlag(currentFlag, cmd.Flags().Lookup(currentFlag))
		},
		Run: func(cmd *cobra.Command, args []string) {
			c := viper.Get(currentFlag).(bool)

			v := verzion.FindVersion(c, verzion.RepositoryPath{})
			fmt.Println(v)
		},
	}

	cmd.Flags().BoolP(currentFlag, "c", false, "current version")

	return cmd
}

func init() {
	rootCmd.AddCommand(versionCmd())
}

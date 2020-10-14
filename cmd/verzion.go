package cmd

import (
	"fmt"
	"github.com/jsalinaspolo/verzion/pkg/verzion"
	"github.com/jsalinaspolo/verzion/zersion"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var help1 = `* Zersion prints the *next* semantic version to be tagged.

    Your Zersion (current directory):`

var help2 = `
* It's mostly so that you don't have to update your VERSION file each release.

* It looks at the local git tags and VERSION file, compares them,
  and prints out a sensible semantic version (https://semver.org).

* By default, running zersion increments the patch number, e.g. 1.1.1 -> 1.1.2
  To print the current version instead, use 'zersion -c'.

* Zersions are printed in the following format:
  [Major].[Minor].[Patch]-[Branch]-[Misc]-[Sha]

* Your VERSION file should be in the format [Major].[Minor]
  Patch numbers in VERSION files are ignored.`

var helpMessage = func() string {
	var help = "%s\n      - From tags: %s\n      - From packed tags: %s\n      - From VERSION file: %s\n      - Zersion will output: %s\n%s\n\n"
	f, _ := zersion.FromFileTags(".")
	p, _ := zersion.FromPackedRefs(".")
	v, _ := zersion.FromFile("VERSION")
	z := verzion.FindVersion(false, verzion.RepositoryPath{})
	return fmt.Sprintf(help, help1, f, p, v, z, help2)
}

func versionCmd() *cobra.Command {
	const currentFlag = "current"

	cmd := &cobra.Command{
		Use:   "verzion [command]",
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
	cmd.SetHelpTemplate("hola help")
	cmd.SetHelpTemplate(helpMessage())
	return cmd
}


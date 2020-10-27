package cmd

import (
	"fmt"
	"github.com/jsalinaspolo/verzion/internal/git"
	"github.com/jsalinaspolo/verzion/internal/verzion"
	"github.com/jsalinaspolo/verzion/pkg/verzioner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var help1 = `* Verzion prints the *next* semantic version to be tagged.

    Your Verzion (current directory):`

var help2 = `
* It's mostly so that you don't have to update your VERSION file each release.

* It looks at the local git tags and VERSION file, compares them,
  and prints out a sensible semantic version (https://semver.org).

* By default, running zersion increments the minor number, e.g. 1.1.1 -> 1.2.0
  To print the current version instead, use 'verzion -c'.

* Versions are printed in the following format:
  [Major].[Minor].[Patch]+[Sha]

* Your VERSION file should be in the format [Major].0
  Minor and Patch numbers in VERSION files are ignored.`

var helpMessage = func() string {
	var help = "%s\n      - From tags: %s\n      - From packed tags: %s\n      - From VERSION file: %s\n      - Zersion will output: %s\n%s\n\n"
	f, _ := git.FromFileTags(".")
	p, _ := git.FromPackedRefs(".")
	v, _ := verzion.FromVersionFile("VERSION")
	z, _ := verzioner.FindVersion(false, false, false, "", verzioner.RepositoryPath{Path: "."})
	return fmt.Sprintf(help, help1, f, p, v, z, help2)
}

func versionCmd() *cobra.Command {
	const currentFlag = "current"
	const shaFlag = "sha"
	const branchFlag = "branch"
	const patchFlag = "patch"

	cmd := &cobra.Command{
		Use:   "verzion [command]",
		Short: "Prints the CLI version",
		PreRun: func(cmd *cobra.Command, args []string) {
			_ = viper.BindPFlag(currentFlag, cmd.Flags().Lookup(currentFlag))
			_ = viper.BindPFlag(shaFlag, cmd.Flags().Lookup(shaFlag))
			_ = viper.BindPFlag(branchFlag, cmd.Flags().Lookup(branchFlag))
			_ = viper.BindPFlag(patchFlag, cmd.Flags().Lookup(patchFlag))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			c := viper.Get(currentFlag).(bool)
			sha := viper.Get(shaFlag).(bool)
			branch := viper.Get(branchFlag).(bool)
			patch := viper.Get(patchFlag).(string)

			v, err := verzioner.FindVersion(c, sha, branch, patch, verzioner.RepositoryPath{Path: "."})
			if err != nil {
				return err
			}
			fmt.Println(v)
			return nil
		},
	}

	cmd.Flags().BoolP(currentFlag, "c", false, "current version")
	cmd.Flags().BoolP(shaFlag, "s", false, "append the current commit sha")
	cmd.Flags().BoolP(branchFlag, "b", false, "append the current branch name")
	cmd.Flags().StringP(patchFlag, "p", "", "patch version")
	cmd.SetHelpTemplate(helpMessage())
	return cmd
}

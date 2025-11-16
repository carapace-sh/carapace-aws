package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeSolutionVersionCmd = &cobra.Command{
	Use:   "describe-solution-version",
	Short: "Describes a specific version of a solution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeSolutionVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeSolutionVersionCmd).Standalone()

		personalize_describeSolutionVersionCmd.Flags().String("solution-version-arn", "", "The Amazon Resource Name (ARN) of the solution version.")
		personalize_describeSolutionVersionCmd.MarkFlagRequired("solution-version-arn")
	})
	personalizeCmd.AddCommand(personalize_describeSolutionVersionCmd)
}

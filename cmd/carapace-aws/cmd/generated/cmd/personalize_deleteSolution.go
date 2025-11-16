package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteSolutionCmd = &cobra.Command{
	Use:   "delete-solution",
	Short: "Deletes all versions of a solution and the `Solution` object itself.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteSolutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_deleteSolutionCmd).Standalone()

		personalize_deleteSolutionCmd.Flags().String("solution-arn", "", "The ARN of the solution to delete.")
		personalize_deleteSolutionCmd.MarkFlagRequired("solution-arn")
	})
	personalizeCmd.AddCommand(personalize_deleteSolutionCmd)
}

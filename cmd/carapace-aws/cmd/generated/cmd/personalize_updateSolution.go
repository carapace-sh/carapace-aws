package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_updateSolutionCmd = &cobra.Command{
	Use:   "update-solution",
	Short: "Updates an Amazon Personalize solution to use a different automatic training configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_updateSolutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_updateSolutionCmd).Standalone()

		personalize_updateSolutionCmd.Flags().String("perform-auto-training", "", "Whether the solution uses automatic training to create new solution versions (trained models).")
		personalize_updateSolutionCmd.Flags().String("solution-arn", "", "The Amazon Resource Name (ARN) of the solution to update.")
		personalize_updateSolutionCmd.Flags().String("solution-update-config", "", "The new configuration details of the solution.")
		personalize_updateSolutionCmd.MarkFlagRequired("solution-arn")
	})
	personalizeCmd.AddCommand(personalize_updateSolutionCmd)
}

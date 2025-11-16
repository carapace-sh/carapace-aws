package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeSolutionCmd = &cobra.Command{
	Use:   "describe-solution",
	Short: "Describes a solution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeSolutionCmd).Standalone()

	personalize_describeSolutionCmd.Flags().String("solution-arn", "", "The Amazon Resource Name (ARN) of the solution to describe.")
	personalize_describeSolutionCmd.MarkFlagRequired("solution-arn")
	personalizeCmd.AddCommand(personalize_describeSolutionCmd)
}

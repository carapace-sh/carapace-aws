package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeProblemCmd = &cobra.Command{
	Use:   "describe-problem",
	Short: "Describes an application problem.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeProblemCmd).Standalone()

	applicationInsights_describeProblemCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the owner of the resource group affected by the problem.")
	applicationInsights_describeProblemCmd.Flags().String("problem-id", "", "The ID of the problem.")
	applicationInsights_describeProblemCmd.MarkFlagRequired("problem-id")
	applicationInsightsCmd.AddCommand(applicationInsights_describeProblemCmd)
}

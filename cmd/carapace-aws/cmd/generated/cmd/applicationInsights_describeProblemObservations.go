package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeProblemObservationsCmd = &cobra.Command{
	Use:   "describe-problem-observations",
	Short: "Describes the anomalies or errors associated with the problem.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeProblemObservationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationInsights_describeProblemObservationsCmd).Standalone()

		applicationInsights_describeProblemObservationsCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
		applicationInsights_describeProblemObservationsCmd.Flags().String("problem-id", "", "The ID of the problem.")
		applicationInsights_describeProblemObservationsCmd.MarkFlagRequired("problem-id")
	})
	applicationInsightsCmd.AddCommand(applicationInsights_describeProblemObservationsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationInsights_describeObservationCmd = &cobra.Command{
	Use:   "describe-observation",
	Short: "Describes an anomaly or error with the application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationInsights_describeObservationCmd).Standalone()

	applicationInsights_describeObservationCmd.Flags().String("account-id", "", "The Amazon Web Services account ID for the resource group owner.")
	applicationInsights_describeObservationCmd.Flags().String("observation-id", "", "The ID of the observation.")
	applicationInsights_describeObservationCmd.MarkFlagRequired("observation-id")
	applicationInsightsCmd.AddCommand(applicationInsights_describeObservationCmd)
}

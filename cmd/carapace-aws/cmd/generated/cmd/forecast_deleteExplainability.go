package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteExplainabilityCmd = &cobra.Command{
	Use:   "delete-explainability",
	Short: "Deletes an Explainability resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteExplainabilityCmd).Standalone()

	forecast_deleteExplainabilityCmd.Flags().String("explainability-arn", "", "The Amazon Resource Name (ARN) of the Explainability resource to delete.")
	forecast_deleteExplainabilityCmd.MarkFlagRequired("explainability-arn")
	forecastCmd.AddCommand(forecast_deleteExplainabilityCmd)
}

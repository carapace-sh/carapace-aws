package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteExplainabilityExportCmd = &cobra.Command{
	Use:   "delete-explainability-export",
	Short: "Deletes an Explainability export.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteExplainabilityExportCmd).Standalone()

	forecast_deleteExplainabilityExportCmd.Flags().String("explainability-export-arn", "", "The Amazon Resource Name (ARN) of the Explainability export to delete.")
	forecast_deleteExplainabilityExportCmd.MarkFlagRequired("explainability-export-arn")
	forecastCmd.AddCommand(forecast_deleteExplainabilityExportCmd)
}

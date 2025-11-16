package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createExplainabilityExportCmd = &cobra.Command{
	Use:   "create-explainability-export",
	Short: "Exports an Explainability resource created by the [CreateExplainability]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createExplainabilityExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_createExplainabilityExportCmd).Standalone()

		forecast_createExplainabilityExportCmd.Flags().String("destination", "", "")
		forecast_createExplainabilityExportCmd.Flags().String("explainability-arn", "", "The Amazon Resource Name (ARN) of the Explainability to export.")
		forecast_createExplainabilityExportCmd.Flags().String("explainability-export-name", "", "A unique name for the Explainability export.")
		forecast_createExplainabilityExportCmd.Flags().String("format", "", "The format of the exported data, CSV or PARQUET.")
		forecast_createExplainabilityExportCmd.Flags().String("tags", "", "Optional metadata to help you categorize and organize your resources.")
		forecast_createExplainabilityExportCmd.MarkFlagRequired("destination")
		forecast_createExplainabilityExportCmd.MarkFlagRequired("explainability-arn")
		forecast_createExplainabilityExportCmd.MarkFlagRequired("explainability-export-name")
	})
	forecastCmd.AddCommand(forecast_createExplainabilityExportCmd)
}

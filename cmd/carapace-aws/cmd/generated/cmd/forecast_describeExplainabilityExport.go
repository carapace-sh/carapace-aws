package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeExplainabilityExportCmd = &cobra.Command{
	Use:   "describe-explainability-export",
	Short: "Describes an Explainability export created using the [CreateExplainabilityExport]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeExplainabilityExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describeExplainabilityExportCmd).Standalone()

		forecast_describeExplainabilityExportCmd.Flags().String("explainability-export-arn", "", "The Amazon Resource Name (ARN) of the Explainability export.")
		forecast_describeExplainabilityExportCmd.MarkFlagRequired("explainability-export-arn")
	})
	forecastCmd.AddCommand(forecast_describeExplainabilityExportCmd)
}

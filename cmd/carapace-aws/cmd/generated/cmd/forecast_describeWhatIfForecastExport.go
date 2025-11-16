package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeWhatIfForecastExportCmd = &cobra.Command{
	Use:   "describe-what-if-forecast-export",
	Short: "Describes the what-if forecast export created using the [CreateWhatIfForecastExport]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeWhatIfForecastExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describeWhatIfForecastExportCmd).Standalone()

		forecast_describeWhatIfForecastExportCmd.Flags().String("what-if-forecast-export-arn", "", "The Amazon Resource Name (ARN) of the what-if forecast export that you are interested in.")
		forecast_describeWhatIfForecastExportCmd.MarkFlagRequired("what-if-forecast-export-arn")
	})
	forecastCmd.AddCommand(forecast_describeWhatIfForecastExportCmd)
}

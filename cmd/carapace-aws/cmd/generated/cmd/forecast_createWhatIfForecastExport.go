package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createWhatIfForecastExportCmd = &cobra.Command{
	Use:   "create-what-if-forecast-export",
	Short: "Exports a forecast created by the [CreateWhatIfForecast]() operation to your Amazon Simple Storage Service (Amazon S3) bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createWhatIfForecastExportCmd).Standalone()

	forecast_createWhatIfForecastExportCmd.Flags().String("destination", "", "The location where you want to save the forecast and an Identity and Access Management (IAM) role that Amazon Forecast can assume to access the location.")
	forecast_createWhatIfForecastExportCmd.Flags().String("format", "", "The format of the exported data, CSV or PARQUET.")
	forecast_createWhatIfForecastExportCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/forecast/latest/dg/tagging-forecast-resources.html) to apply to the what if forecast.")
	forecast_createWhatIfForecastExportCmd.Flags().String("what-if-forecast-arns", "", "The list of what-if forecast Amazon Resource Names (ARNs) to export.")
	forecast_createWhatIfForecastExportCmd.Flags().String("what-if-forecast-export-name", "", "The name of the what-if forecast to export.")
	forecast_createWhatIfForecastExportCmd.MarkFlagRequired("destination")
	forecast_createWhatIfForecastExportCmd.MarkFlagRequired("what-if-forecast-arns")
	forecast_createWhatIfForecastExportCmd.MarkFlagRequired("what-if-forecast-export-name")
	forecastCmd.AddCommand(forecast_createWhatIfForecastExportCmd)
}

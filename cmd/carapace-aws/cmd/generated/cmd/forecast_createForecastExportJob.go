package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createForecastExportJobCmd = &cobra.Command{
	Use:   "create-forecast-export-job",
	Short: "Exports a forecast created by the [CreateForecast]() operation to your Amazon Simple Storage Service (Amazon S3) bucket.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createForecastExportJobCmd).Standalone()

	forecast_createForecastExportJobCmd.Flags().String("destination", "", "The location where you want to save the forecast and an Identity and Access Management (IAM) role that Amazon Forecast can assume to access the location.")
	forecast_createForecastExportJobCmd.Flags().String("forecast-arn", "", "The Amazon Resource Name (ARN) of the forecast that you want to export.")
	forecast_createForecastExportJobCmd.Flags().String("forecast-export-job-name", "", "The name for the forecast export job.")
	forecast_createForecastExportJobCmd.Flags().String("format", "", "The format of the exported data, CSV or PARQUET.")
	forecast_createForecastExportJobCmd.Flags().String("tags", "", "The optional metadata that you apply to the forecast export job to help you categorize and organize them.")
	forecast_createForecastExportJobCmd.MarkFlagRequired("destination")
	forecast_createForecastExportJobCmd.MarkFlagRequired("forecast-arn")
	forecast_createForecastExportJobCmd.MarkFlagRequired("forecast-export-job-name")
	forecastCmd.AddCommand(forecast_createForecastExportJobCmd)
}

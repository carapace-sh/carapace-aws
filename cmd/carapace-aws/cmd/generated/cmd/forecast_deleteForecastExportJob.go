package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteForecastExportJobCmd = &cobra.Command{
	Use:   "delete-forecast-export-job",
	Short: "Deletes a forecast export job created using the [CreateForecastExportJob]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteForecastExportJobCmd).Standalone()

	forecast_deleteForecastExportJobCmd.Flags().String("forecast-export-job-arn", "", "The Amazon Resource Name (ARN) of the forecast export job to delete.")
	forecast_deleteForecastExportJobCmd.MarkFlagRequired("forecast-export-job-arn")
	forecastCmd.AddCommand(forecast_deleteForecastExportJobCmd)
}

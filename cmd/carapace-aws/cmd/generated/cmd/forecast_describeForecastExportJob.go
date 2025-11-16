package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeForecastExportJobCmd = &cobra.Command{
	Use:   "describe-forecast-export-job",
	Short: "Describes a forecast export job created using the [CreateForecastExportJob]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeForecastExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describeForecastExportJobCmd).Standalone()

		forecast_describeForecastExportJobCmd.Flags().String("forecast-export-job-arn", "", "The Amazon Resource Name (ARN) of the forecast export job.")
		forecast_describeForecastExportJobCmd.MarkFlagRequired("forecast-export-job-arn")
	})
	forecastCmd.AddCommand(forecast_describeForecastExportJobCmd)
}

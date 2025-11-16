package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteForecastCmd = &cobra.Command{
	Use:   "delete-forecast",
	Short: "Deletes a forecast created using the [CreateForecast]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteForecastCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_deleteForecastCmd).Standalone()

		forecast_deleteForecastCmd.Flags().String("forecast-arn", "", "The Amazon Resource Name (ARN) of the forecast to delete.")
		forecast_deleteForecastCmd.MarkFlagRequired("forecast-arn")
	})
	forecastCmd.AddCommand(forecast_deleteForecastCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createForecastCmd = &cobra.Command{
	Use:   "create-forecast",
	Short: "Creates a forecast for each item in the `TARGET_TIME_SERIES` dataset that was used to train the predictor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createForecastCmd).Standalone()

	forecast_createForecastCmd.Flags().String("forecast-name", "", "A name for the forecast.")
	forecast_createForecastCmd.Flags().String("forecast-types", "", "The quantiles at which probabilistic forecasts are generated.")
	forecast_createForecastCmd.Flags().String("predictor-arn", "", "The Amazon Resource Name (ARN) of the predictor to use to generate the forecast.")
	forecast_createForecastCmd.Flags().String("tags", "", "The optional metadata that you apply to the forecast to help you categorize and organize them.")
	forecast_createForecastCmd.Flags().String("time-series-selector", "", "Defines the set of time series that are used to create the forecasts in a `TimeSeriesIdentifiers` object.")
	forecast_createForecastCmd.MarkFlagRequired("forecast-name")
	forecast_createForecastCmd.MarkFlagRequired("predictor-arn")
	forecastCmd.AddCommand(forecast_createForecastCmd)
}

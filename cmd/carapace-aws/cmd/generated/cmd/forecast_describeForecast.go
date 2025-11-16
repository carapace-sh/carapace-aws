package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeForecastCmd = &cobra.Command{
	Use:   "describe-forecast",
	Short: "Describes a forecast created using the [CreateForecast]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeForecastCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describeForecastCmd).Standalone()

		forecast_describeForecastCmd.Flags().String("forecast-arn", "", "The Amazon Resource Name (ARN) of the forecast.")
		forecast_describeForecastCmd.MarkFlagRequired("forecast-arn")
	})
	forecastCmd.AddCommand(forecast_describeForecastCmd)
}

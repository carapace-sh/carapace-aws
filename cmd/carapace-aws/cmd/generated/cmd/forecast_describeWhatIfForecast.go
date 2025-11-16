package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeWhatIfForecastCmd = &cobra.Command{
	Use:   "describe-what-if-forecast",
	Short: "Describes the what-if forecast created using the [CreateWhatIfForecast]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeWhatIfForecastCmd).Standalone()

	forecast_describeWhatIfForecastCmd.Flags().String("what-if-forecast-arn", "", "The Amazon Resource Name (ARN) of the what-if forecast that you are interested in.")
	forecast_describeWhatIfForecastCmd.MarkFlagRequired("what-if-forecast-arn")
	forecastCmd.AddCommand(forecast_describeWhatIfForecastCmd)
}

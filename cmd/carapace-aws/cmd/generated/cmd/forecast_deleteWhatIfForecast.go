package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteWhatIfForecastCmd = &cobra.Command{
	Use:   "delete-what-if-forecast",
	Short: "Deletes a what-if forecast created using the [CreateWhatIfForecast]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteWhatIfForecastCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_deleteWhatIfForecastCmd).Standalone()

		forecast_deleteWhatIfForecastCmd.Flags().String("what-if-forecast-arn", "", "The Amazon Resource Name (ARN) of the what-if forecast that you want to delete.")
		forecast_deleteWhatIfForecastCmd.MarkFlagRequired("what-if-forecast-arn")
	})
	forecastCmd.AddCommand(forecast_deleteWhatIfForecastCmd)
}

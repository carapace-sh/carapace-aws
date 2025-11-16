package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listForecastsCmd = &cobra.Command{
	Use:   "list-forecasts",
	Short: "Returns a list of forecasts created using the [CreateForecast]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listForecastsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_listForecastsCmd).Standalone()

		forecast_listForecastsCmd.Flags().String("filters", "", "An array of filters.")
		forecast_listForecastsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
		forecast_listForecastsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	})
	forecastCmd.AddCommand(forecast_listForecastsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listWhatIfForecastsCmd = &cobra.Command{
	Use:   "list-what-if-forecasts",
	Short: "Returns a list of what-if forecasts created using the [CreateWhatIfForecast]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listWhatIfForecastsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_listWhatIfForecastsCmd).Standalone()

		forecast_listWhatIfForecastsCmd.Flags().String("filters", "", "An array of filters.")
		forecast_listWhatIfForecastsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
		forecast_listWhatIfForecastsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	})
	forecastCmd.AddCommand(forecast_listWhatIfForecastsCmd)
}

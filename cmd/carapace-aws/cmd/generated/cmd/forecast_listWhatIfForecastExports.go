package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listWhatIfForecastExportsCmd = &cobra.Command{
	Use:   "list-what-if-forecast-exports",
	Short: "Returns a list of what-if forecast exports created using the [CreateWhatIfForecastExport]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listWhatIfForecastExportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_listWhatIfForecastExportsCmd).Standalone()

		forecast_listWhatIfForecastExportsCmd.Flags().String("filters", "", "An array of filters.")
		forecast_listWhatIfForecastExportsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
		forecast_listWhatIfForecastExportsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	})
	forecastCmd.AddCommand(forecast_listWhatIfForecastExportsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listForecastExportJobsCmd = &cobra.Command{
	Use:   "list-forecast-export-jobs",
	Short: "Returns a list of forecast export jobs created using the [CreateForecastExportJob]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listForecastExportJobsCmd).Standalone()

	forecast_listForecastExportJobsCmd.Flags().String("filters", "", "An array of filters.")
	forecast_listForecastExportJobsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
	forecast_listForecastExportJobsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	forecastCmd.AddCommand(forecast_listForecastExportJobsCmd)
}

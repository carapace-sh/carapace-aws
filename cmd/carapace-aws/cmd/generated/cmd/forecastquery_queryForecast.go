package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecastquery_queryForecastCmd = &cobra.Command{
	Use:   "query-forecast",
	Short: "Retrieves a forecast for a single item, filtered by the supplied criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecastquery_queryForecastCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecastquery_queryForecastCmd).Standalone()

		forecastquery_queryForecastCmd.Flags().String("end-date", "", "The end date for the forecast.")
		forecastquery_queryForecastCmd.Flags().String("filters", "", "The filtering criteria to apply when retrieving the forecast.")
		forecastquery_queryForecastCmd.Flags().String("forecast-arn", "", "The Amazon Resource Name (ARN) of the forecast to query.")
		forecastquery_queryForecastCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
		forecastquery_queryForecastCmd.Flags().String("start-date", "", "The start date for the forecast.")
		forecastquery_queryForecastCmd.MarkFlagRequired("filters")
		forecastquery_queryForecastCmd.MarkFlagRequired("forecast-arn")
	})
	forecastqueryCmd.AddCommand(forecastquery_queryForecastCmd)
}

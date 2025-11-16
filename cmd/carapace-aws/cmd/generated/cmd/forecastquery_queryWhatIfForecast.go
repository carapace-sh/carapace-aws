package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecastquery_queryWhatIfForecastCmd = &cobra.Command{
	Use:   "query-what-if-forecast",
	Short: "Retrieves a what-if forecast.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecastquery_queryWhatIfForecastCmd).Standalone()

	forecastquery_queryWhatIfForecastCmd.Flags().String("end-date", "", "The end date for the what-if forecast.")
	forecastquery_queryWhatIfForecastCmd.Flags().String("filters", "", "The filtering criteria to apply when retrieving the forecast.")
	forecastquery_queryWhatIfForecastCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a `NextToken`.")
	forecastquery_queryWhatIfForecastCmd.Flags().String("start-date", "", "The start date for the what-if forecast.")
	forecastquery_queryWhatIfForecastCmd.Flags().String("what-if-forecast-arn", "", "The Amazon Resource Name (ARN) of the what-if forecast to query.")
	forecastquery_queryWhatIfForecastCmd.MarkFlagRequired("filters")
	forecastquery_queryWhatIfForecastCmd.MarkFlagRequired("what-if-forecast-arn")
	forecastqueryCmd.AddCommand(forecastquery_queryWhatIfForecastCmd)
}

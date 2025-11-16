package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_listPredictorBacktestExportJobsCmd = &cobra.Command{
	Use:   "list-predictor-backtest-export-jobs",
	Short: "Returns a list of predictor backtest export jobs created using the [CreatePredictorBacktestExportJob]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_listPredictorBacktestExportJobsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_listPredictorBacktestExportJobsCmd).Standalone()

		forecast_listPredictorBacktestExportJobsCmd.Flags().String("filters", "", "An array of filters.")
		forecast_listPredictorBacktestExportJobsCmd.Flags().String("max-results", "", "The number of items to return in the response.")
		forecast_listPredictorBacktestExportJobsCmd.Flags().String("next-token", "", "If the result of the previous request was truncated, the response includes a NextToken.")
	})
	forecastCmd.AddCommand(forecast_listPredictorBacktestExportJobsCmd)
}

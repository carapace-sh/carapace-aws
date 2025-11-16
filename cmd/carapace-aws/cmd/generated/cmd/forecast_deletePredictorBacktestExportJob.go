package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deletePredictorBacktestExportJobCmd = &cobra.Command{
	Use:   "delete-predictor-backtest-export-job",
	Short: "Deletes a predictor backtest export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deletePredictorBacktestExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_deletePredictorBacktestExportJobCmd).Standalone()

		forecast_deletePredictorBacktestExportJobCmd.Flags().String("predictor-backtest-export-job-arn", "", "The Amazon Resource Name (ARN) of the predictor backtest export job to delete.")
		forecast_deletePredictorBacktestExportJobCmd.MarkFlagRequired("predictor-backtest-export-job-arn")
	})
	forecastCmd.AddCommand(forecast_deletePredictorBacktestExportJobCmd)
}

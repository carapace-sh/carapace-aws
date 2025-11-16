package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describePredictorBacktestExportJobCmd = &cobra.Command{
	Use:   "describe-predictor-backtest-export-job",
	Short: "Describes a predictor backtest export job created using the [CreatePredictorBacktestExportJob]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describePredictorBacktestExportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describePredictorBacktestExportJobCmd).Standalone()

		forecast_describePredictorBacktestExportJobCmd.Flags().String("predictor-backtest-export-job-arn", "", "The Amazon Resource Name (ARN) of the predictor backtest export job.")
		forecast_describePredictorBacktestExportJobCmd.MarkFlagRequired("predictor-backtest-export-job-arn")
	})
	forecastCmd.AddCommand(forecast_describePredictorBacktestExportJobCmd)
}

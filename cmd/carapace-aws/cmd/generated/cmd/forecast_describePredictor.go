package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describePredictorCmd = &cobra.Command{
	Use:   "describe-predictor",
	Short: "This operation is only valid for legacy predictors created with CreatePredictor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describePredictorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describePredictorCmd).Standalone()

		forecast_describePredictorCmd.Flags().String("predictor-arn", "", "The Amazon Resource Name (ARN) of the predictor that you want information about.")
		forecast_describePredictorCmd.MarkFlagRequired("predictor-arn")
	})
	forecastCmd.AddCommand(forecast_describePredictorCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeAutoPredictorCmd = &cobra.Command{
	Use:   "describe-auto-predictor",
	Short: "Describes a predictor created using the CreateAutoPredictor operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeAutoPredictorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_describeAutoPredictorCmd).Standalone()

		forecast_describeAutoPredictorCmd.Flags().String("predictor-arn", "", "The Amazon Resource Name (ARN) of the predictor.")
		forecast_describeAutoPredictorCmd.MarkFlagRequired("predictor-arn")
	})
	forecastCmd.AddCommand(forecast_describeAutoPredictorCmd)
}

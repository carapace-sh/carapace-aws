package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deletePredictorCmd = &cobra.Command{
	Use:   "delete-predictor",
	Short: "Deletes a predictor created using the [DescribePredictor]() or [CreatePredictor]() operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deletePredictorCmd).Standalone()

	forecast_deletePredictorCmd.Flags().String("predictor-arn", "", "The Amazon Resource Name (ARN) of the predictor to delete.")
	forecast_deletePredictorCmd.MarkFlagRequired("predictor-arn")
	forecastCmd.AddCommand(forecast_deletePredictorCmd)
}

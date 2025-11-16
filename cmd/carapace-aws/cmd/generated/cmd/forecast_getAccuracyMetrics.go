package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_getAccuracyMetricsCmd = &cobra.Command{
	Use:   "get-accuracy-metrics",
	Short: "Provides metrics on the accuracy of the models that were trained by the [CreatePredictor]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_getAccuracyMetricsCmd).Standalone()

	forecast_getAccuracyMetricsCmd.Flags().String("predictor-arn", "", "The Amazon Resource Name (ARN) of the predictor to get metrics for.")
	forecast_getAccuracyMetricsCmd.MarkFlagRequired("predictor-arn")
	forecastCmd.AddCommand(forecast_getAccuracyMetricsCmd)
}

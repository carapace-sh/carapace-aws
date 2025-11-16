package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createAutoPredictorCmd = &cobra.Command{
	Use:   "create-auto-predictor",
	Short: "Creates an Amazon Forecast predictor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createAutoPredictorCmd).Standalone()

	forecast_createAutoPredictorCmd.Flags().String("data-config", "", "The data configuration for your dataset group and any additional datasets.")
	forecast_createAutoPredictorCmd.Flags().String("encryption-config", "", "")
	forecast_createAutoPredictorCmd.Flags().Bool("explain-predictor", false, "Create an Explainability resource for the predictor.")
	forecast_createAutoPredictorCmd.Flags().String("forecast-dimensions", "", "An array of dimension (field) names that specify how to group the generated forecast.")
	forecast_createAutoPredictorCmd.Flags().String("forecast-frequency", "", "The frequency of predictions in a forecast.")
	forecast_createAutoPredictorCmd.Flags().String("forecast-horizon", "", "The number of time-steps that the model predicts.")
	forecast_createAutoPredictorCmd.Flags().String("forecast-types", "", "The forecast types used to train a predictor.")
	forecast_createAutoPredictorCmd.Flags().String("monitor-config", "", "The configuration details for predictor monitoring.")
	forecast_createAutoPredictorCmd.Flags().Bool("no-explain-predictor", false, "Create an Explainability resource for the predictor.")
	forecast_createAutoPredictorCmd.Flags().String("optimization-metric", "", "The accuracy metric used to optimize the predictor.")
	forecast_createAutoPredictorCmd.Flags().String("predictor-name", "", "A unique name for the predictor")
	forecast_createAutoPredictorCmd.Flags().String("reference-predictor-arn", "", "The ARN of the predictor to retrain or upgrade.")
	forecast_createAutoPredictorCmd.Flags().String("tags", "", "Optional metadata to help you categorize and organize your predictors.")
	forecast_createAutoPredictorCmd.Flags().String("time-alignment-boundary", "", "The time boundary Forecast uses to align and aggregate any data that doesn't align with your forecast frequency.")
	forecast_createAutoPredictorCmd.Flag("no-explain-predictor").Hidden = true
	forecast_createAutoPredictorCmd.MarkFlagRequired("predictor-name")
	forecastCmd.AddCommand(forecast_createAutoPredictorCmd)
}

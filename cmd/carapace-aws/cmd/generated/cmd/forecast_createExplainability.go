package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_createExplainabilityCmd = &cobra.Command{
	Use:   "create-explainability",
	Short: "Explainability is only available for Forecasts and Predictors generated from an AutoPredictor ([CreateAutoPredictor]())\n\nCreates an Amazon Forecast Explainability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_createExplainabilityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(forecast_createExplainabilityCmd).Standalone()

		forecast_createExplainabilityCmd.Flags().String("data-source", "", "")
		forecast_createExplainabilityCmd.Flags().Bool("enable-visualization", false, "Create an Explainability visualization that is viewable within the Amazon Web Services console.")
		forecast_createExplainabilityCmd.Flags().String("end-date-time", "", "If `TimePointGranularity` is set to `SPECIFIC`, define the last time point for the Explainability.")
		forecast_createExplainabilityCmd.Flags().String("explainability-config", "", "The configuration settings that define the granularity of time series and time points for the Explainability.")
		forecast_createExplainabilityCmd.Flags().String("explainability-name", "", "A unique name for the Explainability.")
		forecast_createExplainabilityCmd.Flags().Bool("no-enable-visualization", false, "Create an Explainability visualization that is viewable within the Amazon Web Services console.")
		forecast_createExplainabilityCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Predictor or Forecast used to create the Explainability.")
		forecast_createExplainabilityCmd.Flags().String("schema", "", "")
		forecast_createExplainabilityCmd.Flags().String("start-date-time", "", "If `TimePointGranularity` is set to `SPECIFIC`, define the first point for the Explainability.")
		forecast_createExplainabilityCmd.Flags().String("tags", "", "Optional metadata to help you categorize and organize your resources.")
		forecast_createExplainabilityCmd.MarkFlagRequired("explainability-config")
		forecast_createExplainabilityCmd.MarkFlagRequired("explainability-name")
		forecast_createExplainabilityCmd.Flag("no-enable-visualization").Hidden = true
		forecast_createExplainabilityCmd.MarkFlagRequired("resource-arn")
	})
	forecastCmd.AddCommand(forecast_createExplainabilityCmd)
}

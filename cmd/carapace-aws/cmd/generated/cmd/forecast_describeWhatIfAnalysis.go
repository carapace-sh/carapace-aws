package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_describeWhatIfAnalysisCmd = &cobra.Command{
	Use:   "describe-what-if-analysis",
	Short: "Describes the what-if analysis created using the [CreateWhatIfAnalysis]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_describeWhatIfAnalysisCmd).Standalone()

	forecast_describeWhatIfAnalysisCmd.Flags().String("what-if-analysis-arn", "", "The Amazon Resource Name (ARN) of the what-if analysis that you are interested in.")
	forecast_describeWhatIfAnalysisCmd.MarkFlagRequired("what-if-analysis-arn")
	forecastCmd.AddCommand(forecast_describeWhatIfAnalysisCmd)
}

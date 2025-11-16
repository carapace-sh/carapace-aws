package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forecast_deleteWhatIfAnalysisCmd = &cobra.Command{
	Use:   "delete-what-if-analysis",
	Short: "Deletes a what-if analysis created using the [CreateWhatIfAnalysis]() operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forecast_deleteWhatIfAnalysisCmd).Standalone()

	forecast_deleteWhatIfAnalysisCmd.Flags().String("what-if-analysis-arn", "", "The Amazon Resource Name (ARN) of the what-if analysis that you want to delete.")
	forecast_deleteWhatIfAnalysisCmd.MarkFlagRequired("what-if-analysis-arn")
	forecastCmd.AddCommand(forecast_deleteWhatIfAnalysisCmd)
}

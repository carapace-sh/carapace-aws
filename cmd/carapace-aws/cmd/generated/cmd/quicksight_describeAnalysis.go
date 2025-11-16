package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAnalysisCmd = &cobra.Command{
	Use:   "describe-analysis",
	Short: "Provides a summary of the metadata for an analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAnalysisCmd).Standalone()

	quicksight_describeAnalysisCmd.Flags().String("analysis-id", "", "The ID of the analysis that you're describing.")
	quicksight_describeAnalysisCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the analysis.")
	quicksight_describeAnalysisCmd.MarkFlagRequired("analysis-id")
	quicksight_describeAnalysisCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeAnalysisCmd)
}

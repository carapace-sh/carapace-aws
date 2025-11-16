package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAnalysisDefinitionCmd = &cobra.Command{
	Use:   "describe-analysis-definition",
	Short: "Provides a detailed description of the definition of an analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAnalysisDefinitionCmd).Standalone()

	quicksight_describeAnalysisDefinitionCmd.Flags().String("analysis-id", "", "The ID of the analysis that you're describing.")
	quicksight_describeAnalysisDefinitionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the analysis.")
	quicksight_describeAnalysisDefinitionCmd.MarkFlagRequired("analysis-id")
	quicksight_describeAnalysisDefinitionCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeAnalysisDefinitionCmd)
}

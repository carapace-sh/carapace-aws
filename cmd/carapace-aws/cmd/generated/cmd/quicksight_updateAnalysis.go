package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateAnalysisCmd = &cobra.Command{
	Use:   "update-analysis",
	Short: "Updates an analysis in Amazon Quick Sight",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateAnalysisCmd).Standalone()

	quicksight_updateAnalysisCmd.Flags().String("analysis-id", "", "The ID for the analysis that you're updating.")
	quicksight_updateAnalysisCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the analysis that you're updating.")
	quicksight_updateAnalysisCmd.Flags().String("definition", "", "The definition of an analysis.")
	quicksight_updateAnalysisCmd.Flags().String("name", "", "A descriptive name for the analysis that you're updating.")
	quicksight_updateAnalysisCmd.Flags().String("parameters", "", "The parameter names and override values that you want to use.")
	quicksight_updateAnalysisCmd.Flags().String("source-entity", "", "A source entity to use for the analysis that you're updating.")
	quicksight_updateAnalysisCmd.Flags().String("theme-arn", "", "The Amazon Resource Name (ARN) for the theme to apply to the analysis that you're creating.")
	quicksight_updateAnalysisCmd.Flags().String("validation-strategy", "", "The option to relax the validation needed to update an analysis with definition objects.")
	quicksight_updateAnalysisCmd.MarkFlagRequired("analysis-id")
	quicksight_updateAnalysisCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateAnalysisCmd.MarkFlagRequired("name")
	quicksightCmd.AddCommand(quicksight_updateAnalysisCmd)
}

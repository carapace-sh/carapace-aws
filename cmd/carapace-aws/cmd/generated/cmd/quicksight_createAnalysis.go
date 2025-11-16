package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createAnalysisCmd = &cobra.Command{
	Use:   "create-analysis",
	Short: "Creates an analysis in Amazon Quick Sight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createAnalysisCmd).Standalone()

	quicksight_createAnalysisCmd.Flags().String("analysis-id", "", "The ID for the analysis that you're creating.")
	quicksight_createAnalysisCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account where you are creating an analysis.")
	quicksight_createAnalysisCmd.Flags().String("definition", "", "The definition of an analysis.")
	quicksight_createAnalysisCmd.Flags().String("folder-arns", "", "When you create the analysis, Amazon Quick Sight adds the analysis to these folders.")
	quicksight_createAnalysisCmd.Flags().String("name", "", "A descriptive name for the analysis that you're creating.")
	quicksight_createAnalysisCmd.Flags().String("parameters", "", "The parameter names and override values that you want to use.")
	quicksight_createAnalysisCmd.Flags().String("permissions", "", "A structure that describes the principals and the resource-level permissions on an analysis.")
	quicksight_createAnalysisCmd.Flags().String("source-entity", "", "A source entity to use for the analysis that you're creating.")
	quicksight_createAnalysisCmd.Flags().String("tags", "", "Contains a map of the key-value pairs for the resource tag or tags assigned to the analysis.")
	quicksight_createAnalysisCmd.Flags().String("theme-arn", "", "The ARN for the theme to apply to the analysis that you're creating.")
	quicksight_createAnalysisCmd.Flags().String("validation-strategy", "", "The option to relax the validation needed to create an analysis with definition objects.")
	quicksight_createAnalysisCmd.MarkFlagRequired("analysis-id")
	quicksight_createAnalysisCmd.MarkFlagRequired("aws-account-id")
	quicksight_createAnalysisCmd.MarkFlagRequired("name")
	quicksightCmd.AddCommand(quicksight_createAnalysisCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeAnalysisPermissionsCmd = &cobra.Command{
	Use:   "describe-analysis-permissions",
	Short: "Provides the read and write permissions for an analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeAnalysisPermissionsCmd).Standalone()

	quicksight_describeAnalysisPermissionsCmd.Flags().String("analysis-id", "", "The ID of the analysis whose permissions you're describing.")
	quicksight_describeAnalysisPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the analysis whose permissions you're describing.")
	quicksight_describeAnalysisPermissionsCmd.MarkFlagRequired("analysis-id")
	quicksight_describeAnalysisPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeAnalysisPermissionsCmd)
}

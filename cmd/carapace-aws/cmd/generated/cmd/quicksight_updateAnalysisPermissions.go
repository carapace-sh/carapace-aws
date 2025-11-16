package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateAnalysisPermissionsCmd = &cobra.Command{
	Use:   "update-analysis-permissions",
	Short: "Updates the read and write permissions for an analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateAnalysisPermissionsCmd).Standalone()

	quicksight_updateAnalysisPermissionsCmd.Flags().String("analysis-id", "", "The ID of the analysis whose permissions you're updating.")
	quicksight_updateAnalysisPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the analysis whose permissions you're updating.")
	quicksight_updateAnalysisPermissionsCmd.Flags().String("grant-permissions", "", "A structure that describes the permissions to add and the principal to add them to.")
	quicksight_updateAnalysisPermissionsCmd.Flags().String("revoke-permissions", "", "A structure that describes the permissions to remove and the principal to remove them from.")
	quicksight_updateAnalysisPermissionsCmd.MarkFlagRequired("analysis-id")
	quicksight_updateAnalysisPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_updateAnalysisPermissionsCmd)
}

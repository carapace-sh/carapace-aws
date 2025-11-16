package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_restoreAnalysisCmd = &cobra.Command{
	Use:   "restore-analysis",
	Short: "Restores an analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_restoreAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_restoreAnalysisCmd).Standalone()

		quicksight_restoreAnalysisCmd.Flags().String("analysis-id", "", "The ID of the analysis that you're restoring.")
		quicksight_restoreAnalysisCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the analysis.")
		quicksight_restoreAnalysisCmd.Flags().Bool("no-restore-to-folders", false, "A boolean value that determines if the analysis will be restored to folders that it previously resided in.")
		quicksight_restoreAnalysisCmd.Flags().Bool("restore-to-folders", false, "A boolean value that determines if the analysis will be restored to folders that it previously resided in.")
		quicksight_restoreAnalysisCmd.MarkFlagRequired("analysis-id")
		quicksight_restoreAnalysisCmd.MarkFlagRequired("aws-account-id")
		quicksight_restoreAnalysisCmd.Flag("no-restore-to-folders").Hidden = true
	})
	quicksightCmd.AddCommand(quicksight_restoreAnalysisCmd)
}

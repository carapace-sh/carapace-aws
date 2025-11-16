package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteAnalysisCmd = &cobra.Command{
	Use:   "delete-analysis",
	Short: "Deletes an analysis from Amazon Quick Sight.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteAnalysisCmd).Standalone()

	quicksight_deleteAnalysisCmd.Flags().String("analysis-id", "", "The ID of the analysis that you're deleting.")
	quicksight_deleteAnalysisCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account where you want to delete an analysis.")
	quicksight_deleteAnalysisCmd.Flags().Bool("force-delete-without-recovery", false, "This option defaults to the value `NoForceDeleteWithoutRecovery`.")
	quicksight_deleteAnalysisCmd.Flags().Bool("no-force-delete-without-recovery", false, "This option defaults to the value `NoForceDeleteWithoutRecovery`.")
	quicksight_deleteAnalysisCmd.Flags().String("recovery-window-in-days", "", "A value that specifies the number of days that Amazon Quick Sight waits before it deletes the analysis.")
	quicksight_deleteAnalysisCmd.MarkFlagRequired("analysis-id")
	quicksight_deleteAnalysisCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteAnalysisCmd.Flag("no-force-delete-without-recovery").Hidden = true
	quicksightCmd.AddCommand(quicksight_deleteAnalysisCmd)
}

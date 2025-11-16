package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteRefreshScheduleCmd = &cobra.Command{
	Use:   "delete-refresh-schedule",
	Short: "Deletes a refresh schedule from a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteRefreshScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteRefreshScheduleCmd).Standalone()

		quicksight_deleteRefreshScheduleCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
		quicksight_deleteRefreshScheduleCmd.Flags().String("data-set-id", "", "The ID of the dataset.")
		quicksight_deleteRefreshScheduleCmd.Flags().String("schedule-id", "", "The ID of the refresh schedule.")
		quicksight_deleteRefreshScheduleCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteRefreshScheduleCmd.MarkFlagRequired("data-set-id")
		quicksight_deleteRefreshScheduleCmd.MarkFlagRequired("schedule-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteRefreshScheduleCmd)
}

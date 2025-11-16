package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateRefreshScheduleCmd = &cobra.Command{
	Use:   "update-refresh-schedule",
	Short: "Updates a refresh schedule for a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateRefreshScheduleCmd).Standalone()

	quicksight_updateRefreshScheduleCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_updateRefreshScheduleCmd.Flags().String("data-set-id", "", "The ID of the dataset.")
	quicksight_updateRefreshScheduleCmd.Flags().String("schedule", "", "The refresh schedule.")
	quicksight_updateRefreshScheduleCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateRefreshScheduleCmd.MarkFlagRequired("data-set-id")
	quicksight_updateRefreshScheduleCmd.MarkFlagRequired("schedule")
	quicksightCmd.AddCommand(quicksight_updateRefreshScheduleCmd)
}

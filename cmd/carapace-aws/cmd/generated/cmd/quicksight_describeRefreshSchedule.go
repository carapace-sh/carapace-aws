package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeRefreshScheduleCmd = &cobra.Command{
	Use:   "describe-refresh-schedule",
	Short: "Provides a summary of a refresh schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeRefreshScheduleCmd).Standalone()

	quicksight_describeRefreshScheduleCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_describeRefreshScheduleCmd.Flags().String("data-set-id", "", "The ID of the dataset.")
	quicksight_describeRefreshScheduleCmd.Flags().String("schedule-id", "", "The ID of the refresh schedule.")
	quicksight_describeRefreshScheduleCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeRefreshScheduleCmd.MarkFlagRequired("data-set-id")
	quicksight_describeRefreshScheduleCmd.MarkFlagRequired("schedule-id")
	quicksightCmd.AddCommand(quicksight_describeRefreshScheduleCmd)
}

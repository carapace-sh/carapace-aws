package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_createRefreshScheduleCmd = &cobra.Command{
	Use:   "create-refresh-schedule",
	Short: "Creates a refresh schedule for a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_createRefreshScheduleCmd).Standalone()

	quicksight_createRefreshScheduleCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_createRefreshScheduleCmd.Flags().String("data-set-id", "", "The ID of the dataset.")
	quicksight_createRefreshScheduleCmd.Flags().String("schedule", "", "The refresh schedule.")
	quicksight_createRefreshScheduleCmd.MarkFlagRequired("aws-account-id")
	quicksight_createRefreshScheduleCmd.MarkFlagRequired("data-set-id")
	quicksight_createRefreshScheduleCmd.MarkFlagRequired("schedule")
	quicksightCmd.AddCommand(quicksight_createRefreshScheduleCmd)
}

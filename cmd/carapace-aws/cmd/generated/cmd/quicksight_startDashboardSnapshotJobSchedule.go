package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_startDashboardSnapshotJobScheduleCmd = &cobra.Command{
	Use:   "start-dashboard-snapshot-job-schedule",
	Short: "Starts an asynchronous job that runs an existing dashboard schedule and sends the dashboard snapshot through email.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_startDashboardSnapshotJobScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_startDashboardSnapshotJobScheduleCmd).Standalone()

		quicksight_startDashboardSnapshotJobScheduleCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that the dashboard snapshot job is executed in.")
		quicksight_startDashboardSnapshotJobScheduleCmd.Flags().String("dashboard-id", "", "The ID of the dashboard that you want to start a snapshot job schedule for.")
		quicksight_startDashboardSnapshotJobScheduleCmd.Flags().String("schedule-id", "", "The ID of the schedule that you want to start a snapshot job schedule for.")
		quicksight_startDashboardSnapshotJobScheduleCmd.MarkFlagRequired("aws-account-id")
		quicksight_startDashboardSnapshotJobScheduleCmd.MarkFlagRequired("dashboard-id")
		quicksight_startDashboardSnapshotJobScheduleCmd.MarkFlagRequired("schedule-id")
	})
	quicksightCmd.AddCommand(quicksight_startDashboardSnapshotJobScheduleCmd)
}

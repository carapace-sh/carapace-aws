package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_startDashboardSnapshotJobCmd = &cobra.Command{
	Use:   "start-dashboard-snapshot-job",
	Short: "Starts an asynchronous job that generates a snapshot of a dashboard's output.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_startDashboardSnapshotJobCmd).Standalone()

	quicksight_startDashboardSnapshotJobCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that the dashboard snapshot job is executed in.")
	quicksight_startDashboardSnapshotJobCmd.Flags().String("dashboard-id", "", "The ID of the dashboard that you want to start a snapshot job for.")
	quicksight_startDashboardSnapshotJobCmd.Flags().String("snapshot-configuration", "", "A structure that describes the configuration of the dashboard snapshot.")
	quicksight_startDashboardSnapshotJobCmd.Flags().String("snapshot-job-id", "", "An ID for the dashboard snapshot job.")
	quicksight_startDashboardSnapshotJobCmd.Flags().String("user-configuration", "", "A structure that contains information about the anonymous users that the generated snapshot is for.")
	quicksight_startDashboardSnapshotJobCmd.MarkFlagRequired("aws-account-id")
	quicksight_startDashboardSnapshotJobCmd.MarkFlagRequired("dashboard-id")
	quicksight_startDashboardSnapshotJobCmd.MarkFlagRequired("snapshot-configuration")
	quicksight_startDashboardSnapshotJobCmd.MarkFlagRequired("snapshot-job-id")
	quicksight_startDashboardSnapshotJobCmd.MarkFlagRequired("user-configuration")
	quicksightCmd.AddCommand(quicksight_startDashboardSnapshotJobCmd)
}

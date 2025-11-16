package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDashboardSnapshotJobCmd = &cobra.Command{
	Use:   "describe-dashboard-snapshot-job",
	Short: "Describes an existing snapshot job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDashboardSnapshotJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeDashboardSnapshotJobCmd).Standalone()

		quicksight_describeDashboardSnapshotJobCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that the dashboard snapshot job is executed in.")
		quicksight_describeDashboardSnapshotJobCmd.Flags().String("dashboard-id", "", "The ID of the dashboard that you have started a snapshot job for.")
		quicksight_describeDashboardSnapshotJobCmd.Flags().String("snapshot-job-id", "", "The ID of the job to be described.")
		quicksight_describeDashboardSnapshotJobCmd.MarkFlagRequired("aws-account-id")
		quicksight_describeDashboardSnapshotJobCmd.MarkFlagRequired("dashboard-id")
		quicksight_describeDashboardSnapshotJobCmd.MarkFlagRequired("snapshot-job-id")
	})
	quicksightCmd.AddCommand(quicksight_describeDashboardSnapshotJobCmd)
}

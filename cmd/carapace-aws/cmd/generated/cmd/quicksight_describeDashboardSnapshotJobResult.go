package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDashboardSnapshotJobResultCmd = &cobra.Command{
	Use:   "describe-dashboard-snapshot-job-result",
	Short: "Describes the result of an existing snapshot job that has finished running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDashboardSnapshotJobResultCmd).Standalone()

	quicksight_describeDashboardSnapshotJobResultCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that the dashboard snapshot job is executed in.")
	quicksight_describeDashboardSnapshotJobResultCmd.Flags().String("dashboard-id", "", "The ID of the dashboard that you have started a snapshot job for.")
	quicksight_describeDashboardSnapshotJobResultCmd.Flags().String("snapshot-job-id", "", "The ID of the job to be described.")
	quicksight_describeDashboardSnapshotJobResultCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeDashboardSnapshotJobResultCmd.MarkFlagRequired("dashboard-id")
	quicksight_describeDashboardSnapshotJobResultCmd.MarkFlagRequired("snapshot-job-id")
	quicksightCmd.AddCommand(quicksight_describeDashboardSnapshotJobResultCmd)
}

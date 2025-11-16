package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_deleteApplicationSnapshotCmd = &cobra.Command{
	Use:   "delete-application-snapshot",
	Short: "Deletes a snapshot of application state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_deleteApplicationSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_deleteApplicationSnapshotCmd).Standalone()

		kinesisanalyticsv2_deleteApplicationSnapshotCmd.Flags().String("application-name", "", "The name of an existing application.")
		kinesisanalyticsv2_deleteApplicationSnapshotCmd.Flags().String("snapshot-creation-timestamp", "", "The creation timestamp of the application snapshot to delete.")
		kinesisanalyticsv2_deleteApplicationSnapshotCmd.Flags().String("snapshot-name", "", "The identifier for the snapshot delete.")
		kinesisanalyticsv2_deleteApplicationSnapshotCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_deleteApplicationSnapshotCmd.MarkFlagRequired("snapshot-creation-timestamp")
		kinesisanalyticsv2_deleteApplicationSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_deleteApplicationSnapshotCmd)
}

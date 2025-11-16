package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_createApplicationSnapshotCmd = &cobra.Command{
	Use:   "create-application-snapshot",
	Short: "Creates a snapshot of the application's state data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_createApplicationSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisanalyticsv2_createApplicationSnapshotCmd).Standalone()

		kinesisanalyticsv2_createApplicationSnapshotCmd.Flags().String("application-name", "", "The name of an existing application")
		kinesisanalyticsv2_createApplicationSnapshotCmd.Flags().String("snapshot-name", "", "An identifier for the application snapshot.")
		kinesisanalyticsv2_createApplicationSnapshotCmd.MarkFlagRequired("application-name")
		kinesisanalyticsv2_createApplicationSnapshotCmd.MarkFlagRequired("snapshot-name")
	})
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_createApplicationSnapshotCmd)
}

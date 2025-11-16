package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisanalyticsv2_describeApplicationSnapshotCmd = &cobra.Command{
	Use:   "describe-application-snapshot",
	Short: "Returns information about a snapshot of application state data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisanalyticsv2_describeApplicationSnapshotCmd).Standalone()

	kinesisanalyticsv2_describeApplicationSnapshotCmd.Flags().String("application-name", "", "The name of an existing application.")
	kinesisanalyticsv2_describeApplicationSnapshotCmd.Flags().String("snapshot-name", "", "The identifier of an application snapshot.")
	kinesisanalyticsv2_describeApplicationSnapshotCmd.MarkFlagRequired("application-name")
	kinesisanalyticsv2_describeApplicationSnapshotCmd.MarkFlagRequired("snapshot-name")
	kinesisanalyticsv2Cmd.AddCommand(kinesisanalyticsv2_describeApplicationSnapshotCmd)
}

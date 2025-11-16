package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_restoreVolumeFromSnapshotCmd = &cobra.Command{
	Use:   "restore-volume-from-snapshot",
	Short: "Returns an Amazon FSx for OpenZFS volume to the state saved by the specified snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_restoreVolumeFromSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_restoreVolumeFromSnapshotCmd).Standalone()

		fsx_restoreVolumeFromSnapshotCmd.Flags().String("client-request-token", "", "")
		fsx_restoreVolumeFromSnapshotCmd.Flags().String("options", "", "The settings used when restoring the specified volume from snapshot.")
		fsx_restoreVolumeFromSnapshotCmd.Flags().String("snapshot-id", "", "The ID of the source snapshot.")
		fsx_restoreVolumeFromSnapshotCmd.Flags().String("volume-id", "", "The ID of the volume that you are restoring.")
		fsx_restoreVolumeFromSnapshotCmd.MarkFlagRequired("snapshot-id")
		fsx_restoreVolumeFromSnapshotCmd.MarkFlagRequired("volume-id")
	})
	fsxCmd.AddCommand(fsx_restoreVolumeFromSnapshotCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_copySnapshotAndUpdateVolumeCmd = &cobra.Command{
	Use:   "copy-snapshot-and-update-volume",
	Short: "Updates an existing volume by using a snapshot from another Amazon FSx for OpenZFS file system.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_copySnapshotAndUpdateVolumeCmd).Standalone()

	fsx_copySnapshotAndUpdateVolumeCmd.Flags().String("client-request-token", "", "")
	fsx_copySnapshotAndUpdateVolumeCmd.Flags().String("copy-strategy", "", "Specifies the strategy to use when copying data from a snapshot to the volume.")
	fsx_copySnapshotAndUpdateVolumeCmd.Flags().String("options", "", "Confirms that you want to delete data on the destination volume that wasn’t there during the previous snapshot replication.")
	fsx_copySnapshotAndUpdateVolumeCmd.Flags().String("source-snapshot-arn", "", "")
	fsx_copySnapshotAndUpdateVolumeCmd.Flags().String("volume-id", "", "Specifies the ID of the volume that you are copying the snapshot to.")
	fsx_copySnapshotAndUpdateVolumeCmd.MarkFlagRequired("source-snapshot-arn")
	fsx_copySnapshotAndUpdateVolumeCmd.MarkFlagRequired("volume-id")
	fsxCmd.AddCommand(fsx_copySnapshotAndUpdateVolumeCmd)
}

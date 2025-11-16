package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createDiskSnapshotCmd = &cobra.Command{
	Use:   "create-disk-snapshot",
	Short: "Creates a snapshot of a block storage disk.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createDiskSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createDiskSnapshotCmd).Standalone()

		lightsail_createDiskSnapshotCmd.Flags().String("disk-name", "", "The unique name of the source disk (`Disk-Virginia-1`).")
		lightsail_createDiskSnapshotCmd.Flags().String("disk-snapshot-name", "", "The name of the destination disk snapshot (`my-disk-snapshot`) based on the source disk.")
		lightsail_createDiskSnapshotCmd.Flags().String("instance-name", "", "The unique name of the source instance (`Amazon_Linux-512MB-Virginia-1`).")
		lightsail_createDiskSnapshotCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
		lightsail_createDiskSnapshotCmd.MarkFlagRequired("disk-snapshot-name")
	})
	lightsailCmd.AddCommand(lightsail_createDiskSnapshotCmd)
}

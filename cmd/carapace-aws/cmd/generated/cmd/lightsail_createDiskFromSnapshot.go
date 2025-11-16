package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createDiskFromSnapshotCmd = &cobra.Command{
	Use:   "create-disk-from-snapshot",
	Short: "Creates a block storage disk from a manual or automatic snapshot of a disk.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createDiskFromSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createDiskFromSnapshotCmd).Standalone()

		lightsail_createDiskFromSnapshotCmd.Flags().String("add-ons", "", "An array of objects that represent the add-ons to enable for the new disk.")
		lightsail_createDiskFromSnapshotCmd.Flags().String("availability-zone", "", "The Availability Zone where you want to create the disk (`us-east-2a`).")
		lightsail_createDiskFromSnapshotCmd.Flags().String("disk-name", "", "The unique Lightsail disk name (`my-disk`).")
		lightsail_createDiskFromSnapshotCmd.Flags().String("disk-snapshot-name", "", "The name of the disk snapshot (`my-snapshot`) from which to create the new storage disk.")
		lightsail_createDiskFromSnapshotCmd.Flags().String("restore-date", "", "The date of the automatic snapshot to use for the new disk.")
		lightsail_createDiskFromSnapshotCmd.Flags().String("size-in-gb", "", "The size of the disk in GB (`32`).")
		lightsail_createDiskFromSnapshotCmd.Flags().String("source-disk-name", "", "The name of the source disk from which the source automatic snapshot was created.")
		lightsail_createDiskFromSnapshotCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
		lightsail_createDiskFromSnapshotCmd.Flags().String("use-latest-restorable-auto-snapshot", "", "A Boolean value to indicate whether to use the latest available automatic snapshot.")
		lightsail_createDiskFromSnapshotCmd.MarkFlagRequired("availability-zone")
		lightsail_createDiskFromSnapshotCmd.MarkFlagRequired("disk-name")
		lightsail_createDiskFromSnapshotCmd.MarkFlagRequired("size-in-gb")
	})
	lightsailCmd.AddCommand(lightsail_createDiskFromSnapshotCmd)
}

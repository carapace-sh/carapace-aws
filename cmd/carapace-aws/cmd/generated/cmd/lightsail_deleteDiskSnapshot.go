package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteDiskSnapshotCmd = &cobra.Command{
	Use:   "delete-disk-snapshot",
	Short: "Deletes the specified disk snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteDiskSnapshotCmd).Standalone()

	lightsail_deleteDiskSnapshotCmd.Flags().String("disk-snapshot-name", "", "The name of the disk snapshot you want to delete (`my-disk-snapshot`).")
	lightsail_deleteDiskSnapshotCmd.MarkFlagRequired("disk-snapshot-name")
	lightsailCmd.AddCommand(lightsail_deleteDiskSnapshotCmd)
}

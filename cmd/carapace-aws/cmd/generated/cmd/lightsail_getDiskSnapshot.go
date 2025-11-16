package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getDiskSnapshotCmd = &cobra.Command{
	Use:   "get-disk-snapshot",
	Short: "Returns information about a specific block storage disk snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getDiskSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getDiskSnapshotCmd).Standalone()

		lightsail_getDiskSnapshotCmd.Flags().String("disk-snapshot-name", "", "The name of the disk snapshot (`my-disk-snapshot`).")
		lightsail_getDiskSnapshotCmd.MarkFlagRequired("disk-snapshot-name")
	})
	lightsailCmd.AddCommand(lightsail_getDiskSnapshotCmd)
}

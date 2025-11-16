package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_attachDiskCmd = &cobra.Command{
	Use:   "attach-disk",
	Short: "Attaches a block storage disk to a running or stopped Lightsail instance and exposes it to the instance with the specified disk name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_attachDiskCmd).Standalone()

	lightsail_attachDiskCmd.Flags().String("auto-mounting", "", "A Boolean value used to determine the automatic mounting of a storage volume to a virtual computer.")
	lightsail_attachDiskCmd.Flags().String("disk-name", "", "The unique Lightsail disk name (`my-disk`).")
	lightsail_attachDiskCmd.Flags().String("disk-path", "", "The disk path to expose to the instance (`/dev/xvdf`).")
	lightsail_attachDiskCmd.Flags().String("instance-name", "", "The name of the Lightsail instance where you want to utilize the storage disk.")
	lightsail_attachDiskCmd.MarkFlagRequired("disk-name")
	lightsail_attachDiskCmd.MarkFlagRequired("disk-path")
	lightsail_attachDiskCmd.MarkFlagRequired("instance-name")
	lightsailCmd.AddCommand(lightsail_attachDiskCmd)
}

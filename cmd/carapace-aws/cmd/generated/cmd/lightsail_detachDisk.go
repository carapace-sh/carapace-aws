package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_detachDiskCmd = &cobra.Command{
	Use:   "detach-disk",
	Short: "Detaches a stopped block storage disk from a Lightsail instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_detachDiskCmd).Standalone()

	lightsail_detachDiskCmd.Flags().String("disk-name", "", "The unique name of the disk you want to detach from your instance (`my-disk`).")
	lightsail_detachDiskCmd.MarkFlagRequired("disk-name")
	lightsailCmd.AddCommand(lightsail_detachDiskCmd)
}

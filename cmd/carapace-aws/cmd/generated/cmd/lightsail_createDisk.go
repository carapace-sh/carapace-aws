package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createDiskCmd = &cobra.Command{
	Use:   "create-disk",
	Short: "Creates a block storage disk that can be attached to an Amazon Lightsail instance in the same Availability Zone (`us-east-2a`).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createDiskCmd).Standalone()

	lightsail_createDiskCmd.Flags().String("add-ons", "", "An array of objects that represent the add-ons to enable for the new disk.")
	lightsail_createDiskCmd.Flags().String("availability-zone", "", "The Availability Zone where you want to create the disk (`us-east-2a`).")
	lightsail_createDiskCmd.Flags().String("disk-name", "", "The unique Lightsail disk name (`my-disk`).")
	lightsail_createDiskCmd.Flags().String("size-in-gb", "", "The size of the disk in GB (`32`).")
	lightsail_createDiskCmd.Flags().String("tags", "", "The tag keys and optional values to add to the resource during create.")
	lightsail_createDiskCmd.MarkFlagRequired("availability-zone")
	lightsail_createDiskCmd.MarkFlagRequired("disk-name")
	lightsail_createDiskCmd.MarkFlagRequired("size-in-gb")
	lightsailCmd.AddCommand(lightsail_createDiskCmd)
}

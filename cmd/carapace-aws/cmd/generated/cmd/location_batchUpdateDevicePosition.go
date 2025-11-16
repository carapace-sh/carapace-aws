package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_batchUpdateDevicePositionCmd = &cobra.Command{
	Use:   "batch-update-device-position",
	Short: "Uploads position update data for one or more devices to a tracker resource (up to 10 devices per batch).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_batchUpdateDevicePositionCmd).Standalone()

	location_batchUpdateDevicePositionCmd.Flags().String("tracker-name", "", "The name of the tracker resource to update.")
	location_batchUpdateDevicePositionCmd.Flags().String("updates", "", "Contains the position update details for each device, up to 10 devices.")
	location_batchUpdateDevicePositionCmd.MarkFlagRequired("tracker-name")
	location_batchUpdateDevicePositionCmd.MarkFlagRequired("updates")
	locationCmd.AddCommand(location_batchUpdateDevicePositionCmd)
}

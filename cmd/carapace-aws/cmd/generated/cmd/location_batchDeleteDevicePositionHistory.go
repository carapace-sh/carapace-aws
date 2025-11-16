package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_batchDeleteDevicePositionHistoryCmd = &cobra.Command{
	Use:   "batch-delete-device-position-history",
	Short: "Deletes the position history of one or more devices from a tracker resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_batchDeleteDevicePositionHistoryCmd).Standalone()

	location_batchDeleteDevicePositionHistoryCmd.Flags().String("device-ids", "", "Devices whose position history you want to delete.")
	location_batchDeleteDevicePositionHistoryCmd.Flags().String("tracker-name", "", "The name of the tracker resource to delete the device position history from.")
	location_batchDeleteDevicePositionHistoryCmd.MarkFlagRequired("device-ids")
	location_batchDeleteDevicePositionHistoryCmd.MarkFlagRequired("tracker-name")
	locationCmd.AddCommand(location_batchDeleteDevicePositionHistoryCmd)
}

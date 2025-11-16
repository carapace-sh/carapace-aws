package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_batchGetDevicePositionCmd = &cobra.Command{
	Use:   "batch-get-device-position",
	Short: "Lists the latest device positions for requested devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_batchGetDevicePositionCmd).Standalone()

	location_batchGetDevicePositionCmd.Flags().String("device-ids", "", "Devices whose position you want to retrieve.")
	location_batchGetDevicePositionCmd.Flags().String("tracker-name", "", "The tracker resource retrieving the device position.")
	location_batchGetDevicePositionCmd.MarkFlagRequired("device-ids")
	location_batchGetDevicePositionCmd.MarkFlagRequired("tracker-name")
	locationCmd.AddCommand(location_batchGetDevicePositionCmd)
}

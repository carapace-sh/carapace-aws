package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_verifyDevicePositionCmd = &cobra.Command{
	Use:   "verify-device-position",
	Short: "Verifies the integrity of the device's position by determining if it was reported behind a proxy, and by comparing it to an inferred position estimated based on the device's state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_verifyDevicePositionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_verifyDevicePositionCmd).Standalone()

		location_verifyDevicePositionCmd.Flags().String("device-state", "", "The device's state, including position, IP address, cell signals and Wi-Fi access points.")
		location_verifyDevicePositionCmd.Flags().String("distance-unit", "", "The distance unit for the verification request.")
		location_verifyDevicePositionCmd.Flags().String("tracker-name", "", "The name of the tracker resource to be associated with verification request.")
		location_verifyDevicePositionCmd.MarkFlagRequired("device-state")
		location_verifyDevicePositionCmd.MarkFlagRequired("tracker-name")
	})
	locationCmd.AddCommand(location_verifyDevicePositionCmd)
}

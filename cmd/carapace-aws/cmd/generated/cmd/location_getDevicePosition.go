package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_getDevicePositionCmd = &cobra.Command{
	Use:   "get-device-position",
	Short: "Retrieves a device's most recent position according to its sample time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_getDevicePositionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_getDevicePositionCmd).Standalone()

		location_getDevicePositionCmd.Flags().String("device-id", "", "The device whose position you want to retrieve.")
		location_getDevicePositionCmd.Flags().String("tracker-name", "", "The tracker resource receiving the position update.")
		location_getDevicePositionCmd.MarkFlagRequired("device-id")
		location_getDevicePositionCmd.MarkFlagRequired("tracker-name")
	})
	locationCmd.AddCommand(location_getDevicePositionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createWirelessDeviceCmd = &cobra.Command{
	Use:   "create-wireless-device",
	Short: "Provisions a wireless device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createWirelessDeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_createWirelessDeviceCmd).Standalone()

		iotwireless_createWirelessDeviceCmd.Flags().String("client-request-token", "", "Each resource must have a unique client request token.")
		iotwireless_createWirelessDeviceCmd.Flags().String("description", "", "The description of the new resource.")
		iotwireless_createWirelessDeviceCmd.Flags().String("destination-name", "", "The name of the destination to assign to the new wireless device.")
		iotwireless_createWirelessDeviceCmd.Flags().String("lo-ra-wan", "", "The device configuration information to use to create the wireless device.")
		iotwireless_createWirelessDeviceCmd.Flags().String("name", "", "The name of the new resource.")
		iotwireless_createWirelessDeviceCmd.Flags().String("positioning", "", "The integration status of the Device Location feature for LoRaWAN and Sidewalk devices.")
		iotwireless_createWirelessDeviceCmd.Flags().String("sidewalk", "", "The device configuration information to use to create the Sidewalk device.")
		iotwireless_createWirelessDeviceCmd.Flags().String("tags", "", "The tags to attach to the new wireless device.")
		iotwireless_createWirelessDeviceCmd.Flags().String("type", "", "The wireless device type.")
		iotwireless_createWirelessDeviceCmd.MarkFlagRequired("destination-name")
		iotwireless_createWirelessDeviceCmd.MarkFlagRequired("type")
	})
	iotwirelessCmd.AddCommand(iotwireless_createWirelessDeviceCmd)
}

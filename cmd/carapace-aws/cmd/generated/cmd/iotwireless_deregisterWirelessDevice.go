package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deregisterWirelessDeviceCmd = &cobra.Command{
	Use:   "deregister-wireless-device",
	Short: "Deregister a wireless device from AWS IoT Wireless.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deregisterWirelessDeviceCmd).Standalone()

	iotwireless_deregisterWirelessDeviceCmd.Flags().String("identifier", "", "The identifier of the wireless device to deregister from AWS IoT Wireless.")
	iotwireless_deregisterWirelessDeviceCmd.Flags().String("wireless-device-type", "", "The type of wireless device to deregister from AWS IoT Wireless, which can be `LoRaWAN` or `Sidewalk`.")
	iotwireless_deregisterWirelessDeviceCmd.MarkFlagRequired("identifier")
	iotwirelessCmd.AddCommand(iotwireless_deregisterWirelessDeviceCmd)
}

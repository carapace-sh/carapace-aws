package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateWirelessDeviceCmd = &cobra.Command{
	Use:   "update-wireless-device",
	Short: "Updates properties of a wireless device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateWirelessDeviceCmd).Standalone()

	iotwireless_updateWirelessDeviceCmd.Flags().String("description", "", "A new description of the resource.")
	iotwireless_updateWirelessDeviceCmd.Flags().String("destination-name", "", "The name of the new destination for the device.")
	iotwireless_updateWirelessDeviceCmd.Flags().String("id", "", "The ID of the resource to update.")
	iotwireless_updateWirelessDeviceCmd.Flags().String("lo-ra-wan", "", "The updated wireless device's configuration.")
	iotwireless_updateWirelessDeviceCmd.Flags().String("name", "", "The new name of the resource.")
	iotwireless_updateWirelessDeviceCmd.Flags().String("positioning", "", "The integration status of the Device Location feature for LoRaWAN and Sidewalk devices.")
	iotwireless_updateWirelessDeviceCmd.Flags().String("sidewalk", "", "The updated sidewalk properties.")
	iotwireless_updateWirelessDeviceCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_updateWirelessDeviceCmd)
}

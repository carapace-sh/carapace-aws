package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessDeviceCmd = &cobra.Command{
	Use:   "get-wireless-device",
	Short: "Gets information about a wireless device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessDeviceCmd).Standalone()

	iotwireless_getWirelessDeviceCmd.Flags().String("identifier", "", "The identifier of the wireless device to get.")
	iotwireless_getWirelessDeviceCmd.Flags().String("identifier-type", "", "The type of identifier used in `identifier`.")
	iotwireless_getWirelessDeviceCmd.MarkFlagRequired("identifier")
	iotwireless_getWirelessDeviceCmd.MarkFlagRequired("identifier-type")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessDeviceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_associateWirelessDeviceWithMulticastGroupCmd = &cobra.Command{
	Use:   "associate-wireless-device-with-multicast-group",
	Short: "Associates a wireless device with a multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_associateWirelessDeviceWithMulticastGroupCmd).Standalone()

	iotwireless_associateWirelessDeviceWithMulticastGroupCmd.Flags().String("id", "", "")
	iotwireless_associateWirelessDeviceWithMulticastGroupCmd.Flags().String("wireless-device-id", "", "")
	iotwireless_associateWirelessDeviceWithMulticastGroupCmd.MarkFlagRequired("id")
	iotwireless_associateWirelessDeviceWithMulticastGroupCmd.MarkFlagRequired("wireless-device-id")
	iotwirelessCmd.AddCommand(iotwireless_associateWirelessDeviceWithMulticastGroupCmd)
}

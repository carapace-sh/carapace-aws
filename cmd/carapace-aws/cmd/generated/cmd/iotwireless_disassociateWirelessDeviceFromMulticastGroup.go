package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_disassociateWirelessDeviceFromMulticastGroupCmd = &cobra.Command{
	Use:   "disassociate-wireless-device-from-multicast-group",
	Short: "Disassociates a wireless device from a multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_disassociateWirelessDeviceFromMulticastGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_disassociateWirelessDeviceFromMulticastGroupCmd).Standalone()

		iotwireless_disassociateWirelessDeviceFromMulticastGroupCmd.Flags().String("id", "", "")
		iotwireless_disassociateWirelessDeviceFromMulticastGroupCmd.Flags().String("wireless-device-id", "", "")
		iotwireless_disassociateWirelessDeviceFromMulticastGroupCmd.MarkFlagRequired("id")
		iotwireless_disassociateWirelessDeviceFromMulticastGroupCmd.MarkFlagRequired("wireless-device-id")
	})
	iotwirelessCmd.AddCommand(iotwireless_disassociateWirelessDeviceFromMulticastGroupCmd)
}

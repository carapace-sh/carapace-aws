package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_startBulkAssociateWirelessDeviceWithMulticastGroupCmd = &cobra.Command{
	Use:   "start-bulk-associate-wireless-device-with-multicast-group",
	Short: "Starts a bulk association of all qualifying wireless devices with a multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_startBulkAssociateWirelessDeviceWithMulticastGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_startBulkAssociateWirelessDeviceWithMulticastGroupCmd).Standalone()

		iotwireless_startBulkAssociateWirelessDeviceWithMulticastGroupCmd.Flags().String("id", "", "")
		iotwireless_startBulkAssociateWirelessDeviceWithMulticastGroupCmd.Flags().String("query-string", "", "")
		iotwireless_startBulkAssociateWirelessDeviceWithMulticastGroupCmd.Flags().String("tags", "", "")
		iotwireless_startBulkAssociateWirelessDeviceWithMulticastGroupCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_startBulkAssociateWirelessDeviceWithMulticastGroupCmd)
}

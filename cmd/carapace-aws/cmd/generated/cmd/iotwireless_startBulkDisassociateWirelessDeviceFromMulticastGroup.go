package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_startBulkDisassociateWirelessDeviceFromMulticastGroupCmd = &cobra.Command{
	Use:   "start-bulk-disassociate-wireless-device-from-multicast-group",
	Short: "Starts a bulk disassociatin of all qualifying wireless devices from a multicast group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_startBulkDisassociateWirelessDeviceFromMulticastGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_startBulkDisassociateWirelessDeviceFromMulticastGroupCmd).Standalone()

		iotwireless_startBulkDisassociateWirelessDeviceFromMulticastGroupCmd.Flags().String("id", "", "")
		iotwireless_startBulkDisassociateWirelessDeviceFromMulticastGroupCmd.Flags().String("query-string", "", "")
		iotwireless_startBulkDisassociateWirelessDeviceFromMulticastGroupCmd.Flags().String("tags", "", "")
		iotwireless_startBulkDisassociateWirelessDeviceFromMulticastGroupCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_startBulkDisassociateWirelessDeviceFromMulticastGroupCmd)
}

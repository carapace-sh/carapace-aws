package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_disassociateWirelessDeviceFromThingCmd = &cobra.Command{
	Use:   "disassociate-wireless-device-from-thing",
	Short: "Disassociates a wireless device from its currently associated thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_disassociateWirelessDeviceFromThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_disassociateWirelessDeviceFromThingCmd).Standalone()

		iotwireless_disassociateWirelessDeviceFromThingCmd.Flags().String("id", "", "The ID of the resource to update.")
		iotwireless_disassociateWirelessDeviceFromThingCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_disassociateWirelessDeviceFromThingCmd)
}

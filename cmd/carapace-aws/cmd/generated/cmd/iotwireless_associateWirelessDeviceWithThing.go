package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_associateWirelessDeviceWithThingCmd = &cobra.Command{
	Use:   "associate-wireless-device-with-thing",
	Short: "Associates a wireless device with a thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_associateWirelessDeviceWithThingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_associateWirelessDeviceWithThingCmd).Standalone()

		iotwireless_associateWirelessDeviceWithThingCmd.Flags().String("id", "", "The ID of the resource to update.")
		iotwireless_associateWirelessDeviceWithThingCmd.Flags().String("thing-arn", "", "The ARN of the thing to associate with the wireless device.")
		iotwireless_associateWirelessDeviceWithThingCmd.MarkFlagRequired("id")
		iotwireless_associateWirelessDeviceWithThingCmd.MarkFlagRequired("thing-arn")
	})
	iotwirelessCmd.AddCommand(iotwireless_associateWirelessDeviceWithThingCmd)
}

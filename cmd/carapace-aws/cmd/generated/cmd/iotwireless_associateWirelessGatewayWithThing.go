package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_associateWirelessGatewayWithThingCmd = &cobra.Command{
	Use:   "associate-wireless-gateway-with-thing",
	Short: "Associates a wireless gateway with a thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_associateWirelessGatewayWithThingCmd).Standalone()

	iotwireless_associateWirelessGatewayWithThingCmd.Flags().String("id", "", "The ID of the resource to update.")
	iotwireless_associateWirelessGatewayWithThingCmd.Flags().String("thing-arn", "", "The ARN of the thing to associate with the wireless gateway.")
	iotwireless_associateWirelessGatewayWithThingCmd.MarkFlagRequired("id")
	iotwireless_associateWirelessGatewayWithThingCmd.MarkFlagRequired("thing-arn")
	iotwirelessCmd.AddCommand(iotwireless_associateWirelessGatewayWithThingCmd)
}

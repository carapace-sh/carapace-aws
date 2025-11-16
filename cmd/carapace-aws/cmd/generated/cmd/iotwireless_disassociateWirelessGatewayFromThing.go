package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_disassociateWirelessGatewayFromThingCmd = &cobra.Command{
	Use:   "disassociate-wireless-gateway-from-thing",
	Short: "Disassociates a wireless gateway from its currently associated thing.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_disassociateWirelessGatewayFromThingCmd).Standalone()

	iotwireless_disassociateWirelessGatewayFromThingCmd.Flags().String("id", "", "The ID of the resource to update.")
	iotwireless_disassociateWirelessGatewayFromThingCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_disassociateWirelessGatewayFromThingCmd)
}

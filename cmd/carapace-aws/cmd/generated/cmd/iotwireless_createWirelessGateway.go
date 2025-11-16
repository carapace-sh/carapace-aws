package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createWirelessGatewayCmd = &cobra.Command{
	Use:   "create-wireless-gateway",
	Short: "Provisions a wireless gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createWirelessGatewayCmd).Standalone()

	iotwireless_createWirelessGatewayCmd.Flags().String("client-request-token", "", "Each resource must have a unique client request token.")
	iotwireless_createWirelessGatewayCmd.Flags().String("description", "", "The description of the new resource.")
	iotwireless_createWirelessGatewayCmd.Flags().String("lo-ra-wan", "", "The gateway configuration information to use to create the wireless gateway.")
	iotwireless_createWirelessGatewayCmd.Flags().String("name", "", "The name of the new resource.")
	iotwireless_createWirelessGatewayCmd.Flags().String("tags", "", "The tags to attach to the new wireless gateway.")
	iotwireless_createWirelessGatewayCmd.MarkFlagRequired("lo-ra-wan")
	iotwirelessCmd.AddCommand(iotwireless_createWirelessGatewayCmd)
}

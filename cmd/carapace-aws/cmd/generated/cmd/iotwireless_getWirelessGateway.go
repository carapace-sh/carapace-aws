package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessGatewayCmd = &cobra.Command{
	Use:   "get-wireless-gateway",
	Short: "Gets information about a wireless gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessGatewayCmd).Standalone()

	iotwireless_getWirelessGatewayCmd.Flags().String("identifier", "", "The identifier of the wireless gateway to get.")
	iotwireless_getWirelessGatewayCmd.Flags().String("identifier-type", "", "The type of identifier used in `identifier`.")
	iotwireless_getWirelessGatewayCmd.MarkFlagRequired("identifier")
	iotwireless_getWirelessGatewayCmd.MarkFlagRequired("identifier-type")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessGatewayCmd)
}

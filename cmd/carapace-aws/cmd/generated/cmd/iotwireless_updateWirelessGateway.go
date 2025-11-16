package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_updateWirelessGatewayCmd = &cobra.Command{
	Use:   "update-wireless-gateway",
	Short: "Updates properties of a wireless gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_updateWirelessGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_updateWirelessGatewayCmd).Standalone()

		iotwireless_updateWirelessGatewayCmd.Flags().String("description", "", "A new description of the resource.")
		iotwireless_updateWirelessGatewayCmd.Flags().String("id", "", "The ID of the resource to update.")
		iotwireless_updateWirelessGatewayCmd.Flags().String("join-eui-filters", "", "")
		iotwireless_updateWirelessGatewayCmd.Flags().String("max-eirp", "", "The MaxEIRP value.")
		iotwireless_updateWirelessGatewayCmd.Flags().String("name", "", "The new name of the resource.")
		iotwireless_updateWirelessGatewayCmd.Flags().String("net-id-filters", "", "")
		iotwireless_updateWirelessGatewayCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_updateWirelessGatewayCmd)
}

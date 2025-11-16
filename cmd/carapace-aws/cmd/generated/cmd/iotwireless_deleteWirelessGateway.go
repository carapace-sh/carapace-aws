package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteWirelessGatewayCmd = &cobra.Command{
	Use:   "delete-wireless-gateway",
	Short: "Deletes a wireless gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteWirelessGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_deleteWirelessGatewayCmd).Standalone()

		iotwireless_deleteWirelessGatewayCmd.Flags().String("id", "", "The ID of the resource to delete.")
		iotwireless_deleteWirelessGatewayCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_deleteWirelessGatewayCmd)
}

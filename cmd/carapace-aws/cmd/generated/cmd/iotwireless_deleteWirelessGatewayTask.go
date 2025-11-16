package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteWirelessGatewayTaskCmd = &cobra.Command{
	Use:   "delete-wireless-gateway-task",
	Short: "Deletes a wireless gateway task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteWirelessGatewayTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_deleteWirelessGatewayTaskCmd).Standalone()

		iotwireless_deleteWirelessGatewayTaskCmd.Flags().String("id", "", "The ID of the resource to delete.")
		iotwireless_deleteWirelessGatewayTaskCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_deleteWirelessGatewayTaskCmd)
}

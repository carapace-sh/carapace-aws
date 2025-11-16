package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createWirelessGatewayTaskCmd = &cobra.Command{
	Use:   "create-wireless-gateway-task",
	Short: "Creates a task for a wireless gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createWirelessGatewayTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_createWirelessGatewayTaskCmd).Standalone()

		iotwireless_createWirelessGatewayTaskCmd.Flags().String("id", "", "The ID of the resource to update.")
		iotwireless_createWirelessGatewayTaskCmd.Flags().String("wireless-gateway-task-definition-id", "", "The ID of the WirelessGatewayTaskDefinition.")
		iotwireless_createWirelessGatewayTaskCmd.MarkFlagRequired("id")
		iotwireless_createWirelessGatewayTaskCmd.MarkFlagRequired("wireless-gateway-task-definition-id")
	})
	iotwirelessCmd.AddCommand(iotwireless_createWirelessGatewayTaskCmd)
}

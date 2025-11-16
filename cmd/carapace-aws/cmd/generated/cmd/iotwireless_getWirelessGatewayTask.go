package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessGatewayTaskCmd = &cobra.Command{
	Use:   "get-wireless-gateway-task",
	Short: "Gets information about a wireless gateway task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessGatewayTaskCmd).Standalone()

	iotwireless_getWirelessGatewayTaskCmd.Flags().String("id", "", "The ID of the resource to get.")
	iotwireless_getWirelessGatewayTaskCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessGatewayTaskCmd)
}

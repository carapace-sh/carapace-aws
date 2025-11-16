package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_getWirelessGatewayTaskDefinitionCmd = &cobra.Command{
	Use:   "get-wireless-gateway-task-definition",
	Short: "Gets information about a wireless gateway task definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_getWirelessGatewayTaskDefinitionCmd).Standalone()

	iotwireless_getWirelessGatewayTaskDefinitionCmd.Flags().String("id", "", "The ID of the resource to get.")
	iotwireless_getWirelessGatewayTaskDefinitionCmd.MarkFlagRequired("id")
	iotwirelessCmd.AddCommand(iotwireless_getWirelessGatewayTaskDefinitionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_createWirelessGatewayTaskDefinitionCmd = &cobra.Command{
	Use:   "create-wireless-gateway-task-definition",
	Short: "Creates a gateway task definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_createWirelessGatewayTaskDefinitionCmd).Standalone()

	iotwireless_createWirelessGatewayTaskDefinitionCmd.Flags().String("auto-create-tasks", "", "Whether to automatically create tasks using this task definition for all gateways with the specified current version.")
	iotwireless_createWirelessGatewayTaskDefinitionCmd.Flags().String("client-request-token", "", "Each resource must have a unique client request token.")
	iotwireless_createWirelessGatewayTaskDefinitionCmd.Flags().String("name", "", "The name of the new resource.")
	iotwireless_createWirelessGatewayTaskDefinitionCmd.Flags().String("tags", "", "The tags to attach to the specified resource.")
	iotwireless_createWirelessGatewayTaskDefinitionCmd.Flags().String("update", "", "Information about the gateways to update.")
	iotwireless_createWirelessGatewayTaskDefinitionCmd.MarkFlagRequired("auto-create-tasks")
	iotwirelessCmd.AddCommand(iotwireless_createWirelessGatewayTaskDefinitionCmd)
}

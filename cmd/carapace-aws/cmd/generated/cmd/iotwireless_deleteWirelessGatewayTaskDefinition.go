package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_deleteWirelessGatewayTaskDefinitionCmd = &cobra.Command{
	Use:   "delete-wireless-gateway-task-definition",
	Short: "Deletes a wireless gateway task definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_deleteWirelessGatewayTaskDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_deleteWirelessGatewayTaskDefinitionCmd).Standalone()

		iotwireless_deleteWirelessGatewayTaskDefinitionCmd.Flags().String("id", "", "The ID of the resource to delete.")
		iotwireless_deleteWirelessGatewayTaskDefinitionCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_deleteWirelessGatewayTaskDefinitionCmd)
}

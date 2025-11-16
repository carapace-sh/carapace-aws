package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_updateDirectConnectGatewayCmd = &cobra.Command{
	Use:   "update-direct-connect-gateway",
	Short: "Updates the name of a current Direct Connect gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_updateDirectConnectGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_updateDirectConnectGatewayCmd).Standalone()

		directconnect_updateDirectConnectGatewayCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway to update.")
		directconnect_updateDirectConnectGatewayCmd.Flags().String("new-direct-connect-gateway-name", "", "The new name for the Direct Connect gateway.")
		directconnect_updateDirectConnectGatewayCmd.MarkFlagRequired("direct-connect-gateway-id")
		directconnect_updateDirectConnectGatewayCmd.MarkFlagRequired("new-direct-connect-gateway-name")
	})
	directconnectCmd.AddCommand(directconnect_updateDirectConnectGatewayCmd)
}

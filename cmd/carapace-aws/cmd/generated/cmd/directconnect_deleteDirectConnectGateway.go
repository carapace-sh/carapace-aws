package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_deleteDirectConnectGatewayCmd = &cobra.Command{
	Use:   "delete-direct-connect-gateway",
	Short: "Deletes the specified Direct Connect gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_deleteDirectConnectGatewayCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_deleteDirectConnectGatewayCmd).Standalone()

		directconnect_deleteDirectConnectGatewayCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
		directconnect_deleteDirectConnectGatewayCmd.MarkFlagRequired("direct-connect-gateway-id")
	})
	directconnectCmd.AddCommand(directconnect_deleteDirectConnectGatewayCmd)
}

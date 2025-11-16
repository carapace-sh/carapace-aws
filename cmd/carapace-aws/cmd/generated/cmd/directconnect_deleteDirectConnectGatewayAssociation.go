package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_deleteDirectConnectGatewayAssociationCmd = &cobra.Command{
	Use:   "delete-direct-connect-gateway-association",
	Short: "Deletes the association between the specified Direct Connect gateway and virtual private gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_deleteDirectConnectGatewayAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_deleteDirectConnectGatewayAssociationCmd).Standalone()

		directconnect_deleteDirectConnectGatewayAssociationCmd.Flags().String("association-id", "", "The ID of the Direct Connect gateway association.")
		directconnect_deleteDirectConnectGatewayAssociationCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
		directconnect_deleteDirectConnectGatewayAssociationCmd.Flags().String("virtual-gateway-id", "", "The ID of the virtual private gateway.")
	})
	directconnectCmd.AddCommand(directconnect_deleteDirectConnectGatewayAssociationCmd)
}

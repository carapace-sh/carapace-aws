package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_updateDirectConnectGatewayAssociationCmd = &cobra.Command{
	Use:   "update-direct-connect-gateway-association",
	Short: "Updates the specified attributes of the Direct Connect gateway association.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_updateDirectConnectGatewayAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_updateDirectConnectGatewayAssociationCmd).Standalone()

		directconnect_updateDirectConnectGatewayAssociationCmd.Flags().String("add-allowed-prefixes-to-direct-connect-gateway", "", "The Amazon VPC prefixes to advertise to the Direct Connect gateway.")
		directconnect_updateDirectConnectGatewayAssociationCmd.Flags().String("association-id", "", "The ID of the Direct Connect gateway association.")
		directconnect_updateDirectConnectGatewayAssociationCmd.Flags().String("remove-allowed-prefixes-to-direct-connect-gateway", "", "The Amazon VPC prefixes to no longer advertise to the Direct Connect gateway.")
	})
	directconnectCmd.AddCommand(directconnect_updateDirectConnectGatewayAssociationCmd)
}

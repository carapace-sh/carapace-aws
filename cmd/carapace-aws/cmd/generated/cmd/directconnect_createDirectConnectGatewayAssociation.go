package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createDirectConnectGatewayAssociationCmd = &cobra.Command{
	Use:   "create-direct-connect-gateway-association",
	Short: "Creates an association between a Direct Connect gateway and a virtual private gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createDirectConnectGatewayAssociationCmd).Standalone()

	directconnect_createDirectConnectGatewayAssociationCmd.Flags().String("add-allowed-prefixes-to-direct-connect-gateway", "", "The Amazon VPC prefixes to advertise to the Direct Connect gateway")
	directconnect_createDirectConnectGatewayAssociationCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
	directconnect_createDirectConnectGatewayAssociationCmd.Flags().String("gateway-id", "", "The ID of the virtual private gateway or transit gateway.")
	directconnect_createDirectConnectGatewayAssociationCmd.Flags().String("virtual-gateway-id", "", "The ID of the virtual private gateway.")
	directconnect_createDirectConnectGatewayAssociationCmd.MarkFlagRequired("direct-connect-gateway-id")
	directconnectCmd.AddCommand(directconnect_createDirectConnectGatewayAssociationCmd)
}

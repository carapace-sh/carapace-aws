package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_createDirectConnectGatewayAssociationProposalCmd = &cobra.Command{
	Use:   "create-direct-connect-gateway-association-proposal",
	Short: "Creates a proposal to associate the specified virtual private gateway or transit gateway with the specified Direct Connect gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_createDirectConnectGatewayAssociationProposalCmd).Standalone()

	directconnect_createDirectConnectGatewayAssociationProposalCmd.Flags().String("add-allowed-prefixes-to-direct-connect-gateway", "", "The Amazon VPC prefixes to advertise to the Direct Connect gateway.")
	directconnect_createDirectConnectGatewayAssociationProposalCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
	directconnect_createDirectConnectGatewayAssociationProposalCmd.Flags().String("direct-connect-gateway-owner-account", "", "The ID of the Amazon Web Services account that owns the Direct Connect gateway.")
	directconnect_createDirectConnectGatewayAssociationProposalCmd.Flags().String("gateway-id", "", "The ID of the virtual private gateway or transit gateway.")
	directconnect_createDirectConnectGatewayAssociationProposalCmd.Flags().String("remove-allowed-prefixes-to-direct-connect-gateway", "", "The Amazon VPC prefixes to no longer advertise to the Direct Connect gateway.")
	directconnect_createDirectConnectGatewayAssociationProposalCmd.MarkFlagRequired("direct-connect-gateway-id")
	directconnect_createDirectConnectGatewayAssociationProposalCmd.MarkFlagRequired("direct-connect-gateway-owner-account")
	directconnect_createDirectConnectGatewayAssociationProposalCmd.MarkFlagRequired("gateway-id")
	directconnectCmd.AddCommand(directconnect_createDirectConnectGatewayAssociationProposalCmd)
}

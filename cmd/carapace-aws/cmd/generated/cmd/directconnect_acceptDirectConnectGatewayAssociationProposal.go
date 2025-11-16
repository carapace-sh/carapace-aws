package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_acceptDirectConnectGatewayAssociationProposalCmd = &cobra.Command{
	Use:   "accept-direct-connect-gateway-association-proposal",
	Short: "Accepts a proposal request to attach a virtual private gateway or transit gateway to a Direct Connect gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_acceptDirectConnectGatewayAssociationProposalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_acceptDirectConnectGatewayAssociationProposalCmd).Standalone()

		directconnect_acceptDirectConnectGatewayAssociationProposalCmd.Flags().String("associated-gateway-owner-account", "", "The ID of the Amazon Web Services account that owns the virtual private gateway or transit gateway.")
		directconnect_acceptDirectConnectGatewayAssociationProposalCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
		directconnect_acceptDirectConnectGatewayAssociationProposalCmd.Flags().String("override-allowed-prefixes-to-direct-connect-gateway", "", "Overrides the Amazon VPC prefixes advertised to the Direct Connect gateway.")
		directconnect_acceptDirectConnectGatewayAssociationProposalCmd.Flags().String("proposal-id", "", "The ID of the request proposal.")
		directconnect_acceptDirectConnectGatewayAssociationProposalCmd.MarkFlagRequired("associated-gateway-owner-account")
		directconnect_acceptDirectConnectGatewayAssociationProposalCmd.MarkFlagRequired("direct-connect-gateway-id")
		directconnect_acceptDirectConnectGatewayAssociationProposalCmd.MarkFlagRequired("proposal-id")
	})
	directconnectCmd.AddCommand(directconnect_acceptDirectConnectGatewayAssociationProposalCmd)
}

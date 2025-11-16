package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeDirectConnectGatewayAssociationProposalsCmd = &cobra.Command{
	Use:   "describe-direct-connect-gateway-association-proposals",
	Short: "Describes one or more association proposals for connection between a virtual private gateway or transit gateway and a Direct Connect gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeDirectConnectGatewayAssociationProposalsCmd).Standalone()

	directconnect_describeDirectConnectGatewayAssociationProposalsCmd.Flags().String("associated-gateway-id", "", "The ID of the associated gateway.")
	directconnect_describeDirectConnectGatewayAssociationProposalsCmd.Flags().String("direct-connect-gateway-id", "", "The ID of the Direct Connect gateway.")
	directconnect_describeDirectConnectGatewayAssociationProposalsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	directconnect_describeDirectConnectGatewayAssociationProposalsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	directconnect_describeDirectConnectGatewayAssociationProposalsCmd.Flags().String("proposal-id", "", "The ID of the proposal.")
	directconnectCmd.AddCommand(directconnect_describeDirectConnectGatewayAssociationProposalsCmd)
}

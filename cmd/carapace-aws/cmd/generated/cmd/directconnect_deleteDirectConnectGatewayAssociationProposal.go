package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_deleteDirectConnectGatewayAssociationProposalCmd = &cobra.Command{
	Use:   "delete-direct-connect-gateway-association-proposal",
	Short: "Deletes the association proposal request between the specified Direct Connect gateway and virtual private gateway or transit gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_deleteDirectConnectGatewayAssociationProposalCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_deleteDirectConnectGatewayAssociationProposalCmd).Standalone()

		directconnect_deleteDirectConnectGatewayAssociationProposalCmd.Flags().String("proposal-id", "", "The ID of the proposal.")
		directconnect_deleteDirectConnectGatewayAssociationProposalCmd.MarkFlagRequired("proposal-id")
	})
	directconnectCmd.AddCommand(directconnect_deleteDirectConnectGatewayAssociationProposalCmd)
}

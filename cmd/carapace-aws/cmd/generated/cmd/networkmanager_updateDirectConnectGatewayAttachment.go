package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_updateDirectConnectGatewayAttachmentCmd = &cobra.Command{
	Use:   "update-direct-connect-gateway-attachment",
	Short: "Updates the edge locations associated with an Amazon Web Services Direct Connect gateway attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_updateDirectConnectGatewayAttachmentCmd).Standalone()

	networkmanager_updateDirectConnectGatewayAttachmentCmd.Flags().String("attachment-id", "", "The ID of the Direct Connect gateway attachment for the updated edge locations.")
	networkmanager_updateDirectConnectGatewayAttachmentCmd.Flags().String("edge-locations", "", "One or more edge locations to update for the Direct Connect gateway attachment.")
	networkmanager_updateDirectConnectGatewayAttachmentCmd.MarkFlagRequired("attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_updateDirectConnectGatewayAttachmentCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getDirectConnectGatewayAttachmentCmd = &cobra.Command{
	Use:   "get-direct-connect-gateway-attachment",
	Short: "Returns information about a specific Amazon Web Services Direct Connect gateway attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getDirectConnectGatewayAttachmentCmd).Standalone()

	networkmanager_getDirectConnectGatewayAttachmentCmd.Flags().String("attachment-id", "", "The ID of the Direct Connect gateway attachment that you want to see details about.")
	networkmanager_getDirectConnectGatewayAttachmentCmd.MarkFlagRequired("attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_getDirectConnectGatewayAttachmentCmd)
}

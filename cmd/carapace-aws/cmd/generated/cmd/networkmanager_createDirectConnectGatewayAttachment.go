package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createDirectConnectGatewayAttachmentCmd = &cobra.Command{
	Use:   "create-direct-connect-gateway-attachment",
	Short: "Creates an Amazon Web Services Direct Connect gateway attachment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createDirectConnectGatewayAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createDirectConnectGatewayAttachmentCmd).Standalone()

		networkmanager_createDirectConnectGatewayAttachmentCmd.Flags().String("client-token", "", "client token")
		networkmanager_createDirectConnectGatewayAttachmentCmd.Flags().String("core-network-id", "", "The ID of the Cloud WAN core network that the Direct Connect gateway attachment should be attached to.")
		networkmanager_createDirectConnectGatewayAttachmentCmd.Flags().String("direct-connect-gateway-arn", "", "The ARN of the Direct Connect gateway attachment.")
		networkmanager_createDirectConnectGatewayAttachmentCmd.Flags().String("edge-locations", "", "One or more core network edge locations that the Direct Connect gateway attachment is associated with.")
		networkmanager_createDirectConnectGatewayAttachmentCmd.Flags().String("tags", "", "The key value tags to apply to the Direct Connect gateway attachment during creation.")
		networkmanager_createDirectConnectGatewayAttachmentCmd.MarkFlagRequired("core-network-id")
		networkmanager_createDirectConnectGatewayAttachmentCmd.MarkFlagRequired("direct-connect-gateway-arn")
		networkmanager_createDirectConnectGatewayAttachmentCmd.MarkFlagRequired("edge-locations")
	})
	networkmanagerCmd.AddCommand(networkmanager_createDirectConnectGatewayAttachmentCmd)
}

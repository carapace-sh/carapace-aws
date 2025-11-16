package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createConnectAttachmentCmd = &cobra.Command{
	Use:   "create-connect-attachment",
	Short: "Creates a core network Connect attachment from a specified core network attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createConnectAttachmentCmd).Standalone()

	networkmanager_createConnectAttachmentCmd.Flags().String("client-token", "", "The client token associated with the request.")
	networkmanager_createConnectAttachmentCmd.Flags().String("core-network-id", "", "The ID of a core network where you want to create the attachment.")
	networkmanager_createConnectAttachmentCmd.Flags().String("edge-location", "", "The Region where the edge is located.")
	networkmanager_createConnectAttachmentCmd.Flags().String("options", "", "Options for creating an attachment.")
	networkmanager_createConnectAttachmentCmd.Flags().String("tags", "", "The list of key-value tags associated with the request.")
	networkmanager_createConnectAttachmentCmd.Flags().String("transport-attachment-id", "", "The ID of the attachment between the two connections.")
	networkmanager_createConnectAttachmentCmd.MarkFlagRequired("core-network-id")
	networkmanager_createConnectAttachmentCmd.MarkFlagRequired("edge-location")
	networkmanager_createConnectAttachmentCmd.MarkFlagRequired("options")
	networkmanager_createConnectAttachmentCmd.MarkFlagRequired("transport-attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_createConnectAttachmentCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_getConnectAttachmentCmd = &cobra.Command{
	Use:   "get-connect-attachment",
	Short: "Returns information about a core network Connect attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_getConnectAttachmentCmd).Standalone()

	networkmanager_getConnectAttachmentCmd.Flags().String("attachment-id", "", "The ID of the attachment.")
	networkmanager_getConnectAttachmentCmd.MarkFlagRequired("attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_getConnectAttachmentCmd)
}

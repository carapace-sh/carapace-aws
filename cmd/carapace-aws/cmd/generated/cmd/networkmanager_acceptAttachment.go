package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_acceptAttachmentCmd = &cobra.Command{
	Use:   "accept-attachment",
	Short: "Accepts a core network attachment request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_acceptAttachmentCmd).Standalone()

	networkmanager_acceptAttachmentCmd.Flags().String("attachment-id", "", "The ID of the attachment.")
	networkmanager_acceptAttachmentCmd.MarkFlagRequired("attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_acceptAttachmentCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_rejectAttachmentCmd = &cobra.Command{
	Use:   "reject-attachment",
	Short: "Rejects a core network attachment request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_rejectAttachmentCmd).Standalone()

	networkmanager_rejectAttachmentCmd.Flags().String("attachment-id", "", "The ID of the attachment.")
	networkmanager_rejectAttachmentCmd.MarkFlagRequired("attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_rejectAttachmentCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_deleteAttachmentCmd = &cobra.Command{
	Use:   "delete-attachment",
	Short: "Deletes an attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_deleteAttachmentCmd).Standalone()

	networkmanager_deleteAttachmentCmd.Flags().String("attachment-id", "", "The ID of the attachment to delete.")
	networkmanager_deleteAttachmentCmd.MarkFlagRequired("attachment-id")
	networkmanagerCmd.AddCommand(networkmanager_deleteAttachmentCmd)
}

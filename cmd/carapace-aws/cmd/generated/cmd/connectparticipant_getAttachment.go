package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_getAttachmentCmd = &cobra.Command{
	Use:   "get-attachment",
	Short: "Provides a pre-signed URL for download of a completed attachment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_getAttachmentCmd).Standalone()

	connectparticipant_getAttachmentCmd.Flags().String("attachment-id", "", "A unique identifier for the attachment.")
	connectparticipant_getAttachmentCmd.Flags().String("connection-token", "", "The authentication token associated with the participant's connection.")
	connectparticipant_getAttachmentCmd.Flags().String("url-expiry-in-seconds", "", "The expiration time of the URL in ISO timestamp.")
	connectparticipant_getAttachmentCmd.MarkFlagRequired("attachment-id")
	connectparticipant_getAttachmentCmd.MarkFlagRequired("connection-token")
	connectparticipantCmd.AddCommand(connectparticipant_getAttachmentCmd)
}

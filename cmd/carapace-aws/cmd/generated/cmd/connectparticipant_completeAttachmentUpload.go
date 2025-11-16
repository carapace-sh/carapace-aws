package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_completeAttachmentUploadCmd = &cobra.Command{
	Use:   "complete-attachment-upload",
	Short: "Allows you to confirm that the attachment has been uploaded using the pre-signed URL provided in StartAttachmentUpload API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_completeAttachmentUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectparticipant_completeAttachmentUploadCmd).Standalone()

		connectparticipant_completeAttachmentUploadCmd.Flags().String("attachment-ids", "", "A list of unique identifiers for the attachments.")
		connectparticipant_completeAttachmentUploadCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connectparticipant_completeAttachmentUploadCmd.Flags().String("connection-token", "", "The authentication token associated with the participant's connection.")
		connectparticipant_completeAttachmentUploadCmd.MarkFlagRequired("attachment-ids")
		connectparticipant_completeAttachmentUploadCmd.MarkFlagRequired("client-token")
		connectparticipant_completeAttachmentUploadCmd.MarkFlagRequired("connection-token")
	})
	connectparticipantCmd.AddCommand(connectparticipant_completeAttachmentUploadCmd)
}

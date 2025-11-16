package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectparticipant_startAttachmentUploadCmd = &cobra.Command{
	Use:   "start-attachment-upload",
	Short: "Provides a pre-signed Amazon S3 URL in response for uploading the file directly to S3.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectparticipant_startAttachmentUploadCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectparticipant_startAttachmentUploadCmd).Standalone()

		connectparticipant_startAttachmentUploadCmd.Flags().String("attachment-name", "", "A case-sensitive name of the attachment being uploaded.")
		connectparticipant_startAttachmentUploadCmd.Flags().String("attachment-size-in-bytes", "", "The size of the attachment in bytes.")
		connectparticipant_startAttachmentUploadCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		connectparticipant_startAttachmentUploadCmd.Flags().String("connection-token", "", "The authentication token associated with the participant's connection.")
		connectparticipant_startAttachmentUploadCmd.Flags().String("content-type", "", "Describes the MIME file type of the attachment.")
		connectparticipant_startAttachmentUploadCmd.MarkFlagRequired("attachment-name")
		connectparticipant_startAttachmentUploadCmd.MarkFlagRequired("attachment-size-in-bytes")
		connectparticipant_startAttachmentUploadCmd.MarkFlagRequired("client-token")
		connectparticipant_startAttachmentUploadCmd.MarkFlagRequired("connection-token")
		connectparticipant_startAttachmentUploadCmd.MarkFlagRequired("content-type")
	})
	connectparticipantCmd.AddCommand(connectparticipant_startAttachmentUploadCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_getCaseAttachmentUploadUrlCmd = &cobra.Command{
	Use:   "get-case-attachment-upload-url",
	Short: "Uploads an attachment to a case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_getCaseAttachmentUploadUrlCmd).Standalone()

	securityIr_getCaseAttachmentUploadUrlCmd.Flags().String("case-id", "", "Required element for GetCaseAttachmentUploadUrl to identify the case ID for uploading an attachment.")
	securityIr_getCaseAttachmentUploadUrlCmd.Flags().String("client-token", "", "The `clientToken` field is an idempotency key used to ensure that repeated attempts for a single action will be ignored by the server during retries.")
	securityIr_getCaseAttachmentUploadUrlCmd.Flags().String("content-length", "", "Required element for GetCaseAttachmentUploadUrl to identify the size of the file attachment.")
	securityIr_getCaseAttachmentUploadUrlCmd.Flags().String("file-name", "", "Required element for GetCaseAttachmentUploadUrl to identify the file name of the attachment to upload.")
	securityIr_getCaseAttachmentUploadUrlCmd.MarkFlagRequired("case-id")
	securityIr_getCaseAttachmentUploadUrlCmd.MarkFlagRequired("content-length")
	securityIr_getCaseAttachmentUploadUrlCmd.MarkFlagRequired("file-name")
	securityIrCmd.AddCommand(securityIr_getCaseAttachmentUploadUrlCmd)
}

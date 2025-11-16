package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityIr_getCaseAttachmentDownloadUrlCmd = &cobra.Command{
	Use:   "get-case-attachment-download-url",
	Short: "Returns a Pre-Signed URL for uploading attachments into a case.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityIr_getCaseAttachmentDownloadUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityIr_getCaseAttachmentDownloadUrlCmd).Standalone()

		securityIr_getCaseAttachmentDownloadUrlCmd.Flags().String("attachment-id", "", "Required element for GetCaseAttachmentDownloadUrl to identify the attachment ID for downloading an attachment.")
		securityIr_getCaseAttachmentDownloadUrlCmd.Flags().String("case-id", "", "Required element for GetCaseAttachmentDownloadUrl to identify the case ID for downloading an attachment from.")
		securityIr_getCaseAttachmentDownloadUrlCmd.MarkFlagRequired("attachment-id")
		securityIr_getCaseAttachmentDownloadUrlCmd.MarkFlagRequired("case-id")
	})
	securityIrCmd.AddCommand(securityIr_getCaseAttachmentDownloadUrlCmd)
}

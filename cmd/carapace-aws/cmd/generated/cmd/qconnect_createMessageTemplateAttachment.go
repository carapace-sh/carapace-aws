package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createMessageTemplateAttachmentCmd = &cobra.Command{
	Use:   "create-message-template-attachment",
	Short: "Uploads an attachment file to the specified Amazon Q in Connect message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createMessageTemplateAttachmentCmd).Standalone()

	qconnect_createMessageTemplateAttachmentCmd.Flags().String("body", "", "The body of the attachment file being uploaded.")
	qconnect_createMessageTemplateAttachmentCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_createMessageTemplateAttachmentCmd.Flags().String("content-disposition", "", "The presentation information for the attachment file.")
	qconnect_createMessageTemplateAttachmentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_createMessageTemplateAttachmentCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
	qconnect_createMessageTemplateAttachmentCmd.Flags().String("name", "", "The name of the attachment file being uploaded.")
	qconnect_createMessageTemplateAttachmentCmd.MarkFlagRequired("body")
	qconnect_createMessageTemplateAttachmentCmd.MarkFlagRequired("content-disposition")
	qconnect_createMessageTemplateAttachmentCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_createMessageTemplateAttachmentCmd.MarkFlagRequired("message-template-id")
	qconnect_createMessageTemplateAttachmentCmd.MarkFlagRequired("name")
	qconnectCmd.AddCommand(qconnect_createMessageTemplateAttachmentCmd)
}

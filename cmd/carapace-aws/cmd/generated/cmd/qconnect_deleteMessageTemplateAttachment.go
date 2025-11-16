package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteMessageTemplateAttachmentCmd = &cobra.Command{
	Use:   "delete-message-template-attachment",
	Short: "Deletes the attachment file from the Amazon Q in Connect message template that is referenced by `$LATEST` qualifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteMessageTemplateAttachmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_deleteMessageTemplateAttachmentCmd).Standalone()

		qconnect_deleteMessageTemplateAttachmentCmd.Flags().String("attachment-id", "", "The identifier of the attachment file.")
		qconnect_deleteMessageTemplateAttachmentCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_deleteMessageTemplateAttachmentCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
		qconnect_deleteMessageTemplateAttachmentCmd.MarkFlagRequired("attachment-id")
		qconnect_deleteMessageTemplateAttachmentCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_deleteMessageTemplateAttachmentCmd.MarkFlagRequired("message-template-id")
	})
	qconnectCmd.AddCommand(qconnect_deleteMessageTemplateAttachmentCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createMessageTemplateVersionCmd = &cobra.Command{
	Use:   "create-message-template-version",
	Short: "Creates a new Amazon Q in Connect message template version from the current content and configuration of a message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createMessageTemplateVersionCmd).Standalone()

	qconnect_createMessageTemplateVersionCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_createMessageTemplateVersionCmd.Flags().String("message-template-content-sha256", "", "The checksum value of the message template content that is referenced by the `$LATEST` qualifier.")
	qconnect_createMessageTemplateVersionCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
	qconnect_createMessageTemplateVersionCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_createMessageTemplateVersionCmd.MarkFlagRequired("message-template-id")
	qconnectCmd.AddCommand(qconnect_createMessageTemplateVersionCmd)
}

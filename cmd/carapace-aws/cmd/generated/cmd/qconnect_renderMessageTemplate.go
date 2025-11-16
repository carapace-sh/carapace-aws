package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_renderMessageTemplateCmd = &cobra.Command{
	Use:   "render-message-template",
	Short: "Renders the Amazon Q in Connect message template based on the attribute values provided and generates the message content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_renderMessageTemplateCmd).Standalone()

	qconnect_renderMessageTemplateCmd.Flags().String("attributes", "", "An object that specifies the values to use for variables in the message template.")
	qconnect_renderMessageTemplateCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_renderMessageTemplateCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
	qconnect_renderMessageTemplateCmd.MarkFlagRequired("attributes")
	qconnect_renderMessageTemplateCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_renderMessageTemplateCmd.MarkFlagRequired("message-template-id")
	qconnectCmd.AddCommand(qconnect_renderMessageTemplateCmd)
}

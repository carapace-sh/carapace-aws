package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_activateMessageTemplateCmd = &cobra.Command{
	Use:   "activate-message-template",
	Short: "Activates a specific version of the Amazon Q in Connect message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_activateMessageTemplateCmd).Standalone()

	qconnect_activateMessageTemplateCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_activateMessageTemplateCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
	qconnect_activateMessageTemplateCmd.Flags().String("version-number", "", "The version number of the message template version to activate.")
	qconnect_activateMessageTemplateCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_activateMessageTemplateCmd.MarkFlagRequired("message-template-id")
	qconnect_activateMessageTemplateCmd.MarkFlagRequired("version-number")
	qconnectCmd.AddCommand(qconnect_activateMessageTemplateCmd)
}

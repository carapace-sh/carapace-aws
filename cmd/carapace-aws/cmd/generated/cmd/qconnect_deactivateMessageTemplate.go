package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deactivateMessageTemplateCmd = &cobra.Command{
	Use:   "deactivate-message-template",
	Short: "Deactivates a specific version of the Amazon Q in Connect message template .",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deactivateMessageTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_deactivateMessageTemplateCmd).Standalone()

		qconnect_deactivateMessageTemplateCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_deactivateMessageTemplateCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
		qconnect_deactivateMessageTemplateCmd.Flags().String("version-number", "", "The version number of the message template version to deactivate.")
		qconnect_deactivateMessageTemplateCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_deactivateMessageTemplateCmd.MarkFlagRequired("message-template-id")
		qconnect_deactivateMessageTemplateCmd.MarkFlagRequired("version-number")
	})
	qconnectCmd.AddCommand(qconnect_deactivateMessageTemplateCmd)
}

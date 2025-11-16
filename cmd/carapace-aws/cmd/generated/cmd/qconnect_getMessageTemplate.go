package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_getMessageTemplateCmd = &cobra.Command{
	Use:   "get-message-template",
	Short: "Retrieves the Amazon Q in Connect message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_getMessageTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_getMessageTemplateCmd).Standalone()

		qconnect_getMessageTemplateCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_getMessageTemplateCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
		qconnect_getMessageTemplateCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_getMessageTemplateCmd.MarkFlagRequired("message-template-id")
	})
	qconnectCmd.AddCommand(qconnect_getMessageTemplateCmd)
}

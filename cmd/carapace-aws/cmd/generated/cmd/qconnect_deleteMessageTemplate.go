package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_deleteMessageTemplateCmd = &cobra.Command{
	Use:   "delete-message-template",
	Short: "Deletes an Amazon Q in Connect message template entirely or a specific version of the message template if version is supplied in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_deleteMessageTemplateCmd).Standalone()

	qconnect_deleteMessageTemplateCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_deleteMessageTemplateCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
	qconnect_deleteMessageTemplateCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_deleteMessageTemplateCmd.MarkFlagRequired("message-template-id")
	qconnectCmd.AddCommand(qconnect_deleteMessageTemplateCmd)
}

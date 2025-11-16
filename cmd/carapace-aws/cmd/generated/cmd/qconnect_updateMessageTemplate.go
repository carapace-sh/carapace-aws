package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_updateMessageTemplateCmd = &cobra.Command{
	Use:   "update-message-template",
	Short: "Updates the Amazon Q in Connect message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_updateMessageTemplateCmd).Standalone()

	qconnect_updateMessageTemplateCmd.Flags().String("content", "", "The content of the message template.")
	qconnect_updateMessageTemplateCmd.Flags().String("default-attributes", "", "An object that specifies the default values to use for variables in the message template.")
	qconnect_updateMessageTemplateCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_updateMessageTemplateCmd.Flags().String("language", "", "The language code value for the language in which the quick response is written.")
	qconnect_updateMessageTemplateCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
	qconnect_updateMessageTemplateCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_updateMessageTemplateCmd.MarkFlagRequired("message-template-id")
	qconnectCmd.AddCommand(qconnect_updateMessageTemplateCmd)
}

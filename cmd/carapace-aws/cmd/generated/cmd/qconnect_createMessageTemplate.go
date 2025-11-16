package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_createMessageTemplateCmd = &cobra.Command{
	Use:   "create-message-template",
	Short: "Creates an Amazon Q in Connect message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_createMessageTemplateCmd).Standalone()

	qconnect_createMessageTemplateCmd.Flags().String("channel-subtype", "", "The channel subtype this message template applies to.")
	qconnect_createMessageTemplateCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	qconnect_createMessageTemplateCmd.Flags().String("content", "", "The content of the message template.")
	qconnect_createMessageTemplateCmd.Flags().String("default-attributes", "", "An object that specifies the default values to use for variables in the message template.")
	qconnect_createMessageTemplateCmd.Flags().String("description", "", "The description of the message template.")
	qconnect_createMessageTemplateCmd.Flags().String("grouping-configuration", "", "")
	qconnect_createMessageTemplateCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_createMessageTemplateCmd.Flags().String("language", "", "The language code value for the language in which the quick response is written.")
	qconnect_createMessageTemplateCmd.Flags().String("name", "", "The name of the message template.")
	qconnect_createMessageTemplateCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	qconnect_createMessageTemplateCmd.MarkFlagRequired("channel-subtype")
	qconnect_createMessageTemplateCmd.MarkFlagRequired("content")
	qconnect_createMessageTemplateCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_createMessageTemplateCmd.MarkFlagRequired("name")
	qconnectCmd.AddCommand(qconnect_createMessageTemplateCmd)
}

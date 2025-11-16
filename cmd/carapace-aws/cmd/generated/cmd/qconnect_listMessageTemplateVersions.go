package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_listMessageTemplateVersionsCmd = &cobra.Command{
	Use:   "list-message-template-versions",
	Short: "Lists all the available versions for the specified Amazon Q in Connect message template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_listMessageTemplateVersionsCmd).Standalone()

	qconnect_listMessageTemplateVersionsCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_listMessageTemplateVersionsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_listMessageTemplateVersionsCmd.Flags().String("message-template-id", "", "The identifier of the message template.")
	qconnect_listMessageTemplateVersionsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_listMessageTemplateVersionsCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_listMessageTemplateVersionsCmd.MarkFlagRequired("message-template-id")
	qconnectCmd.AddCommand(qconnect_listMessageTemplateVersionsCmd)
}

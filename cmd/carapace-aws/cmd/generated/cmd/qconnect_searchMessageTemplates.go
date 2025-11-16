package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_searchMessageTemplatesCmd = &cobra.Command{
	Use:   "search-message-templates",
	Short: "Searches for Amazon Q in Connect message templates in the specified knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_searchMessageTemplatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qconnect_searchMessageTemplatesCmd).Standalone()

		qconnect_searchMessageTemplatesCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
		qconnect_searchMessageTemplatesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		qconnect_searchMessageTemplatesCmd.Flags().String("next-token", "", "The token for the next set of results.")
		qconnect_searchMessageTemplatesCmd.Flags().String("search-expression", "", "The search expression for querying the message template.")
		qconnect_searchMessageTemplatesCmd.MarkFlagRequired("knowledge-base-id")
		qconnect_searchMessageTemplatesCmd.MarkFlagRequired("search-expression")
	})
	qconnectCmd.AddCommand(qconnect_searchMessageTemplatesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qconnect_searchQuickResponsesCmd = &cobra.Command{
	Use:   "search-quick-responses",
	Short: "Searches existing Amazon Q in Connect quick responses in an Amazon Q in Connect knowledge base.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qconnect_searchQuickResponsesCmd).Standalone()

	qconnect_searchQuickResponsesCmd.Flags().String("attributes", "", "The [user-defined Amazon Connect contact attributes](https://docs.aws.amazon.com/connect/latest/adminguide/connect-attrib-list.html#user-defined-attributes) to be resolved when search results are returned.")
	qconnect_searchQuickResponsesCmd.Flags().String("knowledge-base-id", "", "The identifier of the knowledge base.")
	qconnect_searchQuickResponsesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	qconnect_searchQuickResponsesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	qconnect_searchQuickResponsesCmd.Flags().String("search-expression", "", "The search expression for querying the quick response.")
	qconnect_searchQuickResponsesCmd.MarkFlagRequired("knowledge-base-id")
	qconnect_searchQuickResponsesCmd.MarkFlagRequired("search-expression")
	qconnectCmd.AddCommand(qconnect_searchQuickResponsesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_searchRelevantContentCmd = &cobra.Command{
	Use:   "search-relevant-content",
	Short: "Searches for relevant content in a Amazon Q Business application based on a query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_searchRelevantContentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(qbusiness_searchRelevantContentCmd).Standalone()

		qbusiness_searchRelevantContentCmd.Flags().String("application-id", "", "The unique identifier of the Amazon Q Business application to search.")
		qbusiness_searchRelevantContentCmd.Flags().String("attribute-filter", "", "")
		qbusiness_searchRelevantContentCmd.Flags().String("content-source", "", "The source of content to search in.")
		qbusiness_searchRelevantContentCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		qbusiness_searchRelevantContentCmd.Flags().String("next-token", "", "The token for the next set of results.")
		qbusiness_searchRelevantContentCmd.Flags().String("query-text", "", "The text to search for.")
		qbusiness_searchRelevantContentCmd.MarkFlagRequired("application-id")
		qbusiness_searchRelevantContentCmd.MarkFlagRequired("content-source")
		qbusiness_searchRelevantContentCmd.MarkFlagRequired("query-text")
	})
	qbusinessCmd.AddCommand(qbusiness_searchRelevantContentCmd)
}

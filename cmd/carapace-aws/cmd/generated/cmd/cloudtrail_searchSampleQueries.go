package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_searchSampleQueriesCmd = &cobra.Command{
	Use:   "search-sample-queries",
	Short: "Searches sample queries and returns a list of sample queries that are sorted by relevance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_searchSampleQueriesCmd).Standalone()

	cloudtrail_searchSampleQueriesCmd.Flags().String("max-results", "", "The maximum number of results to return on a single page.")
	cloudtrail_searchSampleQueriesCmd.Flags().String("next-token", "", "A token you can use to get the next page of results.")
	cloudtrail_searchSampleQueriesCmd.Flags().String("search-phrase", "", "The natural language phrase to use for the semantic search.")
	cloudtrail_searchSampleQueriesCmd.MarkFlagRequired("search-phrase")
	cloudtrailCmd.AddCommand(cloudtrail_searchSampleQueriesCmd)
}

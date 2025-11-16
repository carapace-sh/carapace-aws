package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Searches an index given an input query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_queryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_queryCmd).Standalone()

		kendra_queryCmd.Flags().String("attribute-filter", "", "Filters search results by document fields/attributes.")
		kendra_queryCmd.Flags().String("collapse-configuration", "", "Provides configuration to determine how to group results by document attribute value, and how to display them (collapsed or expanded) under a designated primary document for each group.")
		kendra_queryCmd.Flags().String("document-relevance-override-configurations", "", "Overrides relevance tuning configurations of fields/attributes set at the index level.")
		kendra_queryCmd.Flags().String("facets", "", "An array of documents fields/attributes for faceted search.")
		kendra_queryCmd.Flags().String("index-id", "", "The identifier of the index for the search.")
		kendra_queryCmd.Flags().String("page-number", "", "Query results are returned in pages the size of the `PageSize` parameter.")
		kendra_queryCmd.Flags().String("page-size", "", "Sets the number of results that are returned in each page of results.")
		kendra_queryCmd.Flags().String("query-result-type-filter", "", "Sets the type of query result or response.")
		kendra_queryCmd.Flags().String("query-text", "", "The input query text for the search.")
		kendra_queryCmd.Flags().String("requested-document-attributes", "", "An array of document fields/attributes to include in the response.")
		kendra_queryCmd.Flags().String("sorting-configuration", "", "Provides information that determines how the results of the query are sorted.")
		kendra_queryCmd.Flags().String("sorting-configurations", "", "Provides configuration information to determine how the results of a query are sorted.")
		kendra_queryCmd.Flags().String("spell-correction-configuration", "", "Enables suggested spell corrections for queries.")
		kendra_queryCmd.Flags().String("user-context", "", "The user context token or user and group information.")
		kendra_queryCmd.Flags().String("visitor-id", "", "Provides an identifier for a specific user.")
		kendra_queryCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_queryCmd)
}

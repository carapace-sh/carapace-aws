package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudsearchdomain_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Retrieves a list of documents that match the specified search criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudsearchdomain_searchCmd).Standalone()

	cloudsearchdomain_searchCmd.Flags().String("cursor", "", "Retrieves a cursor value you can use to page through large result sets.")
	cloudsearchdomain_searchCmd.Flags().String("expr", "", "Defines one or more numeric expressions that can be used to sort results or specify search or filter criteria.")
	cloudsearchdomain_searchCmd.Flags().String("facet", "", "Specifies one or more fields for which to get facet information, and options that control how the facet information is returned.")
	cloudsearchdomain_searchCmd.Flags().String("filter-query", "", "Specifies a structured query that filters the results of a search without affecting how the results are scored and sorted.")
	cloudsearchdomain_searchCmd.Flags().String("highlight", "", "Retrieves highlights for matches in the specified `text` or `text-array` fields.")
	cloudsearchdomain_searchCmd.Flags().String("partial", "", "Enables partial results to be returned if one or more index partitions are unavailable.")
	cloudsearchdomain_searchCmd.Flags().String("query", "", "Specifies the search criteria for the request.")
	cloudsearchdomain_searchCmd.Flags().String("query-options", "", "Configures options for the query parser specified in the `queryParser` parameter.")
	cloudsearchdomain_searchCmd.Flags().String("query-parser", "", "Specifies which query parser to use to process the request.")
	cloudsearchdomain_searchCmd.Flags().String("return", "", "Specifies the field and expression values to include in the response.")
	cloudsearchdomain_searchCmd.Flags().String("size", "", "Specifies the maximum number of search hits to include in the response.")
	cloudsearchdomain_searchCmd.Flags().String("sort", "", "Specifies the fields or custom expressions to use to sort the search results.")
	cloudsearchdomain_searchCmd.Flags().String("start", "", "Specifies the offset of the first search hit you want to return.")
	cloudsearchdomain_searchCmd.Flags().String("stats", "", "Specifies one or more fields for which to get statistics information.")
	cloudsearchdomain_searchCmd.MarkFlagRequired("query")
	cloudsearchdomainCmd.AddCommand(cloudsearchdomain_searchCmd)
}

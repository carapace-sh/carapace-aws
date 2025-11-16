package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_searchListingsCmd = &cobra.Command{
	Use:   "search-listings",
	Short: "Searches listings in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_searchListingsCmd).Standalone()

	datazone_searchListingsCmd.Flags().String("additional-attributes", "", "Specifies additional attributes for the search.")
	datazone_searchListingsCmd.Flags().String("aggregations", "", "Enables you to specify one or more attributes to compute and return counts grouped by field values.")
	datazone_searchListingsCmd.Flags().String("domain-identifier", "", "The identifier of the domain in which to search listings.")
	datazone_searchListingsCmd.Flags().String("filters", "", "Specifies the filters for the search of listings.")
	datazone_searchListingsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call to `SearchListings`.")
	datazone_searchListingsCmd.Flags().String("next-token", "", "When the number of results is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of results, the response includes a pagination token named `NextToken`.")
	datazone_searchListingsCmd.Flags().String("search-in", "", "The details of the search.")
	datazone_searchListingsCmd.Flags().String("search-text", "", "Specifies the text for which to search.")
	datazone_searchListingsCmd.Flags().String("sort", "", "Specifies the way for sorting the search results.")
	datazone_searchListingsCmd.MarkFlagRequired("domain-identifier")
	datazoneCmd.AddCommand(datazone_searchListingsCmd)
}

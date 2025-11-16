package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Searches for assets in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_searchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_searchCmd).Standalone()

		datazone_searchCmd.Flags().String("additional-attributes", "", "Specifies additional attributes for the `Search` action.")
		datazone_searchCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
		datazone_searchCmd.Flags().String("filters", "", "Specifies the search filters.")
		datazone_searchCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call to `Search`.")
		datazone_searchCmd.Flags().String("next-token", "", "When the number of results is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of results, the response includes a pagination token named `NextToken`.")
		datazone_searchCmd.Flags().String("owning-project-identifier", "", "The identifier of the owning project specified for the search.")
		datazone_searchCmd.Flags().String("search-in", "", "The details of the search.")
		datazone_searchCmd.Flags().String("search-scope", "", "The scope of the search.")
		datazone_searchCmd.Flags().String("search-text", "", "Specifies the text for which to search.")
		datazone_searchCmd.Flags().String("sort", "", "Specifies the way in which the search results are to be sorted.")
		datazone_searchCmd.MarkFlagRequired("domain-identifier")
		datazone_searchCmd.MarkFlagRequired("search-scope")
	})
	datazoneCmd.AddCommand(datazone_searchCmd)
}

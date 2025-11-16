package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_searchTypesCmd = &cobra.Command{
	Use:   "search-types",
	Short: "Searches for types in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_searchTypesCmd).Standalone()

	datazone_searchTypesCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain in which to invoke the `SearchTypes` action.")
	datazone_searchTypesCmd.Flags().String("filters", "", "The filters for the `SearchTypes` action.")
	datazone_searchTypesCmd.Flags().Bool("managed", false, "Specifies whether the search is managed.")
	datazone_searchTypesCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call to `SearchTypes`.")
	datazone_searchTypesCmd.Flags().String("next-token", "", "When the number of results is greater than the default value for the `MaxResults` parameter, or if you explicitly specify a value for `MaxResults` that is less than the number of results, the response includes a pagination token named `NextToken`.")
	datazone_searchTypesCmd.Flags().Bool("no-managed", false, "Specifies whether the search is managed.")
	datazone_searchTypesCmd.Flags().String("search-in", "", "The details of the search.")
	datazone_searchTypesCmd.Flags().String("search-scope", "", "Specifies the scope of the search for types.")
	datazone_searchTypesCmd.Flags().String("search-text", "", "Specifies the text for which to search.")
	datazone_searchTypesCmd.Flags().String("sort", "", "The specifies the way to sort the `SearchTypes` results.")
	datazone_searchTypesCmd.MarkFlagRequired("domain-identifier")
	datazone_searchTypesCmd.MarkFlagRequired("managed")
	datazone_searchTypesCmd.Flag("no-managed").Hidden = true
	datazone_searchTypesCmd.MarkFlagRequired("no-managed")
	datazone_searchTypesCmd.MarkFlagRequired("search-scope")
	datazoneCmd.AddCommand(datazone_searchTypesCmd)
}

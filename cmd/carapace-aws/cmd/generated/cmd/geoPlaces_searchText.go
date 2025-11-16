package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoPlaces_searchTextCmd = &cobra.Command{
	Use:   "search-text",
	Short: "`SearchText` searches for geocode and place information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoPlaces_searchTextCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoPlaces_searchTextCmd).Standalone()

		geoPlaces_searchTextCmd.Flags().String("additional-features", "", "A list of optional additional parameters, such as time zone, that can be requested for each result.")
		geoPlaces_searchTextCmd.Flags().String("bias-position", "", "The position, in longitude and latitude, that the results should be close to.")
		geoPlaces_searchTextCmd.Flags().String("filter", "", "A structure which contains a set of inclusion/exclusion properties that results must possess in order to be returned as a result.")
		geoPlaces_searchTextCmd.Flags().String("intended-use", "", "Indicates if the results will be stored.")
		geoPlaces_searchTextCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
		geoPlaces_searchTextCmd.Flags().String("language", "", "A list of [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language codes for the results to be rendered in.")
		geoPlaces_searchTextCmd.Flags().String("max-results", "", "An optional limit for the number of results returned in a single call.")
		geoPlaces_searchTextCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
		geoPlaces_searchTextCmd.Flags().String("political-view", "", "The alpha-2 or alpha-3 character code for the political view of a country.")
		geoPlaces_searchTextCmd.Flags().String("query-id", "", "The query Id returned by the suggest API.")
		geoPlaces_searchTextCmd.Flags().String("query-text", "", "The free-form text query to match addresses against.")
	})
	geoPlacesCmd.AddCommand(geoPlaces_searchTextCmd)
}

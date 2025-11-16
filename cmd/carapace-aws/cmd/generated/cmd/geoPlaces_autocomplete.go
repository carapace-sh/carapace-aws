package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoPlaces_autocompleteCmd = &cobra.Command{
	Use:   "autocomplete",
	Short: "`Autocomplete` completes potential places and addresses as the user types, based on the partial input.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoPlaces_autocompleteCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoPlaces_autocompleteCmd).Standalone()

		geoPlaces_autocompleteCmd.Flags().String("additional-features", "", "A list of optional additional parameters that can be requested for each result.")
		geoPlaces_autocompleteCmd.Flags().String("bias-position", "", "The position in longitude and latitude that the results should be close to.")
		geoPlaces_autocompleteCmd.Flags().String("filter", "", "A structure which contains a set of inclusion/exclusion properties that results must possess in order to be returned as a result.")
		geoPlaces_autocompleteCmd.Flags().String("intended-use", "", "Indicates if the results will be stored.")
		geoPlaces_autocompleteCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
		geoPlaces_autocompleteCmd.Flags().String("language", "", "A list of [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language codes for the results to be rendered in.")
		geoPlaces_autocompleteCmd.Flags().String("max-results", "", "An optional limit for the number of results returned in a single call.")
		geoPlaces_autocompleteCmd.Flags().String("political-view", "", "The alpha-2 or alpha-3 character code for the political view of a country.")
		geoPlaces_autocompleteCmd.Flags().String("postal-code-mode", "", "The `PostalCodeMode` affects how postal code results are returned.")
		geoPlaces_autocompleteCmd.Flags().String("query-text", "", "The free-form text query to match addresses against.")
		geoPlaces_autocompleteCmd.MarkFlagRequired("query-text")
	})
	geoPlacesCmd.AddCommand(geoPlaces_autocompleteCmd)
}

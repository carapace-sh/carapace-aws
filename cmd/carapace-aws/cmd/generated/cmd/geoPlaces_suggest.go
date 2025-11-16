package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoPlaces_suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "`Suggest` provides intelligent predictions or recommendations based on the user's input or context, such as relevant places, points of interest, query terms or search category.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoPlaces_suggestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoPlaces_suggestCmd).Standalone()

		geoPlaces_suggestCmd.Flags().String("additional-features", "", "A list of optional additional parameters, such as time zone, that can be requested for each result.")
		geoPlaces_suggestCmd.Flags().String("bias-position", "", "The position, in longitude and latitude, that the results should be close to.")
		geoPlaces_suggestCmd.Flags().String("filter", "", "A structure which contains a set of inclusion/exclusion properties that results must possess in order to be returned as a result.")
		geoPlaces_suggestCmd.Flags().String("intended-use", "", "Indicates if the results will be stored.")
		geoPlaces_suggestCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
		geoPlaces_suggestCmd.Flags().String("language", "", "A list of [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language codes for the results to be rendered in.")
		geoPlaces_suggestCmd.Flags().String("max-query-refinements", "", "Maximum number of query terms to be returned for use with a search text query.")
		geoPlaces_suggestCmd.Flags().String("max-results", "", "An optional limit for the number of results returned in a single call.")
		geoPlaces_suggestCmd.Flags().String("political-view", "", "The alpha-2 or alpha-3 character code for the political view of a country.")
		geoPlaces_suggestCmd.Flags().String("query-text", "", "The free-form text query to match addresses against.")
		geoPlaces_suggestCmd.MarkFlagRequired("query-text")
	})
	geoPlacesCmd.AddCommand(geoPlaces_suggestCmd)
}

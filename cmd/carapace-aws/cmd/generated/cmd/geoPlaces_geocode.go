package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoPlaces_geocodeCmd = &cobra.Command{
	Use:   "geocode",
	Short: "`Geocode` converts a textual address or place into geographic coordinates.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoPlaces_geocodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoPlaces_geocodeCmd).Standalone()

		geoPlaces_geocodeCmd.Flags().String("additional-features", "", "A list of optional additional parameters, such as time zone, that can be requested for each result.")
		geoPlaces_geocodeCmd.Flags().String("bias-position", "", "The position, in longitude and latitude, that the results should be close to.")
		geoPlaces_geocodeCmd.Flags().String("filter", "", "A structure which contains a set of inclusion/exclusion properties that results must possess in order to be returned as a result.")
		geoPlaces_geocodeCmd.Flags().String("intended-use", "", "Indicates if the results will be stored.")
		geoPlaces_geocodeCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
		geoPlaces_geocodeCmd.Flags().String("language", "", "A list of [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language codes for the results to be rendered in.")
		geoPlaces_geocodeCmd.Flags().String("max-results", "", "An optional limit for the number of results returned in a single call.")
		geoPlaces_geocodeCmd.Flags().String("political-view", "", "The alpha-2 or alpha-3 character code for the political view of a country.")
		geoPlaces_geocodeCmd.Flags().String("query-components", "", "")
		geoPlaces_geocodeCmd.Flags().String("query-text", "", "The free-form text query to match addresses against.")
	})
	geoPlacesCmd.AddCommand(geoPlaces_geocodeCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoPlaces_reverseGeocodeCmd = &cobra.Command{
	Use:   "reverse-geocode",
	Short: "`ReverseGeocode` converts geographic coordinates into a human-readable address or place.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoPlaces_reverseGeocodeCmd).Standalone()

	geoPlaces_reverseGeocodeCmd.Flags().String("additional-features", "", "A list of optional additional parameters, such as time zone that can be requested for each result.")
	geoPlaces_reverseGeocodeCmd.Flags().String("filter", "", "A structure which contains a set of inclusion/exclusion properties that results must possess in order to be returned as a result.")
	geoPlaces_reverseGeocodeCmd.Flags().String("intended-use", "", "Indicates if the results will be stored.")
	geoPlaces_reverseGeocodeCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
	geoPlaces_reverseGeocodeCmd.Flags().String("language", "", "A list of [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language codes for the results to be rendered in.")
	geoPlaces_reverseGeocodeCmd.Flags().String("max-results", "", "An optional limit for the number of results returned in a single call.")
	geoPlaces_reverseGeocodeCmd.Flags().String("political-view", "", "The alpha-2 or alpha-3 character code for the political view of a country.")
	geoPlaces_reverseGeocodeCmd.Flags().String("query-position", "", "The position, in `[lng, lat]` for which you are querying nearby results for.")
	geoPlaces_reverseGeocodeCmd.Flags().String("query-radius", "", "The maximum distance in meters from the QueryPosition from which a result will be returned.")
	geoPlaces_reverseGeocodeCmd.MarkFlagRequired("query-position")
	geoPlacesCmd.AddCommand(geoPlaces_reverseGeocodeCmd)
}

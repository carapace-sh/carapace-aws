package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoPlaces_searchNearbyCmd = &cobra.Command{
	Use:   "search-nearby",
	Short: "`SearchNearby` queries for points of interest within a radius from a central coordinates, returning place results with optional filters such as categories, business chains, food types and more.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoPlaces_searchNearbyCmd).Standalone()

	geoPlaces_searchNearbyCmd.Flags().String("additional-features", "", "A list of optional additional parameters, such as time zone, that can be requested for each result.")
	geoPlaces_searchNearbyCmd.Flags().String("filter", "", "A structure which contains a set of inclusion/exclusion properties that results must possess in order to be returned as a result.")
	geoPlaces_searchNearbyCmd.Flags().String("intended-use", "", "Indicates if the results will be stored.")
	geoPlaces_searchNearbyCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
	geoPlaces_searchNearbyCmd.Flags().String("language", "", "A list of [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language codes for the results to be rendered in.")
	geoPlaces_searchNearbyCmd.Flags().String("max-results", "", "An optional limit for the number of results returned in a single call.")
	geoPlaces_searchNearbyCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	geoPlaces_searchNearbyCmd.Flags().String("political-view", "", "The alpha-2 or alpha-3 character code for the political view of a country.")
	geoPlaces_searchNearbyCmd.Flags().String("query-position", "", "The position, in `[lng, lat]` for which you are querying nearby results for.")
	geoPlaces_searchNearbyCmd.Flags().String("query-radius", "", "The maximum distance in meters from the QueryPosition from which a result will be returned.")
	geoPlaces_searchNearbyCmd.MarkFlagRequired("query-position")
	geoPlacesCmd.AddCommand(geoPlaces_searchNearbyCmd)
}

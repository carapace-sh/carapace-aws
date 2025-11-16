package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoPlaces_getPlaceCmd = &cobra.Command{
	Use:   "get-place",
	Short: "`GetPlace` finds a place by its unique ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoPlaces_getPlaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(geoPlaces_getPlaceCmd).Standalone()

		geoPlaces_getPlaceCmd.Flags().String("additional-features", "", "A list of optional additional parameters such as time zone that can be requested for each result.")
		geoPlaces_getPlaceCmd.Flags().String("intended-use", "", "Indicates if the results will be stored.")
		geoPlaces_getPlaceCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
		geoPlaces_getPlaceCmd.Flags().String("language", "", "A list of [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language codes for the results to be rendered in.")
		geoPlaces_getPlaceCmd.Flags().String("place-id", "", "The `PlaceId` of the place you wish to receive the information for.")
		geoPlaces_getPlaceCmd.Flags().String("political-view", "", "The alpha-2 or alpha-3 character code for the political view of a country.")
		geoPlaces_getPlaceCmd.MarkFlagRequired("place-id")
	})
	geoPlacesCmd.AddCommand(geoPlaces_getPlaceCmd)
}

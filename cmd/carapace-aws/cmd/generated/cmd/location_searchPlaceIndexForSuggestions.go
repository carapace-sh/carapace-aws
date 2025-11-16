package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_searchPlaceIndexForSuggestionsCmd = &cobra.Command{
	Use:   "search-place-index-for-suggestions",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_searchPlaceIndexForSuggestionsCmd).Standalone()

	location_searchPlaceIndexForSuggestionsCmd.Flags().String("bias-position", "", "An optional parameter that indicates a preference for place suggestions that are closer to a specified position.")
	location_searchPlaceIndexForSuggestionsCmd.Flags().String("filter-bbox", "", "An optional parameter that limits the search results by returning only suggestions within a specified bounding box.")
	location_searchPlaceIndexForSuggestionsCmd.Flags().String("filter-categories", "", "A list of one or more Amazon Location categories to filter the returned places.")
	location_searchPlaceIndexForSuggestionsCmd.Flags().String("filter-countries", "", "An optional parameter that limits the search results by returning only suggestions within the provided list of countries.")
	location_searchPlaceIndexForSuggestionsCmd.Flags().String("index-name", "", "The name of the place index resource you want to use for the search.")
	location_searchPlaceIndexForSuggestionsCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
	location_searchPlaceIndexForSuggestionsCmd.Flags().String("language", "", "The preferred language used to return results.")
	location_searchPlaceIndexForSuggestionsCmd.Flags().String("max-results", "", "An optional parameter.")
	location_searchPlaceIndexForSuggestionsCmd.Flags().String("text", "", "The free-form partial text to use to generate place suggestions.")
	location_searchPlaceIndexForSuggestionsCmd.MarkFlagRequired("index-name")
	location_searchPlaceIndexForSuggestionsCmd.MarkFlagRequired("text")
	locationCmd.AddCommand(location_searchPlaceIndexForSuggestionsCmd)
}

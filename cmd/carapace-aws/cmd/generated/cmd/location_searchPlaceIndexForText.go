package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_searchPlaceIndexForTextCmd = &cobra.Command{
	Use:   "search-place-index-for-text",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_searchPlaceIndexForTextCmd).Standalone()

	location_searchPlaceIndexForTextCmd.Flags().String("bias-position", "", "An optional parameter that indicates a preference for places that are closer to a specified position.")
	location_searchPlaceIndexForTextCmd.Flags().String("filter-bbox", "", "An optional parameter that limits the search results by returning only places that are within the provided bounding box.")
	location_searchPlaceIndexForTextCmd.Flags().String("filter-categories", "", "A list of one or more Amazon Location categories to filter the returned places.")
	location_searchPlaceIndexForTextCmd.Flags().String("filter-countries", "", "An optional parameter that limits the search results by returning only places that are in a specified list of countries.")
	location_searchPlaceIndexForTextCmd.Flags().String("index-name", "", "The name of the place index resource you want to use for the search.")
	location_searchPlaceIndexForTextCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
	location_searchPlaceIndexForTextCmd.Flags().String("language", "", "The preferred language used to return results.")
	location_searchPlaceIndexForTextCmd.Flags().String("max-results", "", "An optional parameter.")
	location_searchPlaceIndexForTextCmd.Flags().String("text", "", "The address, name, city, or region to be used in the search in free-form text format.")
	location_searchPlaceIndexForTextCmd.MarkFlagRequired("index-name")
	location_searchPlaceIndexForTextCmd.MarkFlagRequired("text")
	locationCmd.AddCommand(location_searchPlaceIndexForTextCmd)
}

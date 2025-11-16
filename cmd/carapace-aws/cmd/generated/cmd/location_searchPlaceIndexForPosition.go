package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_searchPlaceIndexForPositionCmd = &cobra.Command{
	Use:   "search-place-index-for-position",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_searchPlaceIndexForPositionCmd).Standalone()

	location_searchPlaceIndexForPositionCmd.Flags().String("index-name", "", "The name of the place index resource you want to use for the search.")
	location_searchPlaceIndexForPositionCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
	location_searchPlaceIndexForPositionCmd.Flags().String("language", "", "The preferred language used to return results.")
	location_searchPlaceIndexForPositionCmd.Flags().String("max-results", "", "An optional parameter.")
	location_searchPlaceIndexForPositionCmd.Flags().String("position", "", "Specifies the longitude and latitude of the position to query.")
	location_searchPlaceIndexForPositionCmd.MarkFlagRequired("index-name")
	location_searchPlaceIndexForPositionCmd.MarkFlagRequired("position")
	locationCmd.AddCommand(location_searchPlaceIndexForPositionCmd)
}

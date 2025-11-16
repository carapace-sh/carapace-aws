package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_getPlaceCmd = &cobra.Command{
	Use:   "get-place",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_getPlaceCmd).Standalone()

	location_getPlaceCmd.Flags().String("index-name", "", "The name of the place index resource that you want to use for the search.")
	location_getPlaceCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
	location_getPlaceCmd.Flags().String("language", "", "The preferred language used to return results.")
	location_getPlaceCmd.Flags().String("place-id", "", "The identifier of the place to find.")
	location_getPlaceCmd.MarkFlagRequired("index-name")
	location_getPlaceCmd.MarkFlagRequired("place-id")
	locationCmd.AddCommand(location_getPlaceCmd)
}

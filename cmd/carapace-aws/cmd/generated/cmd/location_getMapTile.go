package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_getMapTileCmd = &cobra.Command{
	Use:   "get-map-tile",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_getMapTileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_getMapTileCmd).Standalone()

		location_getMapTileCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
		location_getMapTileCmd.Flags().String("map-name", "", "The map resource to retrieve the map tiles from.")
		location_getMapTileCmd.Flags().String("x", "", "The X axis value for the map tile.")
		location_getMapTileCmd.Flags().String("y", "", "The Y axis value for the map tile.")
		location_getMapTileCmd.Flags().String("z", "", "The zoom value for the map tile.")
		location_getMapTileCmd.MarkFlagRequired("map-name")
		location_getMapTileCmd.MarkFlagRequired("x")
		location_getMapTileCmd.MarkFlagRequired("y")
		location_getMapTileCmd.MarkFlagRequired("z")
	})
	locationCmd.AddCommand(location_getMapTileCmd)
}

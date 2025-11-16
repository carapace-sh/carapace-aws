package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoMaps_getTileCmd = &cobra.Command{
	Use:   "get-tile",
	Short: "`GetTile` returns a tile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoMaps_getTileCmd).Standalone()

	geoMaps_getTileCmd.Flags().String("additional-features", "", "A list of optional additional parameters such as map styles that can be requested for each result.")
	geoMaps_getTileCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
	geoMaps_getTileCmd.Flags().String("tileset", "", "Specifies the desired tile set.")
	geoMaps_getTileCmd.Flags().String("x", "", "The X axis value for the map tile.")
	geoMaps_getTileCmd.Flags().String("y", "", "The Y axis value for the map tile.")
	geoMaps_getTileCmd.Flags().String("z", "", "The zoom value for the map tile.")
	geoMaps_getTileCmd.MarkFlagRequired("tileset")
	geoMaps_getTileCmd.MarkFlagRequired("x")
	geoMaps_getTileCmd.MarkFlagRequired("y")
	geoMaps_getTileCmd.MarkFlagRequired("z")
	geoMapsCmd.AddCommand(geoMaps_getTileCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoMaps_getSpritesCmd = &cobra.Command{
	Use:   "get-sprites",
	Short: "`GetSprites` returns the map's sprites.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoMaps_getSpritesCmd).Standalone()

	geoMaps_getSpritesCmd.Flags().String("color-scheme", "", "Sets color tone for map such as dark and light for specific map styles.")
	geoMaps_getSpritesCmd.Flags().String("file-name", "", "`Sprites` API: The name of the sprite ﬁle to retrieve, following pattern `sprites(@2x)?\\.(png|json)`.")
	geoMaps_getSpritesCmd.Flags().String("style", "", "Style specifies the desired map style for the `Sprites` APIs.")
	geoMaps_getSpritesCmd.Flags().String("variant", "", "Optimizes map styles for specific use case or industry.")
	geoMaps_getSpritesCmd.MarkFlagRequired("color-scheme")
	geoMaps_getSpritesCmd.MarkFlagRequired("file-name")
	geoMaps_getSpritesCmd.MarkFlagRequired("style")
	geoMaps_getSpritesCmd.MarkFlagRequired("variant")
	geoMapsCmd.AddCommand(geoMaps_getSpritesCmd)
}

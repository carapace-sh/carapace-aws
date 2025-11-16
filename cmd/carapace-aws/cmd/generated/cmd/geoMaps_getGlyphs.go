package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoMaps_getGlyphsCmd = &cobra.Command{
	Use:   "get-glyphs",
	Short: "`GetGlyphs` returns the map's glyphs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoMaps_getGlyphsCmd).Standalone()

	geoMaps_getGlyphsCmd.Flags().String("font-stack", "", "Name of the `FontStack` to retrieve.")
	geoMaps_getGlyphsCmd.Flags().String("font-unicode-range", "", "A Unicode range of characters to download glyphs for.")
	geoMaps_getGlyphsCmd.MarkFlagRequired("font-stack")
	geoMaps_getGlyphsCmd.MarkFlagRequired("font-unicode-range")
	geoMapsCmd.AddCommand(geoMaps_getGlyphsCmd)
}

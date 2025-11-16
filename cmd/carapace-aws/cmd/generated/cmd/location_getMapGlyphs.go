package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_getMapGlyphsCmd = &cobra.Command{
	Use:   "get-map-glyphs",
	Short: "This operation is no longer current and may be deprecated in the future.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_getMapGlyphsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(location_getMapGlyphsCmd).Standalone()

		location_getMapGlyphsCmd.Flags().String("font-stack", "", "A comma-separated list of fonts to load glyphs from in order of preference.")
		location_getMapGlyphsCmd.Flags().String("font-unicode-range", "", "A Unicode range of characters to download glyphs for.")
		location_getMapGlyphsCmd.Flags().String("key", "", "The optional [API key](https://docs.aws.amazon.com/location/previous/developerguide/using-apikeys.html) to authorize the request.")
		location_getMapGlyphsCmd.Flags().String("map-name", "", "The map resource associated with the glyph ﬁle.")
		location_getMapGlyphsCmd.MarkFlagRequired("font-stack")
		location_getMapGlyphsCmd.MarkFlagRequired("font-unicode-range")
		location_getMapGlyphsCmd.MarkFlagRequired("map-name")
	})
	locationCmd.AddCommand(location_getMapGlyphsCmd)
}

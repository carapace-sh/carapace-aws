package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoMaps_getStaticMapCmd = &cobra.Command{
	Use:   "get-static-map",
	Short: "`GetStaticMap` provides high-quality static map images with customizable options.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoMaps_getStaticMapCmd).Standalone()

	geoMaps_getStaticMapCmd.Flags().String("bounded-positions", "", "Takes in two or more pair of coordinates in World Geodetic System (WGS 84) format: \\[longitude, latitude], with each coordinate separated by a comma.")
	geoMaps_getStaticMapCmd.Flags().String("bounding-box", "", "Takes in two pairs of coordinates in World Geodetic System (WGS 84) format: \\[longitude, latitude], denoting south-westerly and north-easterly edges of the image.")
	geoMaps_getStaticMapCmd.Flags().String("center", "", "Takes in a pair of coordinates in World Geodetic System (WGS 84) format: \\[longitude, latitude], which becomes the center point of the image.")
	geoMaps_getStaticMapCmd.Flags().String("color-scheme", "", "Sets color tone for map, such as dark and light for specific map styles.")
	geoMaps_getStaticMapCmd.Flags().String("compact-overlay", "", "Takes in a string to draw geometries on the image.")
	geoMaps_getStaticMapCmd.Flags().Bool("crop-labels", false, "It is a flag that takes in true or false.")
	geoMaps_getStaticMapCmd.Flags().String("file-name", "", "The map scaling parameter to size the image, icons, and labels.")
	geoMaps_getStaticMapCmd.Flags().String("geo-json-overlay", "", "Takes in a string to draw geometries on the image.")
	geoMaps_getStaticMapCmd.Flags().String("height", "", "Specifies the height of the map image.")
	geoMaps_getStaticMapCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
	geoMaps_getStaticMapCmd.Flags().String("label-size", "", "Overrides the label size auto-calculated by `FileName`.")
	geoMaps_getStaticMapCmd.Flags().String("language", "", "Specifies the language on the map labels using the BCP 47 language tag, limited to ISO 639-1 two-letter language codes.")
	geoMaps_getStaticMapCmd.Flags().Bool("no-crop-labels", false, "It is a flag that takes in true or false.")
	geoMaps_getStaticMapCmd.Flags().String("padding", "", "Applies additional space (in pixels) around overlay feature to prevent them from being cut or obscured.")
	geoMaps_getStaticMapCmd.Flags().String("points-of-interests", "", "Determines if the result image will display icons representing points of interest on the map.")
	geoMaps_getStaticMapCmd.Flags().String("political-view", "", "Specifies the political view, using ISO 3166-2 or ISO 3166-3 country code format.")
	geoMaps_getStaticMapCmd.Flags().String("radius", "", "Used with center parameter, it specifies the zoom of the image where you can control it on a granular level.")
	geoMaps_getStaticMapCmd.Flags().String("scale-bar-unit", "", "Displays a scale on the bottom right of the map image with the unit specified in the input.")
	geoMaps_getStaticMapCmd.Flags().String("style", "", "`Style` specifies the desired map style.")
	geoMaps_getStaticMapCmd.Flags().String("width", "", "Specifies the width of the map image.")
	geoMaps_getStaticMapCmd.Flags().String("zoom", "", "Specifies the zoom level of the map image.")
	geoMaps_getStaticMapCmd.MarkFlagRequired("file-name")
	geoMaps_getStaticMapCmd.MarkFlagRequired("height")
	geoMaps_getStaticMapCmd.Flag("no-crop-labels").Hidden = true
	geoMaps_getStaticMapCmd.MarkFlagRequired("width")
	geoMapsCmd.AddCommand(geoMaps_getStaticMapCmd)
}

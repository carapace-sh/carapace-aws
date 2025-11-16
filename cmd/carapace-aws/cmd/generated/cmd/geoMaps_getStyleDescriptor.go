package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var geoMaps_getStyleDescriptorCmd = &cobra.Command{
	Use:   "get-style-descriptor",
	Short: "`GetStyleDescriptor` returns information about the style.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(geoMaps_getStyleDescriptorCmd).Standalone()

	geoMaps_getStyleDescriptorCmd.Flags().String("color-scheme", "", "Sets color tone for map such as dark and light for specific map styles.")
	geoMaps_getStyleDescriptorCmd.Flags().String("contour-density", "", "Displays the shape and steepness of terrain features using elevation lines.")
	geoMaps_getStyleDescriptorCmd.Flags().String("key", "", "Optional: The API key to be used for authorization.")
	geoMaps_getStyleDescriptorCmd.Flags().String("political-view", "", "Specifies the political view using ISO 3166-2 or ISO 3166-3 country code format.")
	geoMaps_getStyleDescriptorCmd.Flags().String("style", "", "Style specifies the desired map style.")
	geoMaps_getStyleDescriptorCmd.Flags().String("terrain", "", "Adjusts how physical terrain details are rendered on the map.")
	geoMaps_getStyleDescriptorCmd.Flags().String("traffic", "", "Displays real-time traffic information overlay on map, such as incident events and flow events.")
	geoMaps_getStyleDescriptorCmd.Flags().String("travel-modes", "", "Renders additional map information relevant to selected travel modes.")
	geoMaps_getStyleDescriptorCmd.MarkFlagRequired("style")
	geoMapsCmd.AddCommand(geoMaps_getStyleDescriptorCmd)
}

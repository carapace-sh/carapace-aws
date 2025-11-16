package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_getTileCmd = &cobra.Command{
	Use:   "get-tile",
	Short: "Gets a web mercator tile for the given Earth Observation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_getTileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_getTileCmd).Standalone()

		sagemakerGeospatial_getTileCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the tile operation.")
		sagemakerGeospatial_getTileCmd.Flags().String("execution-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that you specify.")
		sagemakerGeospatial_getTileCmd.Flags().String("image-assets", "", "The particular assets or bands to tile.")
		sagemakerGeospatial_getTileCmd.Flags().Bool("image-mask", false, "Determines whether or not to return a valid data mask.")
		sagemakerGeospatial_getTileCmd.Flags().Bool("no-image-mask", false, "Determines whether or not to return a valid data mask.")
		sagemakerGeospatial_getTileCmd.Flags().String("output-data-type", "", "The output data type of the tile operation.")
		sagemakerGeospatial_getTileCmd.Flags().String("output-format", "", "The data format of the output tile.")
		sagemakerGeospatial_getTileCmd.Flags().String("property-filters", "", "Property filters for the imagery to tile.")
		sagemakerGeospatial_getTileCmd.Flags().String("target", "", "Determines what part of the Earth Observation job to tile.")
		sagemakerGeospatial_getTileCmd.Flags().String("time-range-filter", "", "Time range filter applied to imagery to find the images to tile.")
		sagemakerGeospatial_getTileCmd.Flags().String("x", "", "The x coordinate of the tile input.")
		sagemakerGeospatial_getTileCmd.Flags().String("y", "", "The y coordinate of the tile input.")
		sagemakerGeospatial_getTileCmd.Flags().String("z", "", "The z coordinate of the tile input.")
		sagemakerGeospatial_getTileCmd.MarkFlagRequired("arn")
		sagemakerGeospatial_getTileCmd.MarkFlagRequired("image-assets")
		sagemakerGeospatial_getTileCmd.Flag("no-image-mask").Hidden = true
		sagemakerGeospatial_getTileCmd.MarkFlagRequired("target")
		sagemakerGeospatial_getTileCmd.MarkFlagRequired("x")
		sagemakerGeospatial_getTileCmd.MarkFlagRequired("y")
		sagemakerGeospatial_getTileCmd.MarkFlagRequired("z")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_getTileCmd)
}

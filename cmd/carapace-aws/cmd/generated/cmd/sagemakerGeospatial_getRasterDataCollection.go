package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_getRasterDataCollectionCmd = &cobra.Command{
	Use:   "get-raster-data-collection",
	Short: "Use this operation to get details of a specific raster data collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_getRasterDataCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_getRasterDataCollectionCmd).Standalone()

		sagemakerGeospatial_getRasterDataCollectionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the raster data collection.")
		sagemakerGeospatial_getRasterDataCollectionCmd.MarkFlagRequired("arn")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_getRasterDataCollectionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_searchRasterDataCollectionCmd = &cobra.Command{
	Use:   "search-raster-data-collection",
	Short: "Allows you run image query on a specific raster data collection to get a list of the satellite imagery matching the selected filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_searchRasterDataCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_searchRasterDataCollectionCmd).Standalone()

		sagemakerGeospatial_searchRasterDataCollectionCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the raster data collection.")
		sagemakerGeospatial_searchRasterDataCollectionCmd.Flags().String("next-token", "", "If the previous response was truncated, you receive this token.")
		sagemakerGeospatial_searchRasterDataCollectionCmd.Flags().String("raster-data-collection-query", "", "RasterDataCollectionQuery consisting of [AreaOfInterest(AOI)](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_geospatial_AreaOfInterest.html), [PropertyFilters](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_geospatial_PropertyFilter.html) and [TimeRangeFilterInput](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_geospatial_TimeRangeFilterInput.html) used in [SearchRasterDataCollection](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_geospatial_SearchRasterDataCollection.html).")
		sagemakerGeospatial_searchRasterDataCollectionCmd.MarkFlagRequired("arn")
		sagemakerGeospatial_searchRasterDataCollectionCmd.MarkFlagRequired("raster-data-collection-query")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_searchRasterDataCollectionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemakerGeospatial_listRasterDataCollectionsCmd = &cobra.Command{
	Use:   "list-raster-data-collections",
	Short: "Use this operation to get raster data collections.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemakerGeospatial_listRasterDataCollectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemakerGeospatial_listRasterDataCollectionsCmd).Standalone()

		sagemakerGeospatial_listRasterDataCollectionsCmd.Flags().String("max-results", "", "The total number of items to return.")
		sagemakerGeospatial_listRasterDataCollectionsCmd.Flags().String("next-token", "", "If the previous response was truncated, you receive this token.")
	})
	sagemakerGeospatialCmd.AddCommand(sagemakerGeospatial_listRasterDataCollectionsCmd)
}

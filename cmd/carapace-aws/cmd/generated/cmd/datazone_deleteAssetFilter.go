package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_deleteAssetFilterCmd = &cobra.Command{
	Use:   "delete-asset-filter",
	Short: "Deletes an asset filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_deleteAssetFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_deleteAssetFilterCmd).Standalone()

		datazone_deleteAssetFilterCmd.Flags().String("asset-identifier", "", "The ID of the data asset.")
		datazone_deleteAssetFilterCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to delete an asset filter.")
		datazone_deleteAssetFilterCmd.Flags().String("identifier", "", "The ID of the asset filter that you want to delete.")
		datazone_deleteAssetFilterCmd.MarkFlagRequired("asset-identifier")
		datazone_deleteAssetFilterCmd.MarkFlagRequired("domain-identifier")
		datazone_deleteAssetFilterCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_deleteAssetFilterCmd)
}

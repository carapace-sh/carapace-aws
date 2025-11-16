package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateAssetFilterCmd = &cobra.Command{
	Use:   "update-asset-filter",
	Short: "Updates an asset filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateAssetFilterCmd).Standalone()

	datazone_updateAssetFilterCmd.Flags().String("asset-identifier", "", "The ID of the data asset.")
	datazone_updateAssetFilterCmd.Flags().String("configuration", "", "The configuration of the asset filter.")
	datazone_updateAssetFilterCmd.Flags().String("description", "", "The description of the asset filter.")
	datazone_updateAssetFilterCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to update an asset filter.")
	datazone_updateAssetFilterCmd.Flags().String("identifier", "", "The ID of the asset filter.")
	datazone_updateAssetFilterCmd.Flags().String("name", "", "The name of the asset filter.")
	datazone_updateAssetFilterCmd.MarkFlagRequired("asset-identifier")
	datazone_updateAssetFilterCmd.MarkFlagRequired("domain-identifier")
	datazone_updateAssetFilterCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_updateAssetFilterCmd)
}

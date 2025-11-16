package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createAssetFilterCmd = &cobra.Command{
	Use:   "create-asset-filter",
	Short: "Creates a data asset filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createAssetFilterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createAssetFilterCmd).Standalone()

		datazone_createAssetFilterCmd.Flags().String("asset-identifier", "", "The ID of the data asset.")
		datazone_createAssetFilterCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createAssetFilterCmd.Flags().String("configuration", "", "The configuration of the asset filter.")
		datazone_createAssetFilterCmd.Flags().String("description", "", "The description of the asset filter.")
		datazone_createAssetFilterCmd.Flags().String("domain-identifier", "", "The ID of the domain in which you want to create an asset filter.")
		datazone_createAssetFilterCmd.Flags().String("name", "", "The name of the asset filter.")
		datazone_createAssetFilterCmd.MarkFlagRequired("asset-identifier")
		datazone_createAssetFilterCmd.MarkFlagRequired("configuration")
		datazone_createAssetFilterCmd.MarkFlagRequired("domain-identifier")
		datazone_createAssetFilterCmd.MarkFlagRequired("name")
	})
	datazoneCmd.AddCommand(datazone_createAssetFilterCmd)
}

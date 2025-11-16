package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_getAssetFilterCmd = &cobra.Command{
	Use:   "get-asset-filter",
	Short: "Gets an asset filter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_getAssetFilterCmd).Standalone()

	datazone_getAssetFilterCmd.Flags().String("asset-identifier", "", "The ID of the data asset.")
	datazone_getAssetFilterCmd.Flags().String("domain-identifier", "", "The ID of the domain where you want to get an asset filter.")
	datazone_getAssetFilterCmd.Flags().String("identifier", "", "The ID of the asset filter.")
	datazone_getAssetFilterCmd.MarkFlagRequired("asset-identifier")
	datazone_getAssetFilterCmd.MarkFlagRequired("domain-identifier")
	datazone_getAssetFilterCmd.MarkFlagRequired("identifier")
	datazoneCmd.AddCommand(datazone_getAssetFilterCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateAssetCmd = &cobra.Command{
	Use:   "update-asset",
	Short: "Updates an asset's name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateAssetCmd).Standalone()

	iotsitewise_updateAssetCmd.Flags().String("asset-description", "", "A description for the asset.")
	iotsitewise_updateAssetCmd.Flags().String("asset-external-id", "", "An external ID to assign to the asset.")
	iotsitewise_updateAssetCmd.Flags().String("asset-id", "", "The ID of the asset to update.")
	iotsitewise_updateAssetCmd.Flags().String("asset-name", "", "A friendly name for the asset.")
	iotsitewise_updateAssetCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_updateAssetCmd.MarkFlagRequired("asset-id")
	iotsitewise_updateAssetCmd.MarkFlagRequired("asset-name")
	iotsitewiseCmd.AddCommand(iotsitewise_updateAssetCmd)
}

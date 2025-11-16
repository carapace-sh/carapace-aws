package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createAssetCmd = &cobra.Command{
	Use:   "create-asset",
	Short: "Creates an asset from an existing asset model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createAssetCmd).Standalone()

	iotsitewise_createAssetCmd.Flags().String("asset-description", "", "A description for the asset.")
	iotsitewise_createAssetCmd.Flags().String("asset-external-id", "", "An external ID to assign to the asset.")
	iotsitewise_createAssetCmd.Flags().String("asset-id", "", "The ID to assign to the asset, if desired.")
	iotsitewise_createAssetCmd.Flags().String("asset-model-id", "", "The ID of the asset model from which to create the asset.")
	iotsitewise_createAssetCmd.Flags().String("asset-name", "", "A friendly name for the asset.")
	iotsitewise_createAssetCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_createAssetCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the asset.")
	iotsitewise_createAssetCmd.MarkFlagRequired("asset-model-id")
	iotsitewise_createAssetCmd.MarkFlagRequired("asset-name")
	iotsitewiseCmd.AddCommand(iotsitewise_createAssetCmd)
}

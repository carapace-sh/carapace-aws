package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createAssetModelCmd = &cobra.Command{
	Use:   "create-asset-model",
	Short: "Creates an asset model from specified property and hierarchy definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createAssetModelCmd).Standalone()

	iotsitewise_createAssetModelCmd.Flags().String("asset-model-composite-models", "", "The composite models that are part of this asset model.")
	iotsitewise_createAssetModelCmd.Flags().String("asset-model-description", "", "A description for the asset model.")
	iotsitewise_createAssetModelCmd.Flags().String("asset-model-external-id", "", "An external ID to assign to the asset model.")
	iotsitewise_createAssetModelCmd.Flags().String("asset-model-hierarchies", "", "The hierarchy definitions of the asset model.")
	iotsitewise_createAssetModelCmd.Flags().String("asset-model-id", "", "The ID to assign to the asset model, if desired.")
	iotsitewise_createAssetModelCmd.Flags().String("asset-model-name", "", "A unique name for the asset model.")
	iotsitewise_createAssetModelCmd.Flags().String("asset-model-properties", "", "The property definitions of the asset model.")
	iotsitewise_createAssetModelCmd.Flags().String("asset-model-type", "", "The type of asset model.")
	iotsitewise_createAssetModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_createAssetModelCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the asset model.")
	iotsitewise_createAssetModelCmd.MarkFlagRequired("asset-model-name")
	iotsitewiseCmd.AddCommand(iotsitewise_createAssetModelCmd)
}

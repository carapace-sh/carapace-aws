package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateAssetModelCmd = &cobra.Command{
	Use:   "update-asset-model",
	Short: "Updates an asset model and all of the assets that were created from the model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateAssetModelCmd).Standalone()

	iotsitewise_updateAssetModelCmd.Flags().String("asset-model-composite-models", "", "The composite models that are part of this asset model.")
	iotsitewise_updateAssetModelCmd.Flags().String("asset-model-description", "", "A description for the asset model.")
	iotsitewise_updateAssetModelCmd.Flags().String("asset-model-external-id", "", "An external ID to assign to the asset model.")
	iotsitewise_updateAssetModelCmd.Flags().String("asset-model-hierarchies", "", "The updated hierarchy definitions of the asset model.")
	iotsitewise_updateAssetModelCmd.Flags().String("asset-model-id", "", "The ID of the asset model to update.")
	iotsitewise_updateAssetModelCmd.Flags().String("asset-model-name", "", "A unique name for the asset model.")
	iotsitewise_updateAssetModelCmd.Flags().String("asset-model-properties", "", "The updated property definitions of the asset model.")
	iotsitewise_updateAssetModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_updateAssetModelCmd.Flags().String("if-match", "", "The expected current entity tag (ETag) for the asset model’s latest or active version (specified using `matchForVersionType`).")
	iotsitewise_updateAssetModelCmd.Flags().String("if-none-match", "", "Accepts **\\*** to reject the update request if an active version (specified using `matchForVersionType` as `ACTIVE`) already exists for the asset model.")
	iotsitewise_updateAssetModelCmd.Flags().String("match-for-version-type", "", "Specifies the asset model version type (`LATEST` or `ACTIVE`) used in conjunction with `If-Match` or `If-None-Match` headers to determine the target ETag for the update operation.")
	iotsitewise_updateAssetModelCmd.MarkFlagRequired("asset-model-id")
	iotsitewise_updateAssetModelCmd.MarkFlagRequired("asset-model-name")
	iotsitewiseCmd.AddCommand(iotsitewise_updateAssetModelCmd)
}

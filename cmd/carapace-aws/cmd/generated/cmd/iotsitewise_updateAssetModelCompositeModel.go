package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateAssetModelCompositeModelCmd = &cobra.Command{
	Use:   "update-asset-model-composite-model",
	Short: "Updates a composite model and all of the assets that were created from the model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateAssetModelCompositeModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_updateAssetModelCompositeModelCmd).Standalone()

		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-description", "", "A description for the composite model.")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-external-id", "", "An external ID to assign to the asset model.")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-id", "", "The ID of a composite model on this asset model.")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-name", "", "A unique name for the composite model.")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-properties", "", "The property definitions of the composite model.")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("asset-model-id", "", "The ID of the asset model, in UUID format.")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("if-match", "", "The expected current entity tag (ETag) for the asset model’s latest or active version (specified using `matchForVersionType`).")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("if-none-match", "", "Accepts **\\*** to reject the update request if an active version (specified using `matchForVersionType` as `ACTIVE`) already exists for the asset model.")
		iotsitewise_updateAssetModelCompositeModelCmd.Flags().String("match-for-version-type", "", "Specifies the asset model version type (`LATEST` or `ACTIVE`) used in conjunction with `If-Match` or `If-None-Match` headers to determine the target ETag for the update operation.")
		iotsitewise_updateAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-composite-model-id")
		iotsitewise_updateAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-composite-model-name")
		iotsitewise_updateAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_updateAssetModelCompositeModelCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createAssetModelCompositeModelCmd = &cobra.Command{
	Use:   "create-asset-model-composite-model",
	Short: "Creates a custom composite model from specified property and hierarchy definitions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createAssetModelCompositeModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_createAssetModelCompositeModelCmd).Standalone()

		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-description", "", "A description for the composite model.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-external-id", "", "An external ID to assign to the composite model.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-id", "", "The ID of the composite model.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-name", "", "A unique name for the composite model.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-properties", "", "The property definitions of the composite model.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-type", "", "The composite model type.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("asset-model-id", "", "The ID of the asset model this composite model is a part of.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("composed-asset-model-id", "", "The ID of a component model which is reused to create this composite model.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("if-match", "", "The expected current entity tag (ETag) for the asset model’s latest or active version (specified using `matchForVersionType`).")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("if-none-match", "", "Accepts **\\*** to reject the create request if an active version (specified using `matchForVersionType` as `ACTIVE`) already exists for the asset model.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("match-for-version-type", "", "Specifies the asset model version type (`LATEST` or `ACTIVE`) used in conjunction with `If-Match` or `If-None-Match` headers to determine the target ETag for the create operation.")
		iotsitewise_createAssetModelCompositeModelCmd.Flags().String("parent-asset-model-composite-model-id", "", "The ID of the parent composite model in this asset model relationship.")
		iotsitewise_createAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-composite-model-name")
		iotsitewise_createAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-composite-model-type")
		iotsitewise_createAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_createAssetModelCompositeModelCmd)
}

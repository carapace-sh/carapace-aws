package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteAssetModelCompositeModelCmd = &cobra.Command{
	Use:   "delete-asset-model-composite-model",
	Short: "Deletes a composite model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteAssetModelCompositeModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_deleteAssetModelCompositeModelCmd).Standalone()

		iotsitewise_deleteAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-id", "", "The ID of a composite model on this asset model.")
		iotsitewise_deleteAssetModelCompositeModelCmd.Flags().String("asset-model-id", "", "The ID of the asset model, in UUID format.")
		iotsitewise_deleteAssetModelCompositeModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_deleteAssetModelCompositeModelCmd.Flags().String("if-match", "", "The expected current entity tag (ETag) for the asset model’s latest or active version (specified using `matchForVersionType`).")
		iotsitewise_deleteAssetModelCompositeModelCmd.Flags().String("if-none-match", "", "Accepts **\\*** to reject the delete request if an active version (specified using `matchForVersionType` as `ACTIVE`) already exists for the asset model.")
		iotsitewise_deleteAssetModelCompositeModelCmd.Flags().String("match-for-version-type", "", "Specifies the asset model version type (`LATEST` or `ACTIVE`) used in conjunction with `If-Match` or `If-None-Match` headers to determine the target ETag for the delete operation.")
		iotsitewise_deleteAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-composite-model-id")
		iotsitewise_deleteAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_deleteAssetModelCompositeModelCmd)
}

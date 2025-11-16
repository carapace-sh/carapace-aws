package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteAssetModelCmd = &cobra.Command{
	Use:   "delete-asset-model",
	Short: "Deletes an asset model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteAssetModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_deleteAssetModelCmd).Standalone()

		iotsitewise_deleteAssetModelCmd.Flags().String("asset-model-id", "", "The ID of the asset model to delete.")
		iotsitewise_deleteAssetModelCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_deleteAssetModelCmd.Flags().String("if-match", "", "The expected current entity tag (ETag) for the asset model’s latest or active version (specified using `matchForVersionType`).")
		iotsitewise_deleteAssetModelCmd.Flags().String("if-none-match", "", "Accepts **\\*** to reject the delete request if an active version (specified using `matchForVersionType` as `ACTIVE`) already exists for the asset model.")
		iotsitewise_deleteAssetModelCmd.Flags().String("match-for-version-type", "", "Specifies the asset model version type (`LATEST` or `ACTIVE`) used in conjunction with `If-Match` or `If-None-Match` headers to determine the target ETag for the delete operation.")
		iotsitewise_deleteAssetModelCmd.MarkFlagRequired("asset-model-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_deleteAssetModelCmd)
}

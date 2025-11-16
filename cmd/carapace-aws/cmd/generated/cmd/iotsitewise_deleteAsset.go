package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteAssetCmd = &cobra.Command{
	Use:   "delete-asset",
	Short: "Deletes an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteAssetCmd).Standalone()

	iotsitewise_deleteAssetCmd.Flags().String("asset-id", "", "The ID of the asset to delete.")
	iotsitewise_deleteAssetCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_deleteAssetCmd.MarkFlagRequired("asset-id")
	iotsitewiseCmd.AddCommand(iotsitewise_deleteAssetCmd)
}

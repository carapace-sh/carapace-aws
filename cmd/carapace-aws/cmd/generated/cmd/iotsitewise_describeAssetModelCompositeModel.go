package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeAssetModelCompositeModelCmd = &cobra.Command{
	Use:   "describe-asset-model-composite-model",
	Short: "Retrieves information about an asset model composite model (also known as an asset model component).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeAssetModelCompositeModelCmd).Standalone()

	iotsitewise_describeAssetModelCompositeModelCmd.Flags().String("asset-model-composite-model-id", "", "The ID of a composite model on this asset model.")
	iotsitewise_describeAssetModelCompositeModelCmd.Flags().String("asset-model-id", "", "The ID of the asset model.")
	iotsitewise_describeAssetModelCompositeModelCmd.Flags().String("asset-model-version", "", "The version alias that specifies the latest or active version of the asset model.")
	iotsitewise_describeAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-composite-model-id")
	iotsitewise_describeAssetModelCompositeModelCmd.MarkFlagRequired("asset-model-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describeAssetModelCompositeModelCmd)
}

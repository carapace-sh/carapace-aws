package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeAssetCompositeModelCmd = &cobra.Command{
	Use:   "describe-asset-composite-model",
	Short: "Retrieves information about an asset composite model (also known as an asset component).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeAssetCompositeModelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeAssetCompositeModelCmd).Standalone()

		iotsitewise_describeAssetCompositeModelCmd.Flags().String("asset-composite-model-id", "", "The ID of a composite model on this asset.")
		iotsitewise_describeAssetCompositeModelCmd.Flags().String("asset-id", "", "The ID of the asset.")
		iotsitewise_describeAssetCompositeModelCmd.MarkFlagRequired("asset-composite-model-id")
		iotsitewise_describeAssetCompositeModelCmd.MarkFlagRequired("asset-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeAssetCompositeModelCmd)
}

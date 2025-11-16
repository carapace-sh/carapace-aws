package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeAssetModelCmd = &cobra.Command{
	Use:   "describe-asset-model",
	Short: "Retrieves information about an asset model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeAssetModelCmd).Standalone()

	iotsitewise_describeAssetModelCmd.Flags().String("asset-model-id", "", "The ID of the asset model.")
	iotsitewise_describeAssetModelCmd.Flags().String("asset-model-version", "", "The version alias that specifies the latest or active version of the asset model.")
	iotsitewise_describeAssetModelCmd.Flags().String("exclude-properties", "", "Whether or not to exclude asset model properties from the response.")
	iotsitewise_describeAssetModelCmd.MarkFlagRequired("asset-model-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describeAssetModelCmd)
}

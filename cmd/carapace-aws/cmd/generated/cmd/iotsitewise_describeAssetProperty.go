package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeAssetPropertyCmd = &cobra.Command{
	Use:   "describe-asset-property",
	Short: "Retrieves information about an asset property.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeAssetPropertyCmd).Standalone()

	iotsitewise_describeAssetPropertyCmd.Flags().String("asset-id", "", "The ID of the asset.")
	iotsitewise_describeAssetPropertyCmd.Flags().String("property-id", "", "The ID of the asset property.")
	iotsitewise_describeAssetPropertyCmd.MarkFlagRequired("asset-id")
	iotsitewise_describeAssetPropertyCmd.MarkFlagRequired("property-id")
	iotsitewiseCmd.AddCommand(iotsitewise_describeAssetPropertyCmd)
}

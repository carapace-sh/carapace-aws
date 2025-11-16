package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeAssetCmd = &cobra.Command{
	Use:   "describe-asset",
	Short: "Retrieves information about an asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeAssetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeAssetCmd).Standalone()

		iotsitewise_describeAssetCmd.Flags().String("asset-id", "", "The ID of the asset.")
		iotsitewise_describeAssetCmd.Flags().String("exclude-properties", "", "Whether or not to exclude asset properties from the response.")
		iotsitewise_describeAssetCmd.MarkFlagRequired("asset-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeAssetCmd)
}

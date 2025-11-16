package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_batchGetAssetPropertyValueCmd = &cobra.Command{
	Use:   "batch-get-asset-property-value",
	Short: "Gets the current value for one or more asset properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_batchGetAssetPropertyValueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_batchGetAssetPropertyValueCmd).Standalone()

		iotsitewise_batchGetAssetPropertyValueCmd.Flags().String("entries", "", "The list of asset property value entries for the batch get request.")
		iotsitewise_batchGetAssetPropertyValueCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_batchGetAssetPropertyValueCmd.MarkFlagRequired("entries")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_batchGetAssetPropertyValueCmd)
}

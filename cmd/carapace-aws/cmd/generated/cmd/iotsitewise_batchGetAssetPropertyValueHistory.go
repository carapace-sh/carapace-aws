package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_batchGetAssetPropertyValueHistoryCmd = &cobra.Command{
	Use:   "batch-get-asset-property-value-history",
	Short: "Gets the historical values for one or more asset properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_batchGetAssetPropertyValueHistoryCmd).Standalone()

	iotsitewise_batchGetAssetPropertyValueHistoryCmd.Flags().String("entries", "", "The list of asset property historical value entries for the batch get request.")
	iotsitewise_batchGetAssetPropertyValueHistoryCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_batchGetAssetPropertyValueHistoryCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewise_batchGetAssetPropertyValueHistoryCmd.MarkFlagRequired("entries")
	iotsitewiseCmd.AddCommand(iotsitewise_batchGetAssetPropertyValueHistoryCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_batchGetAssetPropertyAggregatesCmd = &cobra.Command{
	Use:   "batch-get-asset-property-aggregates",
	Short: "Gets aggregated values (for example, average, minimum, and maximum) for one or more asset properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_batchGetAssetPropertyAggregatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_batchGetAssetPropertyAggregatesCmd).Standalone()

		iotsitewise_batchGetAssetPropertyAggregatesCmd.Flags().String("entries", "", "The list of asset property aggregate entries for the batch get request.")
		iotsitewise_batchGetAssetPropertyAggregatesCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_batchGetAssetPropertyAggregatesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_batchGetAssetPropertyAggregatesCmd.MarkFlagRequired("entries")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_batchGetAssetPropertyAggregatesCmd)
}

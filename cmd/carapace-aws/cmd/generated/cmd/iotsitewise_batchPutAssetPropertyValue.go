package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_batchPutAssetPropertyValueCmd = &cobra.Command{
	Use:   "batch-put-asset-property-value",
	Short: "Sends a list of asset property values to IoT SiteWise.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_batchPutAssetPropertyValueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_batchPutAssetPropertyValueCmd).Standalone()

		iotsitewise_batchPutAssetPropertyValueCmd.Flags().String("enable-partial-entry-processing", "", "This setting enables partial ingestion at entry-level.")
		iotsitewise_batchPutAssetPropertyValueCmd.Flags().String("entries", "", "The list of asset property value entries for the batch put request.")
		iotsitewise_batchPutAssetPropertyValueCmd.MarkFlagRequired("entries")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_batchPutAssetPropertyValueCmd)
}

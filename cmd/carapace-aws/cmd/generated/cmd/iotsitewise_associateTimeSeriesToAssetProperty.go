package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_associateTimeSeriesToAssetPropertyCmd = &cobra.Command{
	Use:   "associate-time-series-to-asset-property",
	Short: "Associates a time series (data stream) with an asset property.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_associateTimeSeriesToAssetPropertyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_associateTimeSeriesToAssetPropertyCmd).Standalone()

		iotsitewise_associateTimeSeriesToAssetPropertyCmd.Flags().String("alias", "", "The alias that identifies the time series.")
		iotsitewise_associateTimeSeriesToAssetPropertyCmd.Flags().String("asset-id", "", "The ID of the asset in which the asset property was created.")
		iotsitewise_associateTimeSeriesToAssetPropertyCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_associateTimeSeriesToAssetPropertyCmd.Flags().String("property-id", "", "The ID of the asset property.")
		iotsitewise_associateTimeSeriesToAssetPropertyCmd.MarkFlagRequired("alias")
		iotsitewise_associateTimeSeriesToAssetPropertyCmd.MarkFlagRequired("asset-id")
		iotsitewise_associateTimeSeriesToAssetPropertyCmd.MarkFlagRequired("property-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_associateTimeSeriesToAssetPropertyCmd)
}

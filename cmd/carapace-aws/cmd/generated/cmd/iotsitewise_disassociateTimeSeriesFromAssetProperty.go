package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd = &cobra.Command{
	Use:   "disassociate-time-series-from-asset-property",
	Short: "Disassociates a time series (data stream) from an asset property.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd).Standalone()

		iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd.Flags().String("alias", "", "The alias that identifies the time series.")
		iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd.Flags().String("asset-id", "", "The ID of the asset in which the asset property was created.")
		iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd.Flags().String("property-id", "", "The ID of the asset property.")
		iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd.MarkFlagRequired("alias")
		iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd.MarkFlagRequired("asset-id")
		iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd.MarkFlagRequired("property-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_disassociateTimeSeriesFromAssetPropertyCmd)
}

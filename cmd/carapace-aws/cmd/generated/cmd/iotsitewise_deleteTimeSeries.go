package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteTimeSeriesCmd = &cobra.Command{
	Use:   "delete-time-series",
	Short: "Deletes a time series (data stream).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteTimeSeriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_deleteTimeSeriesCmd).Standalone()

		iotsitewise_deleteTimeSeriesCmd.Flags().String("alias", "", "The alias that identifies the time series.")
		iotsitewise_deleteTimeSeriesCmd.Flags().String("asset-id", "", "The ID of the asset in which the asset property was created.")
		iotsitewise_deleteTimeSeriesCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_deleteTimeSeriesCmd.Flags().String("property-id", "", "The ID of the asset property.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_deleteTimeSeriesCmd)
}

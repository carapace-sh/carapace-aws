package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_describeTimeSeriesCmd = &cobra.Command{
	Use:   "describe-time-series",
	Short: "Retrieves information about a time series (data stream).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_describeTimeSeriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_describeTimeSeriesCmd).Standalone()

		iotsitewise_describeTimeSeriesCmd.Flags().String("alias", "", "The alias that identifies the time series.")
		iotsitewise_describeTimeSeriesCmd.Flags().String("asset-id", "", "The ID of the asset in which the asset property was created.")
		iotsitewise_describeTimeSeriesCmd.Flags().String("property-id", "", "The ID of the asset property.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_describeTimeSeriesCmd)
}

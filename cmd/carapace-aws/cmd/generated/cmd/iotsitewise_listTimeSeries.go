package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listTimeSeriesCmd = &cobra.Command{
	Use:   "list-time-series",
	Short: "Retrieves a paginated list of time series (data streams).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listTimeSeriesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listTimeSeriesCmd).Standalone()

		iotsitewise_listTimeSeriesCmd.Flags().String("alias-prefix", "", "The alias prefix of the time series.")
		iotsitewise_listTimeSeriesCmd.Flags().String("asset-id", "", "The ID of the asset in which the asset property was created.")
		iotsitewise_listTimeSeriesCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listTimeSeriesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_listTimeSeriesCmd.Flags().String("time-series-type", "", "The type of the time series.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listTimeSeriesCmd)
}

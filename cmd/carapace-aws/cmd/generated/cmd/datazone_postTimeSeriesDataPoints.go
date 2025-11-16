package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_postTimeSeriesDataPointsCmd = &cobra.Command{
	Use:   "post-time-series-data-points",
	Short: "Posts time series data points to Amazon DataZone for the specified asset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_postTimeSeriesDataPointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_postTimeSeriesDataPointsCmd).Standalone()

		datazone_postTimeSeriesDataPointsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_postTimeSeriesDataPointsCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain in which you want to post time series data points.")
		datazone_postTimeSeriesDataPointsCmd.Flags().String("entity-identifier", "", "The ID of the asset for which you want to post time series data points.")
		datazone_postTimeSeriesDataPointsCmd.Flags().String("entity-type", "", "The type of the asset for which you want to post data points.")
		datazone_postTimeSeriesDataPointsCmd.Flags().String("forms", "", "The forms that contain the data points that you want to post.")
		datazone_postTimeSeriesDataPointsCmd.MarkFlagRequired("domain-identifier")
		datazone_postTimeSeriesDataPointsCmd.MarkFlagRequired("entity-identifier")
		datazone_postTimeSeriesDataPointsCmd.MarkFlagRequired("entity-type")
		datazone_postTimeSeriesDataPointsCmd.MarkFlagRequired("forms")
	})
	datazoneCmd.AddCommand(datazone_postTimeSeriesDataPointsCmd)
}

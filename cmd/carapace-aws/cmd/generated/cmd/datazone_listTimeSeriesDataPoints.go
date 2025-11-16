package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listTimeSeriesDataPointsCmd = &cobra.Command{
	Use:   "list-time-series-data-points",
	Short: "Lists time series data points.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listTimeSeriesDataPointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listTimeSeriesDataPointsCmd).Standalone()

		datazone_listTimeSeriesDataPointsCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain that houses the assets for which you want to list time series data points.")
		datazone_listTimeSeriesDataPointsCmd.Flags().String("ended-at", "", "The timestamp at which the data points that you wanted to list ended.")
		datazone_listTimeSeriesDataPointsCmd.Flags().String("entity-identifier", "", "The ID of the asset for which you want to list data points.")
		datazone_listTimeSeriesDataPointsCmd.Flags().String("entity-type", "", "The type of the asset for which you want to list data points.")
		datazone_listTimeSeriesDataPointsCmd.Flags().String("form-name", "", "The name of the time series data points form.")
		datazone_listTimeSeriesDataPointsCmd.Flags().String("max-results", "", "The maximum number of data points to return in a single call to ListTimeSeriesDataPoints.")
		datazone_listTimeSeriesDataPointsCmd.Flags().String("next-token", "", "When the number of data points is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of data points, the response includes a pagination token named NextToken.")
		datazone_listTimeSeriesDataPointsCmd.Flags().String("started-at", "", "The timestamp at which the data points that you want to list started.")
		datazone_listTimeSeriesDataPointsCmd.MarkFlagRequired("domain-identifier")
		datazone_listTimeSeriesDataPointsCmd.MarkFlagRequired("entity-identifier")
		datazone_listTimeSeriesDataPointsCmd.MarkFlagRequired("entity-type")
		datazone_listTimeSeriesDataPointsCmd.MarkFlagRequired("form-name")
	})
	datazoneCmd.AddCommand(datazone_listTimeSeriesDataPointsCmd)
}

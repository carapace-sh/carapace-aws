package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_getAssetPropertyAggregatesCmd = &cobra.Command{
	Use:   "get-asset-property-aggregates",
	Short: "Gets aggregated values for an asset property.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_getAssetPropertyAggregatesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_getAssetPropertyAggregatesCmd).Standalone()

		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("aggregate-types", "", "The data aggregating function.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("asset-id", "", "The ID of the asset, in UUID format.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("end-date", "", "The inclusive end of the range from which to query historical data, expressed in seconds in Unix epoch time.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("property-alias", "", "The alias that identifies the property, such as an OPC-UA server data stream path (for example, `/company/windfarm/3/turbine/7/temperature`).")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("property-id", "", "The ID of the asset property, in UUID format.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("qualities", "", "The quality by which to filter asset data.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("resolution", "", "The time interval over which to aggregate data.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("start-date", "", "The exclusive start of the range from which to query historical data, expressed in seconds in Unix epoch time.")
		iotsitewise_getAssetPropertyAggregatesCmd.Flags().String("time-ordering", "", "The chronological sorting order of the requested information.")
		iotsitewise_getAssetPropertyAggregatesCmd.MarkFlagRequired("aggregate-types")
		iotsitewise_getAssetPropertyAggregatesCmd.MarkFlagRequired("end-date")
		iotsitewise_getAssetPropertyAggregatesCmd.MarkFlagRequired("resolution")
		iotsitewise_getAssetPropertyAggregatesCmd.MarkFlagRequired("start-date")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_getAssetPropertyAggregatesCmd)
}

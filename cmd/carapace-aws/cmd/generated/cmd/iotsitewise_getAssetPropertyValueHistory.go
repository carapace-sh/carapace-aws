package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_getAssetPropertyValueHistoryCmd = &cobra.Command{
	Use:   "get-asset-property-value-history",
	Short: "Gets the history of an asset property's values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_getAssetPropertyValueHistoryCmd).Standalone()

	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("asset-id", "", "The ID of the asset, in UUID format.")
	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("end-date", "", "The inclusive end of the range from which to query historical data, expressed in seconds in Unix epoch time.")
	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("property-alias", "", "The alias that identifies the property, such as an OPC-UA server data stream path (for example, `/company/windfarm/3/turbine/7/temperature`).")
	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("property-id", "", "The ID of the asset property, in UUID format.")
	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("qualities", "", "The quality by which to filter asset data.")
	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("start-date", "", "The exclusive start of the range from which to query historical data, expressed in seconds in Unix epoch time.")
	iotsitewise_getAssetPropertyValueHistoryCmd.Flags().String("time-ordering", "", "The chronological sorting order of the requested information.")
	iotsitewiseCmd.AddCommand(iotsitewise_getAssetPropertyValueHistoryCmd)
}

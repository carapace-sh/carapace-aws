package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_getInterpolatedAssetPropertyValuesCmd = &cobra.Command{
	Use:   "get-interpolated-asset-property-values",
	Short: "Get interpolated values for an asset property for a specified time interval, during a period of time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_getInterpolatedAssetPropertyValuesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_getInterpolatedAssetPropertyValuesCmd).Standalone()

		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("asset-id", "", "The ID of the asset, in UUID format.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("end-time-in-seconds", "", "The inclusive end of the range from which to interpolate data, expressed in seconds in Unix epoch time.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("end-time-offset-in-nanos", "", "The nanosecond offset converted from `endTimeInSeconds`.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("interval-in-seconds", "", "The time interval in seconds over which to interpolate data.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("interval-window-in-seconds", "", "The query interval for the window, in seconds.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("property-alias", "", "The alias that identifies the property, such as an OPC-UA server data stream path (for example, `/company/windfarm/3/turbine/7/temperature`).")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("property-id", "", "The ID of the asset property, in UUID format.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("quality", "", "The quality of the asset property value.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("start-time-in-seconds", "", "The exclusive start of the range from which to interpolate data, expressed in seconds in Unix epoch time.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("start-time-offset-in-nanos", "", "The nanosecond offset converted from `startTimeInSeconds`.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.Flags().String("type", "", "The interpolation type.")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.MarkFlagRequired("end-time-in-seconds")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.MarkFlagRequired("interval-in-seconds")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.MarkFlagRequired("quality")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.MarkFlagRequired("start-time-in-seconds")
		iotsitewise_getInterpolatedAssetPropertyValuesCmd.MarkFlagRequired("type")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_getInterpolatedAssetPropertyValuesCmd)
}

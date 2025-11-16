package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_getMetricDataCmd = &cobra.Command{
	Use:   "get-metric-data",
	Short: "You can use the `GetMetricData` API to retrieve CloudWatch metric values.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_getMetricDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_getMetricDataCmd).Standalone()

		cloudwatch_getMetricDataCmd.Flags().String("end-time", "", "The time stamp indicating the latest data to be returned.")
		cloudwatch_getMetricDataCmd.Flags().String("label-options", "", "This structure includes the `Timezone` parameter, which you can use to specify your time zone so that the labels of returned data display the correct time for your time zone.")
		cloudwatch_getMetricDataCmd.Flags().String("max-datapoints", "", "The maximum number of data points the request should return before paginating.")
		cloudwatch_getMetricDataCmd.Flags().String("metric-data-queries", "", "The metric queries to be returned.")
		cloudwatch_getMetricDataCmd.Flags().String("next-token", "", "Include this value, if it was returned by the previous `GetMetricData` operation, to get the next set of data points.")
		cloudwatch_getMetricDataCmd.Flags().String("scan-by", "", "The order in which data points should be returned.")
		cloudwatch_getMetricDataCmd.Flags().String("start-time", "", "The time stamp indicating the earliest data to be returned.")
		cloudwatch_getMetricDataCmd.MarkFlagRequired("end-time")
		cloudwatch_getMetricDataCmd.MarkFlagRequired("metric-data-queries")
		cloudwatch_getMetricDataCmd.MarkFlagRequired("start-time")
	})
	cloudwatchCmd.AddCommand(cloudwatch_getMetricDataCmd)
}

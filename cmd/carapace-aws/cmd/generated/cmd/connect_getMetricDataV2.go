package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getMetricDataV2Cmd = &cobra.Command{
	Use:   "get-metric-data-v2",
	Short: "Gets metric data from the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getMetricDataV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getMetricDataV2Cmd).Standalone()

		connect_getMetricDataV2Cmd.Flags().String("end-time", "", "The timestamp, in UNIX Epoch time format, at which to end the reporting interval for the retrieval of historical metrics data.")
		connect_getMetricDataV2Cmd.Flags().String("filters", "", "The filters to apply to returned metrics.")
		connect_getMetricDataV2Cmd.Flags().String("groupings", "", "The grouping applied to the metrics that are returned.")
		connect_getMetricDataV2Cmd.Flags().String("interval", "", "The interval period and timezone to apply to returned metrics.")
		connect_getMetricDataV2Cmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_getMetricDataV2Cmd.Flags().String("metrics", "", "The metrics to retrieve.")
		connect_getMetricDataV2Cmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_getMetricDataV2Cmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		connect_getMetricDataV2Cmd.Flags().String("start-time", "", "The timestamp, in UNIX Epoch time format, at which to start the reporting interval for the retrieval of historical metrics data.")
		connect_getMetricDataV2Cmd.MarkFlagRequired("end-time")
		connect_getMetricDataV2Cmd.MarkFlagRequired("filters")
		connect_getMetricDataV2Cmd.MarkFlagRequired("metrics")
		connect_getMetricDataV2Cmd.MarkFlagRequired("resource-arn")
		connect_getMetricDataV2Cmd.MarkFlagRequired("start-time")
	})
	connectCmd.AddCommand(connect_getMetricDataV2Cmd)
}

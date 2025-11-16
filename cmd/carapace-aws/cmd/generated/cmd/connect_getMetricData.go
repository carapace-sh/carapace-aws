package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getMetricDataCmd = &cobra.Command{
	Use:   "get-metric-data",
	Short: "Gets historical metric data from the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getMetricDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_getMetricDataCmd).Standalone()

		connect_getMetricDataCmd.Flags().String("end-time", "", "The timestamp, in UNIX Epoch time format, at which to end the reporting interval for the retrieval of historical metrics data.")
		connect_getMetricDataCmd.Flags().String("filters", "", "The queues, up to 100, or channels, to use to filter the metrics returned.")
		connect_getMetricDataCmd.Flags().String("groupings", "", "The grouping applied to the metrics returned.")
		connect_getMetricDataCmd.Flags().String("historical-metrics", "", "The metrics to retrieve.")
		connect_getMetricDataCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_getMetricDataCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_getMetricDataCmd.Flags().String("next-token", "", "The token for the next set of results.")
		connect_getMetricDataCmd.Flags().String("start-time", "", "The timestamp, in UNIX Epoch time format, at which to start the reporting interval for the retrieval of historical metrics data.")
		connect_getMetricDataCmd.MarkFlagRequired("end-time")
		connect_getMetricDataCmd.MarkFlagRequired("filters")
		connect_getMetricDataCmd.MarkFlagRequired("historical-metrics")
		connect_getMetricDataCmd.MarkFlagRequired("instance-id")
		connect_getMetricDataCmd.MarkFlagRequired("start-time")
	})
	connectCmd.AddCommand(connect_getMetricDataCmd)
}

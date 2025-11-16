package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_getCurrentMetricDataCmd = &cobra.Command{
	Use:   "get-current-metric-data",
	Short: "Gets the real-time metric data from the specified Amazon Connect instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_getCurrentMetricDataCmd).Standalone()

	connect_getCurrentMetricDataCmd.Flags().String("current-metrics", "", "The metrics to retrieve.")
	connect_getCurrentMetricDataCmd.Flags().String("filters", "", "The filters to apply to returned metrics.")
	connect_getCurrentMetricDataCmd.Flags().String("groupings", "", "Defines the level of aggregation for metrics data by a dimension(s).")
	connect_getCurrentMetricDataCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_getCurrentMetricDataCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_getCurrentMetricDataCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_getCurrentMetricDataCmd.Flags().String("sort-criteria", "", "The way to sort the resulting response based on metrics.")
	connect_getCurrentMetricDataCmd.MarkFlagRequired("current-metrics")
	connect_getCurrentMetricDataCmd.MarkFlagRequired("filters")
	connect_getCurrentMetricDataCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_getCurrentMetricDataCmd)
}

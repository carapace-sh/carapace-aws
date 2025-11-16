package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getCapacityManagerMetricDataCmd = &cobra.Command{
	Use:   "get-capacity-manager-metric-data",
	Short: "Retrieves capacity usage metrics for your EC2 resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getCapacityManagerMetricDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getCapacityManagerMetricDataCmd).Standalone()

		ec2_getCapacityManagerMetricDataCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getCapacityManagerMetricDataCmd.Flags().String("end-time", "", "The end time for the metric data query, in ISO 8601 format.")
		ec2_getCapacityManagerMetricDataCmd.Flags().String("filter-by", "", "Conditions to filter the metric data.")
		ec2_getCapacityManagerMetricDataCmd.Flags().String("group-by", "", "The dimensions by which to group the metric data.")
		ec2_getCapacityManagerMetricDataCmd.Flags().String("max-results", "", "The maximum number of data points to return.")
		ec2_getCapacityManagerMetricDataCmd.Flags().String("metric-names", "", "The names of the metrics to retrieve.")
		ec2_getCapacityManagerMetricDataCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getCapacityManagerMetricDataCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getCapacityManagerMetricDataCmd.Flags().String("period", "", "The granularity, in seconds, of the returned data points.")
		ec2_getCapacityManagerMetricDataCmd.Flags().String("start-time", "", "The start time for the metric data query, in ISO 8601 format.")
		ec2_getCapacityManagerMetricDataCmd.MarkFlagRequired("end-time")
		ec2_getCapacityManagerMetricDataCmd.MarkFlagRequired("metric-names")
		ec2_getCapacityManagerMetricDataCmd.Flag("no-dry-run").Hidden = true
		ec2_getCapacityManagerMetricDataCmd.MarkFlagRequired("period")
		ec2_getCapacityManagerMetricDataCmd.MarkFlagRequired("start-time")
	})
	ec2Cmd.AddCommand(ec2_getCapacityManagerMetricDataCmd)
}

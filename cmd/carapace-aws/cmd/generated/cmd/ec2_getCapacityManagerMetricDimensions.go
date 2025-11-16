package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getCapacityManagerMetricDimensionsCmd = &cobra.Command{
	Use:   "get-capacity-manager-metric-dimensions",
	Short: "Retrieves the available dimension values for capacity metrics within a specified time range.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getCapacityManagerMetricDimensionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_getCapacityManagerMetricDimensionsCmd).Standalone()

		ec2_getCapacityManagerMetricDimensionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getCapacityManagerMetricDimensionsCmd.Flags().String("end-time", "", "The end time for the dimension query, in ISO 8601 format.")
		ec2_getCapacityManagerMetricDimensionsCmd.Flags().String("filter-by", "", "Conditions to filter which dimension values are returned.")
		ec2_getCapacityManagerMetricDimensionsCmd.Flags().String("group-by", "", "The dimensions to group by when retrieving available dimension values.")
		ec2_getCapacityManagerMetricDimensionsCmd.Flags().String("max-results", "", "The maximum number of dimension combinations to return.")
		ec2_getCapacityManagerMetricDimensionsCmd.Flags().String("metric-names", "", "The metric names to use as an additional filter when retrieving dimensions.")
		ec2_getCapacityManagerMetricDimensionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_getCapacityManagerMetricDimensionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_getCapacityManagerMetricDimensionsCmd.Flags().String("start-time", "", "The start time for the dimension query, in ISO 8601 format.")
		ec2_getCapacityManagerMetricDimensionsCmd.MarkFlagRequired("end-time")
		ec2_getCapacityManagerMetricDimensionsCmd.MarkFlagRequired("group-by")
		ec2_getCapacityManagerMetricDimensionsCmd.MarkFlagRequired("metric-names")
		ec2_getCapacityManagerMetricDimensionsCmd.Flag("no-dry-run").Hidden = true
		ec2_getCapacityManagerMetricDimensionsCmd.MarkFlagRequired("start-time")
	})
	ec2Cmd.AddCommand(ec2_getCapacityManagerMetricDimensionsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd = &cobra.Command{
	Use:   "describe-aws-network-performance-metric-subscriptions",
	Short: "Describes the current Infrastructure Performance metric subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd).Standalone()

	ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describeAwsNetworkPerformanceMetricSubscriptionsCmd)
}

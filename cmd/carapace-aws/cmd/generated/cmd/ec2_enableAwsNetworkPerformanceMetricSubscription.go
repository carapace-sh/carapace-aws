package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd = &cobra.Command{
	Use:   "enable-aws-network-performance-metric-subscription",
	Short: "Enables Infrastructure Performance subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd).Standalone()

	ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().String("destination", "", "The target Region (like `us-east-2`) or Availability Zone ID (like `use2-az2`) that the metric subscription is enabled for.")
	ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().String("metric", "", "The metric used for the enabled subscription.")
	ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().String("source", "", "The source Region (like `us-east-1`) or Availability Zone ID (like `use1-az1`) that the metric subscription is enabled for.")
	ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().String("statistic", "", "The statistic used for the enabled subscription.")
	ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_enableAwsNetworkPerformanceMetricSubscriptionCmd)
}

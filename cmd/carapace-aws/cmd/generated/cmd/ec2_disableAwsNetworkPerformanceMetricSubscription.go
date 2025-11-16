package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd = &cobra.Command{
	Use:   "disable-aws-network-performance-metric-subscription",
	Short: "Disables Infrastructure Performance metric subscriptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd).Standalone()

		ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().String("destination", "", "The target Region or Availability Zone that the metric subscription is disabled for.")
		ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().String("metric", "", "The metric used for the disabled subscription.")
		ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().String("source", "", "The source Region or Availability Zone that the metric subscription is disabled for.")
		ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd.Flags().String("statistic", "", "The statistic used for the disabled subscription.")
		ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_disableAwsNetworkPerformanceMetricSubscriptionCmd)
}

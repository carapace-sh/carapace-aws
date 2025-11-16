package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_describeLoadBalancerTargetGroupsCmd = &cobra.Command{
	Use:   "describe-load-balancer-target-groups",
	Short: "This API operation is superseded by [DescribeTrafficSources](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html), which can describe multiple traffic sources types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_describeLoadBalancerTargetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_describeLoadBalancerTargetGroupsCmd).Standalone()

		autoscaling_describeLoadBalancerTargetGroupsCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_describeLoadBalancerTargetGroupsCmd.Flags().String("max-records", "", "The maximum number of items to return with this call.")
		autoscaling_describeLoadBalancerTargetGroupsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		autoscaling_describeLoadBalancerTargetGroupsCmd.MarkFlagRequired("auto-scaling-group-name")
	})
	autoscalingCmd.AddCommand(autoscaling_describeLoadBalancerTargetGroupsCmd)
}

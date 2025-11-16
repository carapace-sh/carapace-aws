package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_detachLoadBalancerTargetGroupsCmd = &cobra.Command{
	Use:   "detach-load-balancer-target-groups",
	Short: "This API operation is superseded by [DetachTrafficSources](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DescribeTrafficSources.html), which can detach multiple traffic sources types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_detachLoadBalancerTargetGroupsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_detachLoadBalancerTargetGroupsCmd).Standalone()

		autoscaling_detachLoadBalancerTargetGroupsCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_detachLoadBalancerTargetGroupsCmd.Flags().String("target-group-arns", "", "The Amazon Resource Names (ARN) of the target groups.")
		autoscaling_detachLoadBalancerTargetGroupsCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_detachLoadBalancerTargetGroupsCmd.MarkFlagRequired("target-group-arns")
	})
	autoscalingCmd.AddCommand(autoscaling_detachLoadBalancerTargetGroupsCmd)
}

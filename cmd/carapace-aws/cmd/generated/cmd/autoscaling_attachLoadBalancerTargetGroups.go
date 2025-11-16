package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_attachLoadBalancerTargetGroupsCmd = &cobra.Command{
	Use:   "attach-load-balancer-target-groups",
	Short: "This API operation is superseded by [AttachTrafficSources](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html), which can attach multiple traffic sources types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_attachLoadBalancerTargetGroupsCmd).Standalone()

	autoscaling_attachLoadBalancerTargetGroupsCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
	autoscaling_attachLoadBalancerTargetGroupsCmd.Flags().String("target-group-arns", "", "The Amazon Resource Names (ARNs) of the target groups.")
	autoscaling_attachLoadBalancerTargetGroupsCmd.MarkFlagRequired("auto-scaling-group-name")
	autoscaling_attachLoadBalancerTargetGroupsCmd.MarkFlagRequired("target-group-arns")
	autoscalingCmd.AddCommand(autoscaling_attachLoadBalancerTargetGroupsCmd)
}

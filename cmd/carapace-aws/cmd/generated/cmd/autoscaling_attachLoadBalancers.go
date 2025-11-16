package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_attachLoadBalancersCmd = &cobra.Command{
	Use:   "attach-load-balancers",
	Short: "This API operation is superseded by [https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API\\_AttachTrafficSources.html](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AttachTrafficSources.html), which can attach multiple traffic sources types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_attachLoadBalancersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_attachLoadBalancersCmd).Standalone()

		autoscaling_attachLoadBalancersCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_attachLoadBalancersCmd.Flags().String("load-balancer-names", "", "The names of the load balancers.")
		autoscaling_attachLoadBalancersCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_attachLoadBalancersCmd.MarkFlagRequired("load-balancer-names")
	})
	autoscalingCmd.AddCommand(autoscaling_attachLoadBalancersCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoscaling_detachLoadBalancersCmd = &cobra.Command{
	Use:   "detach-load-balancers",
	Short: "This API operation is superseded by [DetachTrafficSources](https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_DetachTrafficSources.html), which can detach multiple traffic sources types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoscaling_detachLoadBalancersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(autoscaling_detachLoadBalancersCmd).Standalone()

		autoscaling_detachLoadBalancersCmd.Flags().String("auto-scaling-group-name", "", "The name of the Auto Scaling group.")
		autoscaling_detachLoadBalancersCmd.Flags().String("load-balancer-names", "", "The names of the load balancers.")
		autoscaling_detachLoadBalancersCmd.MarkFlagRequired("auto-scaling-group-name")
		autoscaling_detachLoadBalancersCmd.MarkFlagRequired("load-balancer-names")
	})
	autoscalingCmd.AddCommand(autoscaling_detachLoadBalancersCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_detachLoadBalancerFromSubnetsCmd = &cobra.Command{
	Use:   "detach-load-balancer-from-subnets",
	Short: "Removes the specified subnets from the set of configured subnets for the load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_detachLoadBalancerFromSubnetsCmd).Standalone()

	elb_detachLoadBalancerFromSubnetsCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	elb_detachLoadBalancerFromSubnetsCmd.Flags().String("subnets", "", "The IDs of the subnets.")
	elb_detachLoadBalancerFromSubnetsCmd.MarkFlagRequired("load-balancer-name")
	elb_detachLoadBalancerFromSubnetsCmd.MarkFlagRequired("subnets")
	elbCmd.AddCommand(elb_detachLoadBalancerFromSubnetsCmd)
}

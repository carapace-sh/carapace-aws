package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_attachLoadBalancerToSubnetsCmd = &cobra.Command{
	Use:   "attach-load-balancer-to-subnets",
	Short: "Adds one or more subnets to the set of configured subnets for the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_attachLoadBalancerToSubnetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_attachLoadBalancerToSubnetsCmd).Standalone()

		elb_attachLoadBalancerToSubnetsCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_attachLoadBalancerToSubnetsCmd.Flags().String("subnets", "", "The IDs of the subnets to add.")
		elb_attachLoadBalancerToSubnetsCmd.MarkFlagRequired("load-balancer-name")
		elb_attachLoadBalancerToSubnetsCmd.MarkFlagRequired("subnets")
	})
	elbCmd.AddCommand(elb_attachLoadBalancerToSubnetsCmd)
}

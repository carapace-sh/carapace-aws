package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_createLoadBalancerCmd = &cobra.Command{
	Use:   "create-load-balancer",
	Short: "Creates a Classic Load Balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_createLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_createLoadBalancerCmd).Standalone()

		elb_createLoadBalancerCmd.Flags().String("availability-zones", "", "One or more Availability Zones from the same region as the load balancer.")
		elb_createLoadBalancerCmd.Flags().String("listeners", "", "The listeners.")
		elb_createLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_createLoadBalancerCmd.Flags().String("scheme", "", "The type of a load balancer.")
		elb_createLoadBalancerCmd.Flags().String("security-groups", "", "The IDs of the security groups to assign to the load balancer.")
		elb_createLoadBalancerCmd.Flags().String("subnets", "", "The IDs of the subnets in your VPC to attach to the load balancer.")
		elb_createLoadBalancerCmd.Flags().String("tags", "", "A list of tags to assign to the load balancer.")
		elb_createLoadBalancerCmd.MarkFlagRequired("listeners")
		elb_createLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	})
	elbCmd.AddCommand(elb_createLoadBalancerCmd)
}

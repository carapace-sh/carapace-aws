package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_registerInstancesWithLoadBalancerCmd = &cobra.Command{
	Use:   "register-instances-with-load-balancer",
	Short: "Adds the specified instances to the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_registerInstancesWithLoadBalancerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_registerInstancesWithLoadBalancerCmd).Standalone()

		elb_registerInstancesWithLoadBalancerCmd.Flags().String("instances", "", "The IDs of the instances.")
		elb_registerInstancesWithLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_registerInstancesWithLoadBalancerCmd.MarkFlagRequired("instances")
		elb_registerInstancesWithLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	})
	elbCmd.AddCommand(elb_registerInstancesWithLoadBalancerCmd)
}

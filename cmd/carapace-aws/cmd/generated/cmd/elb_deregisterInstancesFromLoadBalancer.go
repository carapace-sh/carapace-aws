package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_deregisterInstancesFromLoadBalancerCmd = &cobra.Command{
	Use:   "deregister-instances-from-load-balancer",
	Short: "Deregisters the specified instances from the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_deregisterInstancesFromLoadBalancerCmd).Standalone()

	elb_deregisterInstancesFromLoadBalancerCmd.Flags().String("instances", "", "The IDs of the instances.")
	elb_deregisterInstancesFromLoadBalancerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	elb_deregisterInstancesFromLoadBalancerCmd.MarkFlagRequired("instances")
	elb_deregisterInstancesFromLoadBalancerCmd.MarkFlagRequired("load-balancer-name")
	elbCmd.AddCommand(elb_deregisterInstancesFromLoadBalancerCmd)
}

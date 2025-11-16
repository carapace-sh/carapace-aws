package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_deleteLoadBalancerListenersCmd = &cobra.Command{
	Use:   "delete-load-balancer-listeners",
	Short: "Deletes the specified listeners from the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_deleteLoadBalancerListenersCmd).Standalone()

	elb_deleteLoadBalancerListenersCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	elb_deleteLoadBalancerListenersCmd.Flags().String("load-balancer-ports", "", "The client port numbers of the listeners.")
	elb_deleteLoadBalancerListenersCmd.MarkFlagRequired("load-balancer-name")
	elb_deleteLoadBalancerListenersCmd.MarkFlagRequired("load-balancer-ports")
	elbCmd.AddCommand(elb_deleteLoadBalancerListenersCmd)
}

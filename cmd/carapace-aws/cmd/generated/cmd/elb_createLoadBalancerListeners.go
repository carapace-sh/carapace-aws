package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_createLoadBalancerListenersCmd = &cobra.Command{
	Use:   "create-load-balancer-listeners",
	Short: "Creates one or more listeners for the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_createLoadBalancerListenersCmd).Standalone()

	elb_createLoadBalancerListenersCmd.Flags().String("listeners", "", "The listeners.")
	elb_createLoadBalancerListenersCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	elb_createLoadBalancerListenersCmd.MarkFlagRequired("listeners")
	elb_createLoadBalancerListenersCmd.MarkFlagRequired("load-balancer-name")
	elbCmd.AddCommand(elb_createLoadBalancerListenersCmd)
}

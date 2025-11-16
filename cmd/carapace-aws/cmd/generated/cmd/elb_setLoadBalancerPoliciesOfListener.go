package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_setLoadBalancerPoliciesOfListenerCmd = &cobra.Command{
	Use:   "set-load-balancer-policies-of-listener",
	Short: "Replaces the current set of policies for the specified load balancer port with the specified set of policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_setLoadBalancerPoliciesOfListenerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_setLoadBalancerPoliciesOfListenerCmd).Standalone()

		elb_setLoadBalancerPoliciesOfListenerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_setLoadBalancerPoliciesOfListenerCmd.Flags().String("load-balancer-port", "", "The external port of the load balancer.")
		elb_setLoadBalancerPoliciesOfListenerCmd.Flags().String("policy-names", "", "The names of the policies.")
		elb_setLoadBalancerPoliciesOfListenerCmd.MarkFlagRequired("load-balancer-name")
		elb_setLoadBalancerPoliciesOfListenerCmd.MarkFlagRequired("load-balancer-port")
		elb_setLoadBalancerPoliciesOfListenerCmd.MarkFlagRequired("policy-names")
	})
	elbCmd.AddCommand(elb_setLoadBalancerPoliciesOfListenerCmd)
}

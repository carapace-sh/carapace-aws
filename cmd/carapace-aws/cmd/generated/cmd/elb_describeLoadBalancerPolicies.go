package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_describeLoadBalancerPoliciesCmd = &cobra.Command{
	Use:   "describe-load-balancer-policies",
	Short: "Describes the specified policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_describeLoadBalancerPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_describeLoadBalancerPoliciesCmd).Standalone()

		elb_describeLoadBalancerPoliciesCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_describeLoadBalancerPoliciesCmd.Flags().String("policy-names", "", "The names of the policies.")
	})
	elbCmd.AddCommand(elb_describeLoadBalancerPoliciesCmd)
}

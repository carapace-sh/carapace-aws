package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_describeLoadBalancerPolicyTypesCmd = &cobra.Command{
	Use:   "describe-load-balancer-policy-types",
	Short: "Describes the specified load balancer policy types or all load balancer policy types.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_describeLoadBalancerPolicyTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_describeLoadBalancerPolicyTypesCmd).Standalone()

		elb_describeLoadBalancerPolicyTypesCmd.Flags().String("policy-type-names", "", "The names of the policy types.")
	})
	elbCmd.AddCommand(elb_describeLoadBalancerPolicyTypesCmd)
}

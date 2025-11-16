package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_createLoadBalancerPolicyCmd = &cobra.Command{
	Use:   "create-load-balancer-policy",
	Short: "Creates a policy with the specified attributes for the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_createLoadBalancerPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_createLoadBalancerPolicyCmd).Standalone()

		elb_createLoadBalancerPolicyCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_createLoadBalancerPolicyCmd.Flags().String("policy-attributes", "", "The policy attributes.")
		elb_createLoadBalancerPolicyCmd.Flags().String("policy-name", "", "The name of the load balancer policy to be created.")
		elb_createLoadBalancerPolicyCmd.Flags().String("policy-type-name", "", "The name of the base policy type.")
		elb_createLoadBalancerPolicyCmd.MarkFlagRequired("load-balancer-name")
		elb_createLoadBalancerPolicyCmd.MarkFlagRequired("policy-name")
		elb_createLoadBalancerPolicyCmd.MarkFlagRequired("policy-type-name")
	})
	elbCmd.AddCommand(elb_createLoadBalancerPolicyCmd)
}

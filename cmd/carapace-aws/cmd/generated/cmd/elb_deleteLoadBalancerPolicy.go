package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_deleteLoadBalancerPolicyCmd = &cobra.Command{
	Use:   "delete-load-balancer-policy",
	Short: "Deletes the specified policy from the specified load balancer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_deleteLoadBalancerPolicyCmd).Standalone()

	elb_deleteLoadBalancerPolicyCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
	elb_deleteLoadBalancerPolicyCmd.Flags().String("policy-name", "", "The name of the policy.")
	elb_deleteLoadBalancerPolicyCmd.MarkFlagRequired("load-balancer-name")
	elb_deleteLoadBalancerPolicyCmd.MarkFlagRequired("policy-name")
	elbCmd.AddCommand(elb_deleteLoadBalancerPolicyCmd)
}

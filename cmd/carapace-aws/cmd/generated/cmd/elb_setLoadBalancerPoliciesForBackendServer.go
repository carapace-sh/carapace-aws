package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elb_setLoadBalancerPoliciesForBackendServerCmd = &cobra.Command{
	Use:   "set-load-balancer-policies-for-backend-server",
	Short: "Replaces the set of policies associated with the specified port on which the EC2 instance is listening with a new set of policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elb_setLoadBalancerPoliciesForBackendServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elb_setLoadBalancerPoliciesForBackendServerCmd).Standalone()

		elb_setLoadBalancerPoliciesForBackendServerCmd.Flags().String("instance-port", "", "The port number associated with the EC2 instance.")
		elb_setLoadBalancerPoliciesForBackendServerCmd.Flags().String("load-balancer-name", "", "The name of the load balancer.")
		elb_setLoadBalancerPoliciesForBackendServerCmd.Flags().String("policy-names", "", "The names of the policies.")
		elb_setLoadBalancerPoliciesForBackendServerCmd.MarkFlagRequired("instance-port")
		elb_setLoadBalancerPoliciesForBackendServerCmd.MarkFlagRequired("load-balancer-name")
		elb_setLoadBalancerPoliciesForBackendServerCmd.MarkFlagRequired("policy-names")
	})
	elbCmd.AddCommand(elb_setLoadBalancerPoliciesForBackendServerCmd)
}

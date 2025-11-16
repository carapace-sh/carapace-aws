package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getLoadBalancerTlsPoliciesCmd = &cobra.Command{
	Use:   "get-load-balancer-tls-policies",
	Short: "Returns a list of TLS security policies that you can apply to Lightsail load balancers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getLoadBalancerTlsPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_getLoadBalancerTlsPoliciesCmd).Standalone()

		lightsail_getLoadBalancerTlsPoliciesCmd.Flags().String("page-token", "", "The token to advance to the next page of results from your request.")
	})
	lightsailCmd.AddCommand(lightsail_getLoadBalancerTlsPoliciesCmd)
}

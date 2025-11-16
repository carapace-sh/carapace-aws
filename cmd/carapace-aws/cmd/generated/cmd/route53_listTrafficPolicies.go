package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listTrafficPoliciesCmd = &cobra.Command{
	Use:   "list-traffic-policies",
	Short: "Gets information about the latest version for every traffic policy that is associated with the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listTrafficPoliciesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_listTrafficPoliciesCmd).Standalone()

		route53_listTrafficPoliciesCmd.Flags().String("max-items", "", "(Optional) The maximum number of traffic policies that you want Amazon Route 53 to return in response to this request.")
		route53_listTrafficPoliciesCmd.Flags().String("traffic-policy-id-marker", "", "(Conditional) For your first request to `ListTrafficPolicies`, don't include the `TrafficPolicyIdMarker` parameter.")
	})
	route53Cmd.AddCommand(route53_listTrafficPoliciesCmd)
}

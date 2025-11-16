package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_updateTrafficPolicyInstanceCmd = &cobra.Command{
	Use:   "update-traffic-policy-instance",
	Short: "After you submit a `UpdateTrafficPolicyInstance` request, there's a brief delay while Route\u00a053 creates the resource record sets that are specified in the traffic policy definition.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_updateTrafficPolicyInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_updateTrafficPolicyInstanceCmd).Standalone()

		route53_updateTrafficPolicyInstanceCmd.Flags().String("id", "", "The ID of the traffic policy instance that you want to update.")
		route53_updateTrafficPolicyInstanceCmd.Flags().String("traffic-policy-id", "", "The ID of the traffic policy that you want Amazon Route 53 to use to update resource record sets for the specified traffic policy instance.")
		route53_updateTrafficPolicyInstanceCmd.Flags().String("traffic-policy-version", "", "The version of the traffic policy that you want Amazon Route 53 to use to update resource record sets for the specified traffic policy instance.")
		route53_updateTrafficPolicyInstanceCmd.Flags().String("ttl", "", "The TTL that you want Amazon Route 53 to assign to all of the updated resource record sets.")
		route53_updateTrafficPolicyInstanceCmd.MarkFlagRequired("id")
		route53_updateTrafficPolicyInstanceCmd.MarkFlagRequired("traffic-policy-id")
		route53_updateTrafficPolicyInstanceCmd.MarkFlagRequired("traffic-policy-version")
		route53_updateTrafficPolicyInstanceCmd.MarkFlagRequired("ttl")
	})
	route53Cmd.AddCommand(route53_updateTrafficPolicyInstanceCmd)
}

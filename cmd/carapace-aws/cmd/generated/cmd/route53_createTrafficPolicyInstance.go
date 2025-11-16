package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_createTrafficPolicyInstanceCmd = &cobra.Command{
	Use:   "create-traffic-policy-instance",
	Short: "Creates resource record sets in a specified hosted zone based on the settings in a specified traffic policy version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_createTrafficPolicyInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_createTrafficPolicyInstanceCmd).Standalone()

		route53_createTrafficPolicyInstanceCmd.Flags().String("hosted-zone-id", "", "The ID of the hosted zone that you want Amazon Route 53 to create resource record sets in by using the configuration in a traffic policy.")
		route53_createTrafficPolicyInstanceCmd.Flags().String("name", "", "The domain name (such as example.com) or subdomain name (such as www.example.com) for which Amazon Route 53 responds to DNS queries by using the resource record sets that Route 53 creates for this traffic policy instance.")
		route53_createTrafficPolicyInstanceCmd.Flags().String("traffic-policy-id", "", "The ID of the traffic policy that you want to use to create resource record sets in the specified hosted zone.")
		route53_createTrafficPolicyInstanceCmd.Flags().String("traffic-policy-version", "", "The version of the traffic policy that you want to use to create resource record sets in the specified hosted zone.")
		route53_createTrafficPolicyInstanceCmd.Flags().String("ttl", "", "(Optional) The TTL that you want Amazon Route 53 to assign to all of the resource record sets that it creates in the specified hosted zone.")
		route53_createTrafficPolicyInstanceCmd.MarkFlagRequired("hosted-zone-id")
		route53_createTrafficPolicyInstanceCmd.MarkFlagRequired("name")
		route53_createTrafficPolicyInstanceCmd.MarkFlagRequired("traffic-policy-id")
		route53_createTrafficPolicyInstanceCmd.MarkFlagRequired("traffic-policy-version")
		route53_createTrafficPolicyInstanceCmd.MarkFlagRequired("ttl")
	})
	route53Cmd.AddCommand(route53_createTrafficPolicyInstanceCmd)
}

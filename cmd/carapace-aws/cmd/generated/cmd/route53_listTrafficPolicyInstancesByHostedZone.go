package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listTrafficPolicyInstancesByHostedZoneCmd = &cobra.Command{
	Use:   "list-traffic-policy-instances-by-hosted-zone",
	Short: "Gets information about the traffic policy instances that you created in a specified hosted zone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listTrafficPolicyInstancesByHostedZoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_listTrafficPolicyInstancesByHostedZoneCmd).Standalone()

		route53_listTrafficPolicyInstancesByHostedZoneCmd.Flags().String("hosted-zone-id", "", "The ID of the hosted zone that you want to list traffic policy instances for.")
		route53_listTrafficPolicyInstancesByHostedZoneCmd.Flags().String("max-items", "", "The maximum number of traffic policy instances to be included in the response body for this request.")
		route53_listTrafficPolicyInstancesByHostedZoneCmd.Flags().String("traffic-policy-instance-name-marker", "", "If the value of `IsTruncated` in the previous response is true, you have more traffic policy instances.")
		route53_listTrafficPolicyInstancesByHostedZoneCmd.Flags().String("traffic-policy-instance-type-marker", "", "If the value of `IsTruncated` in the previous response is true, you have more traffic policy instances.")
		route53_listTrafficPolicyInstancesByHostedZoneCmd.MarkFlagRequired("hosted-zone-id")
	})
	route53Cmd.AddCommand(route53_listTrafficPolicyInstancesByHostedZoneCmd)
}

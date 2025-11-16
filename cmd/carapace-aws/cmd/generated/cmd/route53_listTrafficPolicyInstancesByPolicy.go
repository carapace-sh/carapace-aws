package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listTrafficPolicyInstancesByPolicyCmd = &cobra.Command{
	Use:   "list-traffic-policy-instances-by-policy",
	Short: "Gets information about the traffic policy instances that you created by using a specify traffic policy version.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listTrafficPolicyInstancesByPolicyCmd).Standalone()

	route53_listTrafficPolicyInstancesByPolicyCmd.Flags().String("hosted-zone-id-marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more traffic policy instances.")
	route53_listTrafficPolicyInstancesByPolicyCmd.Flags().String("max-items", "", "The maximum number of traffic policy instances to be included in the response body for this request.")
	route53_listTrafficPolicyInstancesByPolicyCmd.Flags().String("traffic-policy-id", "", "The ID of the traffic policy for which you want to list traffic policy instances.")
	route53_listTrafficPolicyInstancesByPolicyCmd.Flags().String("traffic-policy-instance-name-marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more traffic policy instances.")
	route53_listTrafficPolicyInstancesByPolicyCmd.Flags().String("traffic-policy-instance-type-marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more traffic policy instances.")
	route53_listTrafficPolicyInstancesByPolicyCmd.Flags().String("traffic-policy-version", "", "The version of the traffic policy for which you want to list traffic policy instances.")
	route53_listTrafficPolicyInstancesByPolicyCmd.MarkFlagRequired("traffic-policy-id")
	route53_listTrafficPolicyInstancesByPolicyCmd.MarkFlagRequired("traffic-policy-version")
	route53Cmd.AddCommand(route53_listTrafficPolicyInstancesByPolicyCmd)
}

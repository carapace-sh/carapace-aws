package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listTrafficPolicyInstancesCmd = &cobra.Command{
	Use:   "list-traffic-policy-instances",
	Short: "Gets information about the traffic policy instances that you created by using the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listTrafficPolicyInstancesCmd).Standalone()

	route53_listTrafficPolicyInstancesCmd.Flags().String("hosted-zone-id-marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more traffic policy instances.")
	route53_listTrafficPolicyInstancesCmd.Flags().String("max-items", "", "The maximum number of traffic policy instances that you want Amazon Route 53 to return in response to a `ListTrafficPolicyInstances` request.")
	route53_listTrafficPolicyInstancesCmd.Flags().String("traffic-policy-instance-name-marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more traffic policy instances.")
	route53_listTrafficPolicyInstancesCmd.Flags().String("traffic-policy-instance-type-marker", "", "If the value of `IsTruncated` in the previous response was `true`, you have more traffic policy instances.")
	route53Cmd.AddCommand(route53_listTrafficPolicyInstancesCmd)
}

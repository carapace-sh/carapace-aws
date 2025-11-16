package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53_listTrafficPolicyVersionsCmd = &cobra.Command{
	Use:   "list-traffic-policy-versions",
	Short: "Gets information about all of the versions for a specified traffic policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53_listTrafficPolicyVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53_listTrafficPolicyVersionsCmd).Standalone()

		route53_listTrafficPolicyVersionsCmd.Flags().String("id", "", "Specify the value of `Id` of the traffic policy for which you want to list all versions.")
		route53_listTrafficPolicyVersionsCmd.Flags().String("max-items", "", "The maximum number of traffic policy versions that you want Amazon Route 53 to include in the response body for this request.")
		route53_listTrafficPolicyVersionsCmd.Flags().String("traffic-policy-version-marker", "", "For your first request to `ListTrafficPolicyVersions`, don't include the `TrafficPolicyVersionMarker` parameter.")
		route53_listTrafficPolicyVersionsCmd.MarkFlagRequired("id")
	})
	route53Cmd.AddCommand(route53_listTrafficPolicyVersionsCmd)
}

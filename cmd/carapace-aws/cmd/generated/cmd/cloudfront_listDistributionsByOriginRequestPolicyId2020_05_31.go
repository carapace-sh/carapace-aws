package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByOriginRequestPolicyId2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-origin-request-policy-id2020_05_31",
	Short: "Gets a list of distribution IDs for distributions that have a cache behavior that's associated with the specified origin request policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByOriginRequestPolicyId2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listDistributionsByOriginRequestPolicyId2020_05_31Cmd).Standalone()

		cloudfront_listDistributionsByOriginRequestPolicyId2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of distribution IDs.")
		cloudfront_listDistributionsByOriginRequestPolicyId2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distribution IDs that you want in the response.")
		cloudfront_listDistributionsByOriginRequestPolicyId2020_05_31Cmd.Flags().String("origin-request-policy-id", "", "The ID of the origin request policy whose associated distribution IDs you want to list.")
		cloudfront_listDistributionsByOriginRequestPolicyId2020_05_31Cmd.MarkFlagRequired("origin-request-policy-id")
	})
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByOriginRequestPolicyId2020_05_31Cmd)
}

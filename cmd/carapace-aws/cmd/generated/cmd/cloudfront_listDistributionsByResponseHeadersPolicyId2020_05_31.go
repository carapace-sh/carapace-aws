package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listDistributionsByResponseHeadersPolicyId2020_05_31Cmd = &cobra.Command{
	Use:   "list-distributions-by-response-headers-policy-id2020_05_31",
	Short: "Gets a list of distribution IDs for distributions that have a cache behavior that's associated with the specified response headers policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listDistributionsByResponseHeadersPolicyId2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listDistributionsByResponseHeadersPolicyId2020_05_31Cmd).Standalone()

		cloudfront_listDistributionsByResponseHeadersPolicyId2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of distribution IDs.")
		cloudfront_listDistributionsByResponseHeadersPolicyId2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of distribution IDs that you want to get in the response.")
		cloudfront_listDistributionsByResponseHeadersPolicyId2020_05_31Cmd.Flags().String("response-headers-policy-id", "", "The ID of the response headers policy whose associated distribution IDs you want to list.")
		cloudfront_listDistributionsByResponseHeadersPolicyId2020_05_31Cmd.MarkFlagRequired("response-headers-policy-id")
	})
	cloudfrontCmd.AddCommand(cloudfront_listDistributionsByResponseHeadersPolicyId2020_05_31Cmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listOriginRequestPolicies2020_05_31Cmd = &cobra.Command{
	Use:   "list-origin-request-policies2020_05_31",
	Short: "Gets a list of origin request policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listOriginRequestPolicies2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listOriginRequestPolicies2020_05_31Cmd).Standalone()

		cloudfront_listOriginRequestPolicies2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of origin request policies.")
		cloudfront_listOriginRequestPolicies2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of origin request policies that you want in the response.")
		cloudfront_listOriginRequestPolicies2020_05_31Cmd.Flags().String("type", "", "A filter to return only the specified kinds of origin request policies.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listOriginRequestPolicies2020_05_31Cmd)
}

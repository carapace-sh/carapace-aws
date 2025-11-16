package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listResponseHeadersPolicies2020_05_31Cmd = &cobra.Command{
	Use:   "list-response-headers-policies2020_05_31",
	Short: "Gets a list of response headers policies.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listResponseHeadersPolicies2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listResponseHeadersPolicies2020_05_31Cmd).Standalone()

		cloudfront_listResponseHeadersPolicies2020_05_31Cmd.Flags().String("marker", "", "Use this field when paginating results to indicate where to begin in your list of response headers policies.")
		cloudfront_listResponseHeadersPolicies2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of response headers policies that you want to get in the response.")
		cloudfront_listResponseHeadersPolicies2020_05_31Cmd.Flags().String("type", "", "A filter to get only the specified kind of response headers policies.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listResponseHeadersPolicies2020_05_31Cmd)
}

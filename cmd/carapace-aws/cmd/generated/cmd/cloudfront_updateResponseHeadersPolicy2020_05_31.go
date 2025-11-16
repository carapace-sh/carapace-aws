package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateResponseHeadersPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "update-response-headers-policy2020_05_31",
	Short: "Updates a response headers policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateResponseHeadersPolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_updateResponseHeadersPolicy2020_05_31Cmd).Standalone()

		cloudfront_updateResponseHeadersPolicy2020_05_31Cmd.Flags().String("id", "", "The identifier for the response headers policy that you are updating.")
		cloudfront_updateResponseHeadersPolicy2020_05_31Cmd.Flags().String("if-match", "", "The version of the response headers policy that you are updating.")
		cloudfront_updateResponseHeadersPolicy2020_05_31Cmd.Flags().String("response-headers-policy-config", "", "A response headers policy configuration.")
		cloudfront_updateResponseHeadersPolicy2020_05_31Cmd.MarkFlagRequired("id")
		cloudfront_updateResponseHeadersPolicy2020_05_31Cmd.MarkFlagRequired("response-headers-policy-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_updateResponseHeadersPolicy2020_05_31Cmd)
}

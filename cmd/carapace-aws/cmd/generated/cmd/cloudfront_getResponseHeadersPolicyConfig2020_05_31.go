package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getResponseHeadersPolicyConfig2020_05_31Cmd = &cobra.Command{
	Use:   "get-response-headers-policy-config2020_05_31",
	Short: "Gets a response headers policy configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getResponseHeadersPolicyConfig2020_05_31Cmd).Standalone()

	cloudfront_getResponseHeadersPolicyConfig2020_05_31Cmd.Flags().String("id", "", "The identifier for the response headers policy.")
	cloudfront_getResponseHeadersPolicyConfig2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_getResponseHeadersPolicyConfig2020_05_31Cmd)
}

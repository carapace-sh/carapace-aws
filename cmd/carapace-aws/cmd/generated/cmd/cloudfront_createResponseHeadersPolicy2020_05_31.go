package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createResponseHeadersPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "create-response-headers-policy2020_05_31",
	Short: "Creates a response headers policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createResponseHeadersPolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createResponseHeadersPolicy2020_05_31Cmd).Standalone()

		cloudfront_createResponseHeadersPolicy2020_05_31Cmd.Flags().String("response-headers-policy-config", "", "Contains metadata about the response headers policy, and a set of configurations that specify the HTTP headers.")
		cloudfront_createResponseHeadersPolicy2020_05_31Cmd.MarkFlagRequired("response-headers-policy-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_createResponseHeadersPolicy2020_05_31Cmd)
}

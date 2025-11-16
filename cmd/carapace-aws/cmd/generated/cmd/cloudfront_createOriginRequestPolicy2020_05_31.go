package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createOriginRequestPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "create-origin-request-policy2020_05_31",
	Short: "Creates an origin request policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createOriginRequestPolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createOriginRequestPolicy2020_05_31Cmd).Standalone()

		cloudfront_createOriginRequestPolicy2020_05_31Cmd.Flags().String("origin-request-policy-config", "", "An origin request policy configuration.")
		cloudfront_createOriginRequestPolicy2020_05_31Cmd.MarkFlagRequired("origin-request-policy-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_createOriginRequestPolicy2020_05_31Cmd)
}

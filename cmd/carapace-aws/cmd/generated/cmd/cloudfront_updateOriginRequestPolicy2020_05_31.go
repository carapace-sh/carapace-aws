package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateOriginRequestPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "update-origin-request-policy2020_05_31",
	Short: "Updates an origin request policy configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateOriginRequestPolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_updateOriginRequestPolicy2020_05_31Cmd).Standalone()

		cloudfront_updateOriginRequestPolicy2020_05_31Cmd.Flags().String("id", "", "The unique identifier for the origin request policy that you are updating.")
		cloudfront_updateOriginRequestPolicy2020_05_31Cmd.Flags().String("if-match", "", "The version of the origin request policy that you are updating.")
		cloudfront_updateOriginRequestPolicy2020_05_31Cmd.Flags().String("origin-request-policy-config", "", "An origin request policy configuration.")
		cloudfront_updateOriginRequestPolicy2020_05_31Cmd.MarkFlagRequired("id")
		cloudfront_updateOriginRequestPolicy2020_05_31Cmd.MarkFlagRequired("origin-request-policy-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_updateOriginRequestPolicy2020_05_31Cmd)
}

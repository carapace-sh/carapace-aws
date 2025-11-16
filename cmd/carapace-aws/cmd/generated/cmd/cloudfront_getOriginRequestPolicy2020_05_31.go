package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getOriginRequestPolicy2020_05_31Cmd = &cobra.Command{
	Use:   "get-origin-request-policy2020_05_31",
	Short: "Gets an origin request policy, including the following metadata:\n\n- The policy's identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getOriginRequestPolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getOriginRequestPolicy2020_05_31Cmd).Standalone()

		cloudfront_getOriginRequestPolicy2020_05_31Cmd.Flags().String("id", "", "The unique identifier for the origin request policy.")
		cloudfront_getOriginRequestPolicy2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getOriginRequestPolicy2020_05_31Cmd)
}

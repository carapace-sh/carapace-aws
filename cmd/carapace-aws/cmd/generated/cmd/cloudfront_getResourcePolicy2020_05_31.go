package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getResourcePolicy2020_05_31Cmd = &cobra.Command{
	Use:   "get-resource-policy2020_05_31",
	Short: "Retrieves the resource policy for the specified CloudFront resource that you own and have shared.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getResourcePolicy2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getResourcePolicy2020_05_31Cmd).Standalone()

		cloudfront_getResourcePolicy2020_05_31Cmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudFront resource that is associated with the resource policy.")
		cloudfront_getResourcePolicy2020_05_31Cmd.MarkFlagRequired("resource-arn")
	})
	cloudfrontCmd.AddCommand(cloudfront_getResourcePolicy2020_05_31Cmd)
}

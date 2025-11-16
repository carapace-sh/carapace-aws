package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_untagResource2020_05_31Cmd = &cobra.Command{
	Use:   "untag-resource2020_05_31",
	Short: "Remove tags from a CloudFront resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_untagResource2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_untagResource2020_05_31Cmd).Standalone()

		cloudfront_untagResource2020_05_31Cmd.Flags().String("resource", "", "An ARN of a CloudFront resource.")
		cloudfront_untagResource2020_05_31Cmd.Flags().String("tag-keys", "", "A complex type that contains zero or more `Tag` key elements.")
		cloudfront_untagResource2020_05_31Cmd.MarkFlagRequired("resource")
		cloudfront_untagResource2020_05_31Cmd.MarkFlagRequired("tag-keys")
	})
	cloudfrontCmd.AddCommand(cloudfront_untagResource2020_05_31Cmd)
}

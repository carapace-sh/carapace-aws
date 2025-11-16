package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createKeyGroup2020_05_31Cmd = &cobra.Command{
	Use:   "create-key-group2020_05_31",
	Short: "Creates a key group that you can use with [CloudFront signed URLs and signed cookies](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/PrivateContent.html).\n\nTo create a key group, you must specify at least one public key for the key group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createKeyGroup2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_createKeyGroup2020_05_31Cmd).Standalone()

		cloudfront_createKeyGroup2020_05_31Cmd.Flags().String("key-group-config", "", "A key group configuration.")
		cloudfront_createKeyGroup2020_05_31Cmd.MarkFlagRequired("key-group-config")
	})
	cloudfrontCmd.AddCommand(cloudfront_createKeyGroup2020_05_31Cmd)
}

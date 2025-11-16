package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_tagResource2020_05_31Cmd = &cobra.Command{
	Use:   "tag-resource2020_05_31",
	Short: "Add tags to a CloudFront resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_tagResource2020_05_31Cmd).Standalone()

	cloudfront_tagResource2020_05_31Cmd.Flags().String("resource", "", "An ARN of a CloudFront resource.")
	cloudfront_tagResource2020_05_31Cmd.Flags().String("tags", "", "A complex type that contains zero or more `Tag` elements.")
	cloudfront_tagResource2020_05_31Cmd.MarkFlagRequired("resource")
	cloudfront_tagResource2020_05_31Cmd.MarkFlagRequired("tags")
	cloudfrontCmd.AddCommand(cloudfront_tagResource2020_05_31Cmd)
}

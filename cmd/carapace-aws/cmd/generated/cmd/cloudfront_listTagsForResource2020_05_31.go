package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listTagsForResource2020_05_31Cmd = &cobra.Command{
	Use:   "list-tags-for-resource2020_05_31",
	Short: "List tags for a CloudFront resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listTagsForResource2020_05_31Cmd).Standalone()

	cloudfront_listTagsForResource2020_05_31Cmd.Flags().String("resource", "", "An ARN of a CloudFront resource.")
	cloudfront_listTagsForResource2020_05_31Cmd.MarkFlagRequired("resource")
	cloudfrontCmd.AddCommand(cloudfront_listTagsForResource2020_05_31Cmd)
}

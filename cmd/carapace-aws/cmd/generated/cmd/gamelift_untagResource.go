package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRemoves a tag assigned to a Amazon GameLift Servers resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_untagResourceCmd).Standalone()

		gamelift_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name ([ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that uniquely identifies the Amazon GameLift Servers resource that you want to remove tags from.")
		gamelift_untagResourceCmd.Flags().String("tag-keys", "", "A list of one or more tag keys to remove from the specified Amazon GameLift Servers resource.")
		gamelift_untagResourceCmd.MarkFlagRequired("resource-arn")
		gamelift_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	gameliftCmd.AddCommand(gamelift_untagResourceCmd)
}

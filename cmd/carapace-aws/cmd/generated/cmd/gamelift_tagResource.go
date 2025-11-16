package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nAssigns a tag to an Amazon GameLift Servers resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_tagResourceCmd).Standalone()

	gamelift_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name ([ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that uniquely identifies the Amazon GameLift Servers resource that you want to assign tags to.")
	gamelift_tagResourceCmd.Flags().String("tags", "", "A list of one or more tags to assign to the specified Amazon GameLift Servers resource.")
	gamelift_tagResourceCmd.MarkFlagRequired("resource-arn")
	gamelift_tagResourceCmd.MarkFlagRequired("tags")
	gameliftCmd.AddCommand(gamelift_tagResourceCmd)
}

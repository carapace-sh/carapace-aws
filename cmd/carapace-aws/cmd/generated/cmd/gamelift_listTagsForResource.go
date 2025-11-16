package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "**This API works with the following fleet types:** EC2, Anywhere, Container\n\nRetrieves all tags assigned to a Amazon GameLift Servers resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_listTagsForResourceCmd).Standalone()

		gamelift_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name ([ARN](https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-arn-format.html)) that uniquely identifies the Amazon GameLift Servers resource that you want to retrieve tags for.")
		gamelift_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	gameliftCmd.AddCommand(gamelift_listTagsForResourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a Amazon GameLift Streams resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_untagResourceCmd).Standalone()

	gameliftstreams_untagResourceCmd.Flags().String("resource-arn", "", "The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) of the Amazon GameLift Streams resource that you want to remove tags from.")
	gameliftstreams_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the specified Amazon GameLift Streams resource.")
	gameliftstreams_untagResourceCmd.MarkFlagRequired("resource-arn")
	gameliftstreams_untagResourceCmd.MarkFlagRequired("tag-keys")
	gameliftstreamsCmd.AddCommand(gameliftstreams_untagResourceCmd)
}

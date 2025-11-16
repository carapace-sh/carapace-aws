package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags to a Amazon GameLift Streams resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_tagResourceCmd).Standalone()

	gameliftstreams_tagResourceCmd.Flags().String("resource-arn", "", "The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) of the Amazon GameLift Streams resource that you want to apply tags to.")
	gameliftstreams_tagResourceCmd.Flags().String("tags", "", "A list of tags, in the form of key-value pairs, to assign to the specified Amazon GameLift Streams resource.")
	gameliftstreams_tagResourceCmd.MarkFlagRequired("resource-arn")
	gameliftstreams_tagResourceCmd.MarkFlagRequired("tags")
	gameliftstreamsCmd.AddCommand(gameliftstreams_tagResourceCmd)
}

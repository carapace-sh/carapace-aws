package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gameliftstreams_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves all tags assigned to a Amazon GameLift Streams resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gameliftstreams_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gameliftstreams_listTagsForResourceCmd).Standalone()

		gameliftstreams_listTagsForResourceCmd.Flags().String("resource-arn", "", "The [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference-arns.html) that you want to retrieve tags for.")
		gameliftstreams_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	gameliftstreamsCmd.AddCommand(gameliftstreams_listTagsForResourceCmd)
}

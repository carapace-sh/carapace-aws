package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a rescore execution plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_untagResourceCmd).Standalone()

	kendraRanking_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the rescore execution plan to remove the tag.")
	kendraRanking_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the rescore execution plan.")
	kendraRanking_untagResourceCmd.MarkFlagRequired("resource-arn")
	kendraRanking_untagResourceCmd.MarkFlagRequired("tag-keys")
	kendraRankingCmd.AddCommand(kendraRanking_untagResourceCmd)
}

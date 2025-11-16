package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds a specified tag to a specified rescore execution plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendraRanking_tagResourceCmd).Standalone()

		kendraRanking_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the rescore execution plan to tag.")
		kendraRanking_tagResourceCmd.Flags().String("tags", "", "A list of tag keys to add to a rescore execution plan.")
		kendraRanking_tagResourceCmd.MarkFlagRequired("resource-arn")
		kendraRanking_tagResourceCmd.MarkFlagRequired("tags")
	})
	kendraRankingCmd.AddCommand(kendraRanking_tagResourceCmd)
}

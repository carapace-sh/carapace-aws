package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendraRanking_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets a list of tags associated with a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendraRanking_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendraRanking_listTagsForResourceCmd).Standalone()

		kendraRanking_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the rescore execution plan to get a list of tags for.")
		kendraRanking_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	kendraRankingCmd.AddCommand(kendraRanking_listTagsForResourceCmd)
}

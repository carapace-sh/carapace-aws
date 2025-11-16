package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codepipeline_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets the set of key-value pairs (metadata) that are used to manage the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codepipeline_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codepipeline_listTagsForResourceCmd).Standalone()

		codepipeline_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		codepipeline_listTagsForResourceCmd.Flags().String("next-token", "", "The token that was returned from the previous API call, which would be used to return the next page of the list.")
		codepipeline_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to get tags for.")
		codepipeline_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	codepipelineCmd.AddCommand(codepipeline_listTagsForResourceCmd)
}

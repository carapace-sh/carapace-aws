package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags associated with the specified signaling channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_listTagsForResourceCmd).Standalone()

	kinesisvideo_listTagsForResourceCmd.Flags().String("next-token", "", "If you specify this parameter and the result of a `ListTagsForResource` call is truncated, the response includes a token that you can use in the next request to fetch the next batch of tags.")
	kinesisvideo_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the signaling channel for which you want to list tags.")
	kinesisvideo_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	kinesisvideoCmd.AddCommand(kinesisvideo_listTagsForResourceCmd)
}

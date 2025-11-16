package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_listTagsForStreamCmd = &cobra.Command{
	Use:   "list-tags-for-stream",
	Short: "Returns a list of tags associated with the specified stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_listTagsForStreamCmd).Standalone()

	kinesisvideo_listTagsForStreamCmd.Flags().String("next-token", "", "If you specify this parameter and the result of a `ListTagsForStream` call is truncated, the response includes a token that you can use in the next request to fetch the next batch of tags.")
	kinesisvideo_listTagsForStreamCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream that you want to list tags for.")
	kinesisvideo_listTagsForStreamCmd.Flags().String("stream-name", "", "The name of the stream that you want to list tags for.")
	kinesisvideoCmd.AddCommand(kinesisvideo_listTagsForStreamCmd)
}

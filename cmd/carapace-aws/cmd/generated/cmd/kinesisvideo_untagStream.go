package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_untagStreamCmd = &cobra.Command{
	Use:   "untag-stream",
	Short: "Removes one or more tags from a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_untagStreamCmd).Standalone()

	kinesisvideo_untagStreamCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream that you want to remove tags from.")
	kinesisvideo_untagStreamCmd.Flags().String("stream-name", "", "The name of the stream that you want to remove tags from.")
	kinesisvideo_untagStreamCmd.Flags().String("tag-key-list", "", "A list of the keys of the tags that you want to remove.")
	kinesisvideo_untagStreamCmd.MarkFlagRequired("tag-key-list")
	kinesisvideoCmd.AddCommand(kinesisvideo_untagStreamCmd)
}

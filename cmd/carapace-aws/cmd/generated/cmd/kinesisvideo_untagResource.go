package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a signaling channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_untagResourceCmd).Standalone()

	kinesisvideo_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the signaling channel from which you want to remove tags.")
	kinesisvideo_untagResourceCmd.Flags().String("tag-key-list", "", "A list of the keys of the tags that you want to remove.")
	kinesisvideo_untagResourceCmd.MarkFlagRequired("resource-arn")
	kinesisvideo_untagResourceCmd.MarkFlagRequired("tag-key-list")
	kinesisvideoCmd.AddCommand(kinesisvideo_untagResourceCmd)
}

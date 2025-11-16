package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Remove tags from a MediaConvert queue, preset, or job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_untagResourceCmd).Standalone()

		mediaconvert_untagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource that you want to remove tags from.")
		mediaconvert_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags that you want to remove from the resource.")
		mediaconvert_untagResourceCmd.MarkFlagRequired("arn")
	})
	mediaconvertCmd.AddCommand(mediaconvert_untagResourceCmd)
}

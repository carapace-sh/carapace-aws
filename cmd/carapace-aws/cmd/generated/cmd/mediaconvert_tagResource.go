package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add tags to a MediaConvert queue, preset, or job template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_tagResourceCmd).Standalone()

	mediaconvert_tagResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource that you want to tag.")
	mediaconvert_tagResourceCmd.Flags().String("tags", "", "The tags that you want to add to the resource.")
	mediaconvert_tagResourceCmd.MarkFlagRequired("arn")
	mediaconvert_tagResourceCmd.MarkFlagRequired("tags")
	mediaconvertCmd.AddCommand(mediaconvert_tagResourceCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconvert_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieve the tags for a MediaConvert resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconvert_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconvert_listTagsForResourceCmd).Standalone()

		mediaconvert_listTagsForResourceCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the resource that you want to list tags for.")
		mediaconvert_listTagsForResourceCmd.MarkFlagRequired("arn")
	})
	mediaconvertCmd.AddCommand(mediaconvert_listTagsForResourceCmd)
}

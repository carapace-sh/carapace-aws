package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to a signaling channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_tagResourceCmd).Standalone()

		kinesisvideo_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the signaling channel to which you want to add tags.")
		kinesisvideo_tagResourceCmd.Flags().String("tags", "", "A list of tags to associate with the specified signaling channel.")
		kinesisvideo_tagResourceCmd.MarkFlagRequired("resource-arn")
		kinesisvideo_tagResourceCmd.MarkFlagRequired("tags")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_tagResourceCmd)
}

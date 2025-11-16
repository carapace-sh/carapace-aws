package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_tagStreamCmd = &cobra.Command{
	Use:   "tag-stream",
	Short: "Adds one or more tags to a stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_tagStreamCmd).Standalone()

	kinesisvideo_tagStreamCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to add the tag or tags to.")
	kinesisvideo_tagStreamCmd.Flags().String("stream-name", "", "The name of the stream that you want to add the tag or tags to.")
	kinesisvideo_tagStreamCmd.Flags().String("tags", "", "A list of tags to associate with the specified stream.")
	kinesisvideo_tagStreamCmd.MarkFlagRequired("tags")
	kinesisvideoCmd.AddCommand(kinesisvideo_tagStreamCmd)
}

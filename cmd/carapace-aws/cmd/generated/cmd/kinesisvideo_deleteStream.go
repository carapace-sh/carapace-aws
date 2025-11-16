package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_deleteStreamCmd = &cobra.Command{
	Use:   "delete-stream",
	Short: "Deletes a Kinesis video stream and the data contained in the stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_deleteStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_deleteStreamCmd).Standalone()

		kinesisvideo_deleteStreamCmd.Flags().String("current-version", "", "Optional: The version of the stream that you want to delete.")
		kinesisvideo_deleteStreamCmd.Flags().String("stream-arn", "", "The Amazon Resource Name (ARN) of the stream that you want to delete.")
		kinesisvideo_deleteStreamCmd.MarkFlagRequired("stream-arn")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_deleteStreamCmd)
}

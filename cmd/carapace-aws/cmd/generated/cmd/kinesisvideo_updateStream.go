package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_updateStreamCmd = &cobra.Command{
	Use:   "update-stream",
	Short: "Updates stream metadata, such as the device name and media type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_updateStreamCmd).Standalone()

	kinesisvideo_updateStreamCmd.Flags().String("current-version", "", "The version of the stream whose metadata you want to update.")
	kinesisvideo_updateStreamCmd.Flags().String("device-name", "", "The name of the device that is writing to the stream.")
	kinesisvideo_updateStreamCmd.Flags().String("media-type", "", "The stream's media type.")
	kinesisvideo_updateStreamCmd.Flags().String("stream-arn", "", "The ARN of the stream whose metadata you want to update.")
	kinesisvideo_updateStreamCmd.Flags().String("stream-name", "", "The name of the stream whose metadata you want to update.")
	kinesisvideo_updateStreamCmd.MarkFlagRequired("current-version")
	kinesisvideoCmd.AddCommand(kinesisvideo_updateStreamCmd)
}

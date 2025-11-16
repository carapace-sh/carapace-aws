package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_deleteSignalingChannelCmd = &cobra.Command{
	Use:   "delete-signaling-channel",
	Short: "Deletes a specified signaling channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_deleteSignalingChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_deleteSignalingChannelCmd).Standalone()

		kinesisvideo_deleteSignalingChannelCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the signaling channel that you want to delete.")
		kinesisvideo_deleteSignalingChannelCmd.Flags().String("current-version", "", "The current version of the signaling channel that you want to delete.")
		kinesisvideo_deleteSignalingChannelCmd.MarkFlagRequired("channel-arn")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_deleteSignalingChannelCmd)
}

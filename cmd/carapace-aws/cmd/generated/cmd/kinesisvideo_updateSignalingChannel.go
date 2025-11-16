package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_updateSignalingChannelCmd = &cobra.Command{
	Use:   "update-signaling-channel",
	Short: "Updates the existing signaling channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_updateSignalingChannelCmd).Standalone()

	kinesisvideo_updateSignalingChannelCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the signaling channel that you want to update.")
	kinesisvideo_updateSignalingChannelCmd.Flags().String("current-version", "", "The current version of the signaling channel that you want to update.")
	kinesisvideo_updateSignalingChannelCmd.Flags().String("single-master-configuration", "", "The structure containing the configuration for the `SINGLE_MASTER` type of the signaling channel that you want to update.")
	kinesisvideo_updateSignalingChannelCmd.MarkFlagRequired("channel-arn")
	kinesisvideo_updateSignalingChannelCmd.MarkFlagRequired("current-version")
	kinesisvideoCmd.AddCommand(kinesisvideo_updateSignalingChannelCmd)
}

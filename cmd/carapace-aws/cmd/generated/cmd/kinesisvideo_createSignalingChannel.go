package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_createSignalingChannelCmd = &cobra.Command{
	Use:   "create-signaling-channel",
	Short: "Creates a signaling channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_createSignalingChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_createSignalingChannelCmd).Standalone()

		kinesisvideo_createSignalingChannelCmd.Flags().String("channel-name", "", "A name for the signaling channel that you are creating.")
		kinesisvideo_createSignalingChannelCmd.Flags().String("channel-type", "", "A type of the signaling channel that you are creating.")
		kinesisvideo_createSignalingChannelCmd.Flags().String("single-master-configuration", "", "A structure containing the configuration for the `SINGLE_MASTER` channel type.")
		kinesisvideo_createSignalingChannelCmd.Flags().String("tags", "", "A set of tags (key-value pairs) that you want to associate with this channel.")
		kinesisvideo_createSignalingChannelCmd.MarkFlagRequired("channel-name")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_createSignalingChannelCmd)
}

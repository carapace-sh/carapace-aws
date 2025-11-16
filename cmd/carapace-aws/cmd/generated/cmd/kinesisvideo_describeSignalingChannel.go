package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_describeSignalingChannelCmd = &cobra.Command{
	Use:   "describe-signaling-channel",
	Short: "Returns the most current information about the signaling channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_describeSignalingChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_describeSignalingChannelCmd).Standalone()

		kinesisvideo_describeSignalingChannelCmd.Flags().String("channel-arn", "", "The ARN of the signaling channel that you want to describe.")
		kinesisvideo_describeSignalingChannelCmd.Flags().String("channel-name", "", "The name of the signaling channel that you want to describe.")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_describeSignalingChannelCmd)
}

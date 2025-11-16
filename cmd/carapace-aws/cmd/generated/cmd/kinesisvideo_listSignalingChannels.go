package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_listSignalingChannelsCmd = &cobra.Command{
	Use:   "list-signaling-channels",
	Short: "Returns an array of `ChannelInfo` objects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_listSignalingChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisvideo_listSignalingChannelsCmd).Standalone()

		kinesisvideo_listSignalingChannelsCmd.Flags().String("channel-name-condition", "", "Optional: Returns only the channels that satisfy a specific condition.")
		kinesisvideo_listSignalingChannelsCmd.Flags().String("max-results", "", "The maximum number of channels to return in the response.")
		kinesisvideo_listSignalingChannelsCmd.Flags().String("next-token", "", "If you specify this parameter, when the result of a `ListSignalingChannels` operation is truncated, the call returns the `NextToken` in the response.")
	})
	kinesisvideoCmd.AddCommand(kinesisvideo_listSignalingChannelsCmd)
}

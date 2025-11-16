package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoWebrtcStorage_joinStorageSessionCmd = &cobra.Command{
	Use:   "join-storage-session",
	Short: "Before using this API, you must call the `GetSignalingChannelEndpoint` API to request the WEBRTC endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoWebrtcStorage_joinStorageSessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisVideoWebrtcStorage_joinStorageSessionCmd).Standalone()

		kinesisVideoWebrtcStorage_joinStorageSessionCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the signaling channel.")
		kinesisVideoWebrtcStorage_joinStorageSessionCmd.MarkFlagRequired("channel-arn")
	})
	kinesisVideoWebrtcStorageCmd.AddCommand(kinesisVideoWebrtcStorage_joinStorageSessionCmd)
}

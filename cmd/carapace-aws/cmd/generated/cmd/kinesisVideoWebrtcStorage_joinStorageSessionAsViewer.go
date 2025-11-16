package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoWebrtcStorage_joinStorageSessionAsViewerCmd = &cobra.Command{
	Use:   "join-storage-session-as-viewer",
	Short: "Join the ongoing one way-video and/or multi-way audio WebRTC session as a viewer for an input channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoWebrtcStorage_joinStorageSessionAsViewerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisVideoWebrtcStorage_joinStorageSessionAsViewerCmd).Standalone()

		kinesisVideoWebrtcStorage_joinStorageSessionAsViewerCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the signaling channel.")
		kinesisVideoWebrtcStorage_joinStorageSessionAsViewerCmd.Flags().String("client-id", "", "The unique identifier for the sender client.")
		kinesisVideoWebrtcStorage_joinStorageSessionAsViewerCmd.MarkFlagRequired("channel-arn")
		kinesisVideoWebrtcStorage_joinStorageSessionAsViewerCmd.MarkFlagRequired("client-id")
	})
	kinesisVideoWebrtcStorageCmd.AddCommand(kinesisVideoWebrtcStorage_joinStorageSessionAsViewerCmd)
}

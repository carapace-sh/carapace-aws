package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoSignalingCmd = &cobra.Command{
	Use:   "kinesis-video-signaling",
	Short: "Kinesis Video Streams Signaling Service is a intermediate service that establishes a communication channel for discovering peers, transmitting offers and answers in order to establish peer-to-peer connection in webRTC technology.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoSignalingCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kinesisVideoSignalingCmd).Standalone()

	})
	rootCmd.AddCommand(kinesisVideoSignalingCmd)
}

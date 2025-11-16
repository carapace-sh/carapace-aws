package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_updateFlowMediaStreamCmd = &cobra.Command{
	Use:   "update-flow-media-stream",
	Short: "Updates an existing media stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_updateFlowMediaStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_updateFlowMediaStreamCmd).Standalone()

		mediaconnect_updateFlowMediaStreamCmd.Flags().String("attributes", "", "The attributes that you want to assign to the media stream.")
		mediaconnect_updateFlowMediaStreamCmd.Flags().String("clock-rate", "", "The sample rate for the stream.")
		mediaconnect_updateFlowMediaStreamCmd.Flags().String("description", "", "A description that can help you quickly identify what your media stream is used for.")
		mediaconnect_updateFlowMediaStreamCmd.Flags().String("flow-arn", "", "The Amazon Resource Name (ARN) of the flow that is associated with the media stream that you updated.")
		mediaconnect_updateFlowMediaStreamCmd.Flags().String("media-stream-name", "", "The media stream that you updated.")
		mediaconnect_updateFlowMediaStreamCmd.Flags().String("media-stream-type", "", "The type of media stream.")
		mediaconnect_updateFlowMediaStreamCmd.Flags().String("video-format", "", "The resolution of the video.")
		mediaconnect_updateFlowMediaStreamCmd.MarkFlagRequired("flow-arn")
		mediaconnect_updateFlowMediaStreamCmd.MarkFlagRequired("media-stream-name")
	})
	mediaconnectCmd.AddCommand(mediaconnect_updateFlowMediaStreamCmd)
}

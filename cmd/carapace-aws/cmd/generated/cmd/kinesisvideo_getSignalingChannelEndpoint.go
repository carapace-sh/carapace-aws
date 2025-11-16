package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisvideo_getSignalingChannelEndpointCmd = &cobra.Command{
	Use:   "get-signaling-channel-endpoint",
	Short: "Provides an endpoint for the specified signaling channel to send and receive messages.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisvideo_getSignalingChannelEndpointCmd).Standalone()

	kinesisvideo_getSignalingChannelEndpointCmd.Flags().String("channel-arn", "", "The Amazon Resource Name (ARN) of the signalling channel for which you want to get an endpoint.")
	kinesisvideo_getSignalingChannelEndpointCmd.Flags().String("single-master-channel-endpoint-configuration", "", "A structure containing the endpoint configuration for the `SINGLE_MASTER` channel type.")
	kinesisvideo_getSignalingChannelEndpointCmd.MarkFlagRequired("channel-arn")
	kinesisvideoCmd.AddCommand(kinesisvideo_getSignalingChannelEndpointCmd)
}

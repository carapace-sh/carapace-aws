package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesisVideoSignaling_getIceServerConfigCmd = &cobra.Command{
	Use:   "get-ice-server-config",
	Short: "Gets the Interactive Connectivity Establishment (ICE) server configuration information, including URIs, username, and password which can be used to configure the WebRTC connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesisVideoSignaling_getIceServerConfigCmd).Standalone()

	kinesisVideoSignaling_getIceServerConfigCmd.Flags().String("channel-arn", "", "The ARN of the signaling channel to be used for the peer-to-peer connection between configured peers.")
	kinesisVideoSignaling_getIceServerConfigCmd.Flags().String("client-id", "", "Unique identifier for the viewer.")
	kinesisVideoSignaling_getIceServerConfigCmd.Flags().String("service", "", "Specifies the desired service.")
	kinesisVideoSignaling_getIceServerConfigCmd.Flags().String("username", "", "An optional user ID to be associated with the credentials.")
	kinesisVideoSignaling_getIceServerConfigCmd.MarkFlagRequired("channel-arn")
	kinesisVideoSignalingCmd.AddCommand(kinesisVideoSignaling_getIceServerConfigCmd)
}

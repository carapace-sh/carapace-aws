package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_createOriginEndpointCmd = &cobra.Command{
	Use:   "create-origin-endpoint",
	Short: "The endpoint is attached to a channel, and represents the output of the live content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_createOriginEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_createOriginEndpointCmd).Standalone()

		mediapackagev2_createOriginEndpointCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("container-type", "", "The type of container to attach to this origin endpoint.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("dash-manifests", "", "A DASH manifest configuration.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("description", "", "Enter any descriptive text that helps you to identify the origin endpoint.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("force-endpoint-error-configuration", "", "The failover settings for the endpoint.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("hls-manifests", "", "An HTTP live streaming (HLS) manifest configuration.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("low-latency-hls-manifests", "", "A low-latency HLS manifest configuration.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("mss-manifests", "", "A list of Microsoft Smooth Streaming (MSS) manifest configurations for the origin endpoint.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("origin-endpoint-name", "", "The name that describes the origin endpoint.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("segment", "", "The segment configuration, including the segment name, duration, and other configuration values.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("startover-window-seconds", "", "The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing.")
		mediapackagev2_createOriginEndpointCmd.Flags().String("tags", "", "A comma-separated list of tag key:value pairs that you define.")
		mediapackagev2_createOriginEndpointCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_createOriginEndpointCmd.MarkFlagRequired("channel-name")
		mediapackagev2_createOriginEndpointCmd.MarkFlagRequired("container-type")
		mediapackagev2_createOriginEndpointCmd.MarkFlagRequired("origin-endpoint-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_createOriginEndpointCmd)
}

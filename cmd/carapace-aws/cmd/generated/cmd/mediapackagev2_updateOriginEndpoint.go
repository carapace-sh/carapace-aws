package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_updateOriginEndpointCmd = &cobra.Command{
	Use:   "update-origin-endpoint",
	Short: "Update the specified origin endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_updateOriginEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_updateOriginEndpointCmd).Standalone()

		mediapackagev2_updateOriginEndpointCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("container-type", "", "The type of container attached to this origin endpoint.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("dash-manifests", "", "A DASH manifest configuration.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("description", "", "Any descriptive information that you want to add to the origin endpoint for future identification purposes.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("etag", "", "The expected current Entity Tag (ETag) for the resource.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("force-endpoint-error-configuration", "", "The failover settings for the endpoint.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("hls-manifests", "", "An HTTP live streaming (HLS) manifest configuration.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("low-latency-hls-manifests", "", "A low-latency HLS manifest configuration.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("mss-manifests", "", "A list of Microsoft Smooth Streaming (MSS) manifest configurations to update for the origin endpoint.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("origin-endpoint-name", "", "The name that describes the origin endpoint.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("segment", "", "The segment configuration, including the segment name, duration, and other configuration values.")
		mediapackagev2_updateOriginEndpointCmd.Flags().String("startover-window-seconds", "", "The size of the window (in seconds) to create a window of the live stream that's available for on-demand viewing.")
		mediapackagev2_updateOriginEndpointCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_updateOriginEndpointCmd.MarkFlagRequired("channel-name")
		mediapackagev2_updateOriginEndpointCmd.MarkFlagRequired("container-type")
		mediapackagev2_updateOriginEndpointCmd.MarkFlagRequired("origin-endpoint-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_updateOriginEndpointCmd)
}

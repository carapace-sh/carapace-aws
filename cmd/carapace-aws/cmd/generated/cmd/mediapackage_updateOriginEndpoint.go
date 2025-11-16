package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_updateOriginEndpointCmd = &cobra.Command{
	Use:   "update-origin-endpoint",
	Short: "Updates an existing OriginEndpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_updateOriginEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_updateOriginEndpointCmd).Standalone()

		mediapackage_updateOriginEndpointCmd.Flags().String("authorization", "", "")
		mediapackage_updateOriginEndpointCmd.Flags().String("cmaf-package", "", "")
		mediapackage_updateOriginEndpointCmd.Flags().String("dash-package", "", "")
		mediapackage_updateOriginEndpointCmd.Flags().String("description", "", "A short text description of the OriginEndpoint.")
		mediapackage_updateOriginEndpointCmd.Flags().String("hls-package", "", "")
		mediapackage_updateOriginEndpointCmd.Flags().String("id", "", "The ID of the OriginEndpoint to update.")
		mediapackage_updateOriginEndpointCmd.Flags().String("manifest-name", "", "A short string that will be appended to the end of the Endpoint URL.")
		mediapackage_updateOriginEndpointCmd.Flags().String("mss-package", "", "")
		mediapackage_updateOriginEndpointCmd.Flags().String("origination", "", "Control whether origination of video is allowed for this OriginEndpoint.")
		mediapackage_updateOriginEndpointCmd.Flags().String("startover-window-seconds", "", "Maximum duration (in seconds) of content to retain for startover playback.")
		mediapackage_updateOriginEndpointCmd.Flags().String("time-delay-seconds", "", "Amount of delay (in seconds) to enforce on the playback of live content.")
		mediapackage_updateOriginEndpointCmd.Flags().String("whitelist", "", "A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint.")
		mediapackage_updateOriginEndpointCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_updateOriginEndpointCmd)
}

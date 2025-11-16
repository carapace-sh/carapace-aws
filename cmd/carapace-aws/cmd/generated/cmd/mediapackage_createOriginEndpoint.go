package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_createOriginEndpointCmd = &cobra.Command{
	Use:   "create-origin-endpoint",
	Short: "Creates a new OriginEndpoint record.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_createOriginEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_createOriginEndpointCmd).Standalone()

		mediapackage_createOriginEndpointCmd.Flags().String("authorization", "", "")
		mediapackage_createOriginEndpointCmd.Flags().String("channel-id", "", "The ID of the Channel that the OriginEndpoint will be associated with.")
		mediapackage_createOriginEndpointCmd.Flags().String("cmaf-package", "", "")
		mediapackage_createOriginEndpointCmd.Flags().String("dash-package", "", "")
		mediapackage_createOriginEndpointCmd.Flags().String("description", "", "A short text description of the OriginEndpoint.")
		mediapackage_createOriginEndpointCmd.Flags().String("hls-package", "", "")
		mediapackage_createOriginEndpointCmd.Flags().String("id", "", "The ID of the OriginEndpoint.")
		mediapackage_createOriginEndpointCmd.Flags().String("manifest-name", "", "A short string that will be used as the filename of the OriginEndpoint URL (defaults to \"index\").")
		mediapackage_createOriginEndpointCmd.Flags().String("mss-package", "", "")
		mediapackage_createOriginEndpointCmd.Flags().String("origination", "", "Control whether origination of video is allowed for this OriginEndpoint.")
		mediapackage_createOriginEndpointCmd.Flags().String("startover-window-seconds", "", "Maximum duration (seconds) of content to retain for startover playback.")
		mediapackage_createOriginEndpointCmd.Flags().String("tags", "", "")
		mediapackage_createOriginEndpointCmd.Flags().String("time-delay-seconds", "", "Amount of delay (seconds) to enforce on the playback of live content.")
		mediapackage_createOriginEndpointCmd.Flags().String("whitelist", "", "A list of source IP CIDR blocks that will be allowed to access the OriginEndpoint.")
		mediapackage_createOriginEndpointCmd.MarkFlagRequired("channel-id")
		mediapackage_createOriginEndpointCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_createOriginEndpointCmd)
}

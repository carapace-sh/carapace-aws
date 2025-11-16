package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_getOriginEndpointCmd = &cobra.Command{
	Use:   "get-origin-endpoint",
	Short: "Retrieves the specified origin endpoint that's configured in AWS Elemental MediaPackage to obtain its playback URL and to view the packaging settings that it's currently using.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_getOriginEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_getOriginEndpointCmd).Standalone()

		mediapackagev2_getOriginEndpointCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_getOriginEndpointCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_getOriginEndpointCmd.Flags().String("origin-endpoint-name", "", "The name that describes the origin endpoint.")
		mediapackagev2_getOriginEndpointCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_getOriginEndpointCmd.MarkFlagRequired("channel-name")
		mediapackagev2_getOriginEndpointCmd.MarkFlagRequired("origin-endpoint-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_getOriginEndpointCmd)
}

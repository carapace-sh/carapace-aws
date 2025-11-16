package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_deleteOriginEndpointCmd = &cobra.Command{
	Use:   "delete-origin-endpoint",
	Short: "Origin endpoints can serve content until they're deleted.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_deleteOriginEndpointCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_deleteOriginEndpointCmd).Standalone()

		mediapackagev2_deleteOriginEndpointCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_deleteOriginEndpointCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_deleteOriginEndpointCmd.Flags().String("origin-endpoint-name", "", "The name that describes the origin endpoint.")
		mediapackagev2_deleteOriginEndpointCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_deleteOriginEndpointCmd.MarkFlagRequired("channel-name")
		mediapackagev2_deleteOriginEndpointCmd.MarkFlagRequired("origin-endpoint-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_deleteOriginEndpointCmd)
}

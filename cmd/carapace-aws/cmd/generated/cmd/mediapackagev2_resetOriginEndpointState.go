package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_resetOriginEndpointStateCmd = &cobra.Command{
	Use:   "reset-origin-endpoint-state",
	Short: "Resetting the origin endpoint can help to resolve unexpected behavior and other content packaging issues.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_resetOriginEndpointStateCmd).Standalone()

	mediapackagev2_resetOriginEndpointStateCmd.Flags().String("channel-group-name", "", "The name of the channel group that contains the channel with the origin endpoint that you are resetting.")
	mediapackagev2_resetOriginEndpointStateCmd.Flags().String("channel-name", "", "The name of the channel with the origin endpoint that you are resetting.")
	mediapackagev2_resetOriginEndpointStateCmd.Flags().String("origin-endpoint-name", "", "The name of the origin endpoint that you are resetting.")
	mediapackagev2_resetOriginEndpointStateCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2_resetOriginEndpointStateCmd.MarkFlagRequired("channel-name")
	mediapackagev2_resetOriginEndpointStateCmd.MarkFlagRequired("origin-endpoint-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_resetOriginEndpointStateCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Create a channel to start receiving content streams.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_createChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_createChannelCmd).Standalone()

		mediapackagev2_createChannelCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_createChannelCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_createChannelCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
		mediapackagev2_createChannelCmd.Flags().String("description", "", "Enter any descriptive text that helps you to identify the channel.")
		mediapackagev2_createChannelCmd.Flags().String("input-switch-configuration", "", "The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.")
		mediapackagev2_createChannelCmd.Flags().String("input-type", "", "The input type will be an immutable field which will be used to define whether the channel will allow CMAF ingest or HLS ingest.")
		mediapackagev2_createChannelCmd.Flags().String("output-header-configuration", "", "The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.")
		mediapackagev2_createChannelCmd.Flags().String("tags", "", "A comma-separated list of tag key:value pairs that you define.")
		mediapackagev2_createChannelCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_createChannelCmd.MarkFlagRequired("channel-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_createChannelCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Update the specified channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_updateChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_updateChannelCmd).Standalone()

		mediapackagev2_updateChannelCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_updateChannelCmd.Flags().String("channel-name", "", "The name that describes the channel.")
		mediapackagev2_updateChannelCmd.Flags().String("description", "", "Any descriptive information that you want to add to the channel for future identification purposes.")
		mediapackagev2_updateChannelCmd.Flags().String("etag", "", "The expected current Entity Tag (ETag) for the resource.")
		mediapackagev2_updateChannelCmd.Flags().String("input-switch-configuration", "", "The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.")
		mediapackagev2_updateChannelCmd.Flags().String("output-header-configuration", "", "The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.")
		mediapackagev2_updateChannelCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_updateChannelCmd.MarkFlagRequired("channel-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_updateChannelCmd)
}

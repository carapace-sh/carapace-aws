package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_updateChannelGroupCmd = &cobra.Command{
	Use:   "update-channel-group",
	Short: "Update the specified channel group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_updateChannelGroupCmd).Standalone()

	mediapackagev2_updateChannelGroupCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_updateChannelGroupCmd.Flags().String("description", "", "Any descriptive information that you want to add to the channel group for future identification purposes.")
	mediapackagev2_updateChannelGroupCmd.Flags().String("etag", "", "The expected current Entity Tag (ETag) for the resource.")
	mediapackagev2_updateChannelGroupCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_updateChannelGroupCmd)
}

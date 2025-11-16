package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_createChannelGroupCmd = &cobra.Command{
	Use:   "create-channel-group",
	Short: "Create a channel group to group your channels and origin endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_createChannelGroupCmd).Standalone()

	mediapackagev2_createChannelGroupCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_createChannelGroupCmd.Flags().String("client-token", "", "A unique, case-sensitive token that you provide to ensure the idempotency of the request.")
	mediapackagev2_createChannelGroupCmd.Flags().String("description", "", "Enter any descriptive text that helps you to identify the channel group.")
	mediapackagev2_createChannelGroupCmd.Flags().String("tags", "", "A comma-separated list of tag key:value pairs that you define.")
	mediapackagev2_createChannelGroupCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_createChannelGroupCmd)
}

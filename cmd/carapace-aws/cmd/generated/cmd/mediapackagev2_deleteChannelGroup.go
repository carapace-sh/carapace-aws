package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_deleteChannelGroupCmd = &cobra.Command{
	Use:   "delete-channel-group",
	Short: "Delete a channel group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_deleteChannelGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_deleteChannelGroupCmd).Standalone()

		mediapackagev2_deleteChannelGroupCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
		mediapackagev2_deleteChannelGroupCmd.MarkFlagRequired("channel-group-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_deleteChannelGroupCmd)
}

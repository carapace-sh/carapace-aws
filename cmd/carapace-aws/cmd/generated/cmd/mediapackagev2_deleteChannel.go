package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_deleteChannelCmd = &cobra.Command{
	Use:   "delete-channel",
	Short: "Delete a channel to stop AWS Elemental MediaPackage from receiving further content.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_deleteChannelCmd).Standalone()

	mediapackagev2_deleteChannelCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_deleteChannelCmd.Flags().String("channel-name", "", "The name that describes the channel.")
	mediapackagev2_deleteChannelCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2_deleteChannelCmd.MarkFlagRequired("channel-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_deleteChannelCmd)
}

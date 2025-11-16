package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_getChannelCmd = &cobra.Command{
	Use:   "get-channel",
	Short: "Retrieves the specified channel that's configured in AWS Elemental MediaPackage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_getChannelCmd).Standalone()

	mediapackagev2_getChannelCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_getChannelCmd.Flags().String("channel-name", "", "The name that describes the channel.")
	mediapackagev2_getChannelCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2_getChannelCmd.MarkFlagRequired("channel-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_getChannelCmd)
}

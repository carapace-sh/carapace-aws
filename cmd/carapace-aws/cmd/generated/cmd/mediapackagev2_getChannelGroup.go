package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_getChannelGroupCmd = &cobra.Command{
	Use:   "get-channel-group",
	Short: "Retrieves the specified channel group that's configured in AWS Elemental MediaPackage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_getChannelGroupCmd).Standalone()

	mediapackagev2_getChannelGroupCmd.Flags().String("channel-group-name", "", "The name that describes the channel group.")
	mediapackagev2_getChannelGroupCmd.MarkFlagRequired("channel-group-name")
	mediapackagev2Cmd.AddCommand(mediapackagev2_getChannelGroupCmd)
}

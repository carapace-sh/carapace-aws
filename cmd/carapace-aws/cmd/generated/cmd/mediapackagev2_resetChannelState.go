package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackagev2_resetChannelStateCmd = &cobra.Command{
	Use:   "reset-channel-state",
	Short: "Resetting the channel can help to clear errors from misconfigurations in the encoder.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackagev2_resetChannelStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackagev2_resetChannelStateCmd).Standalone()

		mediapackagev2_resetChannelStateCmd.Flags().String("channel-group-name", "", "The name of the channel group that contains the channel that you are resetting.")
		mediapackagev2_resetChannelStateCmd.Flags().String("channel-name", "", "The name of the channel that you are resetting.")
		mediapackagev2_resetChannelStateCmd.MarkFlagRequired("channel-group-name")
		mediapackagev2_resetChannelStateCmd.MarkFlagRequired("channel-name")
	})
	mediapackagev2Cmd.AddCommand(mediapackagev2_resetChannelStateCmd)
}

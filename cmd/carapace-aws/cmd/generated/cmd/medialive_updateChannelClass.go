package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateChannelClassCmd = &cobra.Command{
	Use:   "update-channel-class",
	Short: "Changes the class of the channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateChannelClassCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_updateChannelClassCmd).Standalone()

		medialive_updateChannelClassCmd.Flags().String("channel-class", "", "The channel class that you wish to update this channel to use.")
		medialive_updateChannelClassCmd.Flags().String("channel-id", "", "Channel Id of the channel whose class should be updated.")
		medialive_updateChannelClassCmd.Flags().String("destinations", "", "A list of output destinations for this channel.")
		medialive_updateChannelClassCmd.MarkFlagRequired("channel-class")
		medialive_updateChannelClassCmd.MarkFlagRequired("channel-id")
	})
	medialiveCmd.AddCommand(medialive_updateChannelClassCmd)
}

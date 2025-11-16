package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteChannelCmd = &cobra.Command{
	Use:   "delete-channel",
	Short: "Starts deletion of channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteChannelCmd).Standalone()

	medialive_deleteChannelCmd.Flags().String("channel-id", "", "Unique ID of the channel.")
	medialive_deleteChannelCmd.MarkFlagRequired("channel-id")
	medialiveCmd.AddCommand(medialive_deleteChannelCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_deleteChannelCmd = &cobra.Command{
	Use:   "delete-channel",
	Short: "Deletes a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_deleteChannelCmd).Standalone()

	mediatailor_deleteChannelCmd.Flags().String("channel-name", "", "The name of the channel.")
	mediatailor_deleteChannelCmd.MarkFlagRequired("channel-name")
	mediatailorCmd.AddCommand(mediatailor_deleteChannelCmd)
}

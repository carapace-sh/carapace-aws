package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_startChannelCmd = &cobra.Command{
	Use:   "start-channel",
	Short: "Starts a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_startChannelCmd).Standalone()

	mediatailor_startChannelCmd.Flags().String("channel-name", "", "The name of the channel.")
	mediatailor_startChannelCmd.MarkFlagRequired("channel-name")
	mediatailorCmd.AddCommand(mediatailor_startChannelCmd)
}

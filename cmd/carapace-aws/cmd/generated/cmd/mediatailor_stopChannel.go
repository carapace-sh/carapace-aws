package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_stopChannelCmd = &cobra.Command{
	Use:   "stop-channel",
	Short: "Stops a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_stopChannelCmd).Standalone()

	mediatailor_stopChannelCmd.Flags().String("channel-name", "", "The name of the channel.")
	mediatailor_stopChannelCmd.MarkFlagRequired("channel-name")
	mediatailorCmd.AddCommand(mediatailor_stopChannelCmd)
}

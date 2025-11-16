package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_updateChannelCmd = &cobra.Command{
	Use:   "update-channel",
	Short: "Updates a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_updateChannelCmd).Standalone()

	mediatailor_updateChannelCmd.Flags().String("audiences", "", "The list of audiences defined in channel.")
	mediatailor_updateChannelCmd.Flags().String("channel-name", "", "The name of the channel.")
	mediatailor_updateChannelCmd.Flags().String("filler-slate", "", "The slate used to fill gaps between programs in the schedule.")
	mediatailor_updateChannelCmd.Flags().String("outputs", "", "The channel's output properties.")
	mediatailor_updateChannelCmd.Flags().String("time-shift-configuration", "", "The time-shifted viewing configuration you want to associate to the channel.")
	mediatailor_updateChannelCmd.MarkFlagRequired("channel-name")
	mediatailor_updateChannelCmd.MarkFlagRequired("outputs")
	mediatailorCmd.AddCommand(mediatailor_updateChannelCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_createChannelCmd = &cobra.Command{
	Use:   "create-channel",
	Short: "Creates a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_createChannelCmd).Standalone()

	mediatailor_createChannelCmd.Flags().String("audiences", "", "The list of audiences defined in channel.")
	mediatailor_createChannelCmd.Flags().String("channel-name", "", "The name of the channel.")
	mediatailor_createChannelCmd.Flags().String("filler-slate", "", "The slate used to fill gaps between programs in the schedule.")
	mediatailor_createChannelCmd.Flags().String("outputs", "", "The channel's output properties.")
	mediatailor_createChannelCmd.Flags().String("playback-mode", "", "The type of playback mode to use for this channel.")
	mediatailor_createChannelCmd.Flags().String("tags", "", "The tags to assign to the channel.")
	mediatailor_createChannelCmd.Flags().String("tier", "", "The tier of the channel.")
	mediatailor_createChannelCmd.Flags().String("time-shift-configuration", "", "The time-shifted viewing configuration you want to associate to the channel.")
	mediatailor_createChannelCmd.MarkFlagRequired("channel-name")
	mediatailor_createChannelCmd.MarkFlagRequired("outputs")
	mediatailor_createChannelCmd.MarkFlagRequired("playback-mode")
	mediatailorCmd.AddCommand(mediatailor_createChannelCmd)
}

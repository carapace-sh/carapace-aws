package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_getChannelScheduleCmd = &cobra.Command{
	Use:   "get-channel-schedule",
	Short: "Retrieves information about your channel's schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_getChannelScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_getChannelScheduleCmd).Standalone()

		mediatailor_getChannelScheduleCmd.Flags().String("audience", "", "The single audience for GetChannelScheduleRequest.")
		mediatailor_getChannelScheduleCmd.Flags().String("channel-name", "", "The name of the channel associated with this Channel Schedule.")
		mediatailor_getChannelScheduleCmd.Flags().String("duration-minutes", "", "The duration in minutes of the channel schedule.")
		mediatailor_getChannelScheduleCmd.Flags().String("max-results", "", "The maximum number of channel schedules that you want MediaTailor to return in response to the current request.")
		mediatailor_getChannelScheduleCmd.Flags().String("next-token", "", "(Optional) If the playback configuration has more than `MaxResults` channel schedules, use `NextToken` to get the second and subsequent pages of results.")
		mediatailor_getChannelScheduleCmd.MarkFlagRequired("channel-name")
	})
	mediatailorCmd.AddCommand(mediatailor_getChannelScheduleCmd)
}

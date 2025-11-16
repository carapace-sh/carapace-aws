package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_createPrefetchScheduleCmd = &cobra.Command{
	Use:   "create-prefetch-schedule",
	Short: "Creates a prefetch schedule for a playback configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_createPrefetchScheduleCmd).Standalone()

	mediatailor_createPrefetchScheduleCmd.Flags().String("consumption", "", "The configuration settings for how and when MediaTailor consumes prefetched ads from the ad decision server for single prefetch schedules.")
	mediatailor_createPrefetchScheduleCmd.Flags().String("name", "", "The name to assign to the schedule request.")
	mediatailor_createPrefetchScheduleCmd.Flags().String("playback-configuration-name", "", "The name to assign to the playback configuration.")
	mediatailor_createPrefetchScheduleCmd.Flags().String("recurring-prefetch-configuration", "", "The configuration that defines how and when MediaTailor performs ad prefetching in a live event.")
	mediatailor_createPrefetchScheduleCmd.Flags().String("retrieval", "", "The configuration settings for retrieval of prefetched ads from the ad decision server.")
	mediatailor_createPrefetchScheduleCmd.Flags().String("schedule-type", "", "The frequency that MediaTailor creates prefetch schedules.")
	mediatailor_createPrefetchScheduleCmd.Flags().String("stream-id", "", "An optional stream identifier that MediaTailor uses to prefetch ads for multiple streams that use the same playback configuration.")
	mediatailor_createPrefetchScheduleCmd.MarkFlagRequired("name")
	mediatailor_createPrefetchScheduleCmd.MarkFlagRequired("playback-configuration-name")
	mediatailorCmd.AddCommand(mediatailor_createPrefetchScheduleCmd)
}

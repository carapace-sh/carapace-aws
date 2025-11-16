package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_listPrefetchSchedulesCmd = &cobra.Command{
	Use:   "list-prefetch-schedules",
	Short: "Lists the prefetch schedules for a playback configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_listPrefetchSchedulesCmd).Standalone()

	mediatailor_listPrefetchSchedulesCmd.Flags().String("max-results", "", "The maximum number of prefetch schedules that you want MediaTailor to return in response to the current request.")
	mediatailor_listPrefetchSchedulesCmd.Flags().String("next-token", "", "Pagination token returned by the list request when results exceed the maximum allowed.")
	mediatailor_listPrefetchSchedulesCmd.Flags().String("playback-configuration-name", "", "Retrieves the prefetch schedule(s) for a specific playback configuration.")
	mediatailor_listPrefetchSchedulesCmd.Flags().String("schedule-type", "", "The type of prefetch schedules that you want to list.")
	mediatailor_listPrefetchSchedulesCmd.Flags().String("stream-id", "", "An optional filtering parameter whereby MediaTailor filters the prefetch schedules to include only specific streams.")
	mediatailor_listPrefetchSchedulesCmd.MarkFlagRequired("playback-configuration-name")
	mediatailorCmd.AddCommand(mediatailor_listPrefetchSchedulesCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_getPrefetchScheduleCmd = &cobra.Command{
	Use:   "get-prefetch-schedule",
	Short: "Retrieves a prefetch schedule for a playback configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_getPrefetchScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_getPrefetchScheduleCmd).Standalone()

		mediatailor_getPrefetchScheduleCmd.Flags().String("name", "", "The name of the prefetch schedule.")
		mediatailor_getPrefetchScheduleCmd.Flags().String("playback-configuration-name", "", "Returns information about the prefetch schedule for a specific playback configuration.")
		mediatailor_getPrefetchScheduleCmd.MarkFlagRequired("name")
		mediatailor_getPrefetchScheduleCmd.MarkFlagRequired("playback-configuration-name")
	})
	mediatailorCmd.AddCommand(mediatailor_getPrefetchScheduleCmd)
}

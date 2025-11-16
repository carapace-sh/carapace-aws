package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_deletePrefetchScheduleCmd = &cobra.Command{
	Use:   "delete-prefetch-schedule",
	Short: "Deletes a prefetch schedule for a specific playback configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_deletePrefetchScheduleCmd).Standalone()

	mediatailor_deletePrefetchScheduleCmd.Flags().String("name", "", "The name of the prefetch schedule.")
	mediatailor_deletePrefetchScheduleCmd.Flags().String("playback-configuration-name", "", "The name of the playback configuration for this prefetch schedule.")
	mediatailor_deletePrefetchScheduleCmd.MarkFlagRequired("name")
	mediatailor_deletePrefetchScheduleCmd.MarkFlagRequired("playback-configuration-name")
	mediatailorCmd.AddCommand(mediatailor_deletePrefetchScheduleCmd)
}

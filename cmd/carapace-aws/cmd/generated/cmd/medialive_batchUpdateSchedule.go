package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_batchUpdateScheduleCmd = &cobra.Command{
	Use:   "batch-update-schedule",
	Short: "Update a channel schedule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_batchUpdateScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_batchUpdateScheduleCmd).Standalone()

		medialive_batchUpdateScheduleCmd.Flags().String("channel-id", "", "Id of the channel whose schedule is being updated.")
		medialive_batchUpdateScheduleCmd.Flags().String("creates", "", "Schedule actions to create in the schedule.")
		medialive_batchUpdateScheduleCmd.Flags().String("deletes", "", "Schedule actions to delete from the schedule.")
		medialive_batchUpdateScheduleCmd.MarkFlagRequired("channel-id")
	})
	medialiveCmd.AddCommand(medialive_batchUpdateScheduleCmd)
}

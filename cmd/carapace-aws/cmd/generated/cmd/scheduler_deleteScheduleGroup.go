package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_deleteScheduleGroupCmd = &cobra.Command{
	Use:   "delete-schedule-group",
	Short: "Deletes the specified schedule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_deleteScheduleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_deleteScheduleGroupCmd).Standalone()

		scheduler_deleteScheduleGroupCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
		scheduler_deleteScheduleGroupCmd.Flags().String("name", "", "The name of the schedule group to delete.")
		scheduler_deleteScheduleGroupCmd.MarkFlagRequired("name")
	})
	schedulerCmd.AddCommand(scheduler_deleteScheduleGroupCmd)
}

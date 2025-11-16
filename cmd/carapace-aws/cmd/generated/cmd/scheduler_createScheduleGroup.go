package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_createScheduleGroupCmd = &cobra.Command{
	Use:   "create-schedule-group",
	Short: "Creates the specified schedule group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_createScheduleGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(scheduler_createScheduleGroupCmd).Standalone()

		scheduler_createScheduleGroupCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
		scheduler_createScheduleGroupCmd.Flags().String("name", "", "The name of the schedule group that you are creating.")
		scheduler_createScheduleGroupCmd.Flags().String("tags", "", "The list of tags to associate with the schedule group.")
		scheduler_createScheduleGroupCmd.MarkFlagRequired("name")
	})
	schedulerCmd.AddCommand(scheduler_createScheduleGroupCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduler_deleteScheduleCmd = &cobra.Command{
	Use:   "delete-schedule",
	Short: "Deletes the specified schedule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduler_deleteScheduleCmd).Standalone()

	scheduler_deleteScheduleCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
	scheduler_deleteScheduleCmd.Flags().String("group-name", "", "The name of the schedule group associated with this schedule.")
	scheduler_deleteScheduleCmd.Flags().String("name", "", "The name of the schedule to delete.")
	scheduler_deleteScheduleCmd.MarkFlagRequired("name")
	schedulerCmd.AddCommand(scheduler_deleteScheduleCmd)
}

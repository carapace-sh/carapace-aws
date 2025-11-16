package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateWorkerScheduleCmd = &cobra.Command{
	Use:   "update-worker-schedule",
	Short: "Updates the schedule for a worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateWorkerScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_updateWorkerScheduleCmd).Standalone()

		deadline_updateWorkerScheduleCmd.Flags().String("farm-id", "", "The farm ID to update.")
		deadline_updateWorkerScheduleCmd.Flags().String("fleet-id", "", "The fleet ID to update.")
		deadline_updateWorkerScheduleCmd.Flags().String("updated-session-actions", "", "The session actions associated with the worker schedule to update.")
		deadline_updateWorkerScheduleCmd.Flags().String("worker-id", "", "The worker ID to update.")
		deadline_updateWorkerScheduleCmd.MarkFlagRequired("farm-id")
		deadline_updateWorkerScheduleCmd.MarkFlagRequired("fleet-id")
		deadline_updateWorkerScheduleCmd.MarkFlagRequired("worker-id")
	})
	deadlineCmd.AddCommand(deadline_updateWorkerScheduleCmd)
}

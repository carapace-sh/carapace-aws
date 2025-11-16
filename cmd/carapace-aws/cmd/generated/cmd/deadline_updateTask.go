package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateTaskCmd = &cobra.Command{
	Use:   "update-task",
	Short: "Updates a task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateTaskCmd).Standalone()

	deadline_updateTaskCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_updateTaskCmd.Flags().String("farm-id", "", "The farm ID to update.")
	deadline_updateTaskCmd.Flags().String("job-id", "", "The job ID to update.")
	deadline_updateTaskCmd.Flags().String("queue-id", "", "The queue ID to update.")
	deadline_updateTaskCmd.Flags().String("step-id", "", "The step ID to update.")
	deadline_updateTaskCmd.Flags().String("target-run-status", "", "The run status with which to start the task.")
	deadline_updateTaskCmd.Flags().String("task-id", "", "The task ID to update.")
	deadline_updateTaskCmd.MarkFlagRequired("farm-id")
	deadline_updateTaskCmd.MarkFlagRequired("job-id")
	deadline_updateTaskCmd.MarkFlagRequired("queue-id")
	deadline_updateTaskCmd.MarkFlagRequired("step-id")
	deadline_updateTaskCmd.MarkFlagRequired("target-run-status")
	deadline_updateTaskCmd.MarkFlagRequired("task-id")
	deadlineCmd.AddCommand(deadline_updateTaskCmd)
}

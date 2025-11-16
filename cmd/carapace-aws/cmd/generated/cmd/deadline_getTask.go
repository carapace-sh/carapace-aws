package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getTaskCmd = &cobra.Command{
	Use:   "get-task",
	Short: "Gets a task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getTaskCmd).Standalone()

	deadline_getTaskCmd.Flags().String("farm-id", "", "The farm ID of the farm connected to the task.")
	deadline_getTaskCmd.Flags().String("job-id", "", "The job ID of the job connected to the task.")
	deadline_getTaskCmd.Flags().String("queue-id", "", "The queue ID for the queue connected to the task.")
	deadline_getTaskCmd.Flags().String("step-id", "", "The step ID for the step connected to the task.")
	deadline_getTaskCmd.Flags().String("task-id", "", "The task ID.")
	deadline_getTaskCmd.MarkFlagRequired("farm-id")
	deadline_getTaskCmd.MarkFlagRequired("job-id")
	deadline_getTaskCmd.MarkFlagRequired("queue-id")
	deadline_getTaskCmd.MarkFlagRequired("step-id")
	deadline_getTaskCmd.MarkFlagRequired("task-id")
	deadlineCmd.AddCommand(deadline_getTaskCmd)
}

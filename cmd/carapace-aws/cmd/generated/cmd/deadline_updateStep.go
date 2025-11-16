package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateStepCmd = &cobra.Command{
	Use:   "update-step",
	Short: "Updates a step.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateStepCmd).Standalone()

	deadline_updateStepCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_updateStepCmd.Flags().String("farm-id", "", "The farm ID to update.")
	deadline_updateStepCmd.Flags().String("job-id", "", "The job ID to update.")
	deadline_updateStepCmd.Flags().String("queue-id", "", "The queue ID to update.")
	deadline_updateStepCmd.Flags().String("step-id", "", "The step ID to update.")
	deadline_updateStepCmd.Flags().String("target-task-run-status", "", "The task status to update the step's tasks to.")
	deadline_updateStepCmd.MarkFlagRequired("farm-id")
	deadline_updateStepCmd.MarkFlagRequired("job-id")
	deadline_updateStepCmd.MarkFlagRequired("queue-id")
	deadline_updateStepCmd.MarkFlagRequired("step-id")
	deadline_updateStepCmd.MarkFlagRequired("target-task-run-status")
	deadlineCmd.AddCommand(deadline_updateStepCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateJobCmd = &cobra.Command{
	Use:   "update-job",
	Short: "Updates a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateJobCmd).Standalone()

	deadline_updateJobCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_updateJobCmd.Flags().String("farm-id", "", "The farm ID of the job to update.")
	deadline_updateJobCmd.Flags().String("job-id", "", "The job ID to update.")
	deadline_updateJobCmd.Flags().String("lifecycle-status", "", "The status of a job in its lifecycle.")
	deadline_updateJobCmd.Flags().String("max-failed-tasks-count", "", "The number of task failures before the job stops running and is marked as `FAILED`.")
	deadline_updateJobCmd.Flags().String("max-retries-per-task", "", "The maximum number of retries for a job.")
	deadline_updateJobCmd.Flags().String("max-worker-count", "", "The maximum number of worker hosts that can concurrently process a job.")
	deadline_updateJobCmd.Flags().String("priority", "", "The job priority to update.")
	deadline_updateJobCmd.Flags().String("queue-id", "", "The queue ID of the job to update.")
	deadline_updateJobCmd.Flags().String("target-task-run-status", "", "The task status to update the job's tasks to.")
	deadline_updateJobCmd.MarkFlagRequired("farm-id")
	deadline_updateJobCmd.MarkFlagRequired("job-id")
	deadline_updateJobCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_updateJobCmd)
}

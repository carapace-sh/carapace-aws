package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Creates a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createJobCmd).Standalone()

	deadline_createJobCmd.Flags().String("attachments", "", "The attachments for the job.")
	deadline_createJobCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_createJobCmd.Flags().String("farm-id", "", "The farm ID of the farm to connect to the job.")
	deadline_createJobCmd.Flags().String("max-failed-tasks-count", "", "The number of task failures before the job stops running and is marked as `FAILED`.")
	deadline_createJobCmd.Flags().String("max-retries-per-task", "", "The maximum number of retries for each task.")
	deadline_createJobCmd.Flags().String("max-worker-count", "", "The maximum number of worker hosts that can concurrently process a job.")
	deadline_createJobCmd.Flags().String("parameters", "", "The parameters for the job.")
	deadline_createJobCmd.Flags().String("priority", "", "The priority of the job.")
	deadline_createJobCmd.Flags().String("queue-id", "", "The ID of the queue that the job is submitted to.")
	deadline_createJobCmd.Flags().String("source-job-id", "", "The job ID for the source job.")
	deadline_createJobCmd.Flags().String("storage-profile-id", "", "The storage profile ID for the storage profile to connect to the job.")
	deadline_createJobCmd.Flags().String("target-task-run-status", "", "The initial job status when it is created.")
	deadline_createJobCmd.Flags().String("template", "", "The job template to use for this job.")
	deadline_createJobCmd.Flags().String("template-type", "", "The file type for the job template.")
	deadline_createJobCmd.MarkFlagRequired("farm-id")
	deadline_createJobCmd.MarkFlagRequired("priority")
	deadline_createJobCmd.MarkFlagRequired("queue-id")
	deadlineCmd.AddCommand(deadline_createJobCmd)
}

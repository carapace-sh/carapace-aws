package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_createJobQueueCmd = &cobra.Command{
	Use:   "create-job-queue",
	Short: "Creates an Batch job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_createJobQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_createJobQueueCmd).Standalone()

		batch_createJobQueueCmd.Flags().String("compute-environment-order", "", "The set of compute environments mapped to a job queue and their order relative to each other.")
		batch_createJobQueueCmd.Flags().String("job-queue-name", "", "The name of the job queue.")
		batch_createJobQueueCmd.Flags().String("job-queue-type", "", "The type of job queue.")
		batch_createJobQueueCmd.Flags().String("job-state-time-limit-actions", "", "The set of actions that Batch performs on jobs that remain at the head of the job queue in the specified state longer than specified times.")
		batch_createJobQueueCmd.Flags().String("priority", "", "The priority of the job queue.")
		batch_createJobQueueCmd.Flags().String("scheduling-policy-arn", "", "The Amazon Resource Name (ARN) of the fair-share scheduling policy.")
		batch_createJobQueueCmd.Flags().String("service-environment-order", "", "A list of service environments that this job queue can use to allocate jobs.")
		batch_createJobQueueCmd.Flags().String("state", "", "The state of the job queue.")
		batch_createJobQueueCmd.Flags().String("tags", "", "The tags that you apply to the job queue to help you categorize and organize your resources.")
		batch_createJobQueueCmd.MarkFlagRequired("job-queue-name")
		batch_createJobQueueCmd.MarkFlagRequired("priority")
	})
	batchCmd.AddCommand(batch_createJobQueueCmd)
}

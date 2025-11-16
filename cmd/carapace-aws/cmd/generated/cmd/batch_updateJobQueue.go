package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_updateJobQueueCmd = &cobra.Command{
	Use:   "update-job-queue",
	Short: "Updates a job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_updateJobQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_updateJobQueueCmd).Standalone()

		batch_updateJobQueueCmd.Flags().String("compute-environment-order", "", "Details the set of compute environments mapped to a job queue and their order relative to each other.")
		batch_updateJobQueueCmd.Flags().String("job-queue", "", "The name or the Amazon Resource Name (ARN) of the job queue.")
		batch_updateJobQueueCmd.Flags().String("job-state-time-limit-actions", "", "The set of actions that Batch perform on jobs that remain at the head of the job queue in the specified state longer than specified times.")
		batch_updateJobQueueCmd.Flags().String("priority", "", "The priority of the job queue.")
		batch_updateJobQueueCmd.Flags().String("scheduling-policy-arn", "", "Amazon Resource Name (ARN) of the fair-share scheduling policy.")
		batch_updateJobQueueCmd.Flags().String("service-environment-order", "", "The order of the service environment associated with the job queue.")
		batch_updateJobQueueCmd.Flags().String("state", "", "Describes the queue's ability to accept new jobs.")
		batch_updateJobQueueCmd.MarkFlagRequired("job-queue")
	})
	batchCmd.AddCommand(batch_updateJobQueueCmd)
}

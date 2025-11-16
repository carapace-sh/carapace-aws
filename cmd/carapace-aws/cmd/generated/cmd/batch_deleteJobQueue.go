package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_deleteJobQueueCmd = &cobra.Command{
	Use:   "delete-job-queue",
	Short: "Deletes the specified job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_deleteJobQueueCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_deleteJobQueueCmd).Standalone()

		batch_deleteJobQueueCmd.Flags().String("job-queue", "", "The short name or full Amazon Resource Name (ARN) of the queue to delete.")
		batch_deleteJobQueueCmd.MarkFlagRequired("job-queue")
	})
	batchCmd.AddCommand(batch_deleteJobQueueCmd)
}

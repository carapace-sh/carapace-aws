package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_getJobQueueSnapshotCmd = &cobra.Command{
	Use:   "get-job-queue-snapshot",
	Short: "Provides a list of the first 100 `RUNNABLE` jobs associated to a single job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_getJobQueueSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_getJobQueueSnapshotCmd).Standalone()

		batch_getJobQueueSnapshotCmd.Flags().String("job-queue", "", "The job queue’s name or full queue Amazon Resource Name (ARN).")
		batch_getJobQueueSnapshotCmd.MarkFlagRequired("job-queue")
	})
	batchCmd.AddCommand(batch_getJobQueueSnapshotCmd)
}

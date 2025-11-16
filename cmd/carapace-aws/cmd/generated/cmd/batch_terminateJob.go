package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_terminateJobCmd = &cobra.Command{
	Use:   "terminate-job",
	Short: "Terminates a job in a job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_terminateJobCmd).Standalone()

	batch_terminateJobCmd.Flags().String("job-id", "", "The Batch job ID of the job to terminate.")
	batch_terminateJobCmd.Flags().String("reason", "", "A message to attach to the job that explains the reason for canceling it.")
	batch_terminateJobCmd.MarkFlagRequired("job-id")
	batch_terminateJobCmd.MarkFlagRequired("reason")
	batchCmd.AddCommand(batch_terminateJobCmd)
}

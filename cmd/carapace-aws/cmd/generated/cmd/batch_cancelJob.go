package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batch_cancelJobCmd = &cobra.Command{
	Use:   "cancel-job",
	Short: "Cancels a job in an Batch job queue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batch_cancelJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(batch_cancelJobCmd).Standalone()

		batch_cancelJobCmd.Flags().String("job-id", "", "The Batch job ID of the job to cancel.")
		batch_cancelJobCmd.Flags().String("reason", "", "A message to attach to the job that explains the reason for canceling it.")
		batch_cancelJobCmd.MarkFlagRequired("job-id")
		batch_cancelJobCmd.MarkFlagRequired("reason")
	})
	batchCmd.AddCommand(batch_cancelJobCmd)
}

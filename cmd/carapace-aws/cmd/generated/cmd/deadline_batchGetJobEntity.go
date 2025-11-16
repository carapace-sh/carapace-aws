package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_batchGetJobEntityCmd = &cobra.Command{
	Use:   "batch-get-job-entity",
	Short: "Get batched job details for a worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_batchGetJobEntityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_batchGetJobEntityCmd).Standalone()

		deadline_batchGetJobEntityCmd.Flags().String("farm-id", "", "The farm ID of the worker that's fetching job details.")
		deadline_batchGetJobEntityCmd.Flags().String("fleet-id", "", "The fleet ID of the worker that's fetching job details.")
		deadline_batchGetJobEntityCmd.Flags().String("identifiers", "", "The job identifiers to include within the job entity batch details.")
		deadline_batchGetJobEntityCmd.Flags().String("worker-id", "", "The worker ID of the worker containing the job details to get.")
		deadline_batchGetJobEntityCmd.MarkFlagRequired("farm-id")
		deadline_batchGetJobEntityCmd.MarkFlagRequired("fleet-id")
		deadline_batchGetJobEntityCmd.MarkFlagRequired("identifiers")
		deadline_batchGetJobEntityCmd.MarkFlagRequired("worker-id")
	})
	deadlineCmd.AddCommand(deadline_batchGetJobEntityCmd)
}

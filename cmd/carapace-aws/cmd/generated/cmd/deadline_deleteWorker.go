package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteWorkerCmd = &cobra.Command{
	Use:   "delete-worker",
	Short: "Deletes a worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteWorkerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_deleteWorkerCmd).Standalone()

		deadline_deleteWorkerCmd.Flags().String("farm-id", "", "The farm ID of the worker to delete.")
		deadline_deleteWorkerCmd.Flags().String("fleet-id", "", "The fleet ID of the worker to delete.")
		deadline_deleteWorkerCmd.Flags().String("worker-id", "", "The worker ID of the worker to delete.")
		deadline_deleteWorkerCmd.MarkFlagRequired("farm-id")
		deadline_deleteWorkerCmd.MarkFlagRequired("fleet-id")
		deadline_deleteWorkerCmd.MarkFlagRequired("worker-id")
	})
	deadlineCmd.AddCommand(deadline_deleteWorkerCmd)
}

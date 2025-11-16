package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getWorkerCmd = &cobra.Command{
	Use:   "get-worker",
	Short: "Gets a worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getWorkerCmd).Standalone()

	deadline_getWorkerCmd.Flags().String("farm-id", "", "The farm ID for the worker.")
	deadline_getWorkerCmd.Flags().String("fleet-id", "", "The fleet ID of the worker.")
	deadline_getWorkerCmd.Flags().String("worker-id", "", "The worker ID.")
	deadline_getWorkerCmd.MarkFlagRequired("farm-id")
	deadline_getWorkerCmd.MarkFlagRequired("fleet-id")
	deadline_getWorkerCmd.MarkFlagRequired("worker-id")
	deadlineCmd.AddCommand(deadline_getWorkerCmd)
}

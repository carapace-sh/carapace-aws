package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateWorkerCmd = &cobra.Command{
	Use:   "update-worker",
	Short: "Updates a worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateWorkerCmd).Standalone()

	deadline_updateWorkerCmd.Flags().String("capabilities", "", "The worker capabilities to update.")
	deadline_updateWorkerCmd.Flags().String("farm-id", "", "The farm ID to update.")
	deadline_updateWorkerCmd.Flags().String("fleet-id", "", "The fleet ID to update.")
	deadline_updateWorkerCmd.Flags().String("host-properties", "", "The host properties to update.")
	deadline_updateWorkerCmd.Flags().String("status", "", "The worker status to update.")
	deadline_updateWorkerCmd.Flags().String("worker-id", "", "The worker ID to update.")
	deadline_updateWorkerCmd.MarkFlagRequired("farm-id")
	deadline_updateWorkerCmd.MarkFlagRequired("fleet-id")
	deadline_updateWorkerCmd.MarkFlagRequired("worker-id")
	deadlineCmd.AddCommand(deadline_updateWorkerCmd)
}

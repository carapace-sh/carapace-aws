package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_assumeQueueRoleForWorkerCmd = &cobra.Command{
	Use:   "assume-queue-role-for-worker",
	Short: "Allows a worker to assume a queue role.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_assumeQueueRoleForWorkerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_assumeQueueRoleForWorkerCmd).Standalone()

		deadline_assumeQueueRoleForWorkerCmd.Flags().String("farm-id", "", "The farm ID of the worker assuming the queue role.")
		deadline_assumeQueueRoleForWorkerCmd.Flags().String("fleet-id", "", "The fleet ID of the worker assuming the queue role.")
		deadline_assumeQueueRoleForWorkerCmd.Flags().String("queue-id", "", "The queue ID of the worker assuming the queue role.")
		deadline_assumeQueueRoleForWorkerCmd.Flags().String("worker-id", "", "The worker ID of the worker assuming the queue role.")
		deadline_assumeQueueRoleForWorkerCmd.MarkFlagRequired("farm-id")
		deadline_assumeQueueRoleForWorkerCmd.MarkFlagRequired("fleet-id")
		deadline_assumeQueueRoleForWorkerCmd.MarkFlagRequired("queue-id")
		deadline_assumeQueueRoleForWorkerCmd.MarkFlagRequired("worker-id")
	})
	deadlineCmd.AddCommand(deadline_assumeQueueRoleForWorkerCmd)
}

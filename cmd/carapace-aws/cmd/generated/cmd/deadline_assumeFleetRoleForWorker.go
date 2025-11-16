package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_assumeFleetRoleForWorkerCmd = &cobra.Command{
	Use:   "assume-fleet-role-for-worker",
	Short: "Get credentials from the fleet role for a worker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_assumeFleetRoleForWorkerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_assumeFleetRoleForWorkerCmd).Standalone()

		deadline_assumeFleetRoleForWorkerCmd.Flags().String("farm-id", "", "The farm ID for the fleet's farm.")
		deadline_assumeFleetRoleForWorkerCmd.Flags().String("fleet-id", "", "The fleet ID that contains the worker.")
		deadline_assumeFleetRoleForWorkerCmd.Flags().String("worker-id", "", "The ID of the worker assuming the fleet role.")
		deadline_assumeFleetRoleForWorkerCmd.MarkFlagRequired("farm-id")
		deadline_assumeFleetRoleForWorkerCmd.MarkFlagRequired("fleet-id")
		deadline_assumeFleetRoleForWorkerCmd.MarkFlagRequired("worker-id")
	})
	deadlineCmd.AddCommand(deadline_assumeFleetRoleForWorkerCmd)
}

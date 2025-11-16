package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deregisterTaskFromMaintenanceWindowCmd = &cobra.Command{
	Use:   "deregister-task-from-maintenance-window",
	Short: "Removes a task from a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deregisterTaskFromMaintenanceWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deregisterTaskFromMaintenanceWindowCmd).Standalone()

		ssm_deregisterTaskFromMaintenanceWindowCmd.Flags().String("window-id", "", "The ID of the maintenance window the task should be removed from.")
		ssm_deregisterTaskFromMaintenanceWindowCmd.Flags().String("window-task-id", "", "The ID of the task to remove from the maintenance window.")
		ssm_deregisterTaskFromMaintenanceWindowCmd.MarkFlagRequired("window-id")
		ssm_deregisterTaskFromMaintenanceWindowCmd.MarkFlagRequired("window-task-id")
	})
	ssmCmd.AddCommand(ssm_deregisterTaskFromMaintenanceWindowCmd)
}

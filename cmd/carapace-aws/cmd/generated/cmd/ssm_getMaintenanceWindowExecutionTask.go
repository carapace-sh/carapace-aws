package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getMaintenanceWindowExecutionTaskCmd = &cobra.Command{
	Use:   "get-maintenance-window-execution-task",
	Short: "Retrieves the details about a specific task run as part of a maintenance window execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getMaintenanceWindowExecutionTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getMaintenanceWindowExecutionTaskCmd).Standalone()

		ssm_getMaintenanceWindowExecutionTaskCmd.Flags().String("task-id", "", "The ID of the specific task execution in the maintenance window task that should be retrieved.")
		ssm_getMaintenanceWindowExecutionTaskCmd.Flags().String("window-execution-id", "", "The ID of the maintenance window execution that includes the task.")
		ssm_getMaintenanceWindowExecutionTaskCmd.MarkFlagRequired("task-id")
		ssm_getMaintenanceWindowExecutionTaskCmd.MarkFlagRequired("window-execution-id")
	})
	ssmCmd.AddCommand(ssm_getMaintenanceWindowExecutionTaskCmd)
}

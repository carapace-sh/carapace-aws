package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getMaintenanceWindowExecutionTaskInvocationCmd = &cobra.Command{
	Use:   "get-maintenance-window-execution-task-invocation",
	Short: "Retrieves information about a specific task running on a specific target.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getMaintenanceWindowExecutionTaskInvocationCmd).Standalone()

	ssm_getMaintenanceWindowExecutionTaskInvocationCmd.Flags().String("invocation-id", "", "The invocation ID to retrieve.")
	ssm_getMaintenanceWindowExecutionTaskInvocationCmd.Flags().String("task-id", "", "The ID of the specific task in the maintenance window task that should be retrieved.")
	ssm_getMaintenanceWindowExecutionTaskInvocationCmd.Flags().String("window-execution-id", "", "The ID of the maintenance window execution for which the task is a part.")
	ssm_getMaintenanceWindowExecutionTaskInvocationCmd.MarkFlagRequired("invocation-id")
	ssm_getMaintenanceWindowExecutionTaskInvocationCmd.MarkFlagRequired("task-id")
	ssm_getMaintenanceWindowExecutionTaskInvocationCmd.MarkFlagRequired("window-execution-id")
	ssmCmd.AddCommand(ssm_getMaintenanceWindowExecutionTaskInvocationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd = &cobra.Command{
	Use:   "describe-maintenance-window-execution-task-invocations",
	Short: "Retrieves the individual task executions (one per target) for a particular task run as part of a maintenance window execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd).Standalone()

		ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd.Flags().String("filters", "", "Optional filters used to scope down the returned task invocations.")
		ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd.Flags().String("task-id", "", "The ID of the specific task in the maintenance window task that should be retrieved.")
		ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd.Flags().String("window-execution-id", "", "The ID of the maintenance window execution the task is part of.")
		ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd.MarkFlagRequired("task-id")
		ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd.MarkFlagRequired("window-execution-id")
	})
	ssmCmd.AddCommand(ssm_describeMaintenanceWindowExecutionTaskInvocationsCmd)
}

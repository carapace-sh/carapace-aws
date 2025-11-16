package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeMaintenanceWindowExecutionTasksCmd = &cobra.Command{
	Use:   "describe-maintenance-window-execution-tasks",
	Short: "For a given maintenance window execution, lists the tasks that were run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeMaintenanceWindowExecutionTasksCmd).Standalone()

	ssm_describeMaintenanceWindowExecutionTasksCmd.Flags().String("filters", "", "Optional filters used to scope down the returned tasks.")
	ssm_describeMaintenanceWindowExecutionTasksCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeMaintenanceWindowExecutionTasksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describeMaintenanceWindowExecutionTasksCmd.Flags().String("window-execution-id", "", "The ID of the maintenance window execution whose task executions should be retrieved.")
	ssm_describeMaintenanceWindowExecutionTasksCmd.MarkFlagRequired("window-execution-id")
	ssmCmd.AddCommand(ssm_describeMaintenanceWindowExecutionTasksCmd)
}

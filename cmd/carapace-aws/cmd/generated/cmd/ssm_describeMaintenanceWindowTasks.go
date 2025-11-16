package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeMaintenanceWindowTasksCmd = &cobra.Command{
	Use:   "describe-maintenance-window-tasks",
	Short: "Lists the tasks in a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeMaintenanceWindowTasksCmd).Standalone()

	ssm_describeMaintenanceWindowTasksCmd.Flags().String("filters", "", "Optional filters used to narrow down the scope of the returned tasks.")
	ssm_describeMaintenanceWindowTasksCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeMaintenanceWindowTasksCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describeMaintenanceWindowTasksCmd.Flags().String("window-id", "", "The ID of the maintenance window whose tasks should be retrieved.")
	ssm_describeMaintenanceWindowTasksCmd.MarkFlagRequired("window-id")
	ssmCmd.AddCommand(ssm_describeMaintenanceWindowTasksCmd)
}

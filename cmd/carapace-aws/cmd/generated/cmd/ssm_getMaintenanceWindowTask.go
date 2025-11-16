package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getMaintenanceWindowTaskCmd = &cobra.Command{
	Use:   "get-maintenance-window-task",
	Short: "Retrieves the details of a maintenance window task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getMaintenanceWindowTaskCmd).Standalone()

	ssm_getMaintenanceWindowTaskCmd.Flags().String("window-id", "", "The maintenance window ID that includes the task to retrieve.")
	ssm_getMaintenanceWindowTaskCmd.Flags().String("window-task-id", "", "The maintenance window task ID to retrieve.")
	ssm_getMaintenanceWindowTaskCmd.MarkFlagRequired("window-id")
	ssm_getMaintenanceWindowTaskCmd.MarkFlagRequired("window-task-id")
	ssmCmd.AddCommand(ssm_getMaintenanceWindowTaskCmd)
}

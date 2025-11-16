package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getMaintenanceWindowExecutionCmd = &cobra.Command{
	Use:   "get-maintenance-window-execution",
	Short: "Retrieves details about a specific a maintenance window execution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getMaintenanceWindowExecutionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getMaintenanceWindowExecutionCmd).Standalone()

		ssm_getMaintenanceWindowExecutionCmd.Flags().String("window-execution-id", "", "The ID of the maintenance window execution that includes the task.")
		ssm_getMaintenanceWindowExecutionCmd.MarkFlagRequired("window-execution-id")
	})
	ssmCmd.AddCommand(ssm_getMaintenanceWindowExecutionCmd)
}

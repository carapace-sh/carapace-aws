package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_cancelMaintenanceWindowExecutionCmd = &cobra.Command{
	Use:   "cancel-maintenance-window-execution",
	Short: "Stops a maintenance window execution that is already in progress and cancels any tasks in the window that haven't already starting running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_cancelMaintenanceWindowExecutionCmd).Standalone()

	ssm_cancelMaintenanceWindowExecutionCmd.Flags().String("window-execution-id", "", "The ID of the maintenance window execution to stop.")
	ssm_cancelMaintenanceWindowExecutionCmd.MarkFlagRequired("window-execution-id")
	ssmCmd.AddCommand(ssm_cancelMaintenanceWindowExecutionCmd)
}

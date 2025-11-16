package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deregisterTargetFromMaintenanceWindowCmd = &cobra.Command{
	Use:   "deregister-target-from-maintenance-window",
	Short: "Removes a target from a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deregisterTargetFromMaintenanceWindowCmd).Standalone()

	ssm_deregisterTargetFromMaintenanceWindowCmd.Flags().Bool("no-safe", false, "The system checks if the target is being referenced by a task.")
	ssm_deregisterTargetFromMaintenanceWindowCmd.Flags().Bool("safe", false, "The system checks if the target is being referenced by a task.")
	ssm_deregisterTargetFromMaintenanceWindowCmd.Flags().String("window-id", "", "The ID of the maintenance window the target should be removed from.")
	ssm_deregisterTargetFromMaintenanceWindowCmd.Flags().String("window-target-id", "", "The ID of the target definition to remove.")
	ssm_deregisterTargetFromMaintenanceWindowCmd.Flag("no-safe").Hidden = true
	ssm_deregisterTargetFromMaintenanceWindowCmd.MarkFlagRequired("window-id")
	ssm_deregisterTargetFromMaintenanceWindowCmd.MarkFlagRequired("window-target-id")
	ssmCmd.AddCommand(ssm_deregisterTargetFromMaintenanceWindowCmd)
}

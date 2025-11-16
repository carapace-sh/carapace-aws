package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_registerTargetWithMaintenanceWindowCmd = &cobra.Command{
	Use:   "register-target-with-maintenance-window",
	Short: "Registers a target with a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_registerTargetWithMaintenanceWindowCmd).Standalone()

	ssm_registerTargetWithMaintenanceWindowCmd.Flags().String("client-token", "", "User-provided idempotency token.")
	ssm_registerTargetWithMaintenanceWindowCmd.Flags().String("description", "", "An optional description for the target.")
	ssm_registerTargetWithMaintenanceWindowCmd.Flags().String("name", "", "An optional name for the target.")
	ssm_registerTargetWithMaintenanceWindowCmd.Flags().String("owner-information", "", "User-provided value that will be included in any Amazon CloudWatch Events events raised while running tasks for these targets in this maintenance window.")
	ssm_registerTargetWithMaintenanceWindowCmd.Flags().String("resource-type", "", "The type of target being registered with the maintenance window.")
	ssm_registerTargetWithMaintenanceWindowCmd.Flags().String("targets", "", "The targets to register with the maintenance window.")
	ssm_registerTargetWithMaintenanceWindowCmd.Flags().String("window-id", "", "The ID of the maintenance window the target should be registered with.")
	ssm_registerTargetWithMaintenanceWindowCmd.MarkFlagRequired("resource-type")
	ssm_registerTargetWithMaintenanceWindowCmd.MarkFlagRequired("targets")
	ssm_registerTargetWithMaintenanceWindowCmd.MarkFlagRequired("window-id")
	ssmCmd.AddCommand(ssm_registerTargetWithMaintenanceWindowCmd)
}

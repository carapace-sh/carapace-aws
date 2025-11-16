package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_applyPendingMaintenanceActionCmd = &cobra.Command{
	Use:   "apply-pending-maintenance-action",
	Short: "Applies a pending maintenance action to a resource (for example, to a DB instance).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_applyPendingMaintenanceActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rds_applyPendingMaintenanceActionCmd).Standalone()

		rds_applyPendingMaintenanceActionCmd.Flags().String("apply-action", "", "The pending maintenance action to apply to this resource.")
		rds_applyPendingMaintenanceActionCmd.Flags().String("opt-in-type", "", "A value that specifies the type of opt-in request, or undoes an opt-in request.")
		rds_applyPendingMaintenanceActionCmd.Flags().String("resource-identifier", "", "The RDS Amazon Resource Name (ARN) of the resource that the pending maintenance action applies to.")
		rds_applyPendingMaintenanceActionCmd.MarkFlagRequired("apply-action")
		rds_applyPendingMaintenanceActionCmd.MarkFlagRequired("opt-in-type")
		rds_applyPendingMaintenanceActionCmd.MarkFlagRequired("resource-identifier")
	})
	rdsCmd.AddCommand(rds_applyPendingMaintenanceActionCmd)
}

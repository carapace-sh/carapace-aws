package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_applyPendingMaintenanceActionCmd = &cobra.Command{
	Use:   "apply-pending-maintenance-action",
	Short: "Applies a pending maintenance action to a resource (for example, to a DB instance).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_applyPendingMaintenanceActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_applyPendingMaintenanceActionCmd).Standalone()

		neptune_applyPendingMaintenanceActionCmd.Flags().String("apply-action", "", "The pending maintenance action to apply to this resource.")
		neptune_applyPendingMaintenanceActionCmd.Flags().String("opt-in-type", "", "A value that specifies the type of opt-in request, or undoes an opt-in request.")
		neptune_applyPendingMaintenanceActionCmd.Flags().String("resource-identifier", "", "The Amazon Resource Name (ARN) of the resource that the pending maintenance action applies to.")
		neptune_applyPendingMaintenanceActionCmd.MarkFlagRequired("apply-action")
		neptune_applyPendingMaintenanceActionCmd.MarkFlagRequired("opt-in-type")
		neptune_applyPendingMaintenanceActionCmd.MarkFlagRequired("resource-identifier")
	})
	neptuneCmd.AddCommand(neptune_applyPendingMaintenanceActionCmd)
}

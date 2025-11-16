package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_applyPendingMaintenanceActionCmd = &cobra.Command{
	Use:   "apply-pending-maintenance-action",
	Short: "Applies a pending maintenance action to a resource (for example, to an Amazon DocumentDB instance).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_applyPendingMaintenanceActionCmd).Standalone()

	docdb_applyPendingMaintenanceActionCmd.Flags().String("apply-action", "", "The pending maintenance action to apply to this resource.")
	docdb_applyPendingMaintenanceActionCmd.Flags().String("opt-in-type", "", "A value that specifies the type of opt-in request or undoes an opt-in request.")
	docdb_applyPendingMaintenanceActionCmd.Flags().String("resource-identifier", "", "The Amazon Resource Name (ARN) of the resource that the pending maintenance action applies to.")
	docdb_applyPendingMaintenanceActionCmd.MarkFlagRequired("apply-action")
	docdb_applyPendingMaintenanceActionCmd.MarkFlagRequired("opt-in-type")
	docdb_applyPendingMaintenanceActionCmd.MarkFlagRequired("resource-identifier")
	docdbCmd.AddCommand(docdb_applyPendingMaintenanceActionCmd)
}

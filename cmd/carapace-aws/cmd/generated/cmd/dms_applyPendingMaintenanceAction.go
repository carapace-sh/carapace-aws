package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_applyPendingMaintenanceActionCmd = &cobra.Command{
	Use:   "apply-pending-maintenance-action",
	Short: "Applies a pending maintenance action to a resource (for example, to a replication instance).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_applyPendingMaintenanceActionCmd).Standalone()

	dms_applyPendingMaintenanceActionCmd.Flags().String("apply-action", "", "The pending maintenance action to apply to this resource.")
	dms_applyPendingMaintenanceActionCmd.Flags().String("opt-in-type", "", "A value that specifies the type of opt-in request, or undoes an opt-in request.")
	dms_applyPendingMaintenanceActionCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the DMS resource that the pending maintenance action applies to.")
	dms_applyPendingMaintenanceActionCmd.MarkFlagRequired("apply-action")
	dms_applyPendingMaintenanceActionCmd.MarkFlagRequired("opt-in-type")
	dms_applyPendingMaintenanceActionCmd.MarkFlagRequired("replication-instance-arn")
	dmsCmd.AddCommand(dms_applyPendingMaintenanceActionCmd)
}

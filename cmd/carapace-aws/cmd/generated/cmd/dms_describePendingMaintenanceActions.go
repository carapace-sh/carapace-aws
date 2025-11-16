package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describePendingMaintenanceActionsCmd = &cobra.Command{
	Use:   "describe-pending-maintenance-actions",
	Short: "Returns a list of upcoming maintenance events for replication instances in your account in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describePendingMaintenanceActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describePendingMaintenanceActionsCmd).Standalone()

		dms_describePendingMaintenanceActionsCmd.Flags().String("filters", "", "")
		dms_describePendingMaintenanceActionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		dms_describePendingMaintenanceActionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		dms_describePendingMaintenanceActionsCmd.Flags().String("replication-instance-arn", "", "The Amazon Resource Name (ARN) of the replication instance.")
	})
	dmsCmd.AddCommand(dms_describePendingMaintenanceActionsCmd)
}

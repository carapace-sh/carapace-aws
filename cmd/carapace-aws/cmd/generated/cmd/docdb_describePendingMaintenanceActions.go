package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describePendingMaintenanceActionsCmd = &cobra.Command{
	Use:   "describe-pending-maintenance-actions",
	Short: "Returns a list of resources (for example, instances) that have at least one pending maintenance action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describePendingMaintenanceActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdb_describePendingMaintenanceActionsCmd).Standalone()

		docdb_describePendingMaintenanceActionsCmd.Flags().String("filters", "", "A filter that specifies one or more resources to return pending maintenance actions for.")
		docdb_describePendingMaintenanceActionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
		docdb_describePendingMaintenanceActionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		docdb_describePendingMaintenanceActionsCmd.Flags().String("resource-identifier", "", "The ARN of a resource to return pending maintenance actions for.")
	})
	docdbCmd.AddCommand(docdb_describePendingMaintenanceActionsCmd)
}

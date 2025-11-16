package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describePendingMaintenanceActionsCmd = &cobra.Command{
	Use:   "describe-pending-maintenance-actions",
	Short: "Returns a list of resources (for example, DB instances) that have at least one pending maintenance action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describePendingMaintenanceActionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describePendingMaintenanceActionsCmd).Standalone()

		neptune_describePendingMaintenanceActionsCmd.Flags().String("filters", "", "A filter that specifies one or more resources to return pending maintenance actions for.")
		neptune_describePendingMaintenanceActionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous `DescribePendingMaintenanceActions` request.")
		neptune_describePendingMaintenanceActionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		neptune_describePendingMaintenanceActionsCmd.Flags().String("resource-identifier", "", "The ARN of a resource to return pending maintenance actions for.")
	})
	neptuneCmd.AddCommand(neptune_describePendingMaintenanceActionsCmd)
}

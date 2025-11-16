package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_listPendingMaintenanceActionsCmd = &cobra.Command{
	Use:   "list-pending-maintenance-actions",
	Short: "Retrieves a list of all maintenance actions that are pending.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_listPendingMaintenanceActionsCmd).Standalone()

	docdbElastic_listPendingMaintenanceActionsCmd.Flags().String("max-results", "", "The maximum number of results to include in the response.")
	docdbElastic_listPendingMaintenanceActionsCmd.Flags().String("next-token", "", "An optional pagination token provided by a previous request.")
	docdbElasticCmd.AddCommand(docdbElastic_listPendingMaintenanceActionsCmd)
}

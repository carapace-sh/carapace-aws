package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_getPendingMaintenanceActionCmd = &cobra.Command{
	Use:   "get-pending-maintenance-action",
	Short: "Retrieves all maintenance actions that are pending.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_getPendingMaintenanceActionCmd).Standalone()

	docdbElastic_getPendingMaintenanceActionCmd.Flags().String("resource-arn", "", "Retrieves pending maintenance actions for a specific Amazon Resource Name (ARN).")
	docdbElastic_getPendingMaintenanceActionCmd.MarkFlagRequired("resource-arn")
	docdbElasticCmd.AddCommand(docdbElastic_getPendingMaintenanceActionCmd)
}

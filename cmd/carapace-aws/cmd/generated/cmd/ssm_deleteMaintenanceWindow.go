package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_deleteMaintenanceWindowCmd = &cobra.Command{
	Use:   "delete-maintenance-window",
	Short: "Deletes a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_deleteMaintenanceWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_deleteMaintenanceWindowCmd).Standalone()

		ssm_deleteMaintenanceWindowCmd.Flags().String("window-id", "", "The ID of the maintenance window to delete.")
		ssm_deleteMaintenanceWindowCmd.MarkFlagRequired("window-id")
	})
	ssmCmd.AddCommand(ssm_deleteMaintenanceWindowCmd)
}

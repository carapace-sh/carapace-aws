package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_getMaintenanceWindowCmd = &cobra.Command{
	Use:   "get-maintenance-window",
	Short: "Retrieves a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_getMaintenanceWindowCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_getMaintenanceWindowCmd).Standalone()

		ssm_getMaintenanceWindowCmd.Flags().String("window-id", "", "The ID of the maintenance window for which you want to retrieve information.")
		ssm_getMaintenanceWindowCmd.MarkFlagRequired("window-id")
	})
	ssmCmd.AddCommand(ssm_getMaintenanceWindowCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeMaintenanceWindowsForTargetCmd = &cobra.Command{
	Use:   "describe-maintenance-windows-for-target",
	Short: "Retrieves information about the maintenance window targets or tasks that a managed node is associated with.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeMaintenanceWindowsForTargetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeMaintenanceWindowsForTargetCmd).Standalone()

		ssm_describeMaintenanceWindowsForTargetCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeMaintenanceWindowsForTargetCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeMaintenanceWindowsForTargetCmd.Flags().String("resource-type", "", "The type of resource you want to retrieve information about.")
		ssm_describeMaintenanceWindowsForTargetCmd.Flags().String("targets", "", "The managed node ID or key-value pair to retrieve information about.")
		ssm_describeMaintenanceWindowsForTargetCmd.MarkFlagRequired("resource-type")
		ssm_describeMaintenanceWindowsForTargetCmd.MarkFlagRequired("targets")
	})
	ssmCmd.AddCommand(ssm_describeMaintenanceWindowsForTargetCmd)
}

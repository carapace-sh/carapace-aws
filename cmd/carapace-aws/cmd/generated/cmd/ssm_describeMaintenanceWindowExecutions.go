package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeMaintenanceWindowExecutionsCmd = &cobra.Command{
	Use:   "describe-maintenance-window-executions",
	Short: "Lists the executions of a maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeMaintenanceWindowExecutionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeMaintenanceWindowExecutionsCmd).Standalone()

		ssm_describeMaintenanceWindowExecutionsCmd.Flags().String("filters", "", "Each entry in the array is a structure containing:")
		ssm_describeMaintenanceWindowExecutionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeMaintenanceWindowExecutionsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeMaintenanceWindowExecutionsCmd.Flags().String("window-id", "", "The ID of the maintenance window whose executions should be retrieved.")
		ssm_describeMaintenanceWindowExecutionsCmd.MarkFlagRequired("window-id")
	})
	ssmCmd.AddCommand(ssm_describeMaintenanceWindowExecutionsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeMaintenanceWindowsCmd = &cobra.Command{
	Use:   "describe-maintenance-windows",
	Short: "Retrieves the maintenance windows in an Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeMaintenanceWindowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeMaintenanceWindowsCmd).Standalone()

		ssm_describeMaintenanceWindowsCmd.Flags().String("filters", "", "Optional filters used to narrow down the scope of the returned maintenance windows.")
		ssm_describeMaintenanceWindowsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeMaintenanceWindowsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	ssmCmd.AddCommand(ssm_describeMaintenanceWindowsCmd)
}

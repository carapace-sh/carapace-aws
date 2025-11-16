package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeMaintenanceWindowTargetsCmd = &cobra.Command{
	Use:   "describe-maintenance-window-targets",
	Short: "Lists the targets registered with the maintenance window.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeMaintenanceWindowTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeMaintenanceWindowTargetsCmd).Standalone()

		ssm_describeMaintenanceWindowTargetsCmd.Flags().String("filters", "", "Optional filters that can be used to narrow down the scope of the returned window targets.")
		ssm_describeMaintenanceWindowTargetsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
		ssm_describeMaintenanceWindowTargetsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		ssm_describeMaintenanceWindowTargetsCmd.Flags().String("window-id", "", "The ID of the maintenance window whose targets should be retrieved.")
		ssm_describeMaintenanceWindowTargetsCmd.MarkFlagRequired("window-id")
	})
	ssmCmd.AddCommand(ssm_describeMaintenanceWindowTargetsCmd)
}

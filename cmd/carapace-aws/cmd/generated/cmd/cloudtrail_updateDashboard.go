package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_updateDashboardCmd = &cobra.Command{
	Use:   "update-dashboard",
	Short: "Updates the specified dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_updateDashboardCmd).Standalone()

	cloudtrail_updateDashboardCmd.Flags().String("dashboard-id", "", "The name or ARN of the dashboard.")
	cloudtrail_updateDashboardCmd.Flags().String("refresh-schedule", "", "The refresh schedule configuration for the dashboard.")
	cloudtrail_updateDashboardCmd.Flags().String("termination-protection-enabled", "", "Specifies whether termination protection is enabled for the dashboard.")
	cloudtrail_updateDashboardCmd.Flags().String("widgets", "", "An array of widgets for the dashboard.")
	cloudtrail_updateDashboardCmd.MarkFlagRequired("dashboard-id")
	cloudtrailCmd.AddCommand(cloudtrail_updateDashboardCmd)
}

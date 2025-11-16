package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDashboardPermissionsCmd = &cobra.Command{
	Use:   "describe-dashboard-permissions",
	Short: "Describes read and write permissions for a dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDashboardPermissionsCmd).Standalone()

	quicksight_describeDashboardPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard that you're describing permissions for.")
	quicksight_describeDashboardPermissionsCmd.Flags().String("dashboard-id", "", "The ID for the dashboard, also added to the IAM policy.")
	quicksight_describeDashboardPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeDashboardPermissionsCmd.MarkFlagRequired("dashboard-id")
	quicksightCmd.AddCommand(quicksight_describeDashboardPermissionsCmd)
}

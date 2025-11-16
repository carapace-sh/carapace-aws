package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDashboardPermissionsCmd = &cobra.Command{
	Use:   "update-dashboard-permissions",
	Short: "Updates read and write permissions on a dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDashboardPermissionsCmd).Standalone()

	quicksight_updateDashboardPermissionsCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard whose permissions you're updating.")
	quicksight_updateDashboardPermissionsCmd.Flags().String("dashboard-id", "", "The ID for the dashboard.")
	quicksight_updateDashboardPermissionsCmd.Flags().String("grant-link-permissions", "", "Grants link permissions to all users in a defined namespace.")
	quicksight_updateDashboardPermissionsCmd.Flags().String("grant-permissions", "", "The permissions that you want to grant on this resource.")
	quicksight_updateDashboardPermissionsCmd.Flags().String("revoke-link-permissions", "", "Revokes link permissions from all users in a defined namespace.")
	quicksight_updateDashboardPermissionsCmd.Flags().String("revoke-permissions", "", "The permissions that you want to revoke from this resource.")
	quicksight_updateDashboardPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateDashboardPermissionsCmd.MarkFlagRequired("dashboard-id")
	quicksightCmd.AddCommand(quicksight_updateDashboardPermissionsCmd)
}

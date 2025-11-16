package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteDashboardCmd = &cobra.Command{
	Use:   "delete-dashboard",
	Short: "Deletes a dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteDashboardCmd).Standalone()

		quicksight_deleteDashboardCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard that you're deleting.")
		quicksight_deleteDashboardCmd.Flags().String("dashboard-id", "", "The ID for the dashboard.")
		quicksight_deleteDashboardCmd.Flags().String("version-number", "", "The version number of the dashboard.")
		quicksight_deleteDashboardCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteDashboardCmd.MarkFlagRequired("dashboard-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteDashboardCmd)
}

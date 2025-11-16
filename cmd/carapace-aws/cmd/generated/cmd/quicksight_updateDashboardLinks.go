package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDashboardLinksCmd = &cobra.Command{
	Use:   "update-dashboard-links",
	Short: "Updates the linked analyses on a dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDashboardLinksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateDashboardLinksCmd).Standalone()

		quicksight_updateDashboardLinksCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard whose links you want to update.")
		quicksight_updateDashboardLinksCmd.Flags().String("dashboard-id", "", "The ID for the dashboard.")
		quicksight_updateDashboardLinksCmd.Flags().String("link-entities", "", "list of analysis Amazon Resource Names (ARNs) to be linked to the dashboard.")
		quicksight_updateDashboardLinksCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateDashboardLinksCmd.MarkFlagRequired("dashboard-id")
		quicksight_updateDashboardLinksCmd.MarkFlagRequired("link-entities")
	})
	quicksightCmd.AddCommand(quicksight_updateDashboardLinksCmd)
}

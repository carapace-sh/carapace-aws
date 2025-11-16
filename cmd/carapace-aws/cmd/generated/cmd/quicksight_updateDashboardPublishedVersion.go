package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDashboardPublishedVersionCmd = &cobra.Command{
	Use:   "update-dashboard-published-version",
	Short: "Updates the published version of a dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDashboardPublishedVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_updateDashboardPublishedVersionCmd).Standalone()

		quicksight_updateDashboardPublishedVersionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard that you're updating.")
		quicksight_updateDashboardPublishedVersionCmd.Flags().String("dashboard-id", "", "The ID for the dashboard.")
		quicksight_updateDashboardPublishedVersionCmd.Flags().String("version-number", "", "The version number of the dashboard.")
		quicksight_updateDashboardPublishedVersionCmd.MarkFlagRequired("aws-account-id")
		quicksight_updateDashboardPublishedVersionCmd.MarkFlagRequired("dashboard-id")
		quicksight_updateDashboardPublishedVersionCmd.MarkFlagRequired("version-number")
	})
	quicksightCmd.AddCommand(quicksight_updateDashboardPublishedVersionCmd)
}

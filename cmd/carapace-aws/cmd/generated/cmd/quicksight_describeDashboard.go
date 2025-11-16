package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDashboardCmd = &cobra.Command{
	Use:   "describe-dashboard",
	Short: "Provides a summary for a dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDashboardCmd).Standalone()

	quicksight_describeDashboardCmd.Flags().String("alias-name", "", "The alias name.")
	quicksight_describeDashboardCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard that you're describing.")
	quicksight_describeDashboardCmd.Flags().String("dashboard-id", "", "The ID for the dashboard.")
	quicksight_describeDashboardCmd.Flags().String("version-number", "", "The version number for the dashboard.")
	quicksight_describeDashboardCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeDashboardCmd.MarkFlagRequired("dashboard-id")
	quicksightCmd.AddCommand(quicksight_describeDashboardCmd)
}

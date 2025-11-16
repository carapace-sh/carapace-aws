package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDashboardDefinitionCmd = &cobra.Command{
	Use:   "describe-dashboard-definition",
	Short: "Provides a detailed description of the definition of a dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDashboardDefinitionCmd).Standalone()

	quicksight_describeDashboardDefinitionCmd.Flags().String("alias-name", "", "The alias name.")
	quicksight_describeDashboardDefinitionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard that you're describing.")
	quicksight_describeDashboardDefinitionCmd.Flags().String("dashboard-id", "", "The ID for the dashboard.")
	quicksight_describeDashboardDefinitionCmd.Flags().String("version-number", "", "The version number for the dashboard.")
	quicksight_describeDashboardDefinitionCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeDashboardDefinitionCmd.MarkFlagRequired("dashboard-id")
	quicksightCmd.AddCommand(quicksight_describeDashboardDefinitionCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateDashboardsQaconfigurationCmd = &cobra.Command{
	Use:   "update-dashboards-qaconfiguration",
	Short: "Updates a Dashboard QA configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateDashboardsQaconfigurationCmd).Standalone()

	quicksight_updateDashboardsQaconfigurationCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the dashboard QA configuration that you want to update.")
	quicksight_updateDashboardsQaconfigurationCmd.Flags().String("dashboards-qastatus", "", "The status of dashboards QA configuration that you want to update.")
	quicksight_updateDashboardsQaconfigurationCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateDashboardsQaconfigurationCmd.MarkFlagRequired("dashboards-qastatus")
	quicksightCmd.AddCommand(quicksight_updateDashboardsQaconfigurationCmd)
}

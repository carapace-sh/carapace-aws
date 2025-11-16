package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_updateDashboardCmd = &cobra.Command{
	Use:   "update-dashboard",
	Short: "Updates an existing dashboard's properties, including its name, description, and widget configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_updateDashboardCmd).Standalone()

	bcmDashboards_updateDashboardCmd.Flags().String("arn", "", "The ARN of the dashboard to update.")
	bcmDashboards_updateDashboardCmd.Flags().String("description", "", "The new description for the dashboard.")
	bcmDashboards_updateDashboardCmd.Flags().String("name", "", "The new name for the dashboard.")
	bcmDashboards_updateDashboardCmd.Flags().String("widgets", "", "The updated array of widget configurations for the dashboard.")
	bcmDashboards_updateDashboardCmd.MarkFlagRequired("arn")
	bcmDashboardsCmd.AddCommand(bcmDashboards_updateDashboardCmd)
}

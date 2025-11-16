package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_createDashboardCmd = &cobra.Command{
	Use:   "create-dashboard",
	Short: "Creates a new dashboard that can contain multiple widgets displaying cost and usage data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_createDashboardCmd).Standalone()

	bcmDashboards_createDashboardCmd.Flags().String("description", "", "A description of the dashboard's purpose or contents.")
	bcmDashboards_createDashboardCmd.Flags().String("name", "", "The name of the dashboard.")
	bcmDashboards_createDashboardCmd.Flags().String("resource-tags", "", "The tags to apply to the dashboard resource for organization and management.")
	bcmDashboards_createDashboardCmd.Flags().String("widgets", "", "An array of widget configurations that define the visualizations to be displayed in the dashboard.")
	bcmDashboards_createDashboardCmd.MarkFlagRequired("name")
	bcmDashboards_createDashboardCmd.MarkFlagRequired("widgets")
	bcmDashboardsCmd.AddCommand(bcmDashboards_createDashboardCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_startDashboardRefreshCmd = &cobra.Command{
	Use:   "start-dashboard-refresh",
	Short: "Starts a refresh of the specified dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_startDashboardRefreshCmd).Standalone()

	cloudtrail_startDashboardRefreshCmd.Flags().String("dashboard-id", "", "The name or ARN of the dashboard.")
	cloudtrail_startDashboardRefreshCmd.Flags().String("query-parameter-values", "", "The query parameter values for the dashboard")
	cloudtrail_startDashboardRefreshCmd.MarkFlagRequired("dashboard-id")
	cloudtrailCmd.AddCommand(cloudtrail_startDashboardRefreshCmd)
}

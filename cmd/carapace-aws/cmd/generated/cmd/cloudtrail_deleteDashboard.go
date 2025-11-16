package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_deleteDashboardCmd = &cobra.Command{
	Use:   "delete-dashboard",
	Short: "Deletes the specified dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_deleteDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_deleteDashboardCmd).Standalone()

		cloudtrail_deleteDashboardCmd.Flags().String("dashboard-id", "", "The name or ARN for the dashboard.")
		cloudtrail_deleteDashboardCmd.MarkFlagRequired("dashboard-id")
	})
	cloudtrailCmd.AddCommand(cloudtrail_deleteDashboardCmd)
}

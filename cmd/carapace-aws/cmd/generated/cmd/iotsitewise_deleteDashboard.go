package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_deleteDashboardCmd = &cobra.Command{
	Use:   "delete-dashboard",
	Short: "Deletes a dashboard from IoT SiteWise Monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_deleteDashboardCmd).Standalone()

	iotsitewise_deleteDashboardCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_deleteDashboardCmd.Flags().String("dashboard-id", "", "The ID of the dashboard to delete.")
	iotsitewise_deleteDashboardCmd.MarkFlagRequired("dashboard-id")
	iotsitewiseCmd.AddCommand(iotsitewise_deleteDashboardCmd)
}

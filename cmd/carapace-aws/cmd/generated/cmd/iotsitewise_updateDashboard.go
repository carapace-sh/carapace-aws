package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_updateDashboardCmd = &cobra.Command{
	Use:   "update-dashboard",
	Short: "Updates an IoT SiteWise Monitor dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_updateDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_updateDashboardCmd).Standalone()

		iotsitewise_updateDashboardCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
		iotsitewise_updateDashboardCmd.Flags().String("dashboard-definition", "", "The new dashboard definition, as specified in a JSON literal.")
		iotsitewise_updateDashboardCmd.Flags().String("dashboard-description", "", "A new description for the dashboard.")
		iotsitewise_updateDashboardCmd.Flags().String("dashboard-id", "", "The ID of the dashboard to update.")
		iotsitewise_updateDashboardCmd.Flags().String("dashboard-name", "", "A new friendly name for the dashboard.")
		iotsitewise_updateDashboardCmd.MarkFlagRequired("dashboard-definition")
		iotsitewise_updateDashboardCmd.MarkFlagRequired("dashboard-id")
		iotsitewise_updateDashboardCmd.MarkFlagRequired("dashboard-name")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_updateDashboardCmd)
}

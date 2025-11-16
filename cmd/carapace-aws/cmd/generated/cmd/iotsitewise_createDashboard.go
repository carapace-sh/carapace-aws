package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_createDashboardCmd = &cobra.Command{
	Use:   "create-dashboard",
	Short: "Creates a dashboard in an IoT SiteWise Monitor project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_createDashboardCmd).Standalone()

	iotsitewise_createDashboardCmd.Flags().String("client-token", "", "A unique case-sensitive identifier that you can provide to ensure the idempotency of the request.")
	iotsitewise_createDashboardCmd.Flags().String("dashboard-definition", "", "The dashboard definition specified in a JSON literal.")
	iotsitewise_createDashboardCmd.Flags().String("dashboard-description", "", "A description for the dashboard.")
	iotsitewise_createDashboardCmd.Flags().String("dashboard-name", "", "A friendly name for the dashboard.")
	iotsitewise_createDashboardCmd.Flags().String("project-id", "", "The ID of the project in which to create the dashboard.")
	iotsitewise_createDashboardCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the dashboard.")
	iotsitewise_createDashboardCmd.MarkFlagRequired("dashboard-definition")
	iotsitewise_createDashboardCmd.MarkFlagRequired("dashboard-name")
	iotsitewise_createDashboardCmd.MarkFlagRequired("project-id")
	iotsitewiseCmd.AddCommand(iotsitewise_createDashboardCmd)
}

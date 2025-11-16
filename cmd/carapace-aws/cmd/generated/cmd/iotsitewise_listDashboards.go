package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listDashboardsCmd = &cobra.Command{
	Use:   "list-dashboards",
	Short: "Retrieves a paginated list of dashboards for an IoT SiteWise Monitor project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listDashboardsCmd).Standalone()

	iotsitewise_listDashboardsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_listDashboardsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewise_listDashboardsCmd.Flags().String("project-id", "", "The ID of the project.")
	iotsitewise_listDashboardsCmd.MarkFlagRequired("project-id")
	iotsitewiseCmd.AddCommand(iotsitewise_listDashboardsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listDashboardsCmd = &cobra.Command{
	Use:   "list-dashboards",
	Short: "Returns information about all dashboards in the account, in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listDashboardsCmd).Standalone()

	cloudtrail_listDashboardsCmd.Flags().String("max-results", "", "The maximum number of dashboards to display on a single page.")
	cloudtrail_listDashboardsCmd.Flags().String("name-prefix", "", "Specify a name prefix to filter on.")
	cloudtrail_listDashboardsCmd.Flags().String("next-token", "", "A token you can use to get the next page of dashboard results.")
	cloudtrail_listDashboardsCmd.Flags().String("type", "", "Specify a dashboard type to filter on: `CUSTOM` or `MANAGED`.")
	cloudtrailCmd.AddCommand(cloudtrail_listDashboardsCmd)
}

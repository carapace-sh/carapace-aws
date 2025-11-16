package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_listDashboardsCmd = &cobra.Command{
	Use:   "list-dashboards",
	Short: "Returns a list of all dashboards in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_listDashboardsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDashboards_listDashboardsCmd).Standalone()

		bcmDashboards_listDashboardsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		bcmDashboards_listDashboardsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	bcmDashboardsCmd.AddCommand(bcmDashboards_listDashboardsCmd)
}

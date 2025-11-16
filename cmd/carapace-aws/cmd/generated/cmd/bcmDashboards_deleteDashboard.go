package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_deleteDashboardCmd = &cobra.Command{
	Use:   "delete-dashboard",
	Short: "Deletes a specified dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_deleteDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDashboards_deleteDashboardCmd).Standalone()

		bcmDashboards_deleteDashboardCmd.Flags().String("arn", "", "The ARN of the dashboard to be deleted.")
		bcmDashboards_deleteDashboardCmd.MarkFlagRequired("arn")
	})
	bcmDashboardsCmd.AddCommand(bcmDashboards_deleteDashboardCmd)
}

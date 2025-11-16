package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_getDashboardCmd = &cobra.Command{
	Use:   "get-dashboard",
	Short: "Retrieves the configuration and metadata of a specified dashboard, including its widgets and layout settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_getDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDashboards_getDashboardCmd).Standalone()

		bcmDashboards_getDashboardCmd.Flags().String("arn", "", "The ARN of the dashboard to retrieve.")
		bcmDashboards_getDashboardCmd.MarkFlagRequired("arn")
	})
	bcmDashboardsCmd.AddCommand(bcmDashboards_getDashboardCmd)
}

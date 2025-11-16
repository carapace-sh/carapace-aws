package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_putDashboardCmd = &cobra.Command{
	Use:   "put-dashboard",
	Short: "Creates a dashboard if it does not already exist, or updates an existing dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_putDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_putDashboardCmd).Standalone()

		cloudwatch_putDashboardCmd.Flags().String("dashboard-body", "", "The detailed information about the dashboard in JSON format, including the widgets to include and their location on the dashboard.")
		cloudwatch_putDashboardCmd.Flags().String("dashboard-name", "", "The name of the dashboard.")
		cloudwatch_putDashboardCmd.MarkFlagRequired("dashboard-body")
		cloudwatch_putDashboardCmd.MarkFlagRequired("dashboard-name")
	})
	cloudwatchCmd.AddCommand(cloudwatch_putDashboardCmd)
}

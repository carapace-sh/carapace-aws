package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_getDashboardCmd = &cobra.Command{
	Use:   "get-dashboard",
	Short: "Displays the details of the dashboard that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_getDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_getDashboardCmd).Standalone()

		cloudwatch_getDashboardCmd.Flags().String("dashboard-name", "", "The name of the dashboard to be described.")
		cloudwatch_getDashboardCmd.MarkFlagRequired("dashboard-name")
	})
	cloudwatchCmd.AddCommand(cloudwatch_getDashboardCmd)
}

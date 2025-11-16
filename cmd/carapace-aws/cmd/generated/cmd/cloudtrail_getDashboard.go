package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_getDashboardCmd = &cobra.Command{
	Use:   "get-dashboard",
	Short: "Returns the specified dashboard.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_getDashboardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_getDashboardCmd).Standalone()

		cloudtrail_getDashboardCmd.Flags().String("dashboard-id", "", "The name or ARN for the dashboard.")
		cloudtrail_getDashboardCmd.MarkFlagRequired("dashboard-id")
	})
	cloudtrailCmd.AddCommand(cloudtrail_getDashboardCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_deleteDashboardsCmd = &cobra.Command{
	Use:   "delete-dashboards",
	Short: "Deletes all dashboards that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_deleteDashboardsCmd).Standalone()

	cloudwatch_deleteDashboardsCmd.Flags().String("dashboard-names", "", "The dashboards to be deleted.")
	cloudwatch_deleteDashboardsCmd.MarkFlagRequired("dashboard-names")
	cloudwatchCmd.AddCommand(cloudwatch_deleteDashboardsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_listDashboardsCmd = &cobra.Command{
	Use:   "list-dashboards",
	Short: "Returns a list of the dashboards for your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_listDashboardsCmd).Standalone()

	cloudwatch_listDashboardsCmd.Flags().String("dashboard-name-prefix", "", "If you specify this parameter, only the dashboards with names starting with the specified string are listed.")
	cloudwatch_listDashboardsCmd.Flags().String("next-token", "", "The token returned by a previous call to indicate that there is more data available.")
	cloudwatchCmd.AddCommand(cloudwatch_listDashboardsCmd)
}

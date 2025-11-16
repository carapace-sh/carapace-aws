package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_getQueryStatusMonitorTopContributorsCmd = &cobra.Command{
	Use:   "get-query-status-monitor-top-contributors",
	Short: "Returns the current status of a query for the Network Flow Monitor query interface, for a specified query ID and monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_getQueryStatusMonitorTopContributorsCmd).Standalone()

	networkflowmonitor_getQueryStatusMonitorTopContributorsCmd.Flags().String("monitor-name", "", "The name of the monitor.")
	networkflowmonitor_getQueryStatusMonitorTopContributorsCmd.Flags().String("query-id", "", "The identifier for the query.")
	networkflowmonitor_getQueryStatusMonitorTopContributorsCmd.MarkFlagRequired("monitor-name")
	networkflowmonitor_getQueryStatusMonitorTopContributorsCmd.MarkFlagRequired("query-id")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_getQueryStatusMonitorTopContributorsCmd)
}

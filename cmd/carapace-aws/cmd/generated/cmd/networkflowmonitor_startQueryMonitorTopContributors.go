package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_startQueryMonitorTopContributorsCmd = &cobra.Command{
	Use:   "start-query-monitor-top-contributors",
	Short: "Create a query that you can use with the Network Flow Monitor query interface to return the top contributors for a monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_startQueryMonitorTopContributorsCmd).Standalone()

	networkflowmonitor_startQueryMonitorTopContributorsCmd.Flags().String("destination-category", "", "The category that you want to query top contributors for, for a specific monitor.")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.Flags().String("end-time", "", "The timestamp that is the date and time end of the period that you want to retrieve results for with your query.")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.Flags().String("limit", "", "The maximum number of top contributors to return.")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.Flags().String("metric-name", "", "The metric that you want to query top contributors for.")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.Flags().String("monitor-name", "", "The name of the monitor.")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.Flags().String("start-time", "", "The timestamp that is the date and time that is the beginning of the period that you want to retrieve results for with your query.")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.MarkFlagRequired("destination-category")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.MarkFlagRequired("end-time")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.MarkFlagRequired("metric-name")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.MarkFlagRequired("monitor-name")
	networkflowmonitor_startQueryMonitorTopContributorsCmd.MarkFlagRequired("start-time")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_startQueryMonitorTopContributorsCmd)
}

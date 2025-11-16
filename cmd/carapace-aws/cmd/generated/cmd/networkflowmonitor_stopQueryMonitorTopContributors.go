package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_stopQueryMonitorTopContributorsCmd = &cobra.Command{
	Use:   "stop-query-monitor-top-contributors",
	Short: "Stop a top contributors query for a monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_stopQueryMonitorTopContributorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_stopQueryMonitorTopContributorsCmd).Standalone()

		networkflowmonitor_stopQueryMonitorTopContributorsCmd.Flags().String("monitor-name", "", "The name of the monitor.")
		networkflowmonitor_stopQueryMonitorTopContributorsCmd.Flags().String("query-id", "", "The identifier for the query.")
		networkflowmonitor_stopQueryMonitorTopContributorsCmd.MarkFlagRequired("monitor-name")
		networkflowmonitor_stopQueryMonitorTopContributorsCmd.MarkFlagRequired("query-id")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_stopQueryMonitorTopContributorsCmd)
}

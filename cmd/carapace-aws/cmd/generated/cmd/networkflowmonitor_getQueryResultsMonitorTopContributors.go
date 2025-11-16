package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_getQueryResultsMonitorTopContributorsCmd = &cobra.Command{
	Use:   "get-query-results-monitor-top-contributors",
	Short: "Return the data for a query with the Network Flow Monitor query interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_getQueryResultsMonitorTopContributorsCmd).Standalone()

	networkflowmonitor_getQueryResultsMonitorTopContributorsCmd.Flags().String("max-results", "", "The number of query results that you want to return with this call.")
	networkflowmonitor_getQueryResultsMonitorTopContributorsCmd.Flags().String("monitor-name", "", "The name of the monitor.")
	networkflowmonitor_getQueryResultsMonitorTopContributorsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	networkflowmonitor_getQueryResultsMonitorTopContributorsCmd.Flags().String("query-id", "", "The identifier for the query.")
	networkflowmonitor_getQueryResultsMonitorTopContributorsCmd.MarkFlagRequired("monitor-name")
	networkflowmonitor_getQueryResultsMonitorTopContributorsCmd.MarkFlagRequired("query-id")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_getQueryResultsMonitorTopContributorsCmd)
}

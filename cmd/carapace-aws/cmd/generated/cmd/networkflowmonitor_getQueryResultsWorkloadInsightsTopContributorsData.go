package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd = &cobra.Command{
	Use:   "get-query-results-workload-insights-top-contributors-data",
	Short: "Return the data for a query with the Network Flow Monitor query interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd).Standalone()

		networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd.Flags().String("max-results", "", "The number of query results that you want to return with this call.")
		networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd.Flags().String("next-token", "", "The token for the next set of results.")
		networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd.Flags().String("query-id", "", "The identifier for the query.")
		networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
		networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("query-id")
		networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("scope-id")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsDataCmd)
}

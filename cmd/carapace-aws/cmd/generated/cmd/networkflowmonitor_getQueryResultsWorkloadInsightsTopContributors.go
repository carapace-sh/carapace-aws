package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd = &cobra.Command{
	Use:   "get-query-results-workload-insights-top-contributors",
	Short: "Return the data for a query with the Network Flow Monitor query interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd).Standalone()

	networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd.Flags().String("max-results", "", "The number of query results that you want to return with this call.")
	networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd.Flags().String("query-id", "", "The identifier for the query.")
	networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
	networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd.MarkFlagRequired("query-id")
	networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd.MarkFlagRequired("scope-id")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_getQueryResultsWorkloadInsightsTopContributorsCmd)
}

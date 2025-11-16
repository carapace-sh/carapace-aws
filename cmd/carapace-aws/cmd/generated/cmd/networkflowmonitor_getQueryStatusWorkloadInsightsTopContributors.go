package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsCmd = &cobra.Command{
	Use:   "get-query-status-workload-insights-top-contributors",
	Short: "Return the data for a query with the Network Flow Monitor query interface.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsCmd).Standalone()

		networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsCmd.Flags().String("query-id", "", "The identifier for the query.")
		networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
		networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsCmd.MarkFlagRequired("query-id")
		networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsCmd.MarkFlagRequired("scope-id")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsCmd)
}

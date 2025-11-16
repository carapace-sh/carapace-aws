package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsDataCmd = &cobra.Command{
	Use:   "get-query-status-workload-insights-top-contributors-data",
	Short: "Returns the current status of a query for the Network Flow Monitor query interface, for a specified query ID and monitor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsDataCmd).Standalone()

		networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsDataCmd.Flags().String("query-id", "", "The identifier for the query.")
		networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsDataCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
		networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("query-id")
		networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("scope-id")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_getQueryStatusWorkloadInsightsTopContributorsDataCmd)
}

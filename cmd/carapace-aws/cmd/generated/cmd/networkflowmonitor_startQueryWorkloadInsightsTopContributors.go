package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd = &cobra.Command{
	Use:   "start-query-workload-insights-top-contributors",
	Short: "Create a query with the Network Flow Monitor query interface that you can run to return workload insights top contributors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd).Standalone()

		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.Flags().String("destination-category", "", "The destination category for a top contributors row.")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.Flags().String("end-time", "", "The timestamp that is the date and time end of the period that you want to retrieve results for with your query.")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.Flags().String("limit", "", "The maximum number of top contributors to return.")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.Flags().String("metric-name", "", "The metric that you want to query top contributors for.")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.Flags().String("start-time", "", "The timestamp that is the date and time that is the beginning of the period that you want to retrieve results for with your query.")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.MarkFlagRequired("destination-category")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.MarkFlagRequired("end-time")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.MarkFlagRequired("metric-name")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.MarkFlagRequired("scope-id")
		networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd.MarkFlagRequired("start-time")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_startQueryWorkloadInsightsTopContributorsCmd)
}

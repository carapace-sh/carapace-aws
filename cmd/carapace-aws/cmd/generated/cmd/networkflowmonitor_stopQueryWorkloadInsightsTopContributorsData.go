package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_stopQueryWorkloadInsightsTopContributorsDataCmd = &cobra.Command{
	Use:   "stop-query-workload-insights-top-contributors-data",
	Short: "Stop a top contributors data query for workload insights.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_stopQueryWorkloadInsightsTopContributorsDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkflowmonitor_stopQueryWorkloadInsightsTopContributorsDataCmd).Standalone()

		networkflowmonitor_stopQueryWorkloadInsightsTopContributorsDataCmd.Flags().String("query-id", "", "The identifier for the query.")
		networkflowmonitor_stopQueryWorkloadInsightsTopContributorsDataCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
		networkflowmonitor_stopQueryWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("query-id")
		networkflowmonitor_stopQueryWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("scope-id")
	})
	networkflowmonitorCmd.AddCommand(networkflowmonitor_stopQueryWorkloadInsightsTopContributorsDataCmd)
}

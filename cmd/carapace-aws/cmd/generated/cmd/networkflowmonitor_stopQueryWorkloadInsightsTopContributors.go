package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_stopQueryWorkloadInsightsTopContributorsCmd = &cobra.Command{
	Use:   "stop-query-workload-insights-top-contributors",
	Short: "Stop a top contributors query for workload insights.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_stopQueryWorkloadInsightsTopContributorsCmd).Standalone()

	networkflowmonitor_stopQueryWorkloadInsightsTopContributorsCmd.Flags().String("query-id", "", "The identifier for the query.")
	networkflowmonitor_stopQueryWorkloadInsightsTopContributorsCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
	networkflowmonitor_stopQueryWorkloadInsightsTopContributorsCmd.MarkFlagRequired("query-id")
	networkflowmonitor_stopQueryWorkloadInsightsTopContributorsCmd.MarkFlagRequired("scope-id")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_stopQueryWorkloadInsightsTopContributorsCmd)
}

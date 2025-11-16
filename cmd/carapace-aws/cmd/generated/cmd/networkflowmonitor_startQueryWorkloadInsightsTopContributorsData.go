package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd = &cobra.Command{
	Use:   "start-query-workload-insights-top-contributors-data",
	Short: "Create a query with the Network Flow Monitor query interface that you can run to return data for workload insights top contributors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd).Standalone()

	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.Flags().String("destination-category", "", "The destination category for a top contributors.")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.Flags().String("end-time", "", "The timestamp that is the date and time end of the period that you want to retrieve results for with your query.")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.Flags().String("metric-name", "", "The metric that you want to query top contributors for.")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.Flags().String("scope-id", "", "The identifier for the scope that includes the resources you want to get data results for.")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.Flags().String("start-time", "", "The timestamp that is the date and time that is the beginning of the period that you want to retrieve results for with your query.")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("destination-category")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("end-time")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("metric-name")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("scope-id")
	networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd.MarkFlagRequired("start-time")
	networkflowmonitorCmd.AddCommand(networkflowmonitor_startQueryWorkloadInsightsTopContributorsDataCmd)
}

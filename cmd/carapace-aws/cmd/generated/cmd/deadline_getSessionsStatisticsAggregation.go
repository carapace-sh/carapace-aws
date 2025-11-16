package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getSessionsStatisticsAggregationCmd = &cobra.Command{
	Use:   "get-sessions-statistics-aggregation",
	Short: "Gets a set of statistics for queues or farms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getSessionsStatisticsAggregationCmd).Standalone()

	deadline_getSessionsStatisticsAggregationCmd.Flags().String("aggregation-id", "", "The identifier returned by the `StartSessionsStatisticsAggregation` operation that identifies the aggregated statistics.")
	deadline_getSessionsStatisticsAggregationCmd.Flags().String("farm-id", "", "The identifier of the farm to include in the statistics.")
	deadline_getSessionsStatisticsAggregationCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_getSessionsStatisticsAggregationCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_getSessionsStatisticsAggregationCmd.MarkFlagRequired("aggregation-id")
	deadline_getSessionsStatisticsAggregationCmd.MarkFlagRequired("farm-id")
	deadlineCmd.AddCommand(deadline_getSessionsStatisticsAggregationCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_startSessionsStatisticsAggregationCmd = &cobra.Command{
	Use:   "start-sessions-statistics-aggregation",
	Short: "Starts an asynchronous request for getting aggregated statistics about queues and farms.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_startSessionsStatisticsAggregationCmd).Standalone()

	deadline_startSessionsStatisticsAggregationCmd.Flags().String("end-time", "", "The Linux timestamp of the date and time that the statistics end.")
	deadline_startSessionsStatisticsAggregationCmd.Flags().String("farm-id", "", "The identifier of the farm that contains queues or fleets to return statistics for.")
	deadline_startSessionsStatisticsAggregationCmd.Flags().String("group-by", "", "The field to use to group the statistics.")
	deadline_startSessionsStatisticsAggregationCmd.Flags().String("period", "", "The period to aggregate the statistics.")
	deadline_startSessionsStatisticsAggregationCmd.Flags().String("resource-ids", "", "A list of fleet IDs or queue IDs to gather statistics for.")
	deadline_startSessionsStatisticsAggregationCmd.Flags().String("start-time", "", "The Linux timestamp of the date and time that the statistics start.")
	deadline_startSessionsStatisticsAggregationCmd.Flags().String("statistics", "", "One to four statistics to return.")
	deadline_startSessionsStatisticsAggregationCmd.Flags().String("timezone", "", "The timezone to use for the statistics.")
	deadline_startSessionsStatisticsAggregationCmd.MarkFlagRequired("end-time")
	deadline_startSessionsStatisticsAggregationCmd.MarkFlagRequired("farm-id")
	deadline_startSessionsStatisticsAggregationCmd.MarkFlagRequired("group-by")
	deadline_startSessionsStatisticsAggregationCmd.MarkFlagRequired("resource-ids")
	deadline_startSessionsStatisticsAggregationCmd.MarkFlagRequired("start-time")
	deadline_startSessionsStatisticsAggregationCmd.MarkFlagRequired("statistics")
	deadlineCmd.AddCommand(deadline_startSessionsStatisticsAggregationCmd)
}

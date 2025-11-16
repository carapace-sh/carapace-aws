package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns events related to clusters, cache security groups, and cache parameter groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_describeEventsCmd).Standalone()

	elasticache_describeEventsCmd.Flags().String("duration", "", "The number of minutes worth of events to retrieve.")
	elasticache_describeEventsCmd.Flags().String("end-time", "", "The end of the time interval for which to retrieve events, specified in ISO 8601 format.")
	elasticache_describeEventsCmd.Flags().String("marker", "", "An optional marker returned from a prior request.")
	elasticache_describeEventsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	elasticache_describeEventsCmd.Flags().String("source-identifier", "", "The identifier of the event source for which events are returned.")
	elasticache_describeEventsCmd.Flags().String("source-type", "", "The event source to retrieve events for.")
	elasticache_describeEventsCmd.Flags().String("start-time", "", "The beginning of the time interval to retrieve events for, specified in ISO 8601 format.")
	elasticacheCmd.AddCommand(elasticache_describeEventsCmd)
}

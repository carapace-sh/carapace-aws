package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns events related to DB instances, DB security groups, DB snapshots, and DB parameter groups for the past 14 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_describeEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_describeEventsCmd).Standalone()

		neptune_describeEventsCmd.Flags().String("duration", "", "The number of minutes to retrieve events for.")
		neptune_describeEventsCmd.Flags().String("end-time", "", "The end of the time interval for which to retrieve events, specified in ISO 8601 format.")
		neptune_describeEventsCmd.Flags().String("event-categories", "", "A list of event categories that trigger notifications for a event notification subscription.")
		neptune_describeEventsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
		neptune_describeEventsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous DescribeEvents request.")
		neptune_describeEventsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
		neptune_describeEventsCmd.Flags().String("source-identifier", "", "The identifier of the event source for which events are returned.")
		neptune_describeEventsCmd.Flags().String("source-type", "", "The event source to retrieve events for.")
		neptune_describeEventsCmd.Flags().String("start-time", "", "The beginning of the time interval to retrieve events for, specified in ISO 8601 format.")
	})
	neptuneCmd.AddCommand(neptune_describeEventsCmd)
}

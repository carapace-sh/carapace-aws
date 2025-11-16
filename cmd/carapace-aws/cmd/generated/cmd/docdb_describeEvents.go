package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdb_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns events related to instances, security groups, snapshots, and DB parameter groups for the past 14 days.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdb_describeEventsCmd).Standalone()

	docdb_describeEventsCmd.Flags().String("duration", "", "The number of minutes to retrieve events for.")
	docdb_describeEventsCmd.Flags().String("end-time", "", "The end of the time interval for which to retrieve events, specified in ISO 8601 format.")
	docdb_describeEventsCmd.Flags().String("event-categories", "", "A list of event categories that trigger notifications for an event notification subscription.")
	docdb_describeEventsCmd.Flags().String("filters", "", "This parameter is not currently supported.")
	docdb_describeEventsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	docdb_describeEventsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	docdb_describeEventsCmd.Flags().String("source-identifier", "", "The identifier of the event source for which events are returned.")
	docdb_describeEventsCmd.Flags().String("source-type", "", "The event source to retrieve events for.")
	docdb_describeEventsCmd.Flags().String("start-time", "", "The beginning of the time interval to retrieve events for, specified in ISO 8601 format.")
	docdbCmd.AddCommand(docdb_describeEventsCmd)
}

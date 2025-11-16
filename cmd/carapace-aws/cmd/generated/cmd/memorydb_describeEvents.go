package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeEventsCmd = &cobra.Command{
	Use:   "describe-events",
	Short: "Returns events related to clusters, security groups, and parameter groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_describeEventsCmd).Standalone()

		memorydb_describeEventsCmd.Flags().String("duration", "", "The number of minutes worth of events to retrieve.")
		memorydb_describeEventsCmd.Flags().String("end-time", "", "The end of the time interval for which to retrieve events, specified in ISO 8601 format.")
		memorydb_describeEventsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
		memorydb_describeEventsCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
		memorydb_describeEventsCmd.Flags().String("source-name", "", "The identifier of the event source for which events are returned.")
		memorydb_describeEventsCmd.Flags().String("source-type", "", "The event source to retrieve events for.")
		memorydb_describeEventsCmd.Flags().String("start-time", "", "The beginning of the time interval to retrieve events for, specified in ISO 8601 format.")
	})
	memorydbCmd.AddCommand(memorydb_describeEventsCmd)
}

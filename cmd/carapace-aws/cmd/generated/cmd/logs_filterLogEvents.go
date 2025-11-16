package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_filterLogEventsCmd = &cobra.Command{
	Use:   "filter-log-events",
	Short: "Lists log events from the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_filterLogEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_filterLogEventsCmd).Standalone()

		logs_filterLogEventsCmd.Flags().String("end-time", "", "The end of the time range, expressed as the number of milliseconds after `Jan 1, 1970 00:00:00 UTC`.")
		logs_filterLogEventsCmd.Flags().String("filter-pattern", "", "The filter pattern to use.")
		logs_filterLogEventsCmd.Flags().String("interleaved", "", "If the value is true, the operation attempts to provide responses that contain events from multiple log streams within the log group, interleaved in a single response.")
		logs_filterLogEventsCmd.Flags().String("limit", "", "The maximum number of events to return.")
		logs_filterLogEventsCmd.Flags().String("log-group-identifier", "", "Specify either the name or ARN of the log group to view log events from.")
		logs_filterLogEventsCmd.Flags().String("log-group-name", "", "The name of the log group to search.")
		logs_filterLogEventsCmd.Flags().String("log-stream-name-prefix", "", "Filters the results to include only events from log streams that have names starting with this prefix.")
		logs_filterLogEventsCmd.Flags().String("log-stream-names", "", "Filters the results to only logs from the log streams in this list.")
		logs_filterLogEventsCmd.Flags().String("next-token", "", "The token for the next set of events to return.")
		logs_filterLogEventsCmd.Flags().String("start-time", "", "The start of the time range, expressed as the number of milliseconds after `Jan 1, 1970 00:00:00 UTC`.")
		logs_filterLogEventsCmd.Flags().String("unmask", "", "Specify `true` to display the log event fields with all sensitive data unmasked and visible.")
	})
	logsCmd.AddCommand(logs_filterLogEventsCmd)
}

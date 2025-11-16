package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getLogEventsCmd = &cobra.Command{
	Use:   "get-log-events",
	Short: "Lists log events from the specified log stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getLogEventsCmd).Standalone()

	logs_getLogEventsCmd.Flags().String("end-time", "", "The end of the time range, expressed as the number of milliseconds after `Jan 1, 1970 00:00:00 UTC`.")
	logs_getLogEventsCmd.Flags().String("limit", "", "The maximum number of log events returned.")
	logs_getLogEventsCmd.Flags().String("log-group-identifier", "", "Specify either the name or ARN of the log group to view events from.")
	logs_getLogEventsCmd.Flags().String("log-group-name", "", "The name of the log group.")
	logs_getLogEventsCmd.Flags().String("log-stream-name", "", "The name of the log stream.")
	logs_getLogEventsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	logs_getLogEventsCmd.Flags().String("start-from-head", "", "If the value is true, the earliest log events are returned first.")
	logs_getLogEventsCmd.Flags().String("start-time", "", "The start of the time range, expressed as the number of milliseconds after `Jan 1, 1970 00:00:00 UTC`.")
	logs_getLogEventsCmd.Flags().String("unmask", "", "Specify `true` to display the log event fields with all sensitive data unmasked and visible.")
	logs_getLogEventsCmd.MarkFlagRequired("log-stream-name")
	logsCmd.AddCommand(logs_getLogEventsCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putLogEventsCmd = &cobra.Command{
	Use:   "put-log-events",
	Short: "Uploads a batch of log events to the specified log stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putLogEventsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putLogEventsCmd).Standalone()

		logs_putLogEventsCmd.Flags().String("entity", "", "The entity associated with the log events.")
		logs_putLogEventsCmd.Flags().String("log-events", "", "The log events.")
		logs_putLogEventsCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_putLogEventsCmd.Flags().String("log-stream-name", "", "The name of the log stream.")
		logs_putLogEventsCmd.Flags().String("sequence-token", "", "The sequence token obtained from the response of the previous `PutLogEvents` call.")
		logs_putLogEventsCmd.MarkFlagRequired("log-events")
		logs_putLogEventsCmd.MarkFlagRequired("log-group-name")
		logs_putLogEventsCmd.MarkFlagRequired("log-stream-name")
	})
	logsCmd.AddCommand(logs_putLogEventsCmd)
}

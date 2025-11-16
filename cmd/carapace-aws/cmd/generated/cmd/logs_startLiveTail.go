package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_startLiveTailCmd = &cobra.Command{
	Use:   "start-live-tail",
	Short: "Starts a Live Tail streaming session for one or more log groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_startLiveTailCmd).Standalone()

	logs_startLiveTailCmd.Flags().String("log-event-filter-pattern", "", "An optional pattern to use to filter the results to include only log events that match the pattern.")
	logs_startLiveTailCmd.Flags().String("log-group-identifiers", "", "An array where each item in the array is a log group to include in the Live Tail session.")
	logs_startLiveTailCmd.Flags().String("log-stream-name-prefixes", "", "If you specify this parameter, then only log events in the log streams that have names that start with the prefixes that you specify here are included in the Live Tail session.")
	logs_startLiveTailCmd.Flags().String("log-stream-names", "", "If you specify this parameter, then only log events in the log streams that you specify here are included in the Live Tail session.")
	logs_startLiveTailCmd.MarkFlagRequired("log-group-identifiers")
	logsCmd.AddCommand(logs_startLiveTailCmd)
}

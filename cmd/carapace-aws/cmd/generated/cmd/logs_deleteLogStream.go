package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteLogStreamCmd = &cobra.Command{
	Use:   "delete-log-stream",
	Short: "Deletes the specified log stream and permanently deletes all the archived log events associated with the log stream.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteLogStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteLogStreamCmd).Standalone()

		logs_deleteLogStreamCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_deleteLogStreamCmd.Flags().String("log-stream-name", "", "The name of the log stream.")
		logs_deleteLogStreamCmd.MarkFlagRequired("log-group-name")
		logs_deleteLogStreamCmd.MarkFlagRequired("log-stream-name")
	})
	logsCmd.AddCommand(logs_deleteLogStreamCmd)
}

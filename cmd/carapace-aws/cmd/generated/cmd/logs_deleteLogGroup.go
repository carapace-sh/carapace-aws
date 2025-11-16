package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteLogGroupCmd = &cobra.Command{
	Use:   "delete-log-group",
	Short: "Deletes the specified log group and permanently deletes all the archived log events associated with the log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteLogGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_deleteLogGroupCmd).Standalone()

		logs_deleteLogGroupCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_deleteLogGroupCmd.MarkFlagRequired("log-group-name")
	})
	logsCmd.AddCommand(logs_deleteLogGroupCmd)
}

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_createLogStreamCmd = &cobra.Command{
	Use:   "create-log-stream",
	Short: "Creates a log stream for the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_createLogStreamCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_createLogStreamCmd).Standalone()

		logs_createLogStreamCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_createLogStreamCmd.Flags().String("log-stream-name", "", "The name of the log stream.")
		logs_createLogStreamCmd.MarkFlagRequired("log-group-name")
		logs_createLogStreamCmd.MarkFlagRequired("log-stream-name")
	})
	logsCmd.AddCommand(logs_createLogStreamCmd)
}

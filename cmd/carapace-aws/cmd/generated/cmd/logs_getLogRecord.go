package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getLogRecordCmd = &cobra.Command{
	Use:   "get-log-record",
	Short: "Retrieves all of the fields and values of a single log event.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getLogRecordCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_getLogRecordCmd).Standalone()

		logs_getLogRecordCmd.Flags().String("log-record-pointer", "", "The pointer corresponding to the log event record you want to retrieve.")
		logs_getLogRecordCmd.Flags().String("unmask", "", "Specify `true` to display the log event fields with all sensitive data unmasked and visible.")
		logs_getLogRecordCmd.MarkFlagRequired("log-record-pointer")
	})
	logsCmd.AddCommand(logs_getLogRecordCmd)
}
